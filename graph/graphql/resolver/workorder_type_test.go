// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver_test

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/AlekSi/pointer"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/checklistcategorydefinition"
	"github.com/facebookincubator/symphony/pkg/ent/checklistitemdefinition"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
	pkgmodels "github.com/facebookincubator/symphony/pkg/exporter/models"
	"github.com/facebookincubator/symphony/pkg/viewer/viewertest"
	"github.com/stretchr/testify/require"
)

func TestAddWorkOrderType(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)
	mr, qr := r.Mutation(), r.Query()

	typ, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type"})
	require.NoError(t, err)

	node, err := qr.Node(ctx, typ.ID)
	require.NoError(t, err)
	typ, ok := node.(*ent.WorkOrderType)
	require.True(t, ok)
	require.Equal(t, typ.Name, "example_type", "verifying work order type name")
}

func TestAddWorkOrderTypes(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)
	mr, qr := r.Mutation(), r.Query()

	_, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_a"})
	require.NoError(t, err)
	_, err = mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_b"})
	require.NoError(t, err)

	types, _ := qr.WorkOrderTypes(ctx, nil, nil, nil, nil)
	require.Len(t, types.Edges, 2, "verify the added work order types are fetched properly")
}

func TestNumberOfWorkOrders(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)
	mr, wtr := r.Mutation(), r.WorkOrderType()

	woType, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type"})
	require.NoError(t, err)

	numWO, err := wtr.NumberOfWorkOrders(ctx, woType)
	require.NoError(t, err)
	require.Equal(t, 0, numWO)

	workOrder, err := mr.AddWorkOrder(ctx, models.AddWorkOrderInput{
		Name: "foo", WorkOrderTypeID: woType.ID,
	})
	require.NoError(t, err)

	numWO, err = wtr.NumberOfWorkOrders(ctx, woType)
	require.NoError(t, err)
	require.Equal(t, 1, numWO)

	_, err = mr.RemoveWorkOrder(ctx, workOrder.ID)
	require.NoError(t, err)

	numWO, err = wtr.NumberOfWorkOrders(ctx, woType)
	require.NoError(t, err)
	require.Equal(t, 0, numWO)
}

func TestAddWorkOrderTypeWithDescription(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)
	mr, qr := r.Mutation(), r.Query()

	typ, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{
		Name:        "example_type",
		Description: pointer.ToString("wo_type_desc"),
	})
	require.NoError(t, err)

	node, err := qr.Node(ctx, typ.ID)
	require.NoError(t, err)
	typ, ok := node.(*ent.WorkOrderType)
	require.True(t, ok)
	require.Equal(t, *typ.Description, "wo_type_desc", "verifying work order type description")
}

func TestAddWorkOrderTypeWithProperties(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	mr := r.Mutation()
	strValue, strIndex := "Foo", 7
	intValue, intIndex := 5, 12

	strPropType := pkgmodels.PropertyTypeInput{
		Name:        "str_prop",
		Type:        "string",
		Index:       &strIndex,
		StringValue: &strValue,
	}
	intPropType := pkgmodels.PropertyTypeInput{
		Name:     "int_prop",
		Type:     "int",
		Index:    &intIndex,
		IntValue: &intValue,
	}
	propTypeInputs := []*pkgmodels.PropertyTypeInput{&strPropType, &intPropType}
	woType, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{
		Name:       "example_type_a",
		Properties: propTypeInputs,
	})
	require.NoError(t, err)

	intProp := woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeInt)).OnlyX(ctx)
	strProp := woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeString)).OnlyX(ctx)

	require.Equal(t, "int_prop", intProp.Name, "verifying int property type's name")
	require.Nil(t, intProp.StringVal, "verifying int property type's string value (default as this is an int property)")
	require.Equal(t, intValue, pointer.GetInt(intProp.IntVal), "verifying int property type's int value")
	require.Equal(t, intIndex, intProp.Index, "verifying int property type's index")
	require.Equal(t, "str_prop", strProp.Name, "verifying string property type's name")
	require.Equal(t, strValue, pointer.GetString(strProp.StringVal), "verifying string property type's String value")
	require.Nil(t, strProp.IntVal, "verifying int property type's int value")
	require.Equal(t, strIndex, strProp.Index, "verifying string property type's index")

	pt, err := woType.PropertyTypes(ctx)
	require.NoError(t, err)
	require.Len(t, pt, 2)
}

