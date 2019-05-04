package visualgo

var vertexList []float32 // vertices of each shape to be drawn in the next frame held here

// var (
// 	vertexList = []float32{
// 		0, 0.5, 0,
// 		-0.5, -0.5, 0,
// 		0.5, -0.5, 0,
// 	}
// )

// func (conf *WindowConfig) mapX(x int) float32 {

// }

// func (conf *WindowConfig) mapY(y int) float32 {
// }

func (conf *WindowConfig) Triangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int) {
	// draws triangle with all three vertices specified
	// Placeholder test triangle
	vertexList = append(vertexList, float32(x0))
	vertexList = append(vertexList, float32(y0))
	vertexList = append(vertexList, 0)

	vertexList = append(vertexList, float32(x1))
	vertexList = append(vertexList, float32(y1))
	vertexList = append(vertexList, 0)

	vertexList = append(vertexList, float32(x2))
	vertexList = append(vertexList, float32(y2))
	vertexList = append(vertexList, 0)
}

// func (conf *WindowConfig) rect(x int, y int) {
// 	// Will draw rect from left corner
// }
