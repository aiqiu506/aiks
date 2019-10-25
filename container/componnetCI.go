package container


var ComponentCI ComponentHandel


type ci interface {
	Init(config interface{})
}

type ComponentHandel struct {
	Handel map[string]ci
}
func (ser ComponentHandel) RegisterComponent(name string, c ci) {
	if _, ok := ser.Handel[name]; !ok {
		ser.Handel[name] = c
	}
}
func init()  {
	ComponentCI=ComponentHandel{
		Handel:make(map[string]ci),
	}
}