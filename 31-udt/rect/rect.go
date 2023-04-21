package rect

type Rect struct {
	L, B float32
	A, P float32
}

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

type Square struct {
	S float32
	a float32
}

func (s *Square) Area() float32 {
	s.a = s.S * s.S
	return s.a
}
