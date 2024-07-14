package puppy

import (
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
