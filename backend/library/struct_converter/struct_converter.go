// ref: https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice
// ref: https://stackoverflow.com/questions/38818915/creating-slice-from-reflected-type
// ref: https://stackoverflow.com/questions/19389629/golang-get-the-type-of-slice
// ref: https://stackoverflow.com/questions/23555241/how-to-get-zero-value-of-a-field-type
// ref: https://stackoverflow.com/questions/64211864/setting-nil-pointers-address-with-reflections

package struct_converter

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"fudjie.waizly/backend-test/library/util"
	"github.com/google/uuid"
)

type nillableTypeData struct {
	DataType reflect.Type
	NilValue interface{}
}

var nillableTypeDataList []nillableTypeData

type TypeCustomConverterFunc func(source interface{}) interface{}

type typeCustomConverterData struct {
	SourceType    reflect.Type
	TargetType    reflect.Type
	ConverterFunc TypeCustomConverterFunc
}

var typeCustomConverterDataList []typeCustomConverterData = []typeCustomConverterData{}

type AfterConvertEventFunc func(source interface{}, target TargetModifierWrapper) error

type afterConvertEventData struct {
	SourceType            reflect.Type
	TargetType            reflect.Type
	AfterConvertEventFunc AfterConvertEventFunc
}

var afterConvertEventList []afterConvertEventData = []afterConvertEventData{}

func getStructRef(obj interface{}, isTarget bool) (reflect.Value, error) {
	ref := reflect.ValueOf(obj)

	if isTarget {
		// if ref.Kind() != reflect.Ptr && ref.Kind() != reflect.Interface {
		// 	return ref, errors.New("object must be of type pointer or interface of struct")
		// }

		if ref.Kind() != reflect.Ptr {
			return ref, errors.New("object must be of type pointer of struct")
		}
	}

	// if its an interface, resolve its value
	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	// should double check we now have a struct (could still be anything)
	if ref.Kind() != reflect.Struct {
		return ref, errors.New("object must be of struct type")
	}

	return ref, nil
}

func getSliceRef(obj interface{}) (reflect.Value, error) {
	ref := reflect.ValueOf(obj)

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// should double check we now have a slice (could still be anything)
	if ref.Kind() != reflect.Slice {
		return ref, errors.New("object is not slice")
	}

	return ref, nil
}

func convertToNewType(sourceType reflect.Type, sourceValue interface{}, targetType reflect.Type) (isHandled bool, targetValue interface{}, err error) {
	strSourceType := sourceType.String()
	strTargetType := targetType.String()

	switch {
	case strSourceType == "string" && strTargetType == "uuid.UUID":
		s := sourceValue.(string)
		if len(s) > 0 {
			targetValue, err := uuid.Parse(s)
			if err != nil {
				return false, nil, err
			}
			return true, targetValue, nil
		}
		return true, uuid.Nil, nil

	case strSourceType == "uuid.UUID" && strTargetType == "string":
		return true, sourceValue.(uuid.UUID).String(), nil

	case strSourceType == "time.Time" && strTargetType == "string":
		timeString := sourceValue.(time.Time).Format(time.RFC3339)
		return true, timeString, nil

	case strSourceType == "string" && strTargetType == "time.Time":
		s := sourceValue.(string)
		timeValue, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return false, nil, errors.New("invalid string format, cannot convert to time.Time")
		}
		return true, timeValue, nil

	case sourceType.ConvertibleTo(targetType):
		return true, reflect.ValueOf(sourceValue).Convert(targetType).Interface(), nil

	default:
		return false, nil, nil
	}
}

func RegisterNillableTypeAndValue[T any](nilValue T) error {
	ref := reflect.ValueOf(nilValue)
	nillableType := ref.Type()

	nillableTypeDataList = append(nillableTypeDataList, nillableTypeData{
		DataType: nillableType,
		NilValue: nilValue,
	})

	return nil
}

func isRegisteredNillableTypeNilValue(sourceFieldType reflect.Type, sourceFieldValue interface{}) bool {
	for _, n := range nillableTypeDataList {
		if n.DataType == sourceFieldType {
			return sourceFieldValue == n.NilValue
		}
	}

	return false
}

