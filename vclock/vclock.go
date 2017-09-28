package vclock

import "sort"

type (
	Vclock interface {
		Data() string
		SetData(s string)
		Incr(id string)
		Vector() map[string]int
	}
	vclock struct {
		data string
		m    map[string]int
	}

	ByVectorClock []Vclock
)

func NewVclock(data string) Vclock {
	return &vclock{m: map[string]int{}, data: data}
}

func (v *vclock) Data() string {
	return v.data
}

func (v *vclock) SetData(data string) {
	v.data = data
}

func (v *vclock) Incr(id string) {
	i := v.m[id]
	v.m[id] = i + 1
}

func (v *vclock) Vector() map[string]int {
	return v.m
}

func (a ByVectorClock) Len() int {
	return len(a)
}
func (a ByVectorClock) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByVectorClock) Less(i, j int) bool {
	if len(a[i].Vector()) == len(a[j].Vector()) {
		for k1, v1 := range a[i].Vector() {
			if v2, ok := a[j].Vector()[k1]; !ok {
				//different maps
			} else if v1 != v2 {
				return v1 < v2
			}
		}
		return false //equal
	}
	return len(a[i].Vector()) < len(a[j].Vector())
}

func GetMostRecent(vclocks []Vclock) Vclock {
	sort.Sort(ByVectorClock(vclocks))
	winner := vclocks[len(vclocks)-1].Data()
	res := NewVclock(winner)
	for _, vclock := range vclocks {
		for k, v := range vclock.Vector() {
			if res.Vector()[k] < v {
				res.Vector()[k] = v
			}
		}
	}
	return res
}
