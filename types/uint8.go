package types

type Uint8 uint8

// ToString returns string
func (i *Uint8) ToString() string {
	return string(*i)
}
