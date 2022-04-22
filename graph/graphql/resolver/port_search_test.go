// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentport"
	"github.com/facebookincubator/symphony/pkg/ent/equipmentportdefinition"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
	"github.com/facebookincubator/symphony/pkg/ent/service"
	"github.com/facebookincubator/symphony/pkg/ent/serviceendpointdefinition"
	pkgmodels "github.com/facebookincubator/symphony/pkg/exporter/models"
	"github.com/facebookincubator/symphony/pkg/viewer/viewertest"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/require"
)

const equipmentType1Port1Name = "typ1_p1"
const equipmentType1Port2Name = "typ1_p2"
const equipmentType2Port1Name = "typ2_p1"
const equipmentType2Port2Name = "typ2_p2"

type portSearchDataModels struct {
	typ1 *ent.EquipmentType
	typ2 *ent.EquipmentType
	e1   *ent.Equipment
	e2   *ent.Equipment
	e3   *ent.Equipment
	e4   *ent.Equipment
	loc1 int
	loc2 int
}

/*
	helper: data now is of type:
	loc1:
		e1(type1)[port: typ1_p1]  <--> e2(type1)[port: typ1_p2]
	loc2:
		e3(type2)[port: typ2_p1]
		e4(type2)[port: typ2_p2]
*/
func preparePortData(ctx context.Context, r *TestResolver) portSearchDataModels {
	mr := r.Mutation()
	locType1, _ := mr.AddLocationType(ctx, models.AddLocationTypeInput{
		Name: "loc_type1",
	})

	loc1, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name:       "loc_inst1",
		Type:       locType1.ID,
		ExternalID: pointer.ToString("111"),
	})

	loc2, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name: "loc_inst2",
		Type: locType1.ID,
	})
	ptyp, _ := mr.AddEquipmentPortType(ctx, models.AddEquipmentPortTypeInput{
		Name: "portType1",
		Properties: []*pkgmodels.PropertyTypeInput{
			{
				Name:        "propStr",
				Type:        "string",
				StringValue: pointer.ToString("t1"),
			},
			{
				Name: "propBool",
				Type: "bool",
			},
			{
				Name: "connected_date",
				Type: propertytype.TypeDate,
			},
		},
	})

	strProp := ptyp.QueryPropertyTypes().Where(propertytype.Name("propStr")).OnlyX(ctx)
	boolProp := ptyp.QueryPropertyTypes().Where(propertytype.Name("propBool")).OnlyX(ctx)
	dateProp := ptyp.QueryPropertyTypes().Where(propertytype.Name("connected_date")).OnlyX(ctx)
	equType1, _ := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "eq_type",
		Ports: []*models.EquipmentPortInput{
			{Name: equipmentType1Port1Name, PortTypeID: &ptyp.ID},
			{Name: equipmentType1Port2Name},
		},
	})
	defs1 := equType1.QueryPortDefinitions().AllX(ctx)
	equType2, _ := mr.AddEquipmentType(ctx, models.AddEquipmentTypeInput{
		Name: "eq_type2",
		Ports: []*models.EquipmentPortInput{
			{Name: equipmentType2Port1Name},
			{Name: equipmentType2Port2Name},
		},
	})

	e1, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst1",
		Type:     equType1.ID,
		Location: &loc1.ID,
	})

	def1 := equType1.QueryPortDefinitions().Where(equipmentportdefinition.Name("typ1_p1")).OnlyX(ctx)
	_, _ = mr.EditEquipmentPort(ctx, models.EditEquipmentPortInput{
		Side: &models.LinkSide{
			Equipment: e1.ID,
			Port:      def1.ID,
		},
		Properties: []*models.PropertyInput{
			{
				PropertyTypeID: strProp.ID,
				StringValue:    pointer.ToString("newVal"),
			},
			{
				PropertyTypeID: boolProp.ID,
				BooleanValue:   pointer.ToBool(true),
			},
			{
				PropertyTypeID: dateProp.ID,
				StringValue:    pointer.ToString("1988-03-29"),
			},
		},
	})

	e2, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst2",
		Type:     equType1.ID,
		Location: &loc1.ID,
	})
	e3, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst3",
		Type:     equType2.ID,
		Location: &loc2.ID,
	})
	e4, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     "eq_inst4",
		Type:     equType2.ID,
		Location: &loc2.ID,
	})
	_, _ = mr.AddLink(ctx, models.AddLinkInput{
		Sides: []*models.LinkSide{
			{Equipment: e1.ID, Port: defs1[0].ID},
			{Equipment: e2.ID, Port: defs1[0].ID},
		},
	})
	return portSearchDataModels{
		equType1,
		equType2,
		e1,
		e2,
		e3,
		e4,
		loc1.ID,
		loc2.ID,
	}
}