func TestAddWorkOrderTypeWithCheckListCategories(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	c := r.GraphClient()

	selectionMode := enum.CheckListItemEnumSelectionModeSingle
	woTypeInput := models.AddWorkOrderTypeInput{
		Name: "WO Type",
		CheckListCategories: []*models.CheckListCategoryDefinitionInput{
			{
				Title:       "Category 1",
				Description: pointer.ToString("Category 1"),
				CheckList: []*models.CheckListDefinitionInput{{
					Title: "String",
					Type:  "string",
					Index: pointer.ToInt(0),
				}, {
					Title:             "Single Choice",
					Type:              "enum",
					Index:             pointer.ToInt(1),
					EnumSelectionMode: &selectionMode,
					EnumValues:        pointer.ToString("1,2,3"),
				}},
			},
		},
	}
	var rsp struct {
		AddWorkOrderType struct {
			ID                           string
			CheckListCategoryDefinitions []struct {
				ID                       string
				Title                    string
				ChecklistItemDefinitions []struct {
					ID                string
					Title             string
					Type              enum.CheckListItemType
					EnumValues        *string
					EnumSelectionMode *enum.CheckListItemEnumSelectionMode
				}
			}
		}
	}
	c.MustPost(
		`mutation($input: AddWorkOrderTypeInput!) {
			addWorkOrderType(input: $input) {
				id
				checkListCategoryDefinitions {
					id
					title
					checklistItemDefinitions {
						id
						title
						type
						enumValues
						enumSelectionMode
					}
				}
			}
		}`,
		&rsp,
		client.Var("input", woTypeInput),
	)

	require.Len(t, rsp.AddWorkOrderType.CheckListCategoryDefinitions, 1)
	category := rsp.AddWorkOrderType.CheckListCategoryDefinitions[0]
	require.Equal(t, category.Title, "Category 1")
	require.Len(t, category.ChecklistItemDefinitions, 2)

	for _, item := range category.ChecklistItemDefinitions {
		switch item.Type {
		case enum.CheckListItemTypeString:
			require.Equal(t, item.Title, "String")
		case enum.CheckListItemTypeEnum:
			require.Equal(t, *item.EnumSelectionMode, enum.CheckListItemEnumSelectionModeSingle)
			require.Equal(t, *item.EnumValues, "1,2,3")
		}
	}
}

