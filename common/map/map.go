package zmap

type Map map[string]interface{}

func New() Map {
	return make(Map)
}

func (b Map) Get(key string) interface{} {
	return b[key]
}

func (b Map) Set(key string, value interface{}) {
	b[key] = value
}
