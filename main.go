package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var commonMode bool
var file1, file2 string

func main() {
	commonModeFlag := flag.Bool("common", false, "Set this flag to list only common items")
	file1Flag := flag.String("f1", "", "file 1")
	file2Flag := flag.String("f2", "", "file 2")
	flag.Parse()
	commonMode = *commonModeFlag
	file1 = *file1Flag
	file2 = *file2Flag

	// Check if the required number of arguments is provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [-common] -f1 <file1> -f2 <file2>")
		return
	}

	// Read file contents
	contents1, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Println("Error reading file 1:", err)
		return
	}

	contents2, err := ioutil.ReadFile(file2)
	if err != nil {
		fmt.Println("Error reading file 2:", err)
		return
	}

	// Convert file contents to lists
	packages1 := strings.Split(string(contents1), "\n")
	packages2 := strings.Split(string(contents2), "\n")

	// Create a map to store the packages from file 2
	packagesMap := make(map[string]bool)
	for _, pkg := range packages2 {
		packagesMap[pkg] = true
	}

	outputPackages := []string{}

	if commonMode {
		for _, pkg := range packages1 {
			if _, ok := packagesMap[pkg]; ok {
				outputPackages = append(outputPackages, pkg)
			}
		}
		fmt.Println("Common packages in both files:")
	} else {
		for _, pkg := range packages1 {
			if _, ok := packagesMap[pkg]; !ok {
				outputPackages = append(outputPackages, pkg)
			}
		}
		fmt.Println("Missing packages from file 2:")
	}

	// Print output packages
	for _, pkg := range outputPackages {
		fmt.Println(pkg)
	}
}