func TestEditWorkOrderTypeWithCheckListCategories(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)
	mr := r.Mutation()

	woTypeInput := models.AddWorkOrderTypeInput{
		Name: "WO Type",
		CheckListCategories: []*models.CheckListCategoryDefinitionInput{
			{
				Title:       "Category 1",
				Description: pointer.ToString("Category 1"),
				CheckList: []*models.CheckListDefinitionInput{{
					Title: "String 1",
					Type:  "string",
					Index: pointer.ToInt(0),
				}, {
					Title: "String 2",
					Type:  "string",
					Index: pointer.ToInt(1),
				}},
			},
			{
				Title:       "Category 2",
				Description: pointer.ToString("Category 2"),
			},
		},
	}

	wot, err := mr.AddWorkOrderType(ctx, woTypeInput)
	require.NoError(t, err)

	category1 := wot.QueryCheckListCategoryDefinitions().Where(checklistcategorydefinition.Title("Category 1")).OnlyX(ctx)
	string1 := wot.QueryCheckListCategoryDefinitions().QueryCheckListItemDefinitions().Where(checklistitemdefinition.Title("String 1")).OnlyX(ctx)
	string2 := wot.QueryCheckListCategoryDefinitions().QueryCheckListItemDefinitions().Where(checklistitemdefinition.Title("String 2")).OnlyX(ctx)

	category2 := wot.QueryCheckListCategoryDefinitions().Where(checklistcategorydefinition.Title("Category 2")).OnlyX(ctx)

	// Category 1: 1 item renamed, 1 deleted, 1 new. Category 2 deleted. Category 3 - new.
	editWOTypeInput := models.EditWorkOrderTypeInput{
		ID:   wot.ID,
		Name: "WO Type",
		CheckListCategories: []*models.CheckListCategoryDefinitionInput{
			{
				ID:          &category1.ID,
				Title:       "Category 1 Renamed",
				Description: pointer.ToString("Category 1 Renamed"),
				CheckList: []*models.CheckListDefinitionInput{{
					ID:    &string1.ID,
					Title: "String 1 Renamed",
					Type:  "string",
					Index: pointer.ToInt(0),
				}, {
					Title: "String 3",
					Type:  "string",
					Index: pointer.ToInt(2),
				}},
			},
			{
				Title:       "Category 3",
				Description: pointer.ToString("Category 3"),
			},
		},
	}

	updatedWOT, err := mr.EditWorkOrderType(ctx, editWOTypeInput)
	require.NoError(t, err)

	updatedCategories, err := updatedWOT.QueryCheckListCategoryDefinitions().All(ctx)
	require.NoError(t, err)
	require.Len(t, updatedCategories, 2)

	// Verify Category 1 Renamed
	category1 = updatedWOT.QueryCheckListCategoryDefinitions().Where(checklistcategorydefinition.ID(category1.ID)).OnlyX(ctx)
	require.Equal(t, category1.Title, "Category 1 Renamed")

	// Verify Category 2 deleted
	category2Exists := updatedWOT.QueryCheckListCategoryDefinitions().Where(checklistcategorydefinition.ID(category2.ID)).ExistX(ctx)
	require.False(t, category2Exists)

	// Verify Category 3 created
	category3Exists, err := updatedWOT.QueryCheckListCategoryDefinitions().Where(checklistcategorydefinition.Title("Category 3")).Exist(ctx)
	require.NoError(t, err)
	require.True(t, category3Exists)

	// Verify category 1 items created
	category1Items, err := category1.QueryCheckListItemDefinitions().All(ctx)
	require.NoError(t, err)
	require.Len(t, category1Items, 2)

	// Verify item string1 renamed
	string1Exists := category1.QueryCheckListItemDefinitions().Where(checklistitemdefinition.Title("String 1 Renamed")).Where(checklistitemdefinition.ID(string1.ID)).ExistX(ctx)
	require.True(t, string1Exists)

	// Verify item string2 deleted
	string2Exists := category1.QueryCheckListItemDefinitions().Where(checklistitemdefinition.ID(string2.ID)).ExistX(ctx)
	require.False(t, string2Exists)

	// Verify item string3 created
	string3Exists, err := category1.QueryCheckListItemDefinitions().Where(checklistitemdefinition.Title("String 3")).Exist(ctx)
	require.NoError(t, err)
	require.True(t, string3Exists)
}

func TestAddWorkOrderTypesSameName(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	mr, qr := r.Mutation(), r.Query()
	woType, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_name"})
	require.NoError(t, err)
	require.Equal(t, "example_type_name", woType.Name)
	_, err = mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_name"})
	require.Error(t, err, "adding work order type with an existing work order type name yields an error")
	types, _ := qr.WorkOrderTypes(ctx, nil, nil, nil, nil)
	require.Len(t, types.Edges, 1, "two attempts to create work order types with same name will create one work order type")
}