func TestSearchPortEquipmentName(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	data := preparePortData(ctx, r)
	qr := r.Query()
	limit := 100
	all, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{})
	require.NoError(t, err)
	require.Len(t, all.Edges, 8)
	require.Equal(t, all.TotalCount, 8)
	maxDepth := 2
	f1 := pkgmodels.PortFilterInput{
		FilterType:  enum.PortFilterTypePortInstEquipment,
		Operator:    enum.FilterOperatorContains,
		StringValue: pointer.ToString(data.e1.Name),
		MaxDepth:    &maxDepth,
	}
	res1, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f1})
	require.NoError(t, err)
	ports := res1.Edges
	require.Len(t, ports, 2)
}

func TestSearchPortHasLink(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	preparePortData(ctx, r)
	qr := r.Query()
	limit := 100
	all, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{})
	require.NoError(t, err)
	require.Len(t, all.Edges, 8)
	require.Equal(t, all.TotalCount, 8)
	f1 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypePortInstHasLink,
		Operator:   enum.FilterOperatorIs,
		BoolValue:  pointer.ToBool(false),
	}
	res1, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f1})
	require.NoError(t, err)
	ports := res1.Edges
	require.Len(t, ports, 6)
}

func TestSearchPortDefinition(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	d := preparePortData(ctx, r)

	qr := r.Query()
	limit := 100
	defs := d.typ1.QueryPortDefinitions().AllX(ctx)

	f1 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypePortDef,
		Operator:   enum.FilterOperatorIsOneOf,
		IDSet:      []int{defs[0].ID, defs[1].ID},
	}
	res1, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f1})
	require.NoError(t, err)
	ports := res1.Edges
	require.Len(t, ports, 4)
}

func TestSearchPortLocation(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	d := preparePortData(ctx, r)
	qr := r.Query()
	limit := 100

	f1 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeLocationInst,
		Operator:   enum.FilterOperatorIsOneOf,
		IDSet:      []int{d.loc1},
		MaxDepth:   pointer.ToInt(2),
	}
	res1, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f1})
	require.NoError(t, err)
	ports := res1.Edges
	require.Len(t, ports, 4)

	fExternal := pkgmodels.PortFilterInput{
		FilterType:  enum.PortFilterTypeLocationInstExternalID,
		Operator:    enum.FilterOperatorContains,
		StringValue: pointer.ToString("1"),
	}
	res1, err = qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&fExternal})
	require.NoError(t, err)
	ports = res1.Edges
	require.Len(t, ports, 4)
}

