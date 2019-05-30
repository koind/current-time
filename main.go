package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	tm, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tm)
}
