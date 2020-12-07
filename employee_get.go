package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewEmployeeGetRequest() EmployeeGetRequest {
	r := EmployeeGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewEmployeeGetQueryParams()
	r.pathParams = r.NewEmployeeGetPathParams()
	r.requestBody = r.NewEmployeeGetRequestBody()
	return r
}

type EmployeeGetRequest struct {
	client      *Client
	queryParams *EmployeeGetQueryParams
	pathParams  *EmployeeGetPathParams
	method      string
	headers     http.Header
	requestBody EmployeeGetRequestBody
}

func (r EmployeeGetRequest) NewEmployeeGetQueryParams() *EmployeeGetQueryParams {
	return &EmployeeGetQueryParams{
		From:  0,
		Count: 1000,
	}
}

type EmployeeGetQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p EmployeeGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *EmployeeGetRequest) QueryParams() *EmployeeGetQueryParams {
	return r.queryParams
}

func (r EmployeeGetRequest) NewEmployeeGetPathParams() *EmployeeGetPathParams {
	return &EmployeeGetPathParams{}
}

type EmployeeGetPathParams struct {
}

func (p *EmployeeGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeeGetRequest) PathParams() *EmployeeGetPathParams {
	return r.pathParams
}

func (r *EmployeeGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeeGetRequest) Method() string {
	return r.method
}

func (r EmployeeGetRequest) NewEmployeeGetRequestBody() EmployeeGetRequestBody {
	return EmployeeGetRequestBody{}
}

type EmployeeGetRequestBody struct{}

func (r *EmployeeGetRequest) RequestBody() *EmployeeGetRequestBody {
	return &r.requestBody
}

func (r *EmployeeGetRequest) SetRequestBody(body EmployeeGetRequestBody) {
	r.requestBody = body
}

func (r *EmployeeGetRequest) NewResponseBody() *EmployeeGetResponseBody {
	return &EmployeeGetResponseBody{}
}

type EmployeeGetResponseBody struct {
	FullResultSize int       `json:"fullResultSize"`
	From           int       `json:"from"`
	Count          int       `json:"count"`
	VersionDigest  string    `json:"versionDigest"`
	Values         Employees `json:"values"`
}

func (r *EmployeeGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/employee", r.PathParams())
}

func (r *EmployeeGetRequest) Do() (EmployeeGetResponseBody, error) {
	// fetch a new token if it isn't set already
	if r.client.token == "" {
		var err error
		r.client.token, err = r.client.NewToken()
		if err != nil {
			return *r.NewResponseBody(), err
		}
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *EmployeeGetRequest) All() (EmployeeGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := EmployeeGetResponseBody{}
	concat.Count = resp.Count
	concat.From = resp.From
	concat.FullResultSize = resp.FullResultSize
	concat.Values = resp.Values
	concat.VersionDigest = resp.VersionDigest

	for concat.From+concat.Count < concat.FullResultSize {
		r.QueryParams().From = r.QueryParams().From + r.QueryParams().Count
		resp, err := r.Do()
		if err != nil {
			return resp, err
		}

		concat.Count = resp.Count
		concat.From = resp.From
		concat.FullResultSize = resp.FullResultSize
		concat.Values = append(concat.Values, resp.Values...)
		concat.VersionDigest = resp.VersionDigest
	}

	return concat, nil
}
