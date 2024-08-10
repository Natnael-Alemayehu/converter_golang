package length

import "fmt"

type Kilometer float64
type Yard float64
type Mile float64

func (k Kilometer) String() string {
	return fmt.Sprintf("%g Kms", k)
}

func (m Mile) String() string {
	return fmt.Sprintf("%g mi", m)
}
