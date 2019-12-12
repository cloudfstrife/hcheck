package hcheck

//GenChecker generic checker
type GenChecker struct{}

//Check GenChecker return true and string "OK" as Status result
func (gen *GenChecker) Check() Status {
	return Status{Pass: true, Msg: "OK"}
}

func init() {
	Register("Gen", &GenChecker{})
}
