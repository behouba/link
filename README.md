# HTML link parser package implementation 
## This is an exercice from gophercises.com

This package allow you to parse an HTML file and extract all of the links (`<a href="">...</a>` tags).

## Exemple:

to extract all links from this HTML documents:
```html
<html>
    <body>
    <h1>Hello!</h1>
    <a href="/other-page">A link to another page</a>
    <a href="https://akwabaexpress.ci">Akwaba Express <bold>SARL</bold></a>
    </body>
</html>
```

we use the Parse function as follow: 

```go
package main

import (
	"log"

	"github.com/behouba/link"
)

var html1 = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
  <a href="https://akwabaexpress.ci">Akwaba Express <bold>SARL</bold></a>
</body>
</html>`

func main() {
	links := link.ParseString(html1)
	log.Println(links)
}
```

will give the following output: 

```
[{/other-page A link to another page} {https://akwabaexpress.ci Akwaba Express SARL}]
```
