package main

import (
	"fmt"

	"github.com/wwgberlin/timelord/db"
	"github.com/wwgberlin/timelord/vc"
)

const (
	aliceID = iota
	bobID
	cathyID
	daveID
)

var dbase db.DB

func dump() {
	fmt.Printf("db: %s\n", dbase.String())
}

type person struct {
	id int
	ts vc.Timestamp
}

func newPerson(id, groupSize int) *person {
	return &person{
		id: id,
		ts: vc.New(groupSize, id),
	}
}

func (p *person) suggest(value string) error {
	fmt.Printf("suggesting %s\n", value)
	p.ts.Tick()
	return dbase.Set(value, p.ts)
}

func (p *person) discussWith(ap *person) {
	ap.ts.Merge(p.ts)
}

func main() {
	alice := newPerson(aliceID, 4)
	bob := newPerson(bobID, 4)
	cathy := newPerson(cathyID, 4)
	dave := newPerson(daveID, 4)

	alice.suggest("Wednesday")
	dump()
	alice.discussWith(cathy)
	alice.discussWith(bob)

	bob.suggest("Tuesday")
	dump()
	bob.discussWith(dave)

	dave.suggest("Tuesday")
	dump()

	cathy.suggest("Thursday")
	dump()
}
