package transformer

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"

	enum "fudjie.waizly/backend-test/library/tag/component/transform/enum"
	model "fudjie.waizly/backend-test/library/tag/component/transform/model"
)

// types: float32  float64

type floatTransformer struct{}

func (trf *floatTransformer) Transform(tagArgElems []model.TagArgElem, fieldRef reflect.Value, fieldName string, fieldType reflect.Type, fieldValue interface{}) error {
	var dFieldValue decimal.Decimal
	fieldKind := fieldType.Kind()
	if fieldKind == reflect.Float32 {
		dFieldValue = decimal.NewFromFloat32(fieldValue.(float32))
	} else {
		dFieldValue = decimal.NewFromFloat(fieldValue.(float64))
	}

	for _, tagArgElem := range tagArgElems {
		if tagArgElem.ArgName == enum.TagArgNested {
			continue
		}

		if tagArgElem.ArgName == enum.TagArgDecimal {
			if len(tagArgElem.ArgValue) == 0 {
				return errors.New("decimal tag must have value")
			}

			argValue, err := strconv.Atoi(tagArgElem.ArgValue)
			if err != nil {
				return err
			}

			exponent := dFieldValue.Exponent() * -1
			if exponent > int32(argValue) {
				dFieldValue = dFieldValue.RoundBank(int32(argValue))
				if fieldKind == reflect.Float32 {
					fieldRef.Set(reflect.ValueOf(float32(dFieldValue.InexactFloat64())))
				} else {
					fieldRef.Set(reflect.ValueOf(dFieldValue.InexactFloat64()))
				}
			}
		} else {
			return fmt.Errorf("unsupported tag arg for float transformer. field name: %s. arg: %s", fieldName, tagArgElem.ArgName)
		}
	}

	return nil
}

var FloatTransformer = &floatTransformer{}
