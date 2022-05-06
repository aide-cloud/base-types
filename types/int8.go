package types

import (
	"math"
	"strconv"
)

type Int8 int8

// ToString converts the int8 to a string
func (i *Int8) ToString() string {
	return strconv.Itoa(int(*i))
}

// Toint converts the int8 to an int
func (i *Int8) Toint() int {
	return int(*i)
}

// ToInt converts the int8 to an Int
func (i *Int8) ToInt() Int {
	return Int(*i)
}

// Max returns the maximum value of the int8
func (i *Int8) Max(other Int8) Int8 {
	if *i > other {
		return *i
	}
	return other
}

// Min returns the minimum value of the int8
func (i *Int8) Min(other Int8) Int8 {
	if *i < other {
		return *i
	}
	return other
}

// Square returns the square of the value.
func (i *Int8) Square() Int8 {
	return *i * *i
}

// Cube returns the cube of the value.
func (i *Int8) Cube() Int8 {
	return *i * *i * *i
}

// Sqrt returns the square root of the value.
func (i *Int8) Sqrt() Int8 {
	return Int8(math.Sqrt(float64(*i)))
}

// Cbrt returns the cube root of the value.
func (i *Int8) Cbrt() Int8 {
	return Int8(math.Cbrt(float64(*i)))
}

// Pow returns the value raised to the power of the other value.
func (i *Int8) Pow(other Int8) Int8 {
	return Int8(math.Pow(float64(*i), float64(other)))
}

// Div returns the division of the value by the other value.
func (i *Int8) Div(other Int8) Int8 {
	return Int8(math.Floor(float64(*i) / float64(other)))
}

// Log returns the logarithm of the value.
func (i *Int8) Log(base Int8) Int8 {
	return Int8(math.Log(float64(*i)) / math.Log(float64(base)))
}

// Abs returns the absolute value of the value.
func (i *Int8) Abs() Int8 {
	if *i < 0 {
		return -*i
	}
	return *i
}

// Mod returns the modulo of the value.
func (i *Int8) Mod(other Int8) Int8 {
	if other == 0 {
		return *i
	}
	return *i % other
}

// Inc increments the value by 1.
func (i *Int8) Inc() Int8 {
	return *i + 1
}

// Dec decrements the value by 1.
func (i *Int8) Dec() Int8 {
	return *i - 1
}

// Add returns the sum of the value and the other value.
func (i *Int8) Add(other Int8) Int8 {
	return *i + other
}

// Sub returns the difference of the value and the other value.
func (i *Int8) Sub(other Int8) Int8 {
	return *i - other
}

// Mul returns the product of the value and the other value.
func (i *Int8) Mul(other Int8) Int8 {
	return *i * other
}

// Equal returns true if the value is equal to the other value.
func (i *Int8) Equal(other Int8) bool {
	return *i == other
}

// NotEqual returns true if the value is not equal to the other value.
func (i *Int8) NotEqual(other Int8) bool {
	return *i != other
}

// Greater returns true if the value is greater than the other value.
func (i *Int8) Greater(other Int8) bool {
	return *i > other
}

// GreaterEqual returns true if the value is greater than or equal to the other value.
func (i *Int8) GreaterEqual(other Int8) bool {
	return *i >= other
}

// Fewer returns true if the value is fewer than the other value.
func (i *Int8) Fewer(other Int8) bool {
	return *i < other
}

// FewerEqual returns true if the value is fewer than or equal to the other value.
func (i *Int8) FewerEqual(other Int8) bool {
	return *i <= other
}

// Neg returns the negation of the value.
func (i *Int8) Neg() Int8 {
	return -*i
}

// Lsh returns the value left shifted by the other value.
func (i *Int8) Lsh(other Int8) Int8 {
	return *i << other
}

// Rsh returns the value right shifted by the other value.
func (i *Int8) Rsh(other Int8) Int8 {
	return *i >> other
}

// And returns the bitwise and of the value and the other value.
func (i *Int8) And(other Int8) Int8 {
	return *i & other
}

// Or returns the bitwise or of the value and the other value.
func (i *Int8) Or(other Int8) Int8 {
	return *i | other
}

// Xor returns the bitwise xor of the value and the other value.
func (i *Int8) Xor(other Int8) Int8 {
	return *i ^ other
}

// Not returns the bitwise not of the value.
func (i *Int8) Not() Int8 {
	return ^*i
}
