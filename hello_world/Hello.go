package main

const helloPrefix = "hello, "
const ChineseHelloPrefix = "你好，"
const FrenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Chinese" {
		return ChineseHelloPrefix + name
	}

	if language == "French" {
		return FrenchHelloPrefix + name
	}
	return helloPrefix + name
}
