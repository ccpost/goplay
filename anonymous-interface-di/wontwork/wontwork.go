package wontwork

import (
	"log"
)

type Doer interface {
	Do()
}

type doer struct{}

func (d doer) Do() {
	log.Println("wontwork.doer.Do")
}

type DoerFactory struct{}

func (tf DoerFactory) NewDoer() Doer {
	return doer{}
}