func TestSearchPortProperties(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	preparePortData(ctx, r)

	qr := r.Query()
	limit := 100

	f1 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeProperty,
		Operator:   enum.FilterOperatorIs,
		PropertyValue: &pkgmodels.PropertyTypeInput{
			Name:        "propStr",
			Type:        propertytype.TypeString,
			StringValue: pointer.ToString("t1"),
		},
	}

	res1, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f1})
	require.NoError(t, err)
	ports := res1.Edges
	require.Len(t, ports, 1)

	f2 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeProperty,
		Operator:   enum.FilterOperatorIs,
		PropertyValue: &pkgmodels.PropertyTypeInput{
			Name:        "propStr",
			Type:        propertytype.TypeString,
			StringValue: pointer.ToString("newVal"),
		},
	}

	res2, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f2})
	require.NoError(t, err)
	ports = res2.Edges
	require.Len(t, ports, 1)

	f3 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeProperty,
		Operator:   enum.FilterOperatorIs,
		PropertyValue: &pkgmodels.PropertyTypeInput{
			Name:         "propBool",
			Type:         propertytype.TypeBool,
			BooleanValue: pointer.ToBool(true),
		},
	}

	res3, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f3})
	require.NoError(t, err)
	ports = res3.Edges
	require.Len(t, ports, 1)

	f4 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeProperty,
		Operator:   enum.FilterOperatorIs,
		PropertyValue: &pkgmodels.PropertyTypeInput{
			Name:         "propBool",
			Type:         propertytype.TypeBool,
			BooleanValue: pointer.ToBool(false),
		},
	}

	res4, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f4})
	require.NoError(t, err)
	ports = res4.Edges
	require.Len(t, ports, 0)

	f5 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeProperty,
		Operator:   enum.FilterOperatorDateLessThan,
		PropertyValue: &pkgmodels.PropertyTypeInput{
			Name:        "connected_date",
			Type:        propertytype.TypeDate,
			StringValue: pointer.ToString("2019-01-01"),
		},
	}

	res5, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f5})
	require.NoError(t, err)
	ports = res5.Edges
	require.Len(t, ports, 1)
}

func TestSearchPortsByService(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	data := preparePortData(ctx, r)

	qr, mr := r.Query(), r.Mutation()

	port1, err := data.e1.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.Name(equipmentType1Port1Name))).Only(ctx)
	require.NoError(t, err)
	port2, err := data.e1.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.Name(equipmentType1Port2Name))).Only(ctx)
	require.NoError(t, err)
	port3, err := data.e3.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.Name(equipmentType2Port1Name))).Only(ctx)
	require.NoError(t, err)

	st, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:        "Service Type",
		HasCustomer: false,
		Endpoints: []*models.ServiceEndpointDefinitionInput{
			{
				Name:            "endpoint type1",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ1.ID,
				Index:           0,
			},
			{
				Name:            "endpoint type2",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ2.ID,
				Index:           1,
			},
		}})
	require.NoError(t, err)

	s1, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Service Instance 1",
		ServiceTypeID: st.ID,
		Status:        service.StatusPending,
	})
	require.NoError(t, err)

	ept0 := st.QueryEndpointDefinitions().Where(serviceendpointdefinition.Index(0)).OnlyX(ctx)
	ept1 := st.QueryEndpointDefinitions().Where(serviceendpointdefinition.Index(1)).OnlyX(ctx)

	_, err = mr.AddServiceEndpoint(ctx, models.AddServiceEndpointInput{
		ID:          s1.ID,
		EquipmentID: data.e1.ID,
		PortID:      pointer.ToInt(port1.ID),
		Definition:  ept0.ID,
	})
	require.NoError(t, err)

	s2, err := mr.AddService(ctx, models.ServiceCreateData{
		Name:          "Service Instance 2",
		ServiceTypeID: st.ID,
		Status:        service.StatusPending,
	})
	require.NoError(t, err)
	_, err = mr.AddServiceEndpoint(ctx, models.AddServiceEndpointInput{
		ID:          s2.ID,
		EquipmentID: data.e1.ID,
		PortID:      pointer.ToInt(port1.ID),
		Definition:  ept0.ID,
	})
	require.NoError(t, err)

	_, err = mr.AddServiceEndpoint(ctx, models.AddServiceEndpointInput{
		ID:          s2.ID,
		EquipmentID: data.e1.ID,
		PortID:      pointer.ToInt(port2.ID),
		Definition:  ept0.ID,
	})
	require.NoError(t, err)

	_, err = mr.AddServiceEndpoint(ctx, models.AddServiceEndpointInput{
		ID:          s2.ID,
		EquipmentID: data.e3.ID,
		PortID:      pointer.ToInt(port3.ID),
		Definition:  ept1.ID,
	})
	require.NoError(t, err)

	limit := 100
	all, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{})
	require.NoError(t, err)
	require.Len(t, all.Edges, 8)
	maxDepth := 2

	f1 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeServiceInst,
		Operator:   enum.FilterOperatorIsOneOf,
		IDSet:      []int{s1.ID},
		MaxDepth:   &maxDepth,
	}
	res1, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f1})
	require.NoError(t, err)
	require.Len(t, res1.Edges, 1)
	require.Equal(t, res1.Edges[0].Node.ID, port1.ID)

	f2 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeServiceInst,
		Operator:   enum.FilterOperatorIsOneOf,
		IDSet:      []int{s2.ID},
		MaxDepth:   &maxDepth,
	}
	res2, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f2})
	require.NoError(t, err)
	require.Len(t, res2.Edges, 3)

	f3 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeServiceInst,
		Operator:   enum.FilterOperatorIsNotOneOf,
		IDSet:      []int{s1.ID},
		MaxDepth:   &maxDepth,
	}
	res3, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f3})
	require.NoError(t, err)
	require.Len(t, res3.Edges, 7)

	f4 := pkgmodels.PortFilterInput{
		FilterType: enum.PortFilterTypeServiceInst,
		Operator:   enum.FilterOperatorIsNotOneOf,
		IDSet:      []int{s2.ID},
		MaxDepth:   &maxDepth,
	}
	res4, err := qr.EquipmentPorts(ctx, nil, &limit, nil, nil, []*pkgmodels.PortFilterInput{&f4})
	require.NoError(t, err)
	require.Len(t, res4.Edges, 5)
}

