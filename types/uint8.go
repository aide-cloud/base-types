package types

import "math"

type Uint8 uint8

// ToString returns string
func (i *Uint8) ToString() string {
	return string(*i)
}

// Toint returns int
func (i *Uint8) Toint() int {
	return int(*i)
}

// ToInt returns Int
func (i *Uint8) ToInt() Int {
	return Int(*i)
}

// ToUint returns uint
func (i *Uint8) ToUint() uint {
	return uint(*i)
}

// Greater returns true if 'i' is greater than other
func (i *Uint8) Greater(other Uint8) bool {
	return *i > other
}

// GreaterEqual returns true if 'i' is greater than or equal to other
func (i *Uint8) GreaterEqual(other Uint8) bool {
	return *i >= other
}

// Fewer returns true if 'i' is fewer than other
func (i *Uint8) Fewer(other Uint8) bool {
	return *i < other
}

// FewerEqual returns true if 'i' is fewer than or equal to other
func (i *Uint8) FewerEqual(other Uint8) bool {
	return *i <= other
}

// Equal returns true if i is equal to other
func (i *Uint8) Equal(other Uint8) bool {
	return *i == other
}

// NotEqual returns true if 'i' is not equal to other
func (i *Uint8) NotEqual(other Uint8) bool {
	return *i != other
}

// Add returns i + other
func (i *Uint8) Add(other Uint8) Uint8 {
	return *i + other
}

// Subtract returns i - other
func (i *Uint8) Subtract(other Uint8) Uint8 {
	return *i - other
}

// Inc returns i + 1
func (i *Uint8) Inc() Uint8 {
	return *i + 1
}

// Dec returns i - 1
func (i *Uint8) Dec() Uint8 {
	return *i - 1
}

// Mul returns i * other
func (i *Uint8) Mul(other Uint8) Uint8 {
	return *i * other
}

// Div returns i / other
func (i *Uint8) Div(other Uint8) Uint8 {
	if other == 0 {
		return *i
	}
	return *i / other
}

// Mod returns i % other
func (i *Uint8) Mod(other Uint8) Uint8 {
	if other == 0 {
		return *i
	}
	return *i % other
}

// Pow returns i ^ other
func (i *Uint8) Pow(other Uint8) Uint8 {
	return Uint8(math.Pow(float64(*i), float64(other)))
}

// Max returns the greatest of i and other
func (i *Uint8) Max(other Uint8) Uint8 {
	if *i > other {
		return *i
	}
	return other
}

// Min returns the least of i and other
func (i *Uint8) Min(other Uint8) Uint8 {
	if *i < other {
		return *i
	}
	return other
}

// Sqrt returns the square root of i
func (i *Uint8) Sqrt() Uint8 {
	return Uint8(math.Sqrt(float64(*i)))
}

// Log returns the logarithm of i
func (i *Uint8) Log() Uint8 {
	return Uint8(math.Log(float64(*i)))
}

// Square returns the square of i
func (i *Uint8) Square() Uint8 {
	return *i * *i
}

// Cube returns the cube of i
func (i *Uint8) Cube() Uint8 {
	return *i * *i * *i
}

// Abs returns the absolute value of i
func (i *Uint8) Abs() Uint8 {
	if *i < 0 {
		return *i
	}
	return *i
}

// Cbrt returns the cube root of i
func (i *Uint8) Cbrt() Uint8 {
	return Uint8(math.Cbrt(float64(*i)))
}

// Lsh returns i << other
func (i *Uint8) Lsh(other Uint8) Uint8 {
	return *i << other
}

// Rsh returns i >> other
func (i *Uint8) Rsh(other Uint8) Uint8 {
	return *i >> other
}

// And returns i & other
func (i *Uint8) And(other Uint8) Uint8 {
	return *i & other
}

// Or returns i | other
func (i *Uint8) Or(other Uint8) Uint8 {
	return *i | other
}

// Xor returns i ^ other
func (i *Uint8) Xor(other Uint8) Uint8 {
	return *i ^ other
}

// Not returns the bitwise inverse of i
func (i *Uint8) Not() Uint8 {
	return ^*i
}

// Neg returns the negation of i
func (i *Uint8) Neg() Uint8 {
	return -*i
}
