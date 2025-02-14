package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const SidesCircle = 0
const SidesTriangle = 3
const SidesSquare = 4

func CalcSquare(sideLen float64, sidesNum uint32) float64 {
	var squared float64 = sideLen * sideLen
	switch sidesNum {
	case 0:
		return math.Pi * squared
	case 3:
		return math.Sqrt(3.0) * squared / 4.0
	case 4:
		return squared
	default:
		return 0
	}
}
