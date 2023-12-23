package utils

import "github.com/fatih/color"

// fmt.Println with optional formatting.
func PrintLine(attributes ...color.Attribute) func(a ...interface{}) {
	return color.New(attributes...).PrintlnFunc()
}
