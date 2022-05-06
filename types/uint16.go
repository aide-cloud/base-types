package types

import (
	"math"
	"strconv"
)

type Uint16 uint16

// ToString returns string value of Uint16
func (i *Uint16) ToString() string {
	return strconv.Itoa(int(*i))
}

// Greater returns true if i is greater than j
func (i *Uint16) Greater(j *Uint16) bool {
	return *i > *j
}

// GreaterEqual returns true if i is greater than or equal to j
func (i *Uint16) GreaterEqual(j *Uint16) bool {
	return *i >= *j
}

// Fewer returns true if 'i' is fewer than j
func (i *Uint16) Fewer(j *Uint16) bool {
	return *i < *j
}

// FewerEqual returns true if 'i' is fewer than or equal to j
func (i *Uint16) FewerEqual(j *Uint16) bool {
	return *i <= *j
}

// Equal returns true if 'i' is equal to j
func (i *Uint16) Equal(j *Uint16) bool {
	return *i == *j
}

// NotEqual returns true if 'i' is not equal to j
func (i *Uint16) NotEqual(j *Uint16) bool {
	return *i != *j
}

// Add returns the sum of 'i' and 'j'
func (i *Uint16) Add(j *Uint16) Uint16 {
	return *i + *j
}

// Max returns the greatest of 'i' and 'j'
func (i *Uint16) Max(j Uint16) Uint16 {
	if *i > j {
		return *i
	}
	return j
}

// Min returns the least of 'i' and 'j'
func (i *Uint16) Min(j Uint16) Uint16 {
	if *i < j {
		return *i
	}
	return j
}

// Square returns the square of 'i'
func (i *Uint16) Square() Uint16 {
	return *i * *i
}

// Cube returns the cube of 'i'
func (i *Uint16) Cube() Uint16 {
	return *i * *i * *i
}

// Pow returns the 'i' to the power of 'j'
func (i *Uint16) Pow(j *Uint16) Uint16 {
	return Uint16(math.Pow(float64(*i), float64(*j)))
}

// Abs returns the absolute value of 'i'
func (i *Uint16) Abs() Uint16 {
	if *i < 0 {
		return -*i
	}
	return *i
}

// Sqrt returns the square root of 'i'
func (i *Uint16) Sqrt() Uint16 {
	return Uint16(math.Sqrt(float64(*i)))
}

// Cbrt returns the cube root of 'i'
func (i *Uint16) Cbrt() Uint16 {
	return Uint16(math.Cbrt(float64(*i)))
}

// Mul returns the product of 'i' and 'j'
func (i *Uint16) Mul(j *Uint16) Uint16 {
	return *i * *j
}

// Div returns the quotient of 'i' and 'j'
func (i *Uint16) Div(j *Uint16) Uint16 {
	if *j == 0 {
		return *i
	}
	return *i / *j
}

// Mod returns the remainder of 'i' and 'j'
func (i *Uint16) Mod(j *Uint16) Uint16 {
	if *j == 0 {
		return *i
	}
	return *i % *j
}

// Inc returns the incremented value of 'i'
func (i *Uint16) Inc() Uint16 {
	return *i + 1
}

// Dec returns the decremented value of 'i'
func (i *Uint16) Dec() Uint16 {
	return *i - 1
}

// Neg returns the negated value of 'i'
func (i *Uint16) Neg() Uint16 {
	return -*i
}

// Not returns the bitwise NOT of 'i'
func (i *Uint16) Not() Uint16 {
	return ^*i
}

// Log returns the logarithm of 'i'
func (i *Uint16) Log() Uint16 {
	return Uint16(math.Log(float64(*i)))
}

// Lsh returns the value of 'i' left shifted by 'j'
func (i *Uint16) Lsh(j *Uint16) Uint16 {
	return *i << *j
}

// Rsh returns the value of 'i' right shifted by 'j'
func (i *Uint16) Rsh(j *Uint16) Uint16 {
	return *i >> *j
}

// And returns the bitwise AND of 'i' and 'j'
func (i *Uint16) And(j *Uint16) Uint16 {
	return *i & *j
}

// Or returns the bitwise OR of 'i' and 'j'
func (i *Uint16) Or(j *Uint16) Uint16 {
	return *i | *j
}

// Xor returns the bitwise XOR of 'i' and 'j'
func (i *Uint16) Xor(j *Uint16) Uint16 {
	return *i ^ *j
}
