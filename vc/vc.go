package vc

type Timestamp struct {
	vc   []int
	line int
}

func New(size, line int) Timestamp {
	return Timestamp{
		vc:   make([]int, size),
		line: line,
	}
}

func (t Timestamp) Tick() {
	t.vc[t.line]++
}

func (t Timestamp) Merge(r Timestamp) {
	for i := range t.vc {
		if t.vc[i] < r.vc[i] {
			t.vc[i] = r.vc[i]
		}
	}
	t.vc[t.line]++
}

const (
	Equals        = 0
	HappensBefore = -1
	HappensAfter  = 1
	NotComparable = -100
)

func (t Timestamp) HappensBefore(r Timestamp) int {
	sign := func(a, b int) int {
		switch {
		case a > b:
			return 1
		case a < b:
			return -1
		}
		return 0
	}
	ac := 0
	for i := range t.vc {
		s := sign(t.vc[i], r.vc[i])
		if s == 0 {
			continue
		}
		if ac == 0 {
			ac = s
			continue
		}
		if ac != s {
			return NotComparable
		}
	}
	return ac
}
