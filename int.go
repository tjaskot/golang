package main

import (
	"strconv"
)

func Int(i int) int {
	i, _ = strconv.Atoi(string(strconv.Itoa(i)))
	return i
}
