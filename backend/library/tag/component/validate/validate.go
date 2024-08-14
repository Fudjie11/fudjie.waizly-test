package validate

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/samber/lo"

	enum "fudjie.waizly/backend-test/library/tag/component/validate/enum"
	model "fudjie.waizly/backend-test/library/tag/component/validate/model"
	validator "fudjie.waizly/backend-test/library/tag/component/validate/validator"

	util "fudjie.waizly/backend-test/library/util"
)

var tagName = "validate"

// func ValidateDef(tagDefs []tagmodel.TagDefinition) ([]model.ValidatorResult, error) {
// 	results := []model.ValidatorResult{}
// 	var err error

// 	for _, tagDef := range tagDefs {
// 		if tagDef.TagName != tagName || len(tagDef.TagValue) == 0 {
// 			continue
// 		}

// 		fType := tagDef.FieldType.String()

// 		tagArgs := strings.Split(tagDef.TagValue, ",")
// 		if fType == "string" {
// 			results, err = validator.StringValidator.Validate(results, tagDef.FieldName, fType, tagDef.FieldValue, tagArgs)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else if fType == "time.Time" {
// 			results, err = validator.TimeValidator.Validate(results, tagDef.FieldName, fType, tagDef.FieldValue, tagArgs)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else if fType == "uuid.UUID" {
// 			results, err = validator.UUIDValidator.Validate(results, tagDef.FieldName, fType, tagDef.FieldValue, tagArgs)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else if util.ExistsInArray[string](fType, []string{"int", "int8", "int16", "int32", "int64"}...) {
// 			results, err = validator.IntegerValidator.Validate(results, tagDef.FieldName, fType, tagDef.FieldValue, tagArgs)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else if fType == "float32" || fType == "float64" {
// 			results, err = validator.FloatValidator.Validate(results, tagDef.FieldName, fType, tagDef.FieldValue, tagArgs)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else if tagDef.FieldValue != nil && tagDef.IsStruct {
// 			// handle child struct validation
// 			childTagDefs, err := tagutil.GetTagDefinitionByName([]string{tagName}, tagDef.FieldValue)
// 			if err != nil {
// 				return nil, err
// 			}

// 			childResults, err := ValidateDef(childTagDefs)
// 			if err != nil {
// 				return nil, err
// 			}

// 			results = append(results, childResults...)
// 		} else if tagDef.FieldValue != nil && tagDef.IsSlice {
// 			sliceType := tagDef.FieldRef.Type().Elem()
// 			if tagutil.IsRefBasicType(sliceType) {
// 				continue
// 			}

// 			sliceType.Elem().Kind()

// 			// handle child struct validation
// 			childTagDefs, err := tagutil.GetTagDefinitionByName([]string{tagName}, tagDef.FieldValue)
// 			if err != nil {
// 				return nil, err
// 			}

// 			childResults, err := ValidateDef(childTagDefs)
// 			if err != nil {
// 				return nil, err
// 			}

// 			results = append(results, childResults...)
// 		} else if tagDef.FieldValue != nil {
// 			return nil, fmt.Errorf("unsupported field type for tag validate. field type: %s", fType)
// 		}
// 	}

// 	return results, nil
// }

// func Validate(obj interface{}) ([]model.ValidatorResult, error) {
// 	tagDefs, err := tagutil.GetTagDefinitionByName([]string{tagName}, obj)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ValidateDef(tagDefs)
// }

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

func Validate(obj interface{}) (results []model.ValidatorResult, err error) {
	ref, err := util.GetStructReflectValue(obj)
	if err != nil {
		return results, err
	}

	refNumField := ref.NumField()
	for i := 0; i < refNumField; i++ {
		fStructField := ref.Type().Field(i)
		tagArgElems := getTagArgElems(fStructField.Tag.Get(tagName))

		fName := fStructField.Name
		fRef, fType, fKind, fValue := util.GetUnderlyingReflectValueInfo(ref.Field(i))

		if fKind == reflect.Invalid && len(tagArgElems) > 0 {
			return nil, errors.New("cannot validate. cannot get info of the field kind and type. it might be nil interface")
		}

		if len(tagArgElems) == 0 {
			continue
		}

		hasNestedArg := lo.ContainsBy(tagArgElems, func(t model.TagArgElem) bool {
			return t.ArgName == enum.TagArgNested
		})

		fTypeStr := fType.String()

		if fKind == reflect.String {
			results, err = validator.StringValidator.Validate(results, tagArgElems, fRef, fName, fType, fValue)
			if err != nil {
				return nil, err
			}
		} else if lo.Contains(util.AllIntReflectKind, fKind) {
			results, err = validator.IntegerValidator.Validate(results, tagArgElems, fRef, fName, fType, fValue)
			if err != nil {
				return nil, err
			}
		} else if lo.Contains(util.AllFloatReflectKind, fKind) {
			results, err = validator.FloatValidator.Validate(results, tagArgElems, fRef, fName, fType, fValue)
			if err != nil {
				return nil, err
			}
		} else if fTypeStr == "time.Time" {
			results, err = validator.TimeValidator.Validate(results, tagArgElems, fRef, fName, fType, fValue)
			if err != nil {
				return nil, err
			}
		} else if fTypeStr == "uuid.UUID" {
			results, err = validator.UUIDValidator.Validate(results, tagArgElems, fRef, fName, fType, fValue)
			if err != nil {
				return nil, err
			}
		} else if fKind == reflect.Struct {
			results, err = validator.StructValidator.Validate(results, tagArgElems, fRef, fName, fType, fValue)
			if err != nil {
				return nil, err
			}
		} else if fKind == reflect.Slice {
			results, err = validator.SliceValidator.Validate(results, tagArgElems, fRef, fName, fType, fValue)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unsupported field type for tag validate. field type: %s", fTypeStr)
		}

		if hasNestedArg {
			if len(tagArgElems) > 1 {
				return nil, fmt.Errorf("nested tag arg cannot be followed with any other arg")
			}

			if fKind != reflect.Struct && fKind != reflect.Slice {
				return nil, fmt.Errorf("nested validation only supports struct and slice type")
			}

			if fKind == reflect.Struct && fValue != nil { // handle child struct validation
				childResults, err := Validate(fValue)
				if err != nil {
					return nil, err
				}

				results = append(results, childResults...)
			} else if fKind == reflect.Slice && fRef.Len() > 0 && fValue != nil { // handle child slice validation
				childRefType := fType.Elem()
				if childRefType.Kind() == reflect.Pointer {
					childRefType = childRefType.Elem()
				}

				// we only support struct for now
				if childRefType.Kind() == reflect.Struct {
					for iChildSlice := 0; iChildSlice < fRef.Len(); iChildSlice++ {
						childRef := fRef.Index(iChildSlice)
						if !util.IsNilReflectValue(childRef) {
							childResults, err := Validate(childRef.Interface())
							if err != nil {
								return nil, err
							}

							results = append(results, childResults...)
						}
					}
				}
			}
		}
	}

	return results, nil
}
