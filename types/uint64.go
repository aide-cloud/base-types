package types

import (
	"math"
	"strconv"
)

type Uint64 uint64

// ToString converts uint64 to string.
func (i *Uint64) ToString() string {
	return strconv.FormatUint(uint64(*i), 10)
}

// Toint converts uint64 to int.
func (i *Uint64) Toint() int {
	return int(*i)
}

// ToInt converts uint64 to Int.
func (i *Uint64) ToInt() Int {
	return Int(*i)
}

// Greater checks if uint64 is greater than others.
func (i *Uint64) Greater(other Uint64) bool {
	return *i > other
}

// GreaterEqual checks if uint64 is greater than or equal to others.
func (i *Uint64) GreaterEqual(other Uint64) bool {
	return *i >= other
}

// Fewer checks if uint64 is lesser than others.
func (i *Uint64) Fewer(other Uint64) bool {
	return *i < other
}

// FewerEqual checks if uint64 is lesser than or equal to others.
func (i *Uint64) FewerEqual(other Uint64) bool {
	return *i <= other
}

// Equal checks if uint64 is equal to others.
func (i *Uint64) Equal(other Uint64) bool {
	return *i == other
}

// NotEqual checks if uint64 is not equal to others.
func (i *Uint64) NotEqual(other Uint64) bool {
	return *i != other
}

// Add adds uint64 to others.
func (i *Uint64) Add(other Uint64) Uint64 {
	return *i + other
}

// Sub subtracts other from uint64.
func (i *Uint64) Sub(other Uint64) Uint64 {
	return *i - other
}

// Mul multiplies uint64 by others.
func (i *Uint64) Mul(other Uint64) Uint64 {
	return *i * other
}

// Div divides uint64 by others.
func (i *Uint64) Div(other Uint64) Uint64 {
	if other == 0 {
		return *i
	}
	return *i / other
}

// Max returns the greatest of uint64 and other.
func (i *Uint64) Max(other Uint64) Uint64 {
	if *i > other {
		return *i
	}
	return other
}

// Min returns the least of uint64 and other.
func (i *Uint64) Min(other Uint64) Uint64 {
	if *i < other {
		return *i
	}
	return other
}

// Square returns the square of uint64.
func (i *Uint64) Square() Uint64 {
	return *i * *i
}

// Cube returns the cube of uint64.
func (i *Uint64) Cube() Uint64 {
	return *i * *i * *i
}

// Mod returns the remainder of uint64 divided by others.
func (i *Uint64) Mod(other Uint64) Uint64 {
	if other == 0 {
		return *i
	}
	return *i % other
}

// Pow returns the uint64 to the power of others.
func (i *Uint64) Pow(other Uint64) Uint64 {
	return Uint64(math.Pow(float64(*i), float64(other)))
}

// Abs returns the absolute value of uint64.
func (i *Uint64) Abs() Uint64 {
	if *i < 0 {
		return -*i
	}
	return *i
}

// Inc increments uint64 by 1.
func (i *Uint64) Inc() Uint64 {
	return *i + 1
}

// Dec decrements uint64 by 1.
func (i *Uint64) Dec() Uint64 {
	return *i - 1
}

// Neg returns the negative value of uint64.
func (i *Uint64) Neg() Uint64 {
	return -*i
}

// Lsh returns the uint64 left shifted by others.
func (i *Uint64) Lsh(other Uint64) Uint64 {
	return *i << other
}

// Rsh returns the uint64 right shifted by others.
func (i *Uint64) Rsh(other Uint64) Uint64 {
	return *i >> other
}

// And returns the bitwise and of uint64 and other.
func (i *Uint64) And(other Uint64) Uint64 {
	return *i & other
}

// Or returns the bitwise or of uint64 and other.
func (i *Uint64) Or(other Uint64) Uint64 {
	return *i | other
}

// Xor returns the bitwise xor of uint64 and other.
func (i *Uint64) Xor(other Uint64) Uint64 {
	return *i ^ other
}

// Not returns the bitwise not of uint64.
func (i *Uint64) Not() Uint64 {
	return ^*i
}

// Sqrt returns the square root of uint64.
func (i *Uint64) Sqrt() Uint64 {
	return Uint64(math.Sqrt(float64(*i)))
}
