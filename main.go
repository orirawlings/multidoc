package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	var data map[string]interface{}
	if err := yaml.NewDecoder(os.Stdin).Decode(&data); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	if items, ok := data["items"].([]interface{}); ok {
		for _, item := range items {
			fmt.Fprintf(os.Stdout, "\n---\n")
			if err := yaml.NewEncoder(os.Stdout).Encode(item); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		}
	}
}
