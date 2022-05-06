package types

type Int int

// ToString returns the string
func (i *Int) ToString() string {
	return string(rune(*i))
}
