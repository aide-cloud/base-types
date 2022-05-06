package types

import (
	"math"
	"strconv"
)

type Float64 float64

// ToString returns the string representation of the value.
func (f *Float64) ToString() string {
	return strconv.FormatFloat(float64(*f), 'f', -1, 64)
}

// ToInt returns the Int representation of the value.
func (f *Float64) ToInt() Int {
	return Int(*f)
}

// Tofloat64 returns the float64 representation of the value.
func (f *Float64) Tofloat64() float64 {
	return float64(*f)
}

// Toint returns the int representation of the value.
func (f *Float64) Toint() int {
	return int(*f)
}

// Tofloat32 returns the float32 representation of the value.
func (f *Float64) Tofloat32() float32 {
	return float32(*f)
}

// max returns the maximum of two float64 values.
func (f *Float64) max(other Float64) Float64 {
	if *f > other {
		return *f
	}
	return other
}

// min returns the minimum of two float64 values.
func (f *Float64) min(other Float64) Float64 {
	if *f < other {
		return *f
	}
	return other
}

// Square returns the square of the value.
func (f *Float64) Square() Float64 {
	return *f * *f
}

// Cube returns the cube of the value.
func (f *Float64) Cube() Float64 {
	return *f * *f * *f
}

// Sqrt returns the square root of the value.
func (f *Float64) Sqrt() Float64 {
	return Float64(math.Sqrt(float64(*f)))
}

// Cbrt returns the cube root of the value.
func (f *Float64) Cbrt() Float64 {
	return Float64(math.Cbrt(float64(*f)))
}

// Pow returns the value raised to the power of the other value.
func (f *Float64) Pow(other Float64) Float64 {
	return Float64(math.Pow(float64(*f), float64(other)))
}

// Div returns the division of the value by the other value.
func (f *Float64) Div(other Float64) Float64 {
	if other == 0 {
		return *f
	}
	return *f / other
}

// Log returns the logarithm of the value.
func (f *Float64) Log() Float64 {
	return Float64(math.Log(float64(*f)))
}

// Abs returns the absolute value of the value.
func (f *Float64) Abs() Float64 {
	return Float64(math.Abs(float64(*f)))
}
