package helper

import (
	"context"
	"fmt"

	pkgTagTransform "fudjie.waizly/backend-test/library/tag/component/transform"
	pkgTagValidate "fudjie.waizly/backend-test/library/tag/component/validate"
	pkgTagValidateModel "fudjie.waizly/backend-test/library/tag/component/validate/model"
)

type TagValidationErrorHandler = func(ctx context.Context, validationResults []pkgTagValidateModel.ValidatorResult) error

var tagValidationErrorHandler TagValidationErrorHandler

func TransformByTag(obj interface{}) error {
	if err := pkgTagTransform.Transform(obj); err != nil {
		return err
	}

	return nil
}

func SetTagValidationErrorHandler(handler TagValidationErrorHandler) {
	tagValidationErrorHandler = handler
}

func ValidateByTag(ctx context.Context, obj interface{}) error {
	valResults, err := pkgTagValidate.Validate(obj)
	if err != nil {
		return err
	}

	if len(valResults) == 0 {
		return nil
	}

	if tagValidationErrorHandler != nil {
		return tagValidationErrorHandler(ctx, valResults)
	} else {
		return fmt.Errorf("tag validation failed. field name: %s. tag arg:%s", valResults[0].FieldName, valResults[0].ArgValue)
	}
}

func TransformAndValidateByTag(ctx context.Context, obj interface{}) error {
	err := TransformByTag(obj)
	if err != nil {
		return err
	}

	err = ValidateByTag(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}
