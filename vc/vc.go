package vc

// Timestamp is a represenation of timestamap based on vector clock
type Timestamp struct {
	vc   []int
	line int
}

// New creates a new timestamp in a group timelines. For examle each server or
// actor in distributed system may have its own timeline.
// size is a group size - e.g. number of servers
// line - index of a particular timeline in the group which will be considered as a local
// time for this vector clock (like time in particular server)
func New(size, line int) Timestamp {
	return Timestamp{}
}

// Tick is clock move in local timeline. Whenever a process does work, increment the logical clock value of the node in the vector
func (t Timestamp) Tick() {
}

// Merge happens when actors in 2 different timelines communicate, for example when a message with timestamp is recieved.
// What is reuired:
// - update each element in the vector to be max(local, received)
// - increment the logical clock value representing the current node in the vector
// r remote timeline, recieved as part of a message
func (t Timestamp) Merge(r Timestamp) {
}

const (
	// Equals means that two timestamps are equal
	Equals = 0
	// HappensBefore means that the reciver timestamp happened before
	HappensBefore = -1
	// HappensAfter means that the reciver timestamp happened after
	HappensAfter = 1
	// NotComparable means that "happens before" relatioinship does not exists for 2 given timestamps, they are independent
	NotComparable = -100
)

// HappensBefore detects what kind of happen-before relationship exists
// between two timestamps. If it exists it could be Equals, HappensBefore or
// HappensAfter. If they are from independent timelines NotComparable must be returned
func (t Timestamp) HappensBefore(r Timestamp) int {
	return 0
}
