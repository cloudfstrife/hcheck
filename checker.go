package hcheck

import "sync"

//Checker wrap inspect method
type Checker interface {
	//Check function, do something inspect
	Check() Status
}

//Status wrap check result
type Status struct {
	//Pass flag
	Pass bool
	//Msg check result message
	Msg string
}

var (
	checklist map[string]Checker
	once      sync.Once
)

//Check do all check and return check result
func Check() map[string]Status {
	result := make(map[string]Status, len(checklist))

	cmap := make(map[string]chan Status, len(checklist))

	for k, v := range checklist {
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
	checklist[name] = checker
}

//Unregister regist a checker
func Unregister(name string) {
	delete(checklist, name)
}

//GenChecker generic checker
type GenChecker struct{}

//Check GenChecker return true and string "OK" as Status result
func (gen *GenChecker) Check() Status {
	return Status{Pass: true, Msg: "OK"}
}

func init() {
	once.Do(func() {
		if checklist == nil {
			checklist = make(map[string]Checker)
		}
	})
	checklist["Gen"] = &GenChecker{}
}
