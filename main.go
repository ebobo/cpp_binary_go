package main

import (
	"bytes"
	"fmt"
	"os/exec"
)


func main (){

	fmt.Println("Hello, Qi")
	command := exec.Command("/home/qi/Downloads/calcApp", "212", "30")

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