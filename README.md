# go-realip-in-context

Reads Real IP from request and saves in context.

We use [realip](github.com/tomasen/realip) to find the client's real IP, then
we save it to context. That allow us to get the client's IP in app code that
doesn't have access to the request, only the context.

We expose a middleware and a function.

## Usage

Examples use go-chi.


```go
package main

import (
  // ...
  "github.com/nrfta/go-realip-in-context"
)

// ...
router := chi.NewRouter()

router.Use(realip.Middleware)
// ...
```

```go
func something(ctx context.Context) {
  ip := realip.GetRealIP(ctx)

  fmt.Println(ip)
}
```
## License

This project is licensed under the [MIT License](LICENSE.md).
