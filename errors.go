package main

import (
	"strconv"
)

type syntaxError struct {
	Message string
	Line    int
}

func (e *syntaxError) Error() string {
	return "syntaxError at line " + strconv.Itoa(e.Line) + ": " + e.Message
}

type undefinedError struct {
	Referenced string
	Line       int
}

func (e *undefinedError) Error() string {
	return "undefinedError at line " + strconv.Itoa(e.Line) + ": " + e.Referenced + " does not exist in the current context"
}

type typeError struct {
	Provided string
	Expected string
	Line     int
}

func (e *typeError) Error() string {
	return "typeError at line " + strconv.Itoa(e.Line) + ": Received type " + e.Provided + ", expected " + e.Expected
}

func throw(err error) {
	if err != nil {
		panic(err)
	}
}
