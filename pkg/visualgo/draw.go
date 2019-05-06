package visualgo

var vertexList []float32 // vertices of each shape to be drawn in the next frame held here

func (conf *WindowConfig) Triangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int) {
	// draws triangle with all three vertices specified
	// Placeholder test triangle
	vertexList = append(vertexList, MapX(x0, conf.Width))
	vertexList = append(vertexList, MapY(y0, conf.Height))
	vertexList = append(vertexList, 0)

	vertexList = append(vertexList, MapX(x1, conf.Width))
	vertexList = append(vertexList, MapY(y1, conf.Height))
	vertexList = append(vertexList, 0)

	vertexList = append(vertexList, MapX(x2, conf.Width))
	vertexList = append(vertexList, MapY(y2, conf.Height))
	vertexList = append(vertexList, 0)
}
