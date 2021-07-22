package main

import (
	"os/exec"
	"github.com/consensys/gnark-crypto/field"
	"github.com/consensys/gnark-crypto/field/generator"
)

func main(){
	modulus := "8444461749428370424248824938781546531375899335154063827935233455917409239041"
	Zp, _ := field.NewField("zp", "Element", modulus)
	generator.GenerateFF(Zp, "../zp")
	exec.Command("gofmt", "-s", "-w", "./")
}
