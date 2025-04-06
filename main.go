package main

import(
	"log"
	"fmt"
	"github.com/NoeRicklin/blog_aggregator/internal/config"
)

const userName = "glueckskeks"

func main() {
	c, err := config.Read()
	if err != nil { log.Fatal(err) }

	c.SetUser(userName)

	c, err = config.Read()
	if err != nil { log.Fatal(err) }

	fmt.Println(c)
}
