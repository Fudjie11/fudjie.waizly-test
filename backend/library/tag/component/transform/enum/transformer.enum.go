package enum

const (
	TagArgNested  string = "nested" // command to also transform nested object
	TagArgUpper   string = "upper"
	TagArgLower   string = "lower"
	TagArgTitle   string = "title"
	TagArgTrim    string = "trim"
	TagArgEmail   string = "email"
	TagArgDate    string = "date"
	TagArgNoSpace string = "nospace"
	TagArgDecimal string = "decimal"
)

func IsValidTagArg(tagArg string) bool {
	if tagArg != TagArgNested && tagArg != TagArgUpper && tagArg != TagArgLower && tagArg != TagArgTitle && tagArg != TagArgTrim && tagArg != TagArgEmail && tagArg != TagArgDate && tagArg != TagArgNoSpace && tagArg != TagArgDecimal {
		return false
	}

	return true
}
