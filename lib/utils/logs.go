package utils

import "fmt"

func Debug(Type, Content string) {
	fmt.Println(fmt.Sprintf("%s | %s.", Type, Content))
}
