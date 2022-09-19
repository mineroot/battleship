package ship

import "fmt"

type Variant int

const (
	One Variant = iota
	Two
	Three
	Four
)

func Variants() []Variant {
	return []Variant{
		One,
		Two,
		Three,
		Four,
	}
}

func (v Variant) String() string {
	return fmt.Sprintf("%d", v.Decks())
}

func (v Variant) Decks() int {
	return int(v) + 1
}
