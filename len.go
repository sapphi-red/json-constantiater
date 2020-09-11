package main

const host32bit = ^uint(0)>>32 == 0

func getLenOfSimpleType(typeName string) uint {
	switch typeName {
	case "bool":
		return 5
	case "int":
		if host32bit {
			return 11
		} else {
			return 20
		}
	case "int8":
		return 4
	case "int16":
		return 6
	case "int32":
		return 11
	case "int64":
		return 20
	case "uint":
		if host32bit {
			return 10
		} else {
			return 19
		}
	case "uint8":
		return 3
	case "uint16":
		return 5
	case "uint32":
		return 10
	case "uint64":
		return 19
	case "float32":
		return 22
	case "float64":
		return 24
	}
	return 0
}