func getTypeCustomConverter(sourceType reflect.Type, targetType reflect.Type) TypeCustomConverterFunc {
	for _, c := range typeCustomConverterDataList {
		if c.SourceType == sourceType && c.TargetType == targetType {
			return c.ConverterFunc
		}
	}

	return nil
}

// T is source type and U is target type
func RegisterTypeCustomConverter[T any, U any](converterFunc TypeCustomConverterFunc) error {
	var source T
	var target U

	sourceRef, err := getStructRef(source, false)
	if err != nil {
		return err
	}

	targetRef, err := getStructRef(target, false)
	if err != nil {
		return err
	}

	sourceType := sourceRef.Type()
	targetType := targetRef.Type()

	if sourceType == targetType {
		return fmt.Errorf("cannot register converter for same type. type: %s", sourceType)
	}

	cc := getTypeCustomConverter(sourceType, targetType)
	if cc != nil {
		return fmt.Errorf("custom converter is already registered. source type: %s. target type: %s", sourceType, targetType)
	}

	ccData := typeCustomConverterData{
		SourceType:    sourceType,
		TargetType:    targetType,
		ConverterFunc: converterFunc,
	}

	typeCustomConverterDataList = append(typeCustomConverterDataList, ccData)

	return nil
}

func getAfterConvertEventFunc(sourceType reflect.Type, targetType reflect.Type) AfterConvertEventFunc {
	for _, c := range afterConvertEventList {
		if c.SourceType == sourceType && c.TargetType == targetType {
			return c.AfterConvertEventFunc
		}
	}

	return nil
}

// T is source type and U is target type
func RegisterAfterConvertEvent[T any, U any](afterConvertEventFunc AfterConvertEventFunc) error {
	var source T
	var target U

	sourceRef, err := getStructRef(source, false)
	if err != nil {
		return err
	}

	targetRef, err := getStructRef(target, false)
	if err != nil {
		return err
	}

	sourceType := sourceRef.Type()
	targetType := targetRef.Type()

	if sourceType == targetType {
		return fmt.Errorf("cannot register event for same type. type: %s", sourceType)
	}

	cc := getAfterConvertEventFunc(sourceType, targetType)
	if cc != nil {
		return fmt.Errorf("event is already registered. source type: %s. target type: %s", sourceType, targetType)
	}

	eventData := afterConvertEventData{
		SourceType:            sourceType,
		TargetType:            targetType,
		AfterConvertEventFunc: afterConvertEventFunc,
	}

	afterConvertEventList = append(afterConvertEventList, eventData)

	return nil
}

