package main

import (
	"bytes"
	"fmt"
)

var (
	msg1 = "hello!"
	msg2 = "here you are."
	msg3 = "I am still here waiting you."
)

func main() {

	var out bytes.Buffer

	allMsgGroup := [][]byte{}

	out.Write([]byte(msg1))
	fmt.Println("msg1 is", string(out.Bytes()))
	allMsgGroup = append(allMsgGroup, out.Bytes())
	out.Reset()

	out.Write([]byte(msg2))
	fmt.Println("msg2 is", string(out.Bytes()))
	allMsgGroup = append(allMsgGroup, out.Bytes())
	out.Reset()

	out.Write([]byte(msg3))
	fmt.Println("msg3 is", string(out.Bytes()))
	allMsgGroup = append(allMsgGroup, out.Bytes())
	out.Reset()

	for k, v := range allMsgGroup {
		fmt.Println("")
		fmt.Printf("Index %v Msg one by one is: %v", k, string(v))
	}
}
