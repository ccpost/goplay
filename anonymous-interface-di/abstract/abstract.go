package abstract

type Doer interface {
	Do()
}

type DoerFactory interface {
	NewDoer() Doer
}

type DoerFactoryAnon interface {
	NewDoer() interface{ Doer }
}

func Do(df DoerFactory) {
	df.NewDoer().Do()
}

func DoAnon(df DoerFactoryAnon) {
	df.NewDoer().Do()
}
