package env

import "os"

// GetString returns the value stored in the environment variable
// If the environment variable isn't set the default value will be returned.
func GetString(key, dv string) string {
	s, ok := os.LookupEnv(key)
	if !ok {
		s = dv
	}

	return s
}
