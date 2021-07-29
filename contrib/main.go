package main

import (
	"os/exec"
	"github.com/consensys/gnark-crypto/field"
	"github.com/consensys/gnark-crypto/field/generator"
)

func main(){
	//modulus := "8444461749428370424248824938781546531375899335154063827935233455917409239041"
	modulus := "258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177"
	Zp, _ := field.NewField("zp", "Element", modulus)
	generator.GenerateFF(Zp, "../zp")
	exec.Command("gofmt", "-s", "-w", "./")
}
