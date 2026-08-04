package main

import "interfaces/functions"

var state struct {
	Mydata *functions.DataStore
}

func init() {
	state.Mydata = &functions.DataStore{}
}

func main() {
	state.Mydata.UpdateName("John")

	data := state.Mydata.GetData()

	println(data.Name)
}