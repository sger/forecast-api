# forecast-api

Simple wrapper for Forecast api in Go (work in progress)

Checks the docs here: [https://developer.forecast.io/docs/v2](https://developer.forecast.io/docs/v2)

```
package main

import (
	"fmt"
	forecast "github.com/sger/forecast-api"
	"log"
)

func main() {

	lat := "39.3677979"
	lon := "22.9559652"

	f, err := forecast.Search("api-key", lat, lon)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Forecast %d \n", f)
	fmt.Printf("Summary %s \n", f.Currently.Summary)
}
```


