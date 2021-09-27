package err_handler

import (
	"fmt"
	"os"
)

func HandleError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg, err.Error())
		os.Exit(1)
	}
}
