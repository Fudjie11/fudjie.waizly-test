package helper

import (
	"context"

	"fudjie.waizly/backend-test/internal/model"
)

func getMapValue[T any](m map[string]interface{}, key string) T {
	var result T
	mVal := m[key]
	if mVal != nil {
		return mVal.(T)
	}

	return result
}

func GetContextAppData(ctx context.Context) model.ContextAppData {
	ctxAppData := model.ContextAppData{
		LanguageId: "en",
	}

	c := ctx.Value("APP_DATA")
	if c == nil {
		return ctxAppData
	}

	m := c.(map[string]interface{})

	ctxAppData.LanguageId = getMapValue[string](m, "LanguageId")
	ctxAppData.TimeZoneId = getMapValue[string](m, "TimeZoneId")
	ctxAppData.TimeZoneOffset = getMapValue[int](m, "TimeZoneOffset")

	return ctxAppData
}
