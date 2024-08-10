package weight

import "fmt"

type Killogram float64
type Pound float64

func (k Killogram) String() string { return fmt.Sprintf("%g kgs", k) }

func (p Pound) String() string { return fmt.Sprintf("%g lbs", p) }
