# IMEI

## Table of Content

- [Authors](#authors)
- [Examples](#examples)
- [Links](#links)

## Authors

| Name         | GitHub                                             |
|--------------|----------------------------------------------------|
| Klim Sidorov | [@entrlcom-klim](https://github.com/entrlcom-klim) |

## Examples

```go
package main

import (
	"flida.dev/imei"
)

func main() {
	v, err := imei.NewIMEI("35-209900-176148-1")
	if err != nil {
		// TODO: ...
		return
	}

	// ...
}

```

## Links

- [International Mobile Equipment Identity](https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity)
