// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exporter

import (
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/location"
	"github.com/facebookincubator/symphony/pkg/ent/locationtype"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/property"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
	pkgmodels "github.com/facebookincubator/symphony/pkg/exporter/models"
	"github.com/pkg/errors"
)

func handleLocationFilter(q *ent.LocationQuery, filter *pkgmodels.LocationFilterInput) (*ent.LocationQuery, error) {
	switch filter.FilterType {
	case enum.LocationFilterTypeLocationInst:
		return LocationFilterPredicate(q, filter)
	case enum.LocationFilterTypeLocationInstHasEquipment:
		return locationHasEquipmentFilter(q, filter)
	case enum.LocationFilterTypeLocationInstName:
		return locationNameFilter(q, filter)
	case enum.LocationFilterTypeLocationInstExternalID:
		return locationExternalIDFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func locationExternalIDFilter(q *ent.LocationQuery, filter *pkgmodels.LocationFilterInput) (*ent.LocationQuery, error) {
	switch filter.Operator {
	case enum.FilterOperatorContains:
		return q.Where(location.ExternalIDContainsFold(*filter.StringValue)), nil
	case enum.FilterOperatorIs:
		return q.Where(location.ExternalID(*filter.StringValue)), nil
	}
	return nil, errors.Errorf("operation %s is not supported", filter.Operator)
}

func locationNameFilter(q *ent.LocationQuery, filter *pkgmodels.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.Operator == enum.FilterOperatorIs {
		return q.Where(location.NameEqualFold(*filter.StringValue)), nil
	}
	return nil, errors.Errorf("operation %s is not supported", filter.Operator)
}

func locationHasEquipmentFilter(q *ent.LocationQuery, filter *pkgmodels.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.Operator == enum.FilterOperatorIs {
		var pp predicate.Location
		if *filter.BoolValue {
			pp = location.HasEquipment()
		} else {
			pp = location.Not(location.HasEquipment())
		}
		return q.Where(pp), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func handleLocationTypeFilter(q *ent.LocationQuery, filter *pkgmodels.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.FilterType == enum.LocationFilterTypeLocationType {
		return locationLocationTypeFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func locationLocationTypeFilter(q *ent.LocationQuery, filter *pkgmodels.LocationFilterInput) (*ent.LocationQuery, error) {
	if filter.Operator == enum.FilterOperatorIsOneOf {
		return q.Where(location.HasTypeWith(locationtype.IDIn(filter.IDSet...))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

//nolint: dupl
func handleLocationPropertyFilter(q *ent.LocationQuery, filter *pkgmodels.LocationFilterInput) (*ent.LocationQuery, error) {
	p := filter.PropertyValue
	switch filter.Operator {
	case enum.FilterOperatorIs:
		pred, err := GetPropertyPredicate(*p)
		if err != nil {
			return nil, err
		}
		typePred, err := GetPropertyTypePredicate(*p)
		if err != nil {
			return nil, err
		}

		q = q.Where(location.Or(
			location.HasPropertiesWith(
				property.And(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.TypeEQ(p.Type),
					),
					pred,
				),
			),
			location.And(
				location.HasTypeWith(locationtype.HasPropertyTypesWith(
					propertytype.Name(p.Name),
					propertytype.TypeEQ(p.Type),
					typePred,
				)),
				location.Not(location.HasPropertiesWith(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.TypeEQ(p.Type),
					)),
				))))

		return q, nil
	case enum.FilterOperatorDateLessThan, enum.FilterOperatorDateGreaterThan:
		propPred, propTypePred, err := GetDatePropertyPred(*p, filter.Operator)
		if err != nil {
			return nil, err
		}
		q = q.Where(location.Or(
			location.HasPropertiesWith(
				property.And(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.TypeEQ(p.Type),
					),
					propPred,
				),
			),
			location.And(
				location.HasTypeWith(locationtype.HasPropertyTypesWith(
					propertytype.Name(p.Name),
					propertytype.TypeEQ(p.Type),
					propTypePred,
				)),
				location.Not(location.HasPropertiesWith(
					property.HasTypeWith(
						propertytype.Name(p.Name),
						propertytype.TypeEQ(p.Type),
					)),
				))))
		return q, nil
	default:
		return nil, errors.Errorf("operator %q not supported", filter.Operator)
	}
}
