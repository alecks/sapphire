package main

import (
	"strconv"
	"strings"
)

type moduleFunction func(string, int) (int, error)
type module struct {
	Identifier string
	Function   moduleFunction
}

var modules = []module{{"if", ifStatement}}

func runOp(in string, in2 string, operator string, index int) (int, error) {
	var err error
	var res int

	for _, optr := range operators {
		if operator == optr.Identifier {
			err = nil

			param, err1 := strconv.Atoi(in)
			param2, err2 := strconv.Atoi(in2)

			if err1 != nil || err2 != nil {
				err = &typeError{"string", "int", index + 1}
			} else {
				res = optr.Function(param, param2)
			}

			break
		} else {
			err = &undefinedError{operator, index + 1}
		}
	}

	return res, err
}

func ifStatement(line string, index int) (int, error) {
	var err error
	var res int

	inputs := strings.Split(line, " ")
	if len(inputs) >= 8 {
		runRes, runErr := runOp(inputs[1], inputs[3], inputs[2], index)
		err = runErr

		if err == nil {
			expected, err3 := strconv.Atoi(inputs[5])

			if err3 != nil {
				err = &typeError{"string", "int", index + 1}
			} else {
				if expected == runRes {
					runRes1, runErr1 := runOp(inputs[7], inputs[9], inputs[8], index)
					err = runErr1

					if err == nil {
						res = runRes1
					}
				} else {
					runRes2, runErr2 := runOp(inputs[11], inputs[13], inputs[12], index)
					err = runErr2

					if err == nil {
						res = runRes2
					}
				}
			}
		}
	} else {
		err = &syntaxError{"Expected complete if statement, got " + line, index + 1}
	}

	return res, err
}
