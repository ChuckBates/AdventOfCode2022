package library

import "fmt"

func Println(value string, values ...interface{}) {
	fmt.Printf(value+"\n", values...)
}
