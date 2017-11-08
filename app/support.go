package app

import (
	"fmt"
	"os"
)

func QuitOnError(err error) {
	if err == nil {
		return
	}

	fmt.Println("ERROR:", err)
	os.Exit(1)
}
