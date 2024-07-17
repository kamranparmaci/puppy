package puppy

import (
	"fmt"

	"github.com/kamranparmaci/dog"
)

func Bark() string {
	return "woof!"
}
func Barks() string {
	return "woof! woof! woof!"
}
func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}
func From1() {
	fmt.Println("i am from version 1.1.0")
}
