// Copyright ©2013 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/kittyworks/go-pir-algebra/zp"
)

package ffblas

type FFFloat64 interface {
	FFFloat64Level1
	FFFloat64Level2
	FFFloat64Level3
}

type FFFloat64Level1 interface {
	Ddot(n int, x []zp.Element, incX int, y []zp.Element, incY int) zp.Element
	Dnrm2(n int, x []zp.Element, incX int) zp.Element
	Dasum(n int, x []zp.Element, incX int) zp.Element
	Idamax(n int, x []zp.Element, incX int) int
	Dswap(n int, x []zp.Element, incX int, y []zp.Element, incY int)
	Dcopy(n int, x []zp.Element, incX int, y []zp.Element, incY int)
	Daxpy(n int, alpha zp.Element, x []zp.Element, incX int, y []zp.Element, incY int)
	Drotg(a, b zp.Element) (c, s, r, z zp.Element)
	Drotmg(d1, d2, b1, b2 zp.Element) (p DrotmParams, rd1, rd2, rb1 zp.Element)
	Drot(n int, x []zp.Element, incX int, y []zp.Element, incY int, c zp.Element, s zp.Element)
	Drotm(n int, x []zp.Element, incX int, y []zp.Element, incY int, p DrotmParams)
	Dscal(n int, alpha zp.Element, x []zp.Element, incX int)
}

type FFFloat64Level2 interface {
	Dgemv(tA Transpose, m, n int, alpha zp.Element, a []zp.Element, lda int, x []zp.Element, incX int, beta zp.Element, y []zp.Element, incY int)
	Dgbmv(tA Transpose, m, n, kL, kU int, alpha zp.Element, a []zp.Element, lda int, x []zp.Element, incX int, beta zp.Element, y []zp.Element, incY int)
	Dtrmv(ul Uplo, tA Transpose, d Diag, n int, a []zp.Element, lda int, x []zp.Element, incX int)
	Dtbmv(ul Uplo, tA Transpose, d Diag, n, k int, a []zp.Element, lda int, x []zp.Element, incX int)
	Dtpmv(ul Uplo, tA Transpose, d Diag, n int, ap []zp.Element, x []zp.Element, incX int)
	Dtrsv(ul Uplo, tA Transpose, d Diag, n int, a []zp.Element, lda int, x []zp.Element, incX int)
	Dtbsv(ul Uplo, tA Transpose, d Diag, n, k int, a []zp.Element, lda int, x []zp.Element, incX int)
	Dtpsv(ul Uplo, tA Transpose, d Diag, n int, ap []zp.Element, x []zp.Element, incX int)
	Dsymv(ul Uplo, n int, alpha zp.Element, a []zp.Element, lda int, x []zp.Element, incX int, beta zp.Element, y []zp.Element, incY int)
	Dsbmv(ul Uplo, n, k int, alpha zp.Element, a []zp.Element, lda int, x []zp.Element, incX int, beta zp.Element, y []zp.Element, incY int)
	Dspmv(ul Uplo, n int, alpha zp.Element, ap []zp.Element, x []zp.Element, incX int, beta zp.Element, y []zp.Element, incY int)
	Dger(m, n int, alpha zp.Element, x []zp.Element, incX int, y []zp.Element, incY int, a []zp.Element, lda int)
	Dsyr(ul Uplo, n int, alpha zp.Element, x []zp.Element, incX int, a []zp.Element, lda int)
	Dspr(ul Uplo, n int, alpha zp.Element, x []zp.Element, incX int, ap []zp.Element)
	Dsyr2(ul Uplo, n int, alpha zp.Element, x []zp.Element, incX int, y []zp.Element, incY int, a []zp.Element, lda int)
	Dspr2(ul Uplo, n int, alpha zp.Element, x []zp.Element, incX int, y []zp.Element, incY int, a []zp.Element)
}

type FFFloat64Level3 interface {
	Dgemm(tA, tB Transpose, m, n, k int, alpha zp.Element, a []zp.Element, lda int, b []zp.Element, ldb int, beta zp.Element, c []zp.Element, ldc int)
	Dsymm(s Side, ul Uplo, m, n int, alpha zp.Element, a []zp.Element, lda int, b []zp.Element, ldb int, beta zp.Element, c []zp.Element, ldc int)
	Dsyrk(ul Uplo, t Transpose, n, k int, alpha zp.Element, a []zp.Element, lda int, beta zp.Element, c []zp.Element, ldc int)
	Dsyr2k(ul Uplo, t Transpose, n, k int, alpha zp.Element, a []zp.Element, lda int, b []zp.Element, ldb int, beta zp.Element, c []zp.Element, ldc int)
	Dtrmm(s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha zp.Element, a []zp.Element, lda int, b []zp.Element, ldb int)
	Dtrsm(s Side, ul Uplo, tA Transpose, d Diag, m, n int, alpha zp.Element, a []zp.Element, lda int, b []zp.Element, ldb int)
}
