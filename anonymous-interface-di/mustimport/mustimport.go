package mustimport

import (
	"log"

	"github.com/ccpost/goplay/anonymous-interface-di/abstract"
)

type doer struct{}

func (d doer) Do() {
	log.Println("mustimport.doer.Do")
}

type DoerFactory struct{}

func (tf DoerFactory) NewDoer() abstract.Doer {
	return doer{}
}
