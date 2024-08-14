package helper

import (
	"errors"
	"time"

	"github.com/jinzhu/copier"
)

func Copy(dest interface{}, source interface{}) error {
	ptrStrType := ""
	err := copier.CopyWithOption(dest, source, copier.Option{
		DeepCopy: true,
		Converters: []copier.TypeConverter{
			// time.Time to *string and vice versa
			{
				SrcType: time.Time{},
				DstType: &ptrStrType,
				Fn: func(src interface{}) (dst interface{}, err error) {
					t, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					if t.IsZero() {
						return nil, nil
					}

					s := t.Format(time.RFC3339)
					return &s, nil
				},
			},
			{
				SrcType: &ptrStrType,
				DstType: time.Time{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(*string)

					zeroTime := time.Time{}

					if !ok {
						return zeroTime, errors.New("src type not matching")
					}

					if s == nil {
						return time.Time{}, nil
					}

					time, err := time.Parse(time.RFC3339, *s)
					if err != nil {
						return zeroTime, errors.New("src string is not in valid date format")
					}

					return time, nil
				},
			},
			// time.Time to string and vice versa
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (dst interface{}, err error) {
					t, ok := src.(time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					return t.Format(time.RFC3339), nil
				},
			},
			{
				SrcType: copier.String,
				DstType: time.Time{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					s, ok := src.(string)

					zeroTime := time.Time{}

					if !ok {
						return zeroTime, errors.New("src type not matching")
					}

					time, err := time.Parse(time.RFC3339, s)
					if err != nil {
						return zeroTime, errors.New("src string is not in valid date format")
					}

					return time, nil
				},
			}, // []string to []string deep copy
			{
				SrcType: []string{},
				DstType: []string{},
				Fn: func(src interface{}) (dst interface{}, err error) {
					strSlice, ok := src.([]string)
					if !ok {
						return nil, errors.New("src type not matching")
					}

					// Create a deep copy of the slice
					copiedSlice := make([]string, len(strSlice))
					copy(copiedSlice, strSlice)
					return copiedSlice, nil
				},
			},
		},
	})

	return err
}
