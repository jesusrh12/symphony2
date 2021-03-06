// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"context"
	"strconv"
	"testing"

	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/service"

	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	pkgmodels "github.com/facebookincubator/symphony/pkg/exporter/models"
	"github.com/facebookincubator/symphony/pkg/viewer/viewertest"
	"github.com/stretchr/testify/require"
)

const (
	serviceTypeName  = "serviceType"
	serviceType2Name = "serviceType2"
	serviceType3Name = "serviceType3"
	serviceType4Name = "serviceType4"
)

type serviceIds struct {
	serviceTypeID  int
	serviceTypeID2 int
	serviceTypeID3 int
	serviceTypeID4 int
}

var endpointHeader = [...]string{"Endpoint Definition 1", "Location 1", "Equipment 1",
	"Endpoint Definition 2", "Location 2", "Equipment 2", "Endpoint Definition 3", "Location 3", "Equipment 3",
	"Endpoint Definition 4", "Location 4", "Equipment 4", "Endpoint Definition 5", "Location 5", "Equipment 5",
}

func prepareServiceTypeData(ctx context.Context, t *testing.T, r testImporterResolver) serviceIds {
	mr := r.importer.r.Mutation()

	strDefVal := propDefValue
	propDefInput1 := pkgmodels.PropertyTypeInput{
		Name:        propName1,
		Type:        "string",
		StringValue: &strDefVal,
	}
	propDefInput2 := pkgmodels.PropertyTypeInput{
		Name: propName2,
		Type: "int",
	}
	propDefInput3 := pkgmodels.PropertyTypeInput{
		Name: propName3,
		Type: "date",
	}
	propDefInput4 := pkgmodels.PropertyTypeInput{
		Name: propName4,
		Type: "bool",
	}
	propDefInput5 := pkgmodels.PropertyTypeInput{
		Name: propName5,
		Type: "range",
	}
	propDefInput6 := pkgmodels.PropertyTypeInput{
		Name: propName6,
		Type: "gps_location",
	}
	propDefInput7 := pkgmodels.PropertyTypeInput{
		Name: propName7,
		Type: "node",
	}
	propDefInput8 := pkgmodels.PropertyTypeInput{
		Name: propName8,
		Type: "node",
	}

	serviceType1, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:       serviceTypeName,
		Properties: []*pkgmodels.PropertyTypeInput{&propDefInput1, &propDefInput2},
	})
	require.NoError(t, err)
	serviceType2, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:       serviceType2Name,
		Properties: []*pkgmodels.PropertyTypeInput{&propDefInput3, &propDefInput4},
	})
	require.NoError(t, err)
	serviceType3, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:       serviceType3Name,
		Properties: []*pkgmodels.PropertyTypeInput{&propDefInput5, &propDefInput6},
	})
	require.NoError(t, err)
	serviceType4, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:       serviceType4Name,
		Properties: []*pkgmodels.PropertyTypeInput{&propDefInput7, &propDefInput8},
	})
	require.NoError(t, err)
	return serviceIds{
		serviceTypeID:  serviceType1.ID,
		serviceTypeID2: serviceType2.ID,
		serviceTypeID3: serviceType3.ID,
		serviceTypeID4: serviceType4.ID,
	}
}

