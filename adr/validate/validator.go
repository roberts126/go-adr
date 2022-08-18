package validate

var allFunctions = map[string]Func{
	"match": Match,
}

// Func represents a constructor for a struct that implements the Validator interface.
type Func func(string, ...string) (Validator, error)

// Validator represents an object that can validate a value against a criteria.
type Validator interface {
	Validate() error
	SetMessage(string)
}

// GetFunction looks for the specified key and returns the Func if found.
func GetFunction(key string) Func {
	v, ok := allFunctions[key]
	if !ok {
		return nil
	}

	return v
}

// All validates all of the listed Validator objects.
func All(validators ...Validator) error {
	for _, v := range validators {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
