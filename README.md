# 八一数理测算
  
通过八一数理和康熙字典笔画测算名称和测算手机号码

<https://pkg.go.dev/github.com/cnk3x/bys>

```go
//api 很简单
package main

import "github.com/cnk3x/bys"

func main(){
    nr := bys.AssessName("文天祥")
    fmt.Printf("%+v", nr)
    pr := bys.AssessPhone("13812348765")
    fmt.Printf("%+v", pr)
}

```
