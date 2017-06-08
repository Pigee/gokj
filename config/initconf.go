package config 

import (
 "fmt"
"io"
"os"
"bufio"
)

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
              fmt.Println(line)
              } 
}
