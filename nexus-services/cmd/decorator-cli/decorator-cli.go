package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/juegodynamics/nexus/nexus-services/lib/decorator"
)

func main() {
	// Define command-line flags
	fileFlag := flag.String("file", "", "The file to process")
	dirFlag := flag.String("dir", "", "The directory to process")
	flag.Parse()

	// Register custom decorators
	decorator.RegisterDecorator("graphql", &decorator.GraphQLDecorator{})

	if *fileFlag != "" {
		// Process a single file
		err := decorator.ProcessFile(*fileFlag)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", *fileFlag, err)
			os.Exit(1)
		}
	} else if *dirFlag != "" {
		// Walk through the directory and process each Go file
		err := filepath.Walk(*dirFlag, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".go" {
				err := decorator.ProcessFile(path)
				if err != nil {
					fmt.Printf("Error processing file %s: %v\n", path, err)
					return err
				}
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error walking directory %s: %v\n", *dirFlag, err)
			os.Exit(1)
		}
	} else {
		// Print usage information
		flag.Usage()
	}
}
