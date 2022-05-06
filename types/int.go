package types

import (
	"math"
)

type Int int

// ToString returns the string
func (i *Int) ToString() string {
	return string(rune(*i))
}

// Toint returns the int
func (i *Int) Toint() int {
	return int(*i)
}

// Toint8 returns the int8
func (i *Int) Toint8() int8 {
	return int8(*i)
}

// Toint16 returns the int16
func (i *Int) Toint16() int16 {
	return int16(*i)
}

// Toint32 returns the int32
func (i *Int) Toint32() int32 {
	return int32(*i)
}

// Toint64 returns the int64
func (i *Int) Toint64() int64 {
	return int64(*i)
}

// Touint returns the uint
func (i *Int) Touint() uint {
	return uint(*i)
}

// Touint8 returns the uint8
func (i *Int) Touint8() uint8 {
	return uint8(*i)
}

// Touint16 returns the uint16
func (i *Int) Touint16() uint16 {
	return uint16(*i)
}

// Touint32 returns the uint32
func (i *Int) Touint32() uint32 {
	return uint32(*i)
}

// Touint64 returns the uint64
func (i *Int) Touint64() uint64 {
	return uint64(*i)
}

// Tofloat32 returns the float32
func (i *Int) Tofloat32() float32 {
	return float32(*i)
}

// Tofloat64 returns the float64
func (i *Int) Tofloat64() float64 {
	return float64(*i)
}

// ToInt8 returns the Int8
func (i *Int) ToInt8() Int8 {
	return Int8(*i)
}

// ToInt16 returns the Int16
func (i *Int) ToInt16() Int16 {
	return Int16(*i)
}

// ToInt32 returns the Int32
func (i *Int) ToInt32() Int32 {
	return Int32(*i)
}

// ToInt64 returns the Int64
func (i *Int) ToInt64() Int64 {
	return Int64(*i)
}

// ToUint8 returns the Uint8
func (i *Int) ToUint8() Uint8 {
	return Uint8(*i)
}

// ToUint16 returns the Uint16
func (i *Int) ToUint16() Uint16 {
	return Uint16(*i)
}

// ToUint32 returns the Uint32
func (i *Int) ToUint32() Uint32 {
	return Uint32(*i)
}

// ToUint64 returns the Uint64
func (i *Int) ToUint64() Uint64 {
	return Uint64(*i)
}

// ToFloat32 returns the Float32
func (i *Int) ToFloat32() Float32 {
	return Float32(*i)
}

// ToFloat64 returns the Float64
func (i *Int) ToFloat64() Float64 {
	return Float64(*i)
}

// Greater returns true if 'i' is greater than other
func (i *Int) Greater(y Int) bool {
	return *i > y
}

// GreaterEqual returns true if 'i' is greater than or equal to other
func (i *Int) GreaterEqual(y Int) bool {
	return *i >= y
}

// FewerEqual returns true if 'i' is less than or equal to other
func (i *Int) FewerEqual(y Int) bool {
	return *i <= y
}

// Fewer Returns true if 'i' is less than other
func (i *Int) Fewer(y Int) bool {
	return *i < y
}

// max returns the maximum of two ints
func (i *Int) max(y Int) Int {
	if *i > y {
		return *i
	}
	return y
}

// min returns the minimum of two ints
func (i *Int) min(y Int) Int {
	if *i < y {
		return *i
	}
	return y
}

// Equal returns true if 'i' is equal to other
func (i *Int) Equal(y Int) bool {
	return *i == y
}

// NotEqual returns true if 'i' is not equal to other
func (i *Int) NotEqual(y Int) bool {
	return *i != y
}

// Square returns the square of the number
func (i *Int) Square() Int {
	return *i * *i
}

// Cube returns the cube of the number
func (i *Int) Cube() Int {
	return *i * *i * *i
}

// Pow returns the number raised to the power of the other
func (i *Int) Pow(n Int) Int {
	return Int(math.Pow(float64(*i), float64(n)))
}

// Sqrt returns the square root of the number
func (i *Int) Sqrt() Int {
	return Int(math.Sqrt(float64(*i)))
}

// Cbrt returns the cube root of the number
func (i *Int) Cbrt() Int {
	return Int(math.Cbrt(float64(*i)))
}

// Abs returns the absolute value of the number
func (i *Int) Abs() Int {
	if *i < 0 {
		return -*i
	}
	return *i
}

// Mul multiplies two numbers
func (i *Int) Mul(y Int) Int {
	return *i * y
}

// Div divides two numbers
func (i *Int) Div(y Int) Int {
	if y != 0 {
		return *i / y
	}
	return *i
}

// Mod returns the remainder of two numbers
func (i *Int) Mod(y Int) Int {
	if y != 0 {
		return *i % y
	}
	return *i
}

// Log returns the log of the number
func (i *Int) Log() Int {
	return Int(math.Log(float64(*i)))
}

// Inc increments the number
func (i *Int) Inc() Int {
	return *i + 1
}

// Dec decrements the number
func (i *Int) Dec() Int {
	return *i - 1
}

// Add adds two numbers
func (i *Int) Add(y Int) Int {
	return *i + y
}

// Sub subtracts two numbers
func (i *Int) Sub(y Int) Int {
	return *i - y
}

// Neg returns the negative of the number
func (i *Int) Neg() Int {
	return -*i
}

// Lsh returns the number left shifted by the other
func (i *Int) Lsh(y Int) Int {
	return *i << y
}

// Rsh returns the number right shifted by the other
func (i *Int) Rsh(y Int) Int {
	return *i >> y
}

// And returns the number and the other
func (i *Int) And(y Int) Int {
	return *i & y
}

// Or returns the number or the other
func (i *Int) Or(y Int) Int {
	return *i | y
}

// Xor returns the number xor the other
func (i *Int) Xor(y Int) Int {
	return *i ^ y
}

// Not returns the number not
func (i *Int) Not() Int {
	return ^*i
}
