package entities

// Since a single validation can have multiple errors, creating a custom error type is Go's standard.
// The error type allow us to group every problem with the cause which is good for the user.
// In this case we're using a map so that we can know what field is responsible for that particular error which is good to map back to HTML in the view layer.
// Because of this our functions can return our single custom error type and still adhere to Go's single error return standard and still follow the if err != nil pattern.
type ValidationError struct {
	Fields map[string][]string
}

func (ve ValidationError) Error() string {
	return "validation error"
}

func (ve ValidationError) add(key, value string) {
	if ve.Fields == nil {
		ve.Fields = map[string][]string{}
	}

	ve.Fields[key] = append(ve.Fields[key], value)
}
