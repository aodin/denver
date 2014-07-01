package main

import (
	"flag"
	"fmt"
	"github.com/aodin/denver/crime"
)

func InitSQL() {
	// TODO Actually create the table
	fmt.Println(crime.Codes.Create())
}

func main() {
	var init bool
	flag.BoolVar(&init, "init", false, "print SQL for CREATE TABLE")
	flag.Parse()

	if init {
		InitSQL()
	}
}
