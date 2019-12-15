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
