# hcheck

`hcheck` is a library for checking the health of components.

## go doc 

```
$ go doc -all

package hcheck // import "github.com/cloudfstrife/hcheck"


FUNCTIONS

func Check() map[string]Status
    Check do all check and return check result

func Register(name string, checker Checker)
    Register regist a checker

func Unregister(name string)
    Unregister regist a checker


TYPES

type Checker interface {
        //Check function, do something inspect
        Check() Status
}
    Checker wrap inspect method

type GenChecker struct{}
    GenChecker generic checker

func (gen *GenChecker) Check() Status
    Check GenChecker return true and string "OK" as Status result

type Status struct {
        //Pass flag
        Pass bool
        //Msg check result message
        Msg string
}
    Status wrap check result

```

## usage

**command**

```
mkdir testing
cd testing
go mod init app/testing
go get -u github.com/cloudfstrife/hcheck
mkdir -p checker
touch main.go
touch checker/mq.go
touch checker/db.go
```

**main.go**

```
package main

import (
	"fmt"

	_ "app/testing/checker"

	"github.com/cloudfstrife/hcheck"
)

func main() {
	cr := hcheck.Check()
	for k, v := range cr {
		fmt.Printf("%s\t%t\t%s\n", k, v.Pass, v.Msg)
	}
}

```
**checker/mq.go**

```
package checker

import (
	"github.com/cloudfstrife/hcheck"
)

//MQChecker message queue health checker
type MQChecker struct {
	server   string
	port     int
	username string
	password string
}

//Check message queue connectable
func (dc *MQChecker) Check() hcheck.Status {
	//do  message queue connection check
	return hcheck.Status{
		Pass: false,
		Msg:  "MQ Connection lost",
	}
}

func init() {
	hcheck.Register("MQ", &MQChecker{})
}

```

**checker/db.go**

```
package checker

import "github.com/cloudfstrife/hcheck"

//DBChecker database health checker
type DBChecker struct {
	server   string
	port     int
	username string
	password string
}

//Check the database connectable
func (dc *DBChecker) Check() hcheck.Status {
	//do database connection check
	return hcheck.Status{
		Pass: true,
		Msg:  "MySQL Processlist Count : 100",
	}
}

func init() {
	hcheck.Register("DB", &DBChecker{})
}

```

**build & run**

```
$ go build
$ ./testing.exe
MQ      false   MQ Connection lost
Gen     true    OK
DB      true    MySQL Processlist Count : 100
```
