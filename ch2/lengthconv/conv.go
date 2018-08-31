package lengthconv

func MToF(m Metre) Foot {
	return Foot(m * 3.281)
}

func FToM(f Foot) Metre {
	return Metre(f / 3.281)
}
