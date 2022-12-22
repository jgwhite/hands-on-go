// challenges/packages/begin/proverbs.go
package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() string {
	return proverbs.Random().Saying
}

func main() {
	fmt.Println(getRandomProverb())
}
