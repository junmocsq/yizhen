package main

import (
	"fmt"
	"github.com/junmocsq/yizhen/greetings"
)

func main() {
	message := greetings.Hello("张小盒")
	fmt.Println(message)
}
