package types

import (
	"math"
	"strconv"
)

type Float32 float32

// ToString returns the string representation of the value.
func (f Float32) ToString() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

// ToInt returns the Int representation of the value.
func (f Float32) ToInt() Int {
	return Int(f)
}

// ToFloat64 returns the Float64 representation of the value.
func (f Float32) ToFloat64() Float64 {
	return Float64(f)
}

// Toint returns the integer representation of the value.
func (f Float32) Toint() int {
	return int(f)
}

// Tofloat64 returns the float64 representation of the value.
func (f Float32) Tofloat64() float64 {
	return float64(f)
}

// Tofloat32 returns the float32 representation of the value.
func (f Float32) Tofloat32() float32 {
	return float32(f)
}

// max returns the maximum of two float32 values.
func (f Float32) max(other Float32) Float32 {
	if f > other {
		return f
	}
	return other
}

// min returns the minimum of two float32 values.
func (f Float32) min(other Float32) Float32 {
	if f < other {
		return f
	}
	return other
}

// Square returns the square of the value.
func (f Float32) Square() Float32 {
	return f * f
}

// Cube returns the cube of the value.
func (f Float32) Cube() Float32 {
	return f * f * f
}

// Sqrt returns the square root of the value.
func (f Float32) Sqrt() Float32 {
	return Float32(math.Sqrt(float64(f)))
}

// Cbrt returns the cube root of the value.
func (f Float32) Cbrt() Float32 {
	return Float32(math.Cbrt(float64(f)))
}

// Pow returns the value raised to the power of the other value.
func (f Float32) Pow(other Float32) Float32 {
	return Float32(math.Pow(float64(f), float64(other)))
}

// Div returns the division of the value by the other value.
func (f Float32) Div(other Float32) Float32 {
	if other == 0 {
		return f
	}
	return f / other
}

// Log returns the logarithm of the value.
func (f Float32) Log() Float32 {
	return Float32(math.Log(float64(f)))
}

// Abs returns the absolute value of the value.
func (f Float32) Abs() Float32 {
	if f < 0 {
		return -f
	}
	return f
}
