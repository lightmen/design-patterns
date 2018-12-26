package  adapter


type Target interface {
	Request() string
}


type Adaptee interface {
	SpecificRequest() string
}


type AdapteeImpl struct {}

func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

func (this *AdapteeImpl) SpecificRequest() string{
	return "SpecificRequest"
}

type Adapter struct {
	Adaptee
}

func (this *Adapter) Request() string {
	return this.Adaptee.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target{
	return &Adapter{
		Adaptee: adaptee,
	}
}
