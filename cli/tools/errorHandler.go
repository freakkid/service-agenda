package tools

import (
	"os"
	"fmt"
)


func DealMessage(ok bool, message string) {
	if !ok {
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}	
}