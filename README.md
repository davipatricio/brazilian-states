# brazilian-states

Get informations about Brazilian States.

## Install

```sh
$ go get github.com/davipatricio/brazilian-states
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/davipatricio/brazilian-states"
)

func main() {
    /*
    Returns a slice of all Brazilian States.
    [
        {
            "Name": "Acre",
            "Code": "AC",
            "Capital": "Rio Branco",
            "Region": "Norte",
            "Cities": 22
        },
        ...
        ...
    ]
    */
    fmt.Println(brazilianstates.GetAll())
    fmt.Println("\n")

    /*
    Return data according to parameter. Expects state code.
    {
        "Name": "São Paulo",
        "Code": "SP",
        "Capital": "São Paulo",
        "Region": "Sudeste",
        "Cities": 645
    }
    */
    saoPauloState := brazilianstates.Get("SP")
    fmt.Println(saoPauloState)
    fmt.Println(saoPauloState.Name)
    fmt.Println(saoPauloState.Code)
    fmt.Println(saoPauloState.Capital)
    fmt.Println(saoPauloState.Region)
    fmt.Println(saoPauloState.Cities)
}
```

## Related

[brazilian-states](https://github.com/fernandoporazzi/brazilian-states) - Get informations about Brazilian states

## License
MIT