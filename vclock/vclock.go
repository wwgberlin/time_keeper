package vclock

import (
	"fmt"
	"sort"
)

type (
	VClock interface {
		Data() string
		SetData(s string)
		Incr(id string)
		ToS() string
	}
	vClock struct {
		data string
		m    map[string]int
	}

	ByVectorClock []VClock
)

func NewVclock(data string) VClock {
	return &vClock{m: map[string]int{}, data: data}
}

func (v *vClock) Data() string {
	return v.data
}
func (v *vClock) ToS() string {
	return fmt.Sprintf("%v", v.m)
}

func (v *vClock) SetData(data string) {
	v.data = data
}

func (v *vClock) Incr(id string) {
	i := v.m[id]
	v.m[id] = i + 1
}

func (a ByVectorClock) Len() int {
	return len(a)
}
func (a ByVectorClock) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByVectorClock) Less(i, j int) bool {
	v1 := a[i].(*vClock)
	v2 := a[j].(*vClock)
	if len(v1.m) == len(v2.m) {
		for k1, v1 := range v1.m {
			if v2, ok := v2.m[k1]; !ok {
				//different maps
			} else if v1 != v2 {
				return v1 < v2
			}
		}
		return false //equal
	}
	return len(v1.m) < len(v2.m)
}

func GetMostRecent(vclocks []VClock) VClock {
	sort.Sort(ByVectorClock(vclocks))
	winner := vclocks[len(vclocks)-1].Data()
	res := NewVclock(winner).(*vClock)
	for _, vc := range vclocks {
		v := vc.(*vClock)
		for k, v := range v.m {
			if res.m[k] < v {
				res.m[k] = v
			}
		}
	}
	return res
}
