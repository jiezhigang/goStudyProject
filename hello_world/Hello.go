package main

import "fmt"

const helloPrefix = "hello, "

func Hello(name string) string{
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}

