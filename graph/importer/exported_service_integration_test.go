// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"bytes"
	"context"
	"encoding/csv"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent/property"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/service"
	pkgexporter "github.com/facebookincubator/symphony/pkg/exporter"
	pkgmodels "github.com/facebookincubator/symphony/pkg/exporter/models"
	"github.com/facebookincubator/symphony/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"github.com/facebookincubator/symphony/pkg/log/logtest"

	"github.com/facebookincubator/symphony/pkg/viewer"
	"github.com/facebookincubator/symphony/pkg/viewer/viewertest"

	"github.com/stretchr/testify/require"
)

const (
	serviceName  = "service"
	service2Name = "service2"
	service3Name = "service3"
)

// "Service ID", "Service Name", "Service Type",  "Discovery Method", "Service External ID", "Customer Name", "Customer External ID", "prop1", "prop2", "prop3", "prop4"
func editLine(t *testing.T, line []string, index int) {
	switch index {
	case 1:
		line[1] = "newName"
		line[4] = "D243"
		line[23] = "root"
		line[24] = "20"
	case 2:
		line[5] = "Donald"
		line[6] = "U333"
		line[25] = "22.4"
		line[26] = "true"
	case 3:
		line[1] = "newServiceName"
		line[4] = "D24345"
	default:
		t.Fatal("unexpected index")
	}
}

func writeModifiedCSV(t *testing.T, r *csv.Reader, method pkgexporter.Method, withVerify bool) (*bytes.Buffer, string) {
	var newLine []string
	var lines = make([][]string, 4)
	var buf bytes.Buffer
	bw := multipart.NewWriter(&buf)

	if withVerify {
		err := bw.WriteField("verify_before_commit", "true")
		require.NoError(t, err)
	}
	fileWriter, err := bw.CreateFormFile("file_0", "name1")
	require.Nil(t, err)
	for i := 0; ; i++ {
		line, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			require.Nil(t, err)
		}
		if i == 0 {
			lines[0] = line
		} else {
			switch method {
			case methodAdd:
				newLine = append([]string{""}, line[1:]...)
			case methodEdit:
				newLine = line
			default:
				require.Fail(t, "method should be add or edit")
			}
			editLine(t, newLine, i)
			lines[i] = newLine
		}
	}

	if withVerify {
		failLine := make([]string, len(lines[1]))
		copy(failLine, lines[1])
		lines = append(lines, failLine)
		lines[1][1] = "this"
		lines[1][2] = "should"
		lines[1][3] = "fail"
	}
	for _, l := range lines {
		stringLine := strings.Join(l, ",")
		_, _ = io.WriteString(fileWriter, stringLine+"\n")
	}
	ct := bw.FormDataContentType()
	require.NoError(t, bw.Close())
	return &buf, ct
}

func prepareServiceData(ctx context.Context, t *testing.T, r *testImporterResolver) {
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
		Type: "float",
	}
	propDefInput4 := pkgmodels.PropertyTypeInput{
		Name: propName4,
		Type: "bool",
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

	propertyType1, err := serviceType1.QueryPropertyTypes().Where(propertytype.Name(propName1)).Only(ctx)
	require.NoError(t, err)

	propertyType2, err := serviceType1.QueryPropertyTypes().Where(propertytype.Name(propName2)).Only(ctx)
	require.NoError(t, err)

	propertyType3, err := serviceType2.QueryPropertyTypes().Where(propertytype.Name(propName3)).Only(ctx)
	require.NoError(t, err)

	propertyType4, err := serviceType2.QueryPropertyTypes().Where(propertytype.Name(propName4)).Only(ctx)
	require.NoError(t, err)

	serviceStrProp := models.PropertyInput{PropertyTypeID: propertyType1.ID, StringValue: pointer.ToString("val")}
	serviceIntProp := models.PropertyInput{PropertyTypeID: propertyType2.ID, IntValue: pointer.ToInt(50)}

	service1PropInput := []*models.PropertyInput{&serviceStrProp, &serviceIntProp}

	_, err = mr.AddService(ctx, models.ServiceCreateData{
		Name:          serviceName,
		ServiceTypeID: serviceType1.ID,
		Properties:    service1PropInput,
		Status:        service.StatusPending,
	})
	require.NoError(t, err)

	serviceFloatProp := models.PropertyInput{PropertyTypeID: propertyType3.ID, FloatValue: pointer.ToFloat64(54.6)}
	serviceBoolProp := models.PropertyInput{PropertyTypeID: propertyType4.ID, BooleanValue: pointer.ToBool(false)}

	service2PropInput := []*models.PropertyInput{&serviceFloatProp, &serviceBoolProp}

	_, err = mr.AddService(ctx, models.ServiceCreateData{
		Name:          service2Name,
		ServiceTypeID: serviceType2.ID,
		Properties:    service2PropInput,
		Status:        service.StatusInService,
	})
	require.NoError(t, err)

	// no props case
	serviceType3, err := mr.AddServiceType(ctx, models.ServiceTypeCreateData{
		Name: service3Name,
	})
	require.NoError(t, err)

	_, err = mr.AddService(ctx, models.ServiceCreateData{
		Name:          service3Name,
		ServiceTypeID: serviceType3.ID,
		Status:        service.StatusPending,
	})
	require.NoError(t, err)
}

