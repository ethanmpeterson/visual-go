package visualgo

// Utilities
func MapX(coord int, upper int) float32 { // maps px coordinates to OpenGL coordinates
	newCoord := float32(coord)
	newUpper := float32(upper)
	return newCoord*2/newUpper - 1
}

func MapY(coord int, upper int) float32 { // maps px coordinates to OpenGL coordinates
	newCoord := float32(coord)
	newUpper := float32(upper)
	return (newCoord*2/newUpper - 1) * -1
}

// Exported Types
type WindowConfig struct {
	Width  int
	Height int
	Title  string
}
