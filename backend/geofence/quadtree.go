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

// Checks the bounds to see if this is a point
func (bounds *Bounds) isPoint() bool {
	return bounds.Width == 0 && bounds.Height == 0
}

// Checks to see if bounds intersects with anouth bounds
func (bounds1 *Bounds) Intersects(bounds2 Bounds) bool {
	bounds2MaxX := bounds2.X + bounds2.Width
	bounds2MaxY := bounds2.Y + bounds2.Height
	bounds1MaxX := bounds1.X + bounds1.Width
	bounds1MaxY := bounds1.Y + bounds1.Height

	if bounds2MaxX < bounds1.X || bounds2.X > bounds1MaxX {
		return false
	}

	if bounds2MaxY < bounds1.Y || bounds2.Y > bounds1MaxY {
		return false
	}

	return true
}

// Returns the total number of nodes in a quad tree
func (tree *QuadTree) GetTotalNodes() int {
	total := 0

	if len(tree.Nodes) > 0 {
		for i := 0; i < len(tree.Nodes); i++ {
			total += 1
			total += tree.Nodes[i].GetTotalNodes()
		}
	}

	return total
}
