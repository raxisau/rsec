# RSec - Yet another Security Library

## PassGen

Password Generator - The benefit of this generator is that it does not use any confusing characters such as 1 and l and 0 and O. You can change this by setting the strings befor calling the package. Also maintains a pre-defined ratio of characters to maintains readability and security

```go
package main

import (
    "fmt"

    "github.com/raxisau/rsec"
)

func main() {
    fmt.Println("COMPLEX - has numbers symbols Alpha upper and lower")
    for i := 0; i < 10; i++ {
        fmt.Println(rsec.PassGen(30, rsec.COMPLEX))
    }

    fmt.Println("MEDIUM - has numbers Alpha upper and lower")
    for i := 0; i < 10; i++ {
        fmt.Println(rsec.PassGen(30, rsec.MEDIUM))
    }

    fmt.Println("SIMPLE - has Alpha upper and lower")
    for i := 0; i < 10; i++ {
        fmt.Println(rsec.PassGen(30, rsec.SIMPLE))
    }

    fmt.Println("NUMBERS - numbers only")
    for i := 0; i < 10; i++ {
        fmt.Println(rsec.PassGen(30, rsec.NUMBERS))
    }

    fmt.Println("NUMBERS & SYMBOLS")
    for i := 0; i < 10; i++ {
        fmt.Println(rsec.PassGen(30, rsec.NUMBERS|rsec.SYMBOLS))
    }
}
```
