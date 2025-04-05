package config

import "strings"

type UnknownEnvErr struct {
	env       string
	supported string
}

func NewUnknownEnvErr(env string) error {
	keys := make([]string, len(allEnvs))
	i := 0
	for k := range allEnvs {
		keys[i] = k
		i++
	}

	return UnknownEnvErr{
		env:       env,
		supported: strings.Join(keys, ", "),
	}
}

func (e UnknownEnvErr) Error() string {
	return "unknown env: " + e.env + ", supported: " + e.supported
}