func TestValidatePropertiesForServiceType(t *testing.T) {
	r := newImporterTestResolver(t)
	importer := r.importer
	q := r.importer.r.Query()
	defer r.drv.Close()
	ctx := newImportContext(viewertest.NewContext(context.Background(), r.client))
	data := prepareServiceTypeData(ctx, t, *r)

	mr := r.importer.r.Mutation()
	locType, err := mr.AddLocationType(ctx, models.AddLocationTypeInput{Name: "City"})
	require.NoError(t, err)
	loc, err := mr.AddLocation(ctx, models.AddLocationInput{Name: "New York", Type: locType.ID})
	require.NoError(t, err)

	serviceType, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{Name: "L2 Access", HasCustomer: false})
	require.NoError(t, err)
	svc, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Service23",
		ServiceTypeID: serviceType.ID,
		Status:        service.StatusPending,
	})
	require.NoError(t, err)

	var (
		dataHeader = [...]string{"Service ID", "Service Name", "Service Type", "Discovery Method", "Service External ID", "Customer Name", "Customer External ID", "Status"}
		row1       = []string{"", "s1", serviceTypeName, "MANUAL", "M123", "", "", "IN_SERVICE", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "strVal", "54", "", "", "", "", "", ""}
		row2       = []string{"", "s2", serviceType2Name, "MANUAL", "M456", "", "", "MAINTENANCE", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "29/03/88", "false", "", "", "", ""}
		row3       = []string{"", "s3", serviceType3Name, "MANUAL", "M789", "", "", "DISCONNECTED", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "30.23-50", "45.8,88.9", "", ""}
		row4       = []string{"", "s3", serviceType4Name, "MANUAL", "M789", "", "", "DISCONNECTED", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", strconv.Itoa(loc.ID), strconv.Itoa(svc.ID)}
	)
	titleWithEndpoint := append(dataHeader[:], endpointHeader[:]...)
	titleWithProperties := append(titleWithEndpoint, propName1, propName2, propName3, propName4, propName5, propName6, propName7, propName8)
	fl, _ := NewImportHeader(titleWithProperties, ImportEntityService)
	r1, _ := NewImportRecord(row1, fl)
	require.NoError(t, err)
	node1, err := q.Node(ctx, data.serviceTypeID)
	require.NoError(t, err)
	styp1, ok := node1.(*ent.ServiceType)
	require.True(t, ok)
	ptypes, err := importer.validatePropertiesForServiceType(ctx, r1, styp1)
	require.NoError(t, err)
	require.Len(t, ptypes, 2)
	require.NotEqual(t, ptypes[0].PropertyTypeID, ptypes[1].PropertyTypeID)
	for _, value := range ptypes {
		ptyp := styp1.QueryPropertyTypes().Where(propertytype.ID(value.PropertyTypeID)).OnlyX(ctx)
		switch ptyp.Name {
		case propName1:
			require.Equal(t, *value.StringValue, "strVal")
			require.Equal(t, ptyp.Type, propertytype.TypeString)
		case propName2:
			require.Equal(t, *value.IntValue, 54)
			require.Equal(t, ptyp.Type, propertytype.TypeInt)
		default:
			require.Fail(t, "property type name should be one of the two")
		}
	}
	node2, err := q.Node(ctx, data.serviceTypeID2)
	require.NoError(t, err)
	styp2, ok := node2.(*ent.ServiceType)
	require.True(t, ok)

	r2, _ := NewImportRecord(row2, fl)
	ptypes2, err := importer.validatePropertiesForServiceType(ctx, r2, styp2)
	require.NoError(t, err)
	require.Len(t, ptypes2, 2)
	for _, value := range ptypes2 {
		ptyp := styp2.QueryPropertyTypes().Where(propertytype.ID(value.PropertyTypeID)).OnlyX(ctx)
		switch ptyp.Name {
		case propName3:
			require.Equal(t, *value.StringValue, "29/03/88")
			require.Equal(t, ptyp.Type, propertytype.TypeDate)
		case propName4:
			require.Equal(t, *value.BooleanValue, false)
			require.Equal(t, ptyp.Type, propertytype.TypeBool)
		default:
			require.Fail(t, "property type name should be one of the two")
		}
	}

	node3, err := q.Node(ctx, data.serviceTypeID3)
	require.NoError(t, err)
	styp3, ok := node3.(*ent.ServiceType)
	require.True(t, ok)

	r3, _ := NewImportRecord(row3, fl)
	ptypes3, err := importer.validatePropertiesForServiceType(ctx, r3, styp3)
	require.NoError(t, err)
	require.Len(t, ptypes3, 2)
	require.NotEqual(t, ptypes3[0].PropertyTypeID, ptypes3[1].PropertyTypeID)
	for _, value := range ptypes3 {
		ptyp := styp3.QueryPropertyTypes().Where(propertytype.ID(value.PropertyTypeID)).OnlyX(ctx)
		switch ptyp.Name {
		case propName5:
			require.Equal(t, *value.RangeFromValue, 30.23)
			require.EqualValues(t, *value.RangeToValue, 50)
			require.Equal(t, ptyp.Type, propertytype.TypeRange)
		case propName6:
			require.Equal(t, *value.LatitudeValue, 45.8)
			require.Equal(t, *value.LongitudeValue, 88.9)
			require.Equal(t, ptyp.Type, propertytype.TypeGpsLocation)
		default:
			require.Fail(t, "property type name should be one of the two")
		}
	}

	node4, err := q.Node(ctx, data.serviceTypeID4)
	require.NoError(t, err)
	styp4, ok := node4.(*ent.ServiceType)
	require.True(t, ok)

	r4, _ := NewImportRecord(row4, fl)
	ptypes4, err := importer.validatePropertiesForServiceType(ctx, r4, styp4)
	require.NoError(t, err)
	require.Len(t, ptypes4, 2)
	require.NotEqual(t, ptypes4[0].PropertyTypeID, ptypes4[1].PropertyTypeID)
	for _, value := range ptypes4 {
		ptyp := styp4.QueryPropertyTypes().Where(propertytype.ID(value.PropertyTypeID)).OnlyX(ctx)
		switch ptyp.Name {
		case propName7:
			require.Equal(t, *value.NodeIDValue, loc.ID)
			require.Equal(t, ptyp.Type, propertytype.TypeNode)
		case propName8:
			require.Equal(t, *value.NodeIDValue, svc.ID)
			require.Equal(t, ptyp.Type, propertytype.TypeNode)
		default:
			require.Fail(t, "property type name should be one of the two")
		}
	}
}

func TestValidateForExistingService(t *testing.T) {
	r := newImporterTestResolver(t)
	importer := r.importer
	defer r.drv.Close()
	ctx := newImportContext(viewertest.NewContext(context.Background(), r.client))
	prepareServiceTypeData(ctx, t, *r)

	titleWithProperties := []string{"Service ID", "Service Name", "Service Type", "Discovery Method", "Service External ID", "Customer Name", "Customer External ID", "Status"}
	title, _ := NewImportHeader(titleWithProperties, ImportEntityService)

	serviceType, err := importer.r.Mutation().AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: "type1",
	})
	require.NoError(t, err)
	svc, err := importer.r.Mutation().AddService(ctx, models.ServiceCreateData{
		Name:          "myService",
		ServiceTypeID: serviceType.ID,
		Status:        service.StatusPending,
	})
	require.NoError(t, err)
	var (
		test = []string{strconv.Itoa(svc.ID), "myService", "type1", "", "", "", "", service.StatusPending.String()}
	)
	rec, _ := NewImportRecord(test, title)
	_, err = importer.validateLineForExistingService(ctx, svc.ID, rec)
	require.NoError(t, err)
}
