package validate

import (
	"fmt"
	"reflect"
	"strings"

	"fudjie.waizly/backend-test/library/tag/component/transform/enum"
	"github.com/samber/lo"

	model "fudjie.waizly/backend-test/library/tag/component/transform/model"
	transformer "fudjie.waizly/backend-test/library/tag/component/transform/transformer"

	util "fudjie.waizly/backend-test/library/util"
)

var tagName = "transform"

func getTagArgElems(tagValue string) []model.TagArgElem {
	results := []model.TagArgElem{}
	tValue := strings.TrimSpace(tagValue)

	if len(tValue) == 0 {
		return results
	}

	tagArgs := strings.Split(tValue, ",")
	for _, ta := range tagArgs {
		argElems := strings.Split(ta, "=")
		elem := model.TagArgElem{
			ArgName: strings.TrimSpace(argElems[0]),
		}

		if len(argElems) > 1 {
			argValue := strings.TrimSpace(argElems[1])
			elem.ArgValue = argValue
		}

		results = append(results, elem)
	}

	return results
}

func transform(ref reflect.Value) error {
	refNumField := ref.NumField()
	for i := 0; i < refNumField; i++ {
		fStructField := ref.Type().Field(i)
		tagArgElems := getTagArgElems(fStructField.Tag.Get(tagName))

		fName := fStructField.Name
		fRef, fType, fKind, fValue := util.GetUnderlyingReflectValueInfo(ref.Field(i))

		if util.IsNilReflectValue(fRef) || fKind == reflect.Invalid || fValue == nil {
			continue
		}

		if len(tagArgElems) == 0 {
			continue
		}

		hasNestedArg := lo.ContainsBy(tagArgElems, func(t model.TagArgElem) bool {
			return t.ArgName == enum.TagArgNested
		})

		fTypeStr := fType.String()

		if fKind == reflect.String {
			if err := transformer.StringTransformer.Transform(tagArgElems, fRef, fName, fType, fValue); err != nil {
				return err
			}
			// } else if util.ExistsInArray(fKind, util.AllFloatReflectKind...) {
		} else if lo.Contains(util.AllFloatReflectKind, fKind) {
			if err := transformer.FloatTransformer.Transform(tagArgElems, fRef, fName, fType, fValue); err != nil {
				return err
			}
		} else if fTypeStr == "time.Time" {
			if err := transformer.TimeTransformer.Transform(tagArgElems, fRef, fName, fType, fValue); err != nil {
				return err
			}
		} else {
			if !hasNestedArg {
				return fmt.Errorf("unsupported field type for tag transform. field type: %s", fTypeStr)
			}
		}

		if hasNestedArg {
			if len(tagArgElems) > 1 {
				return fmt.Errorf("nested tag arg cannot be followed with any other arg")
			}

			if fKind != reflect.Struct && fKind != reflect.Slice {
				return fmt.Errorf("nested transformation only supports struct and slice type")
			}

			if fKind == reflect.Struct { // handle child struct validation
				if err := Transform(fValue); err != nil {
					return err
				}
			} else if fKind == reflect.Slice && fRef.Len() > 0 { // handle child slice validation
				childRefType := fType.Elem()
				if childRefType.Kind() == reflect.Pointer {
					childRefType = childRefType.Elem()
				}

				// we only support struct for now
				if childRefType.Kind() == reflect.Struct {
					for iChildSlice := 0; iChildSlice < fRef.Len(); iChildSlice++ {
						childRef := fRef.Index(iChildSlice)
						if childRef.Kind() == reflect.Ptr {
							childRef = reflect.Indirect(childRef)
						}

						if !util.IsNilReflectValue(childRef) {
							if err := transform(childRef); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}

// ref: https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
func Transform(obj interface{}) error {
	ref, err := util.GetStructReflectValue(obj)
	if err != nil {
		return err
	}

	return transform(ref)
}
