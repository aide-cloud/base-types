package types

import "strconv"

type Uint32 uint32

// Tostring returns the string representation of an uint32
func (i *Uint32) String() string {
	return strconv.Itoa(int(*i))
}
