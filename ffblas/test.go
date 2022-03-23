package ffblas_test

import (
	"testing"
	"ffblas"
)

func TestDefine(t *testing.T){
	ffblas.Use(ffblas)
	_ := ffblas.Implementation()
}
