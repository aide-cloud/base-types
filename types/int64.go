package types

import "strconv"

type Int64 int64

// ToString converts the int64 to a string.
func (i *Int64) ToString() string {
	return strconv.FormatInt(int64(*i), 10)
}

// Toint converts the int64 to an int.
func (i *Int64) Toint() int {
	return int(*i)
}

// ToInt converts the int64 to an Int.
func (i *Int64) ToInt() Int {
	return Int(*i)
}

// Greater returns true if the int64 is greater than the given int64.
func (i *Int64) Greater(j Int64) bool {
	return *i > j
}

// GreaterEqual returns true if the int64 is greater than or equal to the given int64.
func (i *Int64) GreaterEqual(j Int64) bool {
	return *i >= j
}

// Fewer returns true if the int64 is less than the given int64.
func (i *Int64) Fewer(j Int64) bool {
	return *i < j
}

// FewerEqual returns true if the int64 is less than or equal to the given int64.
func (i *Int64) FewerEqual(j Int64) bool {
	return *i <= j
}

// Equal returns true if the int64 is equal to the given int64.
func (i *Int64) Equal(j Int64) bool {
	return *i == j
}

// NotEqual returns true if the int64 is not equal to the given int64.
func (i *Int64) NotEqual(j Int64) bool {
	return *i != j
}

// Add returns the sum of the int64 and the given int64.
func (i *Int64) Add(j Int64) Int64 {
	return *i + j
}

// Sub returns the difference of the int64 and the given int64.
func (i *Int64) Sub(j Int64) Int64 {
	return *i - j
}

// Mul returns the product of the int64 and the given int64.
func (i *Int64) Mul(j Int64) Int64 {
	return *i * j
}

// Div returns the quotient of the int64 and the given int64.
func (i *Int64) Div(j Int64) Int64 {
	if j == 0 {
		return *i
	}
	return *i / j
}

// Mod returns the remainder of the int64 and the given int64.
func (i *Int64) Mod(j Int64) Int64 {
	if j == 0 {
		return *i
	}
	return *i % j
}

// Inc returns the int64 incremented by 1.
func (i *Int64) Inc() Int64 {
	return *i + 1
}

// Dec returns the int64 decremented by 1.
func (i *Int64) Dec() Int64 {
	return *i - 1
}

// Neg returns the negative of the int64.
func (i *Int64) Neg() Int64 {
	return -*i
}

// Abs returns the absolute value of the int64.
func (i *Int64) Abs() Int64 {
	if *i < 0 {
		return -*i
	}
	return *i
}

// Lsh returns the int64 left-shifted by the given number of bits.
func (i *Int64) Lsh(j uint) Int64 {
	return *i << j
}

// Rsh returns the int64 right-shifted by the given number of bits.
func (i *Int64) Rsh(j uint) Int64 {
	return *i >> j
}

// And returns the bitwise-and of the int64 and the given int64.
func (i *Int64) And(j Int64) Int64 {
	return *i & j
}

// Xor returns the bitwise-xor of the int64 and the given int64.
func (i *Int64) Xor(j Int64) Int64 {
	return *i ^ j
}

// Or returns the bitwise-or of the int64 and the given int64.
func (i *Int64) Or(j Int64) Int64 {
	return *i | j
}

// Not returns the bitwise-not of the int64.
func (i *Int64) Not() Int64 {
	return ^*i
}

// Max returns the maximum of the int64 and the given int64.
func (i *Int64) Max(j Int64) Int64 {
	if *i > j {
		return *i
	}
	return j
}

// Min returns the minimum of the int64 and the given int64.
func (i *Int64) Min(j Int64) Int64 {
	if *i < j {
		return *i
	}
	return j
}
