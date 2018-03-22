package examples

import (
	"fmt"

	D "github.com/NeoJRotary/describe-go"
)

// Find a string inside slice, modify it
func main() {
	// git branch
	cmdOutput := `
	  newFeature
		master
	*	develop
		hotfix
	`
	list := D.String(cmdOutput).Split("\n").TrimSpace()
	current := D.String(cmdOutput).Split("\n").FindHasPrefix("*").Trim("*").TrimSpace()
	fmt.Println("Current Branch :", current.Get())
	list = current.SetInto(list)
	fmt.Println("Branch List :", list.Get())
}
