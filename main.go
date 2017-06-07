package main
 
import (
  "fmt"
  //"strings"
  "io/ioutil"
)

type Dbinfo struct {
   host string
   database string
   username string
   password string
}

func main() {
    fmt.Println("Hi,KJ")
    myRe,_ := ioutil.ReadFile("config.ini")
    fmt.Println(string(myRe))
}

