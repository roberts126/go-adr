package validate

var allValidators = map[string]ValidatorFunc{
	"match": Match,
}

// ValidatorFunc represents a constructor for a struct that implements the Validator interface.
type ValidatorFunc func(string, ...string) (Validator, error)

// Validator represents an object that can validate a value against a criteria.
type Validator interface {
	Validate() error
	SetMessage(string)
}

// GetValidatorFunction looks for the specified key and returns the ValidatorFunc if found.
func GetValidatorFunction(key string) ValidatorFunc {
	v, ok := allValidators[key]
	if !ok {
		return nil
	}

	return v
}

// ListValidatorFunctions returns a list of function names
func ListValidatorFunctions() []string {
	f := make([]string, len(allValidators))
	i := 0

	for k := range allValidators {
		f[i] = k
		i++
	}

	return f
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
