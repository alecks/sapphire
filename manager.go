package main

import (
	"fmt"
	"strings"
)

func manage(data string) string {
	res := ""

	for index, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)

		for _, mod := range modules {
			if strings.HasPrefix(line, mod.Identifier) {
				res, err := mod.Function(line, index)
				throw(err)

				fmt.Println(res)
			}
		}
	}

	return res
}