func deleteServiceData(ctx context.Context, t *testing.T, r *testImporterResolver) {
	id := r.client.Service.Query().Where(service.Name(serviceName)).OnlyIDX(ctx)
	_, err := r.importer.r.Mutation().RemoveService(ctx, id)
	require.NoError(t, err)

	id2 := r.client.Service.Query().Where(service.Name(service2Name)).OnlyIDX(ctx)
	_, err = r.importer.r.Mutation().RemoveService(ctx, id2)
	require.NoError(t, err)
}

func verifyServiceData(ctx context.Context, t *testing.T, r *testImporterResolver, withVerify bool) {
	s1, err := r.client.Service.Query().Where(service.Name("newName")).Only(ctx)
	if withVerify {
		require.Error(t, err)
		require.Nil(t, s1)
		return
	}
	require.NoError(t, err)

	require.Equal(t, "D243", *s1.ExternalID)
	require.Equal(t, service.StatusPending, s1.Status)

	prop1, err := s1.QueryProperties().Where(property.HasTypeWith(propertytype.TypeEQ(propertytype.TypeString))).Only(ctx)
	require.NoError(t, err)
	require.Equal(t, "root", pointer.GetString(prop1.StringVal))

	prop2, err := s1.QueryProperties().Where(property.HasTypeWith(propertytype.TypeEQ(propertytype.TypeInt))).Only(ctx)
	require.NoError(t, err)
	require.Equal(t, 20, pointer.GetInt(prop2.IntVal))

	s2, err := r.client.Service.Query().Where(service.Name(service2Name)).Only(ctx)
	require.NoError(t, err)
	customer, err := s2.QueryCustomer().Only(ctx)
	require.NoError(t, err)

	require.Equal(t, "Donald", customer.Name)
	require.Equal(t, "U333", *customer.ExternalID)
	require.Equal(t, service.StatusInService, s2.Status)

	prop3, err := s2.QueryProperties().Where(property.HasTypeWith(propertytype.TypeEQ(propertytype.TypeFloat))).Only(ctx)
	require.NoError(t, err)
	require.Equal(t, 22.4, pointer.GetFloat64(prop3.FloatVal))

	prop4, err := s2.QueryProperties().Where(property.HasTypeWith(propertytype.TypeEQ(propertytype.TypeBool))).Only(ctx)
	require.NoError(t, err)
	require.Equal(t, true, pointer.GetBool(prop4.BoolVal))

	s3, err := r.client.Service.Query().Where(service.Name("newServiceName")).Only(ctx)
	require.NoError(t, err)
	require.Equal(t, "D24345", *s3.ExternalID)
	require.Equal(t, service.StatusPending, s3.Status)
}

func exportServiceData(ctx context.Context, t *testing.T, r *testImporterResolver) bytes.Buffer {
	var buf bytes.Buffer
	core, _ := observer.New(zap.DebugLevel)
	log := log.NewDefaultLogger(zap.New(core))
	e := &pkgexporter.Exporter{Log: log, Rower: pkgexporter.ServicesRower{Log: log}}
	rows, err := e.Rower.Rows(ctx, "")
	require.NoError(t, err, "error getting rows")
	err = csv.NewWriter(&buf).WriteAll(rows)
	require.NoError(t, err, "error writing rows")
	return buf
}

func importServiceExportedData(ctx context.Context, t *testing.T, buf bytes.Buffer, contentType string, r *testImporterResolver) int {
	th := viewer.TenancyHandler(
		http.HandlerFunc(r.importer.processExportedService),
		viewer.NewFixedTenancy(r.client),
		logtest.NewTestLogger(t),
	)
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		th.ServeHTTP(w, r.WithContext(ctx))
	})
	server := httptest.NewServer(h)
	defer server.Close()

	req, err := http.NewRequest(http.MethodPost, server.URL, &buf)
	require.NoError(t, err)

	viewertest.SetDefaultViewerHeaders(req)
	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	code := resp.StatusCode
	resp.Body.Close()
	return code
}

func TestServiceImportDataAdd(t *testing.T) {
	for _, withVerify := range []bool{true, false} {
		r := newImporterTestResolver(t)
		ctx := newImportContext(viewertest.NewContext(context.Background(), r.client))
		prepareServiceData(ctx, t, r)
		exportedData := exportServiceData(ctx, t, r)
		deleteServiceData(ctx, t, r)
		readr := csv.NewReader(&exportedData)
		modifiedExportedData, contentType := writeModifiedCSV(t, readr, methodAdd, withVerify)
		code := importServiceExportedData(ctx, t, *modifiedExportedData, contentType, r)
		verifyServiceData(ctx, t, r, withVerify)
		require.Equal(t, http.StatusOK, code)
	}
}

func TestServiceImportDataEdit(t *testing.T) {
	for _, withVerify := range []bool{true, false} {
		r := newImporterTestResolver(t)
		ctx := newImportContext(viewertest.NewContext(context.Background(), r.client))
		prepareServiceData(ctx, t, r)
		exportedData := exportServiceData(ctx, t, r)
		readr := csv.NewReader(&exportedData)
		modifiedExportedData, contentType := writeModifiedCSV(t, readr, methodEdit, withVerify)
		code := importServiceExportedData(ctx, t, *modifiedExportedData, contentType, r)
		verifyServiceData(ctx, t, r, withVerify)
		require.Equal(t, http.StatusOK, code)
	}
}
