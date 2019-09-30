package hcheck

//Checker wrapped inspect method
type Checker interface {
	//Check function, do something inspect
	Check() (string, error)
}

var checklist map[string]Checker

//AllCheck do all check and return check result
func AllCheck() map[string]string {
	result := make(map[string]string, len(checklist))

	for k, v := range checklist {
		cr, err := v.Check()
		if err != nil {
			result[k] = cr + err.Error()
		} else {
			result[k] = cr
		}
	}

	return result
}

//GenChecker generic checker
type GenChecker struct{}

//Check GenChecker return string "OK" and nil error
func (gen *GenChecker) Check() (string, error) {
	return "OK", nil
}

func init() {
	checklist["Gen"] = &GenChecker{}
}
