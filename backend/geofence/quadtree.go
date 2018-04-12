package geo

// QuadTree data structure
type QuadTree struct {
	Bounds    Bounds
	MaxItem   int
	MaxHeight int
	Height    int
	Item      []Bounds
	Nodes     []QuadTree
	Total     int
}

// Bounding box with an x,y origin and width, heigh
type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}
