package main

import (
	"fmt"

	"github.com/be-fixedassetv2/cron/v3"
)

func main() {
	_, err := cron.ParseStandard("13 50 11 30 12 * 2024")
	if err != nil {
		fmt.Println(err)

	}

}
