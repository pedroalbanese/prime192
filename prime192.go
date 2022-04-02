// Parameters for the Prime192v1/2/3 Elliptic curve
package prime192

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once
var prime192v1 *elliptic.CurveParams

func initPrime192v1() {
	prime192v1 = new(elliptic.CurveParams)
	prime192v1.P, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffeffffffffffffffff", 16)
	prime192v1.N, _ = new(big.Int).SetString("ffffffffffffffffffffffff99def836146bc9b1b4d22831", 16)
	prime192v1.B, _ = new(big.Int).SetString("64210519e59c80e70fa7e9ab72243049feb8deecc146b9b1", 16)
	prime192v1.Gx, _ = new(big.Int).SetString("188da80eb03090f67cbf20eb43a18800f4ff0afd82ff1012", 16)
	prime192v1.Gy, _ = new(big.Int).SetString("07192b95ffc8da78631011ed6b24cdd573f977a11e794811", 16)
	prime192v1.BitSize = 192
}

func Prime192v1() elliptic.Curve {
	initonce.Do(initPrime192v1)
	return prime192v1
}

var prime192v2 *elliptic.CurveParams

func initPrime192v2() {
	prime192v2 = new(elliptic.CurveParams)
	prime192v2.P, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffeffffffffffffffff", 16)
	prime192v2.N, _ = new(big.Int).SetString("fffffffffffffffffffffffe5fb1a724dc80418648d8dd31", 16)
	prime192v2.B, _ = new(big.Int).SetString("cc22d6dfb95c6b25e49c0d6364a4e5980c393aa21668d953", 16)
	prime192v2.Gx, _ = new(big.Int).SetString("eea2bae7e1497842f2de7769cfe9c989c072ad696f48034a", 16)
	prime192v2.Gy, _ = new(big.Int).SetString("6574d11d69b6ec7a672bb82a083df2f2b0847de970b2de15", 16)
	prime192v2.BitSize = 192
}

func Prime192v2() elliptic.Curve {
	initonce.Do(initPrime192v2)
	return prime192v2
}

var prime192v3 *elliptic.CurveParams

func initPrime192v3() {
	prime192v3 = new(elliptic.CurveParams)
	prime192v3.P, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffeffffffffffffffff", 16)
	prime192v3.N, _ = new(big.Int).SetString("ffffffffffffffffffffffff7a62d031c83f4294f640ec13", 16)
	prime192v3.B, _ = new(big.Int).SetString("22123dc2395a05caa7423daeccc94760a7d462256bd56916", 16)
	prime192v3.Gx, _ = new(big.Int).SetString("7d29778100c65a1da1783716588dce2b8b4aee8e228f1896", 16)
	prime192v3.Gy, _ = new(big.Int).SetString("38a90f22637337334b49dcb66a6dc8f9978aca7648a943b0", 16)
	prime192v3.BitSize = 192
}

func Prime192v3() elliptic.Curve {
	initonce.Do(initPrime192v3)
	return prime192v3
}
