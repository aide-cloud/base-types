package types

import "strconv"

type Uint16 uint16

// ToString returns string value of Uint16
func (i *Uint16) ToString() string {
	return strconv.Itoa(int(*i))
}
