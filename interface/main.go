package main

import "fmt"

func main() {

	var (
		x  string
		x1 int32
		x2 float64
	)

	fmt.Println(CheckInterface(x))
	fmt.Println(CheckInterface(x1))
	fmt.Println(CheckInterface(x2))
}

func CheckInterface(value interface{}) string {

	switch value.(type) {
	case string:
		return "string"
	case uint, uint16, int, int16, int32, int64:
		return "int"
	case []byte:
		return "byte"
	case float32, float64:
		return "float"
	default:
		return "no"
	}
}
