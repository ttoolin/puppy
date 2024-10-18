package puppy

import (
	"github.com/ttoolin/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof woof woof"
}

func BigBark() string {
	return dog.WhenGrowsUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowsUp(Barks())
}
