package proxy


type Subject interface {
	Request() string
}

type RealSubject struct {}

func (r *RealSubject)Request() string {
	return "RealSubject req"
}

type Proxy struct {
	real *RealSubject
}

func (p *Proxy)PreRequest() string {
	return "Proxy PreReq"
}

func (p *Proxy)Request() string {
	req := p.PreRequest()
	req += ", "
	req += p.real.Request()

	return req
}