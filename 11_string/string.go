package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	name := "Subhadeep Sen"

	fmt.Println("String =", name)
	fmt.Println("Type =", reflect.TypeOf(name))
	fmt.Println("Length =", len(name))
	fmt.Println("Upper =", strings.ToUpper(name))
	fmt.Println("Lower =", strings.ToLower(name))
	fmt.Println("Prefix =", strings.HasPrefix(name, "Su"))
	fmt.Println("Suffix =", strings.HasSuffix(name, "en"))
	strArr := []string{"A", "B", "C", "D"}
	fmt.Println("Join =", strings.Join(strArr, "-"))
	fmt.Println("Repeat =", strings.Repeat(name, 2))
	fmt.Println("Contains =", strings.Contains(name, "deep"))
	fmt.Println("Index =", strings.Index(name, "ep"))
	fmt.Println("Count =", strings.Count(name, "e"))
	fmt.Println("Replace =", strings.Replace(name, "e", "E", 3))
	tokens := strings.Split(name, " ")
	for index, value := range tokens {
		fmt.Printf("Index: %d Value: %s\n", index, value)
	}
	fmt.Println("Compare =", strings.Compare(name, "name"))
	fmt.Println("TrimSpace =", strings.TrimSpace("   Hel lo  "))
	fmt.Println("ContainsAny =", strings.ContainsAny(name, "u & p"))
}
