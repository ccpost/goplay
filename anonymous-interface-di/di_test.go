package di_test

import (
	"testing"

	"github.com/ccpost/goplay/anonymous-interface-di/abstract"
	"github.com/ccpost/goplay/anonymous-interface-di/anonymous"
	"github.com/ccpost/goplay/anonymous-interface-di/mustimport"
)

// This will not compile; un-comment to see error.
// func TestWontWork(t *testing.T) {
//	df := wontwork.DoerFactory{}
//	abstract.Do(df)
//}

func TestMustImport(t *testing.T) {
	df := mustimport.DoerFactory{}
	abstract.Do(df)
}

func TestAnonyous(t *testing.T) {
	df := anonymous.DoerFactory{}
	abstract.DoAnon(df)
}
