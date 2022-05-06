package types

import (
	"math"
	"strconv"
)

type Uint32 uint32

// Tostring returns the string representation of an uint32
func (i Uint32) String() string {
	return strconv.Itoa(int(i))
}

// Toint returns the int representation of an uint32
func (i Uint32) Toint() int {
	return int(i)
}

// ToInt returns the int representation of an uint32
func (i Uint32) ToInt() Int {
	return Int(i)
}

// Greater returns true if i is greater than j
func (i Uint32) Greater(j Uint32) bool {
	return i > j
}

// GreaterEqual returns true if i is greater than or equal to j
func (i Uint32) GreaterEqual(j Uint32) bool {
	return i >= j
}

// Fewer returns true if 'i' is fewer than j
func (i Uint32) Fewer(j Uint32) bool {
	return i < j
}

// FewerEqual returns true if 'i' is fewer than or equal to j
func (i Uint32) FewerEqual(j Uint32) bool {
	return i <= j
}

// Equal returns true if 'i' is equal to j
func (i Uint32) Equal(j Uint32) bool {
	return i == j
}

// NotEqual returns true if 'i' is not equal to j
func (i Uint32) NotEqual(j Uint32) bool {
	return i != j
}

// Max returns the maximum of two uint32
func (i Uint32) Max(j Uint32) Uint32 {
	if i > j {
		return i
	}
	return j
}

// Min returns the minimum of two uint32
func (i Uint32) Min(j Uint32) Uint32 {
	if i < j {
		return i
	}
	return j
}

// Square returns the square of an uint32
func (i Uint32) Square() Uint32 {
	return i * i
}

// Cube returns the cube of an uint32
func (i Uint32) Cube() Uint32 {
	return i * i * i
}

// Pow returns the 'i' in 'i'^'j'
func (i Uint32) Pow(j Uint32) Uint32 {
	return Uint32(math.Pow(float64(i), float64(j)))
}

// Abs returns the absolute value of an uint32
func (i Uint32) Abs() Uint32 {
	if i < 0 {
		return -i
	}
	return i
}

// Mul returns the product of two uint32
func (i Uint32) Mul(j Uint32) Uint32 {
	return i * j
}

// Div returns the division of two uint32
func (i Uint32) Div(j Uint32) Uint32 {
	if j == 0 {
		return i
	}
	return i / j
}

// Mod returns the modulo of two uint32
func (i Uint32) Mod(j Uint32) Uint32 {
	if j == 0 {
		return i
	}
	return i % j
}

// Add returns the addition of two uint32
func (i Uint32) Add(j Uint32) Uint32 {
	return i + j
}

// Sub returns the subtraction of two uint32
func (i Uint32) Sub(j Uint32) Uint32 {
	return i - j
}

// Inc increments an uint32 by 1
func (i Uint32) Inc() Uint32 {
	return i + 1
}

// Dec decrements an uint32 by 1
func (i Uint32) Dec() Uint32 {
	return i - 1
}

// Neg returns the negation of an uint32
func (i Uint32) Neg() Uint32 {
	return -i
}

// Lsh returns the left bitshift of an uint32 by 'r' bits
func (i Uint32) Lsh(r Uint32) Uint32 {
	return i << r
}

// Rsh returns the right bitshift of an uint32 by 'r' bits
func (i Uint32) Rsh(r Uint32) Uint32 {
	return i >> r
}

// And returns the bitwise and of two uint32
func (i Uint32) And(j Uint32) Uint32 {
	return i & j
}

// Or returns the bitwise or of two uint32
func (i Uint32) Or(j Uint32) Uint32 {
	return i | j
}

// Xor returns the bitwise xor of two uint32
func (i Uint32) Xor(j Uint32) Uint32 {
	return i ^ j
}

// Not returns the bitwise not of an uint32
func (i Uint32) Not() Uint32 {
	return ^i
}
