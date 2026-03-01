package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint) {
	if miles >= uint(e.milesLeft()) {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You can not make it! Need to fill up first")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 100)
}
