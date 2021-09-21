package main

import (
	"bytes"
	"fmt"
	"os/exec"
)


func main (){
	fmt.Println("Hello")
	command := exec.Command("./calcApp", "245", "30")

    // set var to get the output
    var out bytes.Buffer

    // set the output to our variable
    command.Stdout = &out
    err := command.Run()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out.String())
} 


