package pencake

import "os"

func Environment() []EnvironmentStruct {
	var result []EnvironmentStruct
	keys := os.Environ()
	for _, key := range keys {
		result = append(result, EnvironmentStruct{
			Key:   key,
			Value: os.Getenv(key),
		})
	}
	return result
}

type EnvironmentStruct struct {
	Key   string
	Value string
}
