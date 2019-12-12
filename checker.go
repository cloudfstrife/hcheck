package hcheck

import "sync"

//Checker wrap inspect method
type Checker interface {
	//Check function, do something inspect
	Check() Status
}

var (
	checkers map[string]Checker
	once     sync.Once
)

func init() {
	once.Do(func() {
		if checkers == nil {
			checkers = make(map[string]Checker)
		}
	})
}

//Check do all check and return check result
func Check() map[string]Status {
	result := make(map[string]Status, len(checkers))

	cmap := make(map[string]chan Status, len(checkers))

	for k, v := range checkers {
		sc := make(chan Status, 1)
		go func(c Checker, cs chan Status) {
			cs <- c.Check()
		}(v, sc)
		cmap[k] = sc
	}
	for k, v := range cmap {
		result[k] = <-v
		close(v)
	}
	return result
}

//Register regist a checker
func Register(name string, checker Checker) {
	checkers[name] = checker
}

//Unregister regist a checker
func Unregister(name string) {
	delete(checkers, name)
}
