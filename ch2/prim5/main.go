package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

type configuration struct {
	Section struct {
		Enabled bool
		Path    string
	}
}

func main() {
	config := configuration{}
	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
