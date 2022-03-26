package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

func Warnf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func Infof(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func PrintObject[T any](object ...T) error {
	var result []string
	for i := 0; i < len(object); i++ {
		content, err := yaml.Marshal(object[i])
		if err != nil {
			return fmt.Errorf("unable to format yaml %w", err)
		}
		result = append(result, string(content))
	}
	fmt.Println(strings.Join(result, "---\n"))

	return nil
}
