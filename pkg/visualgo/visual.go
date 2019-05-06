package visualgo

/*


long map(long x, long in_min, long in_max, long out_min, long out_max) {
  return (x - in_min) * (out_max - out_min) / (in_max - in_min) + out_min;
}

from arduino ref
*/

// Utilities
func Map(coord int, upper int) float32 {
	newCoord := float32(coord)
	newUpper := float32(upper)
	return newCoord*2/newUpper - 1
}

// Exported Types
type WindowConfig struct {
	Width  int
	Height int
	Title  string
}
