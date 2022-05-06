package types

import "strconv"

type Uint uint

// ToString converts the value to a string
func (u *Uint) ToString() string {
	return strconv.Itoa(int(*u))
}

// ToUint8 converts the value to an Uint8
func (u *Uint) ToUint8() Uint8 {
	return Uint8(*u)
}

// ToUint16 converts the value to an Uint16
func (u *Uint) ToUint16() Uint16 {
	return Uint16(*u)
}

// ToUint32 converts the value to an Uint32
func (u *Uint) ToUint32() Uint32 {
	return Uint32(*u)
}

// ToUint64 converts the value to an Uint64
func (u *Uint) ToUint64() Uint64 {
	return Uint64(*u)
}

// Toint converts the value to an int
func (u *Uint) Toint() int {
	return int(*u)
}

// ToInt converts the value to an Int
func (u *Uint) ToInt() Int {
	return Int(*u)
}
