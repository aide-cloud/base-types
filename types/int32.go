package types

import "math"

type Int32 int32

// ToString converts the int32 to a string
func (i Int32) ToString() string {
	return string(i)
}

// Toint converts the int32 to an int
func (i Int32) Toint() int {
	return int(i)
}

// ToInt32 converts the int32 to an Int
func (i Int32) ToInt32() Int {
	return Int(i)
}

// Equal returns true if the int32 is equal to the passed in int32
func (i Int32) Equal(i2 Int32) bool {
	return i == i2
}

// NotEqual returns true if the int32 is not equal to the passed in int32
func (i Int32) NotEqual(i2 Int32) bool {
	return i != i2
}

// Square returns the square of the int32
func (i Int32) Square() Int32 {
	return i * i
}

// Cube returns the cube of the int32
func (i Int32) Cube() Int32 {
	return i * i * i
}

// Pow returns the passed in int32 raised to the power of the int32
func (i Int32) Pow(i2 Int32) Int32 {
	return Int32(math.Pow(float64(i), float64(i2)))
}

// Sqrt returns the square root of the int32
func (i Int32) Sqrt() Int32 {
	return Int32(math.Sqrt(float64(i)))
}

// Cbrt returns the cube root of the int32
func (i Int32) Cbrt() Int32 {
	return Int32(math.Cbrt(float64(i)))
}

// Abs returns the absolute value of the int32
func (i Int32) Abs() Int32 {
	return Int32(math.Abs(float64(i)))
}

// Mul multiplies the int32 by the passed in int32
func (i Int32) Mul(i2 Int32) Int32 {
	return i * i2
}

// Div divides the int32 by the passed in int32
func (i Int32) Div(i2 Int32) Int32 {
	if i2 == 0 {
		return i
	}
	return i / i2
}

// Mod returns the remainder of the int32 divided by the passed in int32
func (i Int32) Mod(i2 Int32) Int32 {
	if i2 == 0 {
		return i
	}
	return i % i2
}

// Add adds the int32 to the passed in int32
func (i Int32) Add(i2 Int32) Int32 {
	return i + i2
}

// Sub subtracts the passed in int32 from the int32
func (i Int32) Sub(i2 Int32) Int32 {
	return i - i2
}

// Inc increments the int32 by 1
func (i Int32) Inc() Int32 {
	return i + 1
}

// Dec decrements the int32 by 1
func (i Int32) Dec() Int32 {
	return i - 1
}

// Neg returns the negative value of the int32
func (i Int32) Neg() Int32 {
	return -i
}

// Lsh returns the int32 left shifted by the passed in int32
func (i Int32) Lsh(i2 Int32) Int32 {
	return i << i2
}

// Rsh returns the int32 right shifted by the passed in int32
func (i Int32) Rsh(i2 Int32) Int32 {
	return i >> i2
}

// And returns the bitwise and of the int32 and the passed in int32
func (i Int32) And(i2 Int32) Int32 {
	return i & i2
}

// Or returns the bitwise or of the int32 and the passed in int32
func (i Int32) Or(i2 Int32) Int32 {
	return i | i2
}

// Xor returns the bitwise xor of the int32 and the passed in int32
func (i Int32) Xor(i2 Int32) Int32 {
	return i ^ i2
}

// Not returns the bitwise not of the int32
func (i Int32) Not() Int32 {
	return ^i
}

// Log returns the natural log of the int32
func (i Int32) Log() Int32 {
	return Int32(math.Log(float64(i)))
}