func getTargetValue(sourceFieldRef reflect.Value, sourceFieldType reflect.Type, sourceFieldValue interface{}, targetFieldRef reflect.Value, targetFieldType reflect.Type) (interface{}, error) {
	// if same type, than directly set the value
	if targetFieldType == sourceFieldType {
		return sourceFieldValue, nil
	}

	// first check if use custom converter. if yes, that let custom converter handle the conversion
	customConverter := getTypeCustomConverter(sourceFieldType, targetFieldType)
	if customConverter != nil {
		tValue := customConverter(sourceFieldValue)
		return tValue, nil
	}

	// if different type, then try to handle conversion
	handled, tValue, err := convertToNewType(sourceFieldType, sourceFieldValue, targetFieldType)
	if err != nil {
		return nil, err
	}

	if handled {
		return tValue, nil
	}

	// if target is pointer, then check if registered nil value. if nil, then return nil
	if targetFieldRef.Kind() == reflect.Ptr && sourceFieldRef.Kind() != reflect.Ptr && sourceFieldRef.Kind() != reflect.Interface {
		if isRegisteredNillableTypeNilValue(sourceFieldType, sourceFieldValue) {
			return nil, nil
		}

		newTargetRef := reflect.New(targetFieldType.Elem())
		newTargetValue, err := getTargetValue(sourceFieldRef, sourceFieldType, sourceFieldValue, newTargetRef.Elem(), targetFieldType.Elem())
		if err != nil {
			return nil, err
		}

		newTargetRef.Elem().Set(reflect.ValueOf(newTargetValue))
		return newTargetRef.Interface(), nil
	}

	if sourceFieldRef.Kind() == reflect.Ptr && targetFieldRef.Kind() != reflect.Ptr && targetFieldRef.Kind() != reflect.Interface {
		if sourceFieldRef.IsNil() {
			newTargetRef := reflect.Indirect(reflect.New(targetFieldType))
			return newTargetRef.Interface(), nil
		} else {
			newSourceFieldRef := reflect.Indirect(sourceFieldRef)

			newTargetValue, err := getTargetValue(newSourceFieldRef, newSourceFieldRef.Type(), newSourceFieldRef.Interface(), targetFieldRef, targetFieldType)
			if err != nil {
				return nil, err
			}

			return newTargetValue, nil
		}
	}

	// if field is struct, then repeat this by calling convertStruct again
	if sourceFieldRef.Kind() == reflect.Struct && targetFieldRef.Kind() == reflect.Struct {
		// we need to create new object with default value for target
		newTargetFieldRef := reflect.Indirect(reflect.New(targetFieldType))

		err = copyValueToTarget(sourceFieldRef, newTargetFieldRef)
		if err != nil {
			return nil, err
		}

		return newTargetFieldRef.Interface(), nil
	}

	return nil, fmt.Errorf("failed to copy field value with different type. source type: %s. target type: %s", sourceFieldType.String(), targetFieldType.String())
}

func copyValueToTarget(sourceRef reflect.Value, targetRef reflect.Value) error {
	if util.IsNilReflectValue(sourceRef) {
		return errors.New("source must not be nil")
	}

	if util.IsNilReflectValue(targetRef) {
		return errors.New("target must not be nil")
	}

	zeroReflectValue := reflect.Value{}

	fieldCount := sourceRef.NumField()
	for i := 0; i < fieldCount; i++ {
		// skip private fields
		if !sourceRef.Type().Field(i).IsExported() {
			continue
		}

		sourceStructField := sourceRef.Type().Field(i)
		sourceFieldRef := sourceRef.Field(i)
		sourceFieldType := sourceFieldRef.Type()
		sourceFieldName := sourceStructField.Name
		sourceFieldValue := sourceFieldRef.Interface()

		targetFieldRef := targetRef.FieldByName(sourceFieldName)

		// target not found
		if targetFieldRef == zeroReflectValue {
			continue
		}

		// skip private fields
		targetStructField, targetStructFieldFound := targetRef.Type().FieldByName(sourceFieldName)
		if !targetStructFieldFound || !targetStructField.IsExported() {
			continue
		}

		targetFieldType := targetFieldRef.Type()

		// if slice then
		if sourceFieldType != targetFieldType && sourceFieldRef.Kind() == reflect.Slice && targetFieldRef.Kind() == reflect.Slice {
			sourceFieldSliceElemType := sourceFieldRef.Type().Elem()
			targetFieldSliceElemType := targetFieldRef.Type().Elem()
			targetFieldSliceRef := reflect.MakeSlice(reflect.SliceOf(targetFieldSliceElemType), 0, 10)

			for islice := 0; islice < sourceFieldRef.Len(); islice++ {
				sourceFieldSliceElemRef := sourceFieldRef.Index(islice)
				// create new target instance
				targetFieldSliceElemRef := reflect.Indirect(reflect.New(targetFieldSliceElemType))

				tValue, err := getTargetValue(sourceFieldSliceElemRef, sourceFieldSliceElemType, sourceFieldSliceElemRef.Interface(), targetFieldSliceElemRef, targetFieldSliceElemType)
				if err != nil {
					return err
				}

				if tValue != nil {
					targetFieldSliceRef = reflect.Append(targetFieldSliceRef, reflect.ValueOf(tValue))
				} else {
					targetFieldSliceRef = reflect.Append(targetFieldSliceRef, reflect.Indirect(reflect.New(targetFieldSliceElemType)))
				}
			}

			targetFieldRef.Set(reflect.ValueOf(targetFieldSliceRef.Interface()))
			continue
		}

		targetFieldValue, err := getTargetValue(sourceFieldRef, sourceFieldType, sourceFieldValue, targetFieldRef, targetFieldType)
		if err != nil {
			return err
		}

		if targetFieldValue == nil {
			targetFieldRef.Set(reflect.Zero(targetFieldRef.Type()))
		} else {
			targetFieldRef.Set(reflect.ValueOf(targetFieldValue))
		}
	}

	// execute after convert event if exists
	afterConvertEventFunc := getAfterConvertEventFunc(sourceRef.Type(), targetRef.Type())
	if afterConvertEventFunc != nil {
		wrapper := TargetModifierWrapper{
			targetRef: targetRef,
		}

		err := afterConvertEventFunc(sourceRef.Interface(), wrapper)
		if err != nil {
			return err
		}
	}

	return nil
}

