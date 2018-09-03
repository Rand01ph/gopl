package main

import (
	"fmt"
	"gopl/ch2/lengthconv"
	"os"
	"strconv"

	"gopl/ch2/tempconv"
	"gopl/ch2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		foot := lengthconv.Foot(t)
		metre := lengthconv.Metre(t)
		fmt.Printf("%s = %s, %s = %s\n",
			foot, lengthconv.FToM(foot), metre, lengthconv.MToF(metre))

		kilogram := weightconv.Kilogram(t)
		pound := weightconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			kilogram, weightconv.KToP(kilogram), pound, weightconv.PToK(pound))

	}
}
