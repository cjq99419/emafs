package errHandle

import "fmt"

func CheckErr(err error, comments string) {
	if err != nil {
		fmt.Printf("%s [%s]", comments, err)
	}
}
