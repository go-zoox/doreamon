package fetch

import (
	"net/http"

	zmap "github.com/go-zoox/doreamon/common/map"
)

type Fetch struct {
	client   *http.Client
	request  *http.Request
	response *http.Response
	//
	headers *zmap.Map
	params  *zmap.Map
	query   *zmap.Map
	body    interface{}
	//
	isSent bool
	//
	// pipelines
}

func New() Fetch {
	client := &http.Client{}
	return Fetch{
		client: client,
	}
}

// an pipeline
func (f *Fetch) Exec() *http.Response {
	if !f.isSent {
		f.doRequest()
	}

	return f.response
}

func (f *Fetch) doRequest() *Fetch {
	if f.isSent {
		return f
	}
	f.isSent = true

	response, err := f.client.Do(f.request)
	if err != nil {
		panic(err)
	}

	f.response = response

	return f
}

func (f *Fetch) Get(url string) *Fetch {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	f.request = request

	return f
}

func (f *Fetch) Params(v *zmap.Map) *Fetch {
	f.params = v
	return f
}

func (f *Fetch) Query(v *zmap.Map) *Fetch {
	f.query = v
	return f
}

func (f *Fetch) Body(v interface{}) *Fetch {
	f.body = v
	return f
}
