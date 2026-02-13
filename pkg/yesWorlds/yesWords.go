package yesWords

import (
	"fmt"
	"os"
	"strings"
)

func IsYes(mtn string) bool {
	yw := os.Getenv("YESWORDS")

	if len(yw) < 1 {
		yw = "yes,ok"
		//	TODO log here
	}

	yws := strings.Split(yw, ",")

	i := false
	for _, c := range yws {
		if strings.Contains(mtn, c) {
			i = true
			fmt.Println(yw, c)

		}
	}

	fmt.Println(i)
	return i

}
