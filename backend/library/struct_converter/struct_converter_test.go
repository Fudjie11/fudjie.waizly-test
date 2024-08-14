package struct_converter

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestConvertToNewType_StringToUUIDAndBack(t *testing.T) {
	sourceType := reflect.TypeOf("")
	targetTypeUUID := reflect.TypeOf(uuid.UUID{})
	sourceValue := "123e4567-e89b-12d3-a456-426614174000"
	expectedUUID, _ := uuid.Parse(sourceValue)

	isHandled, targetValue, err := convertToNewType(sourceType, sourceValue, targetTypeUUID)
	assert.True(t, isHandled)
	assert.Nil(t, err)
	assert.Equal(t, expectedUUID, targetValue)

	// Convert back to string
	targetTypeString := reflect.TypeOf("")
	isHandled, targetValue, err = convertToNewType(targetTypeUUID, expectedUUID, targetTypeString)
	assert.True(t, isHandled)
	assert.Nil(t, err)
	assert.Equal(t, sourceValue, targetValue)
}

func TestCopyValueToTarget_SameFields(t *testing.T) {
	type SourceStruct struct {
		Name string
		Age  int
	}

	type TargetStruct struct {
		Name string
		Age  int
	}

	source := SourceStruct{Name: "John", Age: 30}
	target := TargetStruct{}

	sourceRef := reflect.ValueOf(source)
	targetRef := reflect.Indirect(reflect.New(reflect.TypeOf(target)))

	err := copyValueToTarget(sourceRef, targetRef)
	assert.Nil(t, err)

	result := targetRef.Interface().(TargetStruct)
	assert.Equal(t, source.Name, result.Name)
	assert.Equal(t, source.Age, result.Age)
}

func TestCopyValueToTarget_SliceConversion(t *testing.T) {
	type SourceSliceStruct struct {
		Values []string
	}

	type TargetSliceStruct struct {
		Values []string
	}

	source := SourceSliceStruct{Values: []string{"one", "two", "three"}}
	target := TargetSliceStruct{}

	sourceRef := reflect.ValueOf(source)
	targetRef := reflect.Indirect(reflect.New(reflect.TypeOf(target)))

	// Assuming a custom converter for string to int based on string length is registered
	err := copyValueToTarget(sourceRef, targetRef)
	assert.Nil(t, err)

	result := targetRef.Interface().(TargetSliceStruct)
	assert.Equal(t, 3, len(result.Values))
	assert.Equal(t, "one", result.Values[0])
	assert.Equal(t, "two", result.Values[1])
	assert.Equal(t, "three", result.Values[2])
}

func TestCopyValueToTarget_NestedStructs(t *testing.T) {
	type NestedStruct struct {
		Description string
	}

	type SourceStruct struct {
		Nested NestedStruct
	}

	type TargetStruct struct {
		Nested NestedStruct
	}

	source := SourceStruct{Nested: NestedStruct{Description: "Nested struct"}}
	target := TargetStruct{}

	sourceRef := reflect.ValueOf(source)
	targetRef := reflect.Indirect(reflect.New(reflect.TypeOf(target)))

	err := copyValueToTarget(sourceRef, targetRef)
	assert.Nil(t, err)

	result := targetRef.Interface().(TargetStruct)
	assert.Equal(t, source.Nested.Description, result.Nested.Description)
}

func TestCopyValueToTarget_IncompatibleFieldTypes(t *testing.T) {
	type SourceStruct struct {
		Value int
	}

	type TargetStruct struct {
		Value time.Time // Incompatible type
	}

	source := SourceStruct{Value: 42}
	target := TargetStruct{}

	sourceRef := reflect.ValueOf(source)
	targetRef := reflect.Indirect(reflect.New(reflect.TypeOf(target)))

	err := copyValueToTarget(sourceRef, targetRef)
	assert.NotNil(t, err, "should fail due to incompatible field types")
}

// TestConvertToTargetPointer check if ConvertStruct convert properly variable to pointer variable
func TestConvertToTargetPointer(t *testing.T) {
	type source struct {
		Field1 string
		Field2 time.Time
		Field3 uuid.UUID
		Field4 []string
	}

	type target struct {
		Field1 *string
		Field2 *time.Time
		Field3 *uuid.UUID
		Field4 []*string
	}

	RegisterNillableTypeAndValue(uuid.Nil)
	RegisterNillableTypeAndValue("")
	RegisterNillableTypeAndValue(time.Time{})

	s1 := source{
		Field1: "sadad",
		Field4: []string{"array1", "array2", ""},
	}

	t1, err := ConvertStruct[target](s1)
	fmt.Printf("\nTestConvertToTargetPointer (1): %#v\n", t1)
	assert.Nil(t, err, "should pass")

	s2 := source{}

	t2, err := ConvertStruct[target](s2)
	fmt.Printf("\nTestConvertToTargetPointer (1): %#v\n", t2)
	assert.Nil(t, err, "should pass")

	s3 := source{
		Field1: "abc",
		Field2: time.Now(),
		Field3: uuid.Nil,
	}

	t3, err := ConvertStruct[target](s3)
	fmt.Printf("\nTestConvertToTargetPointer (1): %#v\n", t3)
	assert.Nil(t, err, "should pass")
}

// TestConvertToTargetPointer check if ConvertStruct convert properly variable to pointer variable
func TestConvertFromPointerToNonPointer(t *testing.T) {
	type source struct {
		Field1 *string
		Field2 []*string
		Field3 *time.Time
		field4 *string
	}

	type target struct {
		Field1 string
		Field2 []string
		Field3 time.Time
		field4 string
	}

	RegisterNillableTypeAndValue(uuid.Nil)
	RegisterNillableTypeAndValue("")
	RegisterNillableTypeAndValue(time.Time{})

	// s1 := "test1"
	s2 := "test2"
	s3 := "test3"
	// s4 := "test4"
	time1 := time.Now()

	source1 := source{
		Field1: nil,
		Field2: []*string{&s2, &s3, nil},
		Field3: &time1,
		field4: nil,
	}

	target1, err := ConvertStruct[target](source1)
	fmt.Printf("\nTestConvertFromPointerToNonPointer: %#v\n", target1)
	assert.Nil(t, err, "should pass")
}

// TestConvertToTargetPointer check if ConvertStruct convert properly variable to pointer variable
func TestConvertForSlice(t *testing.T) {
	type source struct {
		Field1 *string
		Field2 []*string
		Field3 *time.Time
		field4 *string
	}

	type target struct {
		Field1 string
		Field2 []string
		Field3 time.Time
		field4 string
	}

	RegisterNillableTypeAndValue(uuid.Nil)
	RegisterNillableTypeAndValue("")
	RegisterNillableTypeAndValue(time.Time{})

	var sourceList []source

	// s1 := "test1"
	s2 := "test2"
	s3 := "test3"
	// s4 := "test4"
	time1 := time.Now()

	sourceList = []source{{
		Field1: nil,
		Field2: []*string{&s2, &s3, nil},
		Field3: &time1,
		field4: nil,
	}}

	targetList, err := ConvertStructSlice[*target](sourceList)
	fmt.Printf("\nTestConvertForSlice: %v\n", targetList)
	assert.Nil(t, err, "should pass")
}
