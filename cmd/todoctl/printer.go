package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func Warnf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func Infof(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

type yamlable interface {
	ToYAML() (string, error)
}

func PrintObject[T yamlable](cmd *cobra.Command, object ...T) error {
	var result []string
	for i := 0; i < len(object); i++ {
		content, err := object[i].ToYAML()
		if err != nil {
			return err
		}
		result = append(result, string(content))
	}

	fmt.Fprintf(cmd.OutOrStdout(), strings.Join(result, "---\n"))

	return nil
}
