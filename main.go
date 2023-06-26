package main

// first we need to create de go.mod file, to do that we go
// to the terminal and type: go mod init <name of the module>
// this will require like a map where first is github.com and the second my name and then the name of the module
// see example in go.mod file

import (
	"github.com/dosorio55/mapas/mapas"
)

func main() {
	mapas.MostrarMapas()
}
