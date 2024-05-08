package main

import "ems/internal"

func main() {
	err := internal.Init()
	if err != nil {
		panic(err)
	}
}
