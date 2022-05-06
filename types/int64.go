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
