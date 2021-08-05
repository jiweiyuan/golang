package cu

import (
	"log"
	"testing"
)

func TestFetch(t *testing.T) {
	wanted, err := Fetch()
	if err != nil {
		log.Fatalln(err)
	}

	println(wanted.USD)
}
