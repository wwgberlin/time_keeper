package main

import (
	"fmt"

	"github.com/wwgberlin/timelord/vclock"
)

type (
	person struct {
		name   string
		client client
		vclock vclock.Vclock
	}

	client struct {
		lastRead int
		msgBoard *msgBoard
	}

	msgBoard struct {
		msgs []vclock.Vclock
	}
)

var board = msgBoard{}

func (p *person) suggest(data string) {
	fmt.Println(p.name, "suggesting", data)
	p.vclock.Incr(p.name)
	p.vclock.SetData(data)
	p.client.dispatch(p.vclock)
}

func (p *person) receive() {
	vclocks := p.client.receive()
	if len(vclocks) != 0 {
		vclock := vclock.GetMostRecent(vclocks)
		p.vclock = vclock
		fmt.Println(p.name, "received", vclock.Data(), vclock.Vector())
	} else {
		fmt.Println("no new messages")
	}
}

func newPerson(name string) person {
	return person{
		name:   name,
		client: client{msgBoard: &board},
		vclock: vclock.NewVclock(""),
	}
}

func (c *client) receive() []vclock.Vclock {
	msgs := c.msgBoard.msgs[c.lastRead:]
	c.lastRead = len(c.msgBoard.msgs)
	return msgs
}

func (c client) dispatch(vclock vclock.Vclock) {
	c.msgBoard.msgs = append(c.msgBoard.msgs, vclock)
}

func main() {

	var alice, dave, cathy, ben = newPerson("Alice"),
		newPerson("Dave"),
		newPerson("Cathy"),
		newPerson("Ben")

	alice.suggest("Wednesday")

	cathy.receive()

	ben.receive()
	ben.suggest("Tuesday")

	dave.receive()
	dave.suggest("Tuesday")

	cathy.suggest("Thursday")

	dave.receive() //conflict resolution happens here
	dave.suggest(dave.vclock.Data())

	alice.receive()
	dave.receive()
	ben.receive()
	cathy.receive()
}
