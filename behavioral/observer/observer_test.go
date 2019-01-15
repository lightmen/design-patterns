package observer

import "testing"

func TestObserver(t *testing.T) {

	subject := NewSubject()

	user1 := NewUser("user1")
	user2 := NewUser("user2")

	subject.Attach(user1)
	subject.Attach(user2)

	subject.UpdateMsg("to all")

	subject.Detach(user2)

	subject.UpdateMsg("just one")

}