func TestRemoveWorkOrderType(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	mr, qr := r.Mutation(), r.Query()
	woType, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_name"})
	require.NoError(t, err)
	require.Equal(t, "example_type_name", woType.Name)

	types, _ := qr.WorkOrderTypes(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	require.Len(t, types.Edges, 1)

	id, err := mr.RemoveWorkOrderType(ctx, woType.ID)
	require.NoError(t, err)
	require.Equal(t, woType.ID, id, "successfully remove work order type")
	types, err = qr.WorkOrderTypes(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	require.Empty(t, types.Edges, "no work order types exist after deletion")
}

func TestEditWorkOrderType(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	mr, qr := r.Mutation(), r.Query()
	woType, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_name"})
	require.NoError(t, err)
	require.Equal(t, "example_type_name", woType.Name)

	newType, err := mr.EditWorkOrderType(ctx, models.EditWorkOrderTypeInput{
		ID:   woType.ID,
		Name: "example_type_name_edited",
	})
	require.NoError(t, err)
	require.Equal(t, "example_type_name_edited", newType.Name, "successfully edited work order type name")

	woType, err = mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_name_2"})
	require.NoError(t, err)
	_, err = mr.EditWorkOrderType(ctx, models.EditWorkOrderTypeInput{
		ID:   woType.ID,
		Name: "example_type_name_edited",
	})
	require.Error(t, err, "duplicate names")

	types, err := qr.WorkOrderTypes(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	require.Len(t, types.Edges, 2)

	node, err := qr.Node(ctx, woType.ID)
	require.NoError(t, err)
	typ, ok := node.(*ent.WorkOrderType)
	require.True(t, ok)
	require.Equal(t, "example_type_name_2", typ.Name)
}

func TestEditWorkOrderTypeWithProperties(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	mr := r.Mutation()
	strValue := "Foo"
	strPropType := pkgmodels.PropertyTypeInput{
		Name:        "str_prop",
		Type:        "string",
		StringValue: &strValue,
	}
	propTypeInput := []*pkgmodels.PropertyTypeInput{&strPropType}
	woType, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_a", Properties: propTypeInput})
	require.NoError(t, err)

	strProp := woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeString)).OnlyX(ctx)
	strValue = "Foo - edited"
	intValue := 5
	strPropType = pkgmodels.PropertyTypeInput{
		ID:          &strProp.ID,
		Name:        "str_prop_new",
		Type:        "string",
		StringValue: &strValue,
	}
	intPropType := pkgmodels.PropertyTypeInput{
		Name:     "int_prop",
		Type:     "int",
		IntValue: &intValue,
	}
	editedPropTypeInput := []*pkgmodels.PropertyTypeInput{&strPropType, &intPropType}
	newType, err := mr.EditWorkOrderType(ctx, models.EditWorkOrderTypeInput{
		ID:         woType.ID,
		Name:       "example_type_a",
		Properties: editedPropTypeInput,
	})
	require.NoError(t, err)
	require.Equal(t, woType.Name, newType.Name, "successfully edited work order type name")

	strProp = woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeString)).OnlyX(ctx)
	require.Equal(t, "str_prop_new", strProp.Name, "successfully edited prop type name")
	require.Equal(t, strValue, pointer.GetString(strProp.StringVal), "successfully edited prop type string value")

	intProp := woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeInt)).OnlyX(ctx)
	require.Equal(t, "int_prop", intProp.Name, "successfully edited prop type name")
	require.Equal(t, intValue, pointer.GetInt(intProp.IntVal), "successfully edited prop type int value")

	intValue = 6
	intPropType = pkgmodels.PropertyTypeInput{
		Name:     "int_prop",
		Type:     "int",
		IntValue: &intValue,
	}
	editedPropTypeInput = []*pkgmodels.PropertyTypeInput{&intPropType}
	_, err = mr.EditWorkOrderType(ctx, models.EditWorkOrderTypeInput{
		ID:         woType.ID,
		Name:       "example_type_a",
		Properties: editedPropTypeInput,
	})
	require.Error(t, err, "duplicate property type names")
}

func TestDeleteWorkOrderTypeProperty(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	mr := r.Mutation()
	strValue := "Foo"
	strPropType := pkgmodels.PropertyTypeInput{
		Name:        "str_prop",
		Type:        "string",
		StringValue: &strValue,
	}
	propTypeInput := []*pkgmodels.PropertyTypeInput{&strPropType}
	woType, err := mr.AddWorkOrderType(ctx, models.AddWorkOrderTypeInput{Name: "example_type_a", Properties: propTypeInput})
	require.NoError(t, err)

	strProp := woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeString)).OnlyX(ctx)
	strPropType = pkgmodels.PropertyTypeInput{
		ID:          &strProp.ID,
		Name:        "str_prop",
		Type:        "string",
		StringValue: &strValue,
		IsDeleted:   pointer.ToBool(true),
	}

	strProp = woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeString)).OnlyX(ctx)
	require.False(t, strProp.Deleted, "successfully edited prop type name")

	editedPropTypeInput := []*pkgmodels.PropertyTypeInput{&strPropType}
	newType, err := mr.EditWorkOrderType(ctx, models.EditWorkOrderTypeInput{
		ID:         woType.ID,
		Name:       "example_type_a",
		Properties: editedPropTypeInput,
	})
	require.NoError(t, err)
	require.Equal(t, woType.Name, newType.Name, "successfully edited work order type name")

	strProp = woType.QueryPropertyTypes().Where(propertytype.TypeEQ(propertytype.TypeString)).OnlyX(ctx)
	require.True(t, strProp.Deleted, "successfully edited prop type name")
}
