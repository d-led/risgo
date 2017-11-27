package app

import (
	"fmt"
	"os"
)

// QuitOnError : display the error and quit if it's not nil
func QuitOnError(err error) {
	if err == nil {
		return
	}

	fmt.Println("ERROR:", err)
	os.Exit(1)
}
