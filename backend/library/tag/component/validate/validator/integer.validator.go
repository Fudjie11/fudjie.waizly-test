package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	enum "fudjie.waizly/backend-test/library/tag/component/validate/enum"
	model "fudjie.waizly/backend-test/library/tag/component/validate/model"
)

// types: int  int8  int16  int32  int64

type integerValidator struct{}

func integerValidatorValidate[T int | int8 | int16 | int32 | int64](results []model.ValidatorResult, tagArgElems []model.TagArgElem, fieldName string, fieldType reflect.Type, fieldValue interface{}) ([]model.ValidatorResult, error) {
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
			cFieldValue := fieldValue.(T)

			if tagArgElem.ArgName == enum.TagArgMin {
				if len(tagArgElem.ArgValue) == 0 {
					return nil, errors.New("min tag must have value")
				}

				argValue, err := strconv.ParseInt(tagArgElem.ArgValue, 10, 64)
				if err != nil {
					return nil, err
				}

				if cFieldValue < T(argValue) {
					addNewError()
				}
			} else if tagArgElem.ArgName == enum.TagArgMax {
				if len(tagArgElem.ArgValue) == 0 {
					return nil, errors.New("max tag must have value")
				}

				argValue, err := strconv.ParseInt(tagArgElem.ArgValue, 10, 64)
				if err != nil {
					return nil, err
				}

				if cFieldValue > T(argValue) {
					addNewError()
				}
			} else {
				return nil, fmt.Errorf("unsupported tag arg for int validator. field name: %s, arg: %s", fieldName, tagArgElem.ArgName)
			}
		}
	}

	return newResults, nil
}

func (vl *integerValidator) Validate(results []model.ValidatorResult, tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) ([]model.ValidatorResult, error) {
	fieldTypeStr := fieldType.String()
	fieldKind := fieldType.Kind()

	if fieldKind == reflect.Int {
		return integerValidatorValidate[int](results, tagArgElems, fieldName, fieldType, fieldValue)
	} else if fieldKind == reflect.Int8 {
		return integerValidatorValidate[int8](results, tagArgElems, fieldName, fieldType, fieldValue)
	} else if fieldKind == reflect.Int16 {
		return integerValidatorValidate[int16](results, tagArgElems, fieldName, fieldType, fieldValue)
	} else if fieldKind == reflect.Int32 {
		return integerValidatorValidate[int32](results, tagArgElems, fieldName, fieldType, fieldValue)
	} else if fieldKind == reflect.Int64 {
		return integerValidatorValidate[int64](results, tagArgElems, fieldName, fieldType, fieldValue)
	} else {
		return nil, fmt.Errorf("invalid field kind for int validator. field tyle: %s", fieldTypeStr)
	}
}

var IntegerValidator = &integerValidator{}
