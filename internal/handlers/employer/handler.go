package employer

type Employer interface {
	SetName(name string)
	GetName() string
}

type employer struct {
	Name string
}

func New(name string) Employer {
	return &employer{
		Name: name,
	}
}

func (e *employer) SetName(name string) {
	e.Name = name
}

func (e *employer) GetName() string {
	return e.Name
}
