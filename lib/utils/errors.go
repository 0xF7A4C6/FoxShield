package utils

import "fmt"

func HandleError(Err error) bool {
	if Err != nil {
		fmt.Println(Err)
		return true
	}

	return false
}
