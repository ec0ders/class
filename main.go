package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	var inherits *string

	if len(os.Args) == 3 {
		inherits = &os.Args[2]
	}

	createClass(os.Args[1], inherits)
}

func createClass(name string, inherits *string) {

	name = camelCase(name)

	filename := name + ".lua"

	if fileExists(filename) {
		log.Fatal("file exists")
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("could not create file")
	}
	defer f.Close()

	class := strings.ReplaceAll(classData, ":class:", name)

	if inherits != nil {
		class = strings.Replace(class, ":inherits:", ": "+camelCase(*inherits), 1)
	} else {
		class = strings.Replace(class, ":inherits:", "", 1)
	}

	_, err = f.WriteString(class)
	if err != nil {
		log.Fatal("could not write class")
	}
}

func camelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	// The user has the responsibility to put the uppercase on multiple words names
	return strings.ToUpper(string(s[0])) + string(s[1:])
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
