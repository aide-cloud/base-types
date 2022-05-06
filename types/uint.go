package types

import (
	"math"
	"strconv"
)

type Uint uint

// ToString converts the value to a string
func (u Uint) ToString() string {
	return strconv.Itoa(int(u))
}

// ToUint8 converts the value to an Uint8
func (u Uint) ToUint8() Uint8 {
	return Uint8(u)
}

// ToUint16 converts the value to an Uint16
func (u Uint) ToUint16() Uint16 {
	return Uint16(u)
}

// ToUint32 converts the value to an Uint32
func (u Uint) ToUint32() Uint32 {
	return Uint32(u)
}

// ToUint64 converts the value to an Uint64
func (u Uint) ToUint64() Uint64 {
	return Uint64(u)
}

// Toint converts the value to an int
func (u Uint) Toint() int {
	return int(u)
}

// ToInt converts the value to an Int
func (u Uint) ToInt() Int {
	return Int(u)
}

// Greater returns true if the value is greater than the other
func (u Uint) Greater(other Uint) bool {
	return u > other
}

// GreaterEqual returns true if the value is greater than or equal to the other
func (u Uint) GreaterEqual(other Uint) bool {
	return u >= other
}

// Fewer returns true if the value is fewer than the other
func (u Uint) Fewer(other Uint) bool {
	return u < other
}

// FewerEqual returns true if the value is fewer than or equal to the other
func (u Uint) FewerEqual(other Uint) bool {
	return u <= other
}

// Equal returns true if the value is equal to the other
func (u Uint) Equal(other Uint) bool {
	return u == other
}

// NotEqual returns true if the value is not equal to the other
func (u Uint) NotEqual(other Uint) bool {
	return u != other
}

// Add adds the other value to the value
func (u Uint) Add(other Uint) Uint {
	return u + other
}

// Sub subtracts the other value from the value
func (u Uint) Sub(other Uint) Uint {
	return u - other
}

// Mul multiplies the value by the other
func (u Uint) Mul(other Uint) Uint {
	return u * other
}

// Div divides the value by the other
func (u Uint) Div(other Uint) Uint {
	if other == 0 {
		return u
	}
	return u / other
}

// Mod returns the remainder of the value divided by the other
func (u Uint) Mod(other Uint) Uint {
	if other == 0 {
		return u
	}
	return u % other
}

// Inc increments the value
func (u Uint) Inc() Uint {
	return u + 1
}

// Dec decrements the value
func (u Uint) Dec() Uint {
	return u - 1
}

// Neg returns the negation of the value
func (u Uint) Neg() Uint {
	return -u
}

// Not returns the bitwise not of the value
func (u Uint) Not() Uint {
	return ^u
}

// And returns the bitwise and of the value
func (u Uint) And(other Uint) Uint {
	return u & other
}

// Or returns the bitwise or of the value
func (u Uint) Or(other Uint) Uint {
	return u | other
}

// Xor returns the bitwise xor of the value
func (u Uint) Xor(other Uint) Uint {
	return u ^ other
}

// Lsh returns the value shifted left by the other
func (u Uint) Lsh(other Uint) Uint {
	return u << other
}

// Rsh returns the value shifted right by the other
func (u Uint) Rsh(other Uint) Uint {
	return u >> other
}

// Square returns the square of the value
func (u Uint) Square() Uint {
	return u * u
}

// Cube returns the cube of the value
func (u Uint) Cube() Uint {
	return u * u * u
}

// Sqrt returns the square root of the value
func (u Uint) Sqrt() Uint {
	return Uint(math.Sqrt(float64(u)))
}

// Cbrt returns the cube root of the value
func (u Uint) Cbrt() Uint {
	return Uint(math.Cbrt(float64(u)))
}

// Abs returns the absolute value of the value
func (u Uint) Abs() Uint {
	return Uint(math.Abs(float64(u)))
}

// Pow returns the value raised to the power of the other
func (u Uint) Pow(other Uint) Uint {
	return Uint(math.Pow(float64(u), float64(other)))
}

// Log returns the natural logarithm of the value
func (u Uint) Log() Uint {
	return Uint(math.Log(float64(u)))
}
