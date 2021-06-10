package main

import (
	"github.com/consensys/gnark-crypto/field"
	"github.com/consensys/gnark-crypto/field/generator"
)

//go:generate go run main.go
func main(){
	config := field.Field{
		PackageName: "FieldMatrix",
		ElementName: "FieldMatrix",
	}
	generator.GenerateFF(&config,"./destinationPath")
}
