package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {

	resp, err := ntp.Query("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting time: ", err)
		os.Exit(1)
	}

	fmt.Println("Current NTP time: ", resp.Time)
}
