package types

import "strconv"

type Uint64 uint64

// ToString converts uint64 to string.
func (i *Uint64) ToString() string {
	return strconv.FormatUint(uint64(*i), 10)
}
