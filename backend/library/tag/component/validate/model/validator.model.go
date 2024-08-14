package model

type ValidatorResult struct {
	FieldName  string
	FieldValue interface{}
	ArgName    string
	ArgValue   interface{}
}

func NewValidatorResult(fieldName string, fieldValue interface{}, argName string, argValue interface{}) ValidatorResult {
	return ValidatorResult{
		FieldName:  fieldName,
		FieldValue: fieldValue,
		ArgName:    argName,
		ArgValue:   argValue,
	}
}

type TagArgElem struct {
	ArgName  string
	ArgValue string
}
