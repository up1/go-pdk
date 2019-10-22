package request

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	ch chan string
}

func NewRequest(ch chan string) *Request {
	return &Request{ch: ch}
}

func (r *Request) GetScheme() string {
	r.ch <- `kong.request.get_scheme`
	return <-r.ch
}

func (r *Request) GetHost() string {
	r.ch <- `kong.request.get_host`
	return <-r.ch
}

func (r *Request) GetPort() string {
	r.ch <- `kong.request.get_port`
	return <-r.ch
}

func (r *Request) GetForwardedScheme() string {
	r.ch <- `kong.request.get_forwarded_scheme`
	return <-r.ch
}

func (r *Request) GetForwardedHost() string {
	r.ch <- `kong.request.get_forwarded_host`
	return <-r.ch
}

func (r *Request) GetForwardedPort() string {
	r.ch <- `kong.request.get_forwarded_port`
	return <-r.ch
}

func (r *Request) GetHttpVersion() string {
	r.ch <- `kong.request.get_http_version`
	return <-r.ch
}

func (r *Request) GetMethod() string {
	r.ch <- `kong.request.get_method`
	return <-r.ch
}

func (r *Request) GetPath() string {
	r.ch <- `kong.request.get_path`
	return <-r.ch
}

func (r *Request) GetPathWithQuery() string {
	r.ch <- `kong.request.get_path_with_query`
	return <-r.ch
}

func (r *Request) GetRawQuery() string {
	r.ch <- `kong.request.get_raw_query`
	return <-r.ch
}

func (r *Request) GetQueryArg() string {
	r.ch <- `kong.request.get_query_arg`
	return <-r.ch
}

func (r *Request) GetQuery(max_args int) map[string]interface{} {
	var method string
	if max_args == -1 {
		method = "kong.request.get_query"
	} else {
		method = fmt.Sprintf(`kong.request.get_query:%d`, max_args)
	}
	r.ch <- method
	reply := <-r.ch
	query := make(map[string]interface{})
	json.Unmarshal([]byte(reply), &query)
	return query
}

func (r *Request) GetHeader(k string) string {
	r.ch <- fmt.Sprintf(`kong.request.get_header:%s`, k)
	return <-r.ch
}

func (r *Request) GetHeaders(max_headers int) map[string]interface{} {
	var method string
	if max_headers == -1 {
		method = `kong.request.get_headers`
	} else {
		method = fmt.Sprintf(`kong.request.get_headers:%d`, max_headers)
	}
	r.ch <- method
	reply := <-r.ch
	headers := make(map[string]interface{})
	json.Unmarshal([]byte(reply), &headers)
	return headers
}

func (r *Request) GetRawBody() string {
	r.ch <- `kong.request.get_raw_body`
	return <-r.ch
}

// TODO get_body
