package observer

import "fmt"

type Subject struct {
	observers []Observer
	msg string
}

func NewSubject() *Subject {
	s := &Subject{
		observers: make([]Observer, 0),
	}

	return s
}

func (s *Subject)Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject)Detach(observer Observer) {
	for idx, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:idx], s.observers[idx+1:]...)
			break
		}
	}
}

func (s *Subject)Notify()  {
	for _, ob := range s.observers {
		ob.Update(s)
	}
}

func (s *Subject)UpdateMsg(msg string) {
	s.msg = msg
	s.Notify()
}

type Observer interface {
	Update(subject *Subject)
}

type User struct {
	user string
}

func NewUser(user string) *User{
	return &User {
		user: user,
	}
}

func (u *User)Update(subject *Subject) {
	fmt.Printf("%s receive msg: %s\n", u.user, subject.msg)
}

