package validator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"

	enum "fudjie.waizly/backend-test/library/tag/component/validate/enum"
	model "fudjie.waizly/backend-test/library/tag/component/validate/model"
)

// types: float32  float64

type floatValidator struct{}

func (vl *floatValidator) Validate(results []model.ValidatorResult, tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) ([]model.ValidatorResult, error) {
	newResults := append([]model.ValidatorResult{}, results...)

	for _, tagArgElem := range tagArgElems {
		if tagArgElem.ArgName == enum.TagArgNested {
			continue
		}

		addNewError := func() {
			newResults = append(newResults, model.NewValidatorResult(fieldName, fieldValue, tagArgElem.ArgName, tagArgElem.ArgValue))
		}

		if tagArgElem.ArgName == enum.TagArgRequired {
			if fieldValue == nil {
				addNewError()
			}
		} else if fieldValue != nil {
			fieldKind := fieldType.Kind()

			var dFieldValue decimal.Decimal
			if fieldKind == reflect.Float32 {
				dFieldValue = decimal.NewFromFloat32(fieldValue.(float32))
			} else {
				dFieldValue = decimal.NewFromFloat(fieldValue.(float64))
			}

			if tagArgElem.ArgName == enum.TagArgMin {
				if len(tagArgElem.ArgValue) == 0 {
					return nil, errors.New("min tag must have value")
				}

				dArgValue, err := decimal.NewFromString(tagArgElem.ArgValue)
				if err != nil {
					return nil, err
				}

				if dFieldValue.LessThan(dArgValue) {
					addNewError()
				}
			} else if tagArgElem.ArgName == enum.TagArgMax {
				if len(tagArgElem.ArgValue) == 0 {
					return nil, errors.New("max tag must have value")
				}

				dArgValue, err := decimal.NewFromString(tagArgElem.ArgValue)
				if err != nil {
					return nil, err
				}

				if dFieldValue.GreaterThan(dArgValue) {
					addNewError()
				}
			} else if tagArgElem.ArgName == enum.TagArgLongitude {
				if fieldKind != reflect.Float64 {
					return nil, errors.New("longitude must be in float64")
				}

			} else if tagArgElem.ArgName == enum.TagArgLatitude {
				if fieldKind != reflect.Float64 {
					return nil, errors.New("latitude must be in float64")
				}

			} else {
				return nil, fmt.Errorf("unsupported tag arg for float validator. field name: %s, arg: %s", fieldName, tagArgElem.ArgName)
			}
		}
	}

	return newResults, nil
}

var FloatValidator = &floatValidator{}
