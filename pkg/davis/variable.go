package davis

import "fmt"

type Variable uint

func (v Variable) String() string {
	switch {
	case v == 1:
		return "Y"
	case v%2 == 0:
		return fmt.Sprintf("X%d", v/2)
	default: // v %2 == 1
		return fmt.Sprintf("Z%d", (v-1)/2)
	}
}
