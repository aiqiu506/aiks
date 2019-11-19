package container

import "sync"

type CI interface {
	Run(locker *sync.WaitGroup, name string)
}

type ServiceHandel struct {
	Handel map[string]CI
}


var Container ServiceHandel
func init(){
	Container= ServiceHandel{
		Handel:make(map[string]CI),
	}
}

func (ser ServiceHandel) Run() {
	lock := &sync.WaitGroup{}
	for k, v := range ser.Handel {
		lock.Add(1)
		go v.Run(lock, k)
	}
	lock.Wait()
}

func (ser ServiceHandel)RunOneByOne()  {
	lock := &sync.WaitGroup{}
	for k, v := range ser.Handel {
		lock.Add(1)
		v.Run(lock, k)
	}
	lock.Wait()
}

func (ser ServiceHandel) RegisterService(name string, ci CI) {
	if _, ok := ser.Handel[name]; !ok {
		ser.Handel[name] = ci
	}
}
