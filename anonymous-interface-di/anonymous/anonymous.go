package anonymous

import (
	"log"
)

type Doer interface {
	Do()
}

type doer struct{}

func (d doer) Do() {
	log.Println("anonymous.doer.Do")
}

type DoerFactory struct{}

func (tf DoerFactory) NewDoer() interface{ Doer } {
	return doer{}
}
