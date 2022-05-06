package types

import (
	"math"
	"strconv"
)

type Int16 int16

// ToString returns a string representation of the value
func (i Int16) ToString() string {
	return strconv.Itoa(int(i))
}

// Toint returns an int representation of the value
func (i Int16) Toint() int {
	return int(i)
}

// ToInt returns an Int representation of the value
func (i Int16) ToInt() Int {
	return Int(i)
}

// Max returns the maximum value in the type
func (i Int16) Max(j Int16) Int16 {
	if i > j {
		return i
	}
	return j
}

// Min returns the minimum value in the type
func (i Int16) Min(j Int16) Int16 {
	if i < j {
		return i
	}
	return j
}

// Equal returns true if the values are equal
func (i Int16) Equal(j Int16) bool {
	return i == j
}

// NotEqual returns true if the values are not equal
func (i Int16) NotEqual(j Int16) bool {
	return i != j
}

// Square returns the square of the value
func (i Int16) Square() Int16 {
	return i * i
}

// Cube returns the cube of the value
func (i Int16) Cube() Int16 {
	return i * i * i
}

// Pow returns the value raised to the power of x
func (i Int16) Pow(x Int16) Int16 {
	return Int16(math.Pow(float64(i), float64(x)))
}

// Sqrt returns the square root of the value
func (i Int16) Sqrt() Int16 {
	return Int16(math.Sqrt(float64(i)))
}

// Abs returns the absolute value of the value
func (i Int16) Abs() Int16 {
	if i < 0 {
		return -i
	}
	return i
}

// Cbrt returns the cube root of the value
func (i Int16) Cbrt() Int16 {
	return Int16(math.Cbrt(float64(i)))
}

// Mul returns the product of the value and x
func (i Int16) Mul(x Int16) Int16 {
	return i * x
}

// Div returns the division of the value and x
func (i Int16) Div(x Int16) Int16 {
	if x == 0 {
		return i
	}
	return i / x
}

// Mod returns the modulo of the value and x
func (i Int16) Mod(x Int16) Int16 {
	if x == 0 {
		return i
	}
	return i % x
}

// Inc increments the value by 1
func (i Int16) Inc() Int16 {
	return i + 1
}

// Dec decrements the value by 1
func (i Int16) Dec() Int16 {
	return i - 1
}

// Neg returns the negation of the value
func (i Int16) Neg() Int16 {
	return -i
}

// Lsh returns the value left shifted by n
func (i Int16) Lsh(n uint) Int16 {
	return i << n
}

// Rsh returns the value right shifted by n
func (i Int16) Rsh(n uint) Int16 {
	return i >> n
}

// And returns the bitwise and of the value and x
func (i Int16) And(x Int16) Int16 {
	return i & x
}

// Or returns the bitwise or of the value and x
func (i Int16) Or(x Int16) Int16 {
	return i | x
}

// Xor returns the bitwise xor of the value and x
func (i Int16) Xor(x Int16) Int16 {
	return i ^ x
}

// Not returns the bitwise not of the value
func (i Int16) Not() Int16 {
	return ^i
}
