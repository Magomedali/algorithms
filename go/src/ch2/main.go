package main

import "fmt"
import "go-packages/gohub/go_operations"

func main() {

	fmt.Println("Welcome to goland!")

	displayIntTypes()
	displayFloatType()
	displayString()
	displayBolean()
	go_operations.IntOperations()
}

func displayIntTypes() [11]string {

	types := [11]string{
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		"uint",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
	}

	fmt.Println("--Know that in golang exists following integer types:")
	fmt.Print("---- ")
	for i := 0; i < len(types); i++ {
		fmt.Print(types[i] + ", ")
	}
	fmt.Println("")

	return types
}


func displayFloatType() {

	fmt.Println("--Know that in golang exists following floats types:")
	fmt.Print("---- ")

	var floats [2]string
	floats = [2]string{
		"float32",
		"float64",
	}

	for _, value := range floats {
		fmt.Print(value + ", ")
	}
	fmt.Println("")
}


func displayString() {
	fmt.Println("--Know that in golang exists also string type")

	fmt.Println("--- string")
}


func displayBolean() {

	fmt.Println("--Know that in golang exists also boolean types")
	fmt.Print("---- ")
	for _, value := range [2]string{"true","false"} {
		fmt.Print(value + ", ")
	}
	fmt.Println("")
}


