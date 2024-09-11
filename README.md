# IMEI

## Table of Content

- [Examples](#examples)
- [License](#license)
- [Links](#links)

## Examples

```go
package main

import (
	"entrlcom.dev/imei"
)

func main() {
	v, err := imei.NewIMEI("35-209900-176148-1")
	if err != nil {
		// TODO: Handle error.
		return
	}

	_ = v.CD().String() // "1".
	_ = v.IsIMEI() // true.
	_ = v.IsIMEISV() // false.
	_ = v.SNR().String() // "176148".
	_ = v.String() // "35 209900 176148 1".
	_ = v.SVN().IsZero() // true.
	_ = v.TAC().RBI().String() // "35".
	_ = v.TAC().ID() // "209900".
}

```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Links

* [International Mobile Equipment Identity](https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity)
