```
package main

import (
	"time"
)

import (
	_ "github.com/wuzhix/go-grace-restart"
)

func main()  {
	time.Sleep(time.Duration(100) * time.Second)
}
```