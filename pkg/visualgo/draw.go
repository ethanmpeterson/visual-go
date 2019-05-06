package visualgo

var vertexList []float32 // vertices of each shape to be drawn in the next frame held here

// func (conf *WindowConfig) mapY(y int) float32 {
// }

func (conf *WindowConfig) Triangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int) {
	// draws triangle with all three vertices specified
	// Placeholder test triangle
	vertexList = append(vertexList, Map(x0, conf.Width))
	vertexList = append(vertexList, float32(y0))
	vertexList = append(vertexList, 0)

	vertexList = append(vertexList, Map(x1, conf.Width))
	vertexList = append(vertexList, float32(y1))
	vertexList = append(vertexList, 0)

	vertexList = append(vertexList, Map(x2, conf.Width))
	vertexList = append(vertexList, float32(y2))
	vertexList = append(vertexList, 0)
}
