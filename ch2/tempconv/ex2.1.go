package tempconv

import (
	"fmt"
)

type Kelvin float64

const (
	AbsoluteZeroK Kelvin = -273.15
)


func (k Kelvin) String() string {return fmt.Sprintf("%g°K", k)}


func CToK(c Celsius) Kelvin {return  Kelvin(c - AbsoluteZeroC)}

func KToC(k Kelvin) Celsius {return  Celsius(k + AbsoluteZeroK)}