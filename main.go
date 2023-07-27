package main

import "greatproject/cmd"

func main() {

	defer cmd.Close()

	cmd.Start()
}
