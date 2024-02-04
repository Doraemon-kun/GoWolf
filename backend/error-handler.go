package backend

import (
	"fmt"
)

var listOfAvailError = map[int]string{
	0: "Executed Successfully without Errors",
	1: "Unknown Error Occurred",
	2: "Syntax Error",
}

func RaiseError(code int) {
	mess := fmt.Sprintf("Error %d: %s", code, listOfAvailError[code])
	fmt.Println(mess)
}