func ConvertStruct[T any](source any) (target T, err error) {
	sourceRef, err := getStructRef(source, false)
	if err != nil {
		return target, err
	}

	sourceType := sourceRef.Type()

	targetType := reflect.ValueOf(target).Type()

	customConverter := getTypeCustomConverter(sourceType, targetType)
	if customConverter != nil {
		return customConverter(source).(T), nil
	}

	targetKind := targetType.Kind()

	var targetRef reflect.Value
	var eTargetValue interface{}

	if targetKind == reflect.Ptr { // special handler if target is a pointer
		eTargetValue = reflect.New(targetType.Elem()).Interface()
		targetRef, err = getStructRef(eTargetValue, true)
		if err != nil {
			return target, err
		}
	} else {
		targetRef, err = getStructRef(&target, true)
		if err != nil {
			return target, err
		}
	}

	err = copyValueToTarget(sourceRef, targetRef)
	if err != nil {
		return target, err
	}

	if targetKind == reflect.Ptr {
		return eTargetValue.(T), nil
	} else {
		return target, nil
	}
}

func MustConvertStruct[T any](source any) T {
	result, err := ConvertStruct[T](source)
	if err != nil {
		panic(err)
	}

	return result
}

func ConvertStructSlice[T any](sources interface{}) (targets []T, err error) {
	sourcesRef, err := getSliceRef(sources)
	if err != nil {
		return nil, err
	}

	for i := 0; i < sourcesRef.Len(); i++ {
		source := sourcesRef.Index(i).Interface()
		target, err := ConvertStruct[T](source)
		if err != nil {
			return nil, err
		}

		targets = append(targets, target)
	}

	return targets, nil
}

func MustConvertStructSlice[T any](sources interface{}) []T {
	targets, err := ConvertStructSlice[T](sources)
	if err != nil {
		panic(err)
	}

	return targets
}

// target modifier wrapper

type TargetModifierWrapper struct {
	targetRef reflect.Value
}

func (m *TargetModifierWrapper) GetValue(fieldName string) (fieldValue interface{}, err error) {
	targetFieldRef := m.targetRef.FieldByName(fieldName)

	zeroReflectValue := reflect.Value{}

	// target not found
	if targetFieldRef == zeroReflectValue {
		return nil, fmt.Errorf("field '%s' not found", fieldName)
	}

	return targetFieldRef.Interface(), nil
}

func (m *TargetModifierWrapper) MustGetValue(fieldName string) interface{} {
	fieldValue, err := m.GetValue(fieldName)
	if err != nil {
		panic(err)
	}

	return fieldValue
}

func (m *TargetModifierWrapper) SetValue(fieldName string, fieldValue interface{}) error {
	targetFieldRef := m.targetRef.FieldByName(fieldName)

	zeroReflectValue := reflect.Value{}

	// target not found
	if targetFieldRef == zeroReflectValue {
		return fmt.Errorf("field '%s' not found", fieldName)
	}

	targetFieldRef.Set(reflect.ValueOf(fieldValue))

	return nil
}

func (m *TargetModifierWrapper) MustSetValue(fieldName string, fieldValue interface{}) {
	err := m.SetValue(fieldName, fieldValue)
	if err != nil {
		panic(err)
	}
}
