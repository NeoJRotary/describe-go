# Describe-Go

Describe your scripts at golang

`go get github.com/NeoJRotary/describe-go`

GoDoc : [Doc](https://godoc.org/github.com/NeoJRotary/describe-go)  

Some issuses : [ISSUES](https://github.com/NeoJRotary/describe-go/blob/master/ISSUES.md)  

Example:
```
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
  * develop
    hotfix
  `
  list := D.String(cmdOutput).Split("\n").ElmTrimSpace()
  current := D.String(cmdOutput).Split("\n").FindHasPrefix("*").TrimPrefix("*").TrimSpace()
  fmt.Println("Current Branch :", current.Get())
  
  list = current.SetInto(list)
  fmt.Println("Branch List :", list.Get())
}

```
