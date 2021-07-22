// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64 || noasm || gccgo || safe
// +build !amd64 noasm gccgo safe

package f64

// DotUnitary is
//  for i, v := range x {
//  	sum += y[i] * v
//  }
//  return sum
func DotUnitary(x, y []zp.Element) (sum zp.Element) {
	for i, v := range x {
		//sum += y[i] * v
		a := zp.Element{}
		a.mul(y[i],v)
		sum.Add(a,sum)
	}
	return sum
}

// DotInc is
//  for i := 0; i < int(n); i++ {
//  	sum += y[iy] * x[ix]
//  	ix += incX
//  	iy += incY
//  }
//  return sum
func DotInc(x, y []zp.Element, n, incX, incY, ix, iy uintptr) (sum zp.Element) {
	for i := 0; i < int(n); i++ {
		sum += y[iy] * x[ix]
		ix += incX
		iy += incY
	}
	return sum
}
