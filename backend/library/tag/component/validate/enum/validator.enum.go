package enum

const (
	TagArgNested    string = "nested" // command to also validate nested object
	TagArgRequired  string = "required"
	TagArgNumber    string = "number"
	TagArgMin       string = "min"
	TagArgMax       string = "max"
	TagArgEmail     string = "email"
	TagArgPhone     string = "phone"
	TagArgLatitude  string = "latitude"
	TagArgLongitude string = "longitude"
)

func IsValidTagArg(tagArg string) bool {
	if tagArg != TagArgNested && tagArg != TagArgRequired && tagArg != TagArgNumber && tagArg != TagArgMin && tagArg != TagArgMax && tagArg != TagArgEmail && tagArg != TagArgPhone &&
		tagArg != TagArgLatitude && tagArg != TagArgLongitude {
		return false
	}

	return true
}
