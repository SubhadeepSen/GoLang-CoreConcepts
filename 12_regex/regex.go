package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(".in")
	fmt.Println(re.FindString("www.amamzon.in"))
	fmt.Println(re.FindString("www.google.in"))
	fmt.Println(re.FindString("www.facebook.com"))

	fmt.Println(re.FindStringIndex("www.google.in"))

	re = regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(re.FindStringSubmatch("flying"))
	fmt.Println(re.FindStringSubmatch("abcfloatingxyz"))
}
