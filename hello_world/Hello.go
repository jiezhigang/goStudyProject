package main

const helloPrefix = "hello, "
const ChineseHelloPrefix = "你好，"
const FrenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := helloPrefix
	
	switch language {
	case "French":
		prefix = FrenchHelloPrefix
	case "Chinese":
		prefix = ChineseHelloPrefix
	}
	return prefix + name
}
