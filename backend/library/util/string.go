package util

import (
	"crypto/rand"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func QueryParamsValueToStringArray(s, separator string) []string {
	s = strings.TrimSpace(s)
	return strings.Split(s, separator)
}

func QueryParamsValueToInt32Array(s, separator string, skipError bool) ([]int32, error) {
	var arrInt32 []int32
	s = strings.TrimSpace(s)
	arrStr := strings.Split(s, separator)
	for _, str := range arrStr {
		i, err := strconv.Atoi(str)
		if err != nil && skipError {
			continue
		}

		if err != nil && !skipError {
			return nil, err
		}

		arrInt32 = append(arrInt32, int32(i))
	}

	return arrInt32, nil
}

var alphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
var numericRegex = regexp.MustCompile(`[^0-9]+`)

func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func RemoveNonAlphaNumericChar(s string) string {
	result := alphanumericRegex.ReplaceAllString(s, "")
	return result
}

func RemoveNonNumericChar(s string) string {
	result := numericRegex.ReplaceAllString(s, "")
	return result
}

var randomNumberTable = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// ref: https://stackoverflow.com/questions/39481826/generate-6-digit-verification-code-with-golang
func GenerateSecureRandomNumberString(numberOfDigit int32) (string, error) {
	b := make([]byte, numberOfDigit)

	n, err := io.ReadAtLeast(rand.Reader, b, int(numberOfDigit))
	if err != nil {
		return "", err
	}

	if n != int(numberOfDigit) {
		return "", errors.New("io failed to read byte")
	}

	for i := 0; i < len(b); i++ {
		b[i] = randomNumberTable[int(b[i])%len(randomNumberTable)]
	}

	return string(b), nil
}

// this code for validate string with number value
func IsValidNumber(str string) bool {
	// Regular expression pattern to match a valid number
	pattern := `^-?\d+$`

	// Compile the regular expression pattern
	reg := regexp.MustCompile(pattern)

	// Match the string against the regular expression
	return reg.MatchString(str)
}

func ConvertUUIDsToStrings(uuidSlice []uuid.UUID) []string {
	return lo.Map(uuidSlice, func(u uuid.UUID, _ int) string {
		return u.String()
	})
}
