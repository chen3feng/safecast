# safecast

Safe Numeric Type Cast Library for Go

English | [简体中文](README_zh.md)

[![License Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-red.svg)](COPYING)
[![Python](https://img.shields.io/badge/Language-go1.18+-blue.svg)](https://www.python.org/)
![Build Status](https://github.com/chen3feng/safecast/actions/workflows/go.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/chen3feng/safecast/badge.svg?branch=master)](https://coveralls.io/github/chen3feng/safecast?branch=master)
[![GoReport](https://goreportcard.com/badge/github.com/securego/gosec)](https://goreportcard.com/report/github.com/chen3feng/safecast)

Safe numeric type cast library. suppoer all integral and floating types, except uintptr.

```go
val, ok := To[type](value)
```

`ok == false` indicates overflow occured. but whatever，`val` is always equals to the result of the type cast (`type(value)`) expression。

This library depends on go generics, which is introduced in 1.18+.

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# safecast

```go
import "github.com/chen3feng/safecast"
```

## Index

- [func To[T numericType, F numericType](value F) (to T, ok bool)](<#func-to>)


## func [To](<https://github.com/chen3feng/safecast/blob/master/generics.go#L9>)

```go
func To[T numericType, F numericType](value F) (to T, ok bool)
```

<details><summary>Example (6alue In Range)</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/chen3feng/safecast"
)

func main() {
	n, ok := safecast.To[uint](1)
	fmt.Print(n, ok)
}
```

#### Output

```
1 true
```

</p>
</details>

<details><summary>Example (Float Overflow)</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/chen3feng/safecast"
	"math"
)

func main() {
	n, ok := safecast.To[float32](math.MaxFloat32 * 2)
	fmt.Print(n, ok)
}
```

#### Output

```
+Inf false
```

</p>
</details>

<details><summary>Example (Int No Overflow)</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/chen3feng/safecast"
)

func main() {
	b, ok := safecast.To[byte](255)
	fmt.Print(b, ok)
}
```

#### Output

```
255 true
```

</p>
</details>

<details><summary>Example (Int Overflow)</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/chen3feng/safecast"
)

func main() {
	b, ok := safecast.To[byte](256)
	fmt.Print(b, ok)
}
```

#### Output

```
0 false
```

</p>
</details>

<details><summary>Example (Value Out Of Range)</summary>
<p>

```go
package main

import (
	"fmt"
	"github.com/chen3feng/safecast"
)

func main() {
	n, ok := safecast.To[uint32](-1)
	fmt.Print(n, ok)
}
```

#### Output

```
4294967295 false
```

</p>
</details>


Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->
