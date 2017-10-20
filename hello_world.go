package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello World!!! \n")
	fmt.Printf("The time is: ")
	fmt.Printf("%02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())

}
