package strings

type Stringer interface {
	Reverse() string
}

type str string

func (s str) Reverse() string {

	ss := ""
	for _, v := range s {
		ss = string(v) + ss
	}
	return ss
}