func TestReorderEndpointDefinitions(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	data := preparePortData(ctx, r)

	mr := r.Mutation()

	st, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name:        "Service Type",
		HasCustomer: false,
		Endpoints: []*models.ServiceEndpointDefinitionInput{
			{
				Name:            "nameX",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ1.ID,
				Index:           0,
			},
			{
				Name:            "nameY",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ2.ID,
				Index:           1,
			},
			{
				Name:            "nameZ",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ1.ID,
				Index:           2,
			},
			{
				Name:            "nameW",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ2.ID,
				Index:           3,
			},
		}})
	require.NoError(t, err)
	epd := st.QueryEndpointDefinitions().Order(ent.Asc(serviceendpointdefinition.FieldIndex)).AllX(ctx)
	for i, v := range epd {
		require.Equal(t, i, v.Index)
	}
	st, err = mr.EditServiceType(ctx, models.ServiceTypeEditData{
		ID:          st.ID,
		Name:        "Service Type",
		HasCustomer: false,
		Endpoints: []*models.ServiceEndpointDefinitionInput{
			{
				ID:              pointer.ToInt(epd[0].ID),
				Name:            "new name 3",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ1.ID,
				Index:           3,
			},
			{
				ID:              pointer.ToInt(epd[1].ID),
				Name:            "new name 2",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ2.ID,
				Index:           2,
			},
			{
				ID:              pointer.ToInt(epd[2].ID),
				Name:            "new name 1",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ1.ID,
				Index:           1,
			},
			{
				ID:              pointer.ToInt(epd[3].ID),
				Name:            "new name 0",
				Role:            pointer.ToString("CONSUMER"),
				EquipmentTypeID: data.typ2.ID,
				Index:           0,
			},
		}})
	require.NoError(t, err)

	epd = st.QueryEndpointDefinitions().AllX(ctx)
	for _, v := range epd {
		require.Contains(t, v.Name, strconv.Itoa(v.Index))
	}
}
