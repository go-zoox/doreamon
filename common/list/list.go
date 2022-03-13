package list

type List []interface{}

func New() List {
	return make(List, 0)
}

func (b List) Get() []interface{} {
	return b
}

func (b List) Add(value ...interface{}) {
	b = append(b, value...)
}
