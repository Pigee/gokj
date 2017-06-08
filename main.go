package main 
 
import (
 "fmt"
"io"
"os"
"bufio"
"strings"
)



type Dbinfo struct {
   host string
   database string
   username string
   password string
}

func main() {

    Initcon()
}






func Initcon() {
 f, err := os.Open("config.ini")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    rd := bufio.NewReader(f)
        for {

   	     line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
       	     if err != nil || io.EOF == err {
             break
              }
             if  strings.HasPrefix(strings.TrimSpace(line),"#")  {
                    continue  //跳过注释
                  }

              fmt.Println(line)
         } 
}
