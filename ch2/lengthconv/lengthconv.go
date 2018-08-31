package lengthconv

import "fmt"

type Foot float64
type Metre float64

const (

)

func (f Foot) String() string {return fmt.Sprintf("%gft", f)}
func (m Metre) String() string {return fmt.Sprintf("%gm", m)}