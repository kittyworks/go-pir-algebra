// Copyright ©2013 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
)

package ffblas

type Element interface {
	//SetUint64(uint64) *Element
	//Set(*Element) *Element
	//SetInterface(interface{}) *Element
	SetZero() *Element
	SetOne() *Element
	Div(*Element,*Element) *Element
	Equal(*Element) bool
	IsZero() bool
	Cmp(*Element) int
	//LexicographicallyLargest() bool
	//SetRandom() (*Element, error)
	MulAssign(*Element) *Element
	AddAssign(*Element) *Element
	SubAssign(*Element) *Element
	Mul(*Element,*Element) *Element
	Square(*Element) *Element
	//FromMont() *Element
	Add(*Element,*Element) *Element
	Double(*Element) *Element
	Sub(*Element,*Element) *Element
	Neg(*Element) *Element
	Exp(*Element, *big.Int) *Element
}

type FF interface {
	FFLevel1
	FFLevel2
	FFLevel3
}

type FFLevel1 interface {
	Ddot(n int, x []Element, incX int, y []Element, incY int) Element
	Dnrm2(n int, x []Element, incX int) Element
	Dasum(n int, x []Element, incX int) Element
	Idamax(n int, x []Element, incX int) int
	Dswap(n int, x []Element, incX int, y []Element, incY int)
	Dcopy(n int, x []Element, incX int, y []Element, incY int)
	Daxpy(n int, alpha Element, x []Element, incX int, y []Element, incY int)
	Drotg(a, b Element) (c, s, r, z Element)
	Drotmg(d1, d2, b1, b2 Element) (p DrotmParams, rd1, rd2, rb1 Element)
	Drot(n int, x []Element, incX int, y []Element, incY int, c Element, s Element)
	Drotm(n int, x []Element, incX int, y []Element, incY int, p DrotmParams)
	Dscal(n int, alpha Element, x []Element, incX int)
}

type FFLevel2 interface {
	Dgemv(tA Transpose, m, n int, alpha Element, a []Element, lda int, x []Element, incX int, beta Element, y []Element, incY int)
	Dgbmv(tA Transpose, m, n, kL, kU int, alpha Element, a []Element, lda int, x []Element, incX int, beta Element, y []Element, incY int)
	Dtrmv(ul Uplo, tA Transpose, d Diag, n int, a []Element, lda int, x []Element, incX int)
	Dtbmv(ul Uplo, tA Transpose, d Diag, n, k int, a []Element, lda int, x []Element, incX int)
	Dtpmv(ul Uplo, tA Transpose, d Diag, n int, ap []Element, x []Element, incX int)
	Dtrsv(ul Uplo, tA Transpose, d Diag, n int, a []Element, lda int, x []Element, incX int)
	Dtbsv(ul Uplo, tA Transpose, d Diag, n, k int, a []Element, lda int, x []Element, incX int)
	Dtpsv(ul Uplo, tA Transpose, d Diag, n int, ap []Element, x []Element, incX int)
	Dsymv(ul Uplo, n int, alpha Element, a []Element, lda int, x []Element, incX int, beta Element, y []Element, incY int)
	Dsbmv(ul Uplo, n, k int, alpha Element, a []Element, lda int, x []Element, incX int, beta Element, y []Element, incY int)
	Dspmv(ul Uplo, n int, alpha Element, ap []Element, x []Element, incX int, beta Element, y []Element, incY int)
	Dger(m, n int, alpha Element, x []Element, incX int, y []Element, incY int, a []Element, lda int)
	Dsyr(ul Uplo, n int, alpha Element, x []Element, incX int, a []Element, lda int)
	Dspr(ul Uplo, n int, alpha Element, x []Element, incX int, ap []Element)
	Dsyr2(ul Uplo, n int, alpha Element, x []Element, incX int, y []Element, incY int, a []Element, lda int)
	Dspr2(ul Uplo, n int, alpha Element, x []Element, incX int, y []Element, incY int, a []Element)
}

type FFLevel3 interface {
	Dgemm(tA, tB Transpose, m, n, k int, alpha Element, a []Element, lda int, b []Element, ldb int, beta Element, c []Element, ldc int)
	Dsymm(s Side, ul Uplo, m, n int, alpha Element, a []Element, lda int, b []Element, ldb int, beta Element, c []Element, ldc int)
	Dsyrk(ul Uplo, t Transpose, n, k int, alpha Element, a []Element, lda int, beta Element, c []Element, ldc int)
	Dsyr2k(ul Uplo, t Transpose, n, k int, alpha Element, a []Element, lda int, b []Element, ldb int, beta Element, c []Element, ldc int)
	Dtrmm(s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha Element, a []Element, lda int, b []Element, ldb int)
	Dtrsm(s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha Element, a []Element, lda int, b []Element, ldb int)
}
