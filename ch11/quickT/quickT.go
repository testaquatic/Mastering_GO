package quickT

type Point2D struct {
	X, Y int
}

func Add(x1, x2 Point2D) Point2D {
	return Point2D{x1.X + x2.X, x1.Y + x2.Y}

}