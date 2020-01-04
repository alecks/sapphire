package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Dhaava\n\nExecute from file: dhaava <filename>\nExecute from command-line: dhaava -e <code>")
	} else if os.Args[1] == "-e" {
		os.Args = arrayRemove(os.Args, 0)
		os.Args = arrayRemove(os.Args, 0)

		fmt.Println(manage(strings.Join(os.Args[:], " ")))
	} else {
		data, err := ioutil.ReadFile(os.Args[1])
		throw(err)

		fmt.Println(manage(string(data)))
	}
}

func arrayRemove(a []string, i int) []string {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]

	return a
}
