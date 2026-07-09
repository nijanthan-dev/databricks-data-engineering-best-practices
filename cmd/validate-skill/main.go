package main

import (
	"fmt"
	"os"

	"databricks-data-engineering-best-practices/internal/skillvalidator"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	result, err := skillvalidator.Validate(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(2)
	}
	if !result.Valid {
		for _, issue := range result.Issues {
			fmt.Printf("FAIL: %s\n", issue)
		}
		os.Exit(1)
	}

	fmt.Println("PASS: skill valid")
}
