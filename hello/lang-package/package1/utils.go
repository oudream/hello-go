package main

import "strconv"

type Utils struct {
	Name string
	Age int
}

var (
	UtilsName string
	count int
)

func (u Utils) ShowName() string {
	count++
	UtilsName = "xxxxxxxxxxxx"
	return u.Name + strconv.Itoa(count)
}
