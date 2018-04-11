# Describe-Go
Describe your scripts at golang

## Intro
- reusable object
- set first, use later

`go get github.com/NeoJRotary/describe-go`

Current Version : 0.1.3   

GoDoc : [Doc](https://godoc.org/github.com/NeoJRotary/describe-go)  

Some issuses : [ISSUES](https://github.com/NeoJRotary/describe-go/blob/master/ISSUES.md)  

## Progess
| Type | Progess |
| --- | --- |
| String | beta |
| StringSlice | beta |
| Byte | scheduled |
| ByteSlice | scheduled |
| Error | beta |
| File | util only |
| JSON | util only |
| Time | util only |
| HTTP | alpha |
| HTTPServer | for test only |

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
