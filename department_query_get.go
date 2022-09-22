package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewDepartmentQueryGetRequest() DepartmentQueryGetRequest {
	r := DepartmentQueryGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewDepartmentQueryGetQueryParams()
	r.pathParams = r.NewDepartmentQueryGetPathParams()
	r.requestBody = r.NewDepartmentQueryGetRequestBody()
	return r
}

type DepartmentQueryGetRequest struct {
	client      *Client
	queryParams *DepartmentQueryGetQueryParams
	pathParams  *DepartmentQueryGetPathParams
	method      string
	headers     http.Header
	requestBody DepartmentQueryGetRequestBody
}

func (r DepartmentQueryGetRequest) NewDepartmentQueryGetQueryParams() *DepartmentQueryGetQueryParams {
	return &DepartmentQueryGetQueryParams{
		From:  0,
		Count: 1000,
	}
}

type DepartmentQueryGetQueryParams struct {
	Query      string `schema:"query,omitempty"`
	Count      int    `schema:"count,omitempty"`
	Fields     string `schema:"fields,omitempty"`
	IsInactive bool   `schema:"isInactive,omitempty"`
	From       int    `schema:"from,omitempty"`
	Sorting    string `schema:"sorting,omitempty"`
}

func (p DepartmentQueryGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DepartmentQueryGetRequest) QueryParams() *DepartmentQueryGetQueryParams {
	return r.queryParams
}

func (r DepartmentQueryGetRequest) NewDepartmentQueryGetPathParams() *DepartmentQueryGetPathParams {
	return &DepartmentQueryGetPathParams{}
}

type DepartmentQueryGetPathParams struct {
}

func (p *DepartmentQueryGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DepartmentQueryGetRequest) PathParams() *DepartmentQueryGetPathParams {
	return r.pathParams
}

func (r *DepartmentQueryGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DepartmentQueryGetRequest) Method() string {
	return r.method
}

func (r DepartmentQueryGetRequest) NewDepartmentQueryGetRequestBody() DepartmentQueryGetRequestBody {
	return DepartmentQueryGetRequestBody{}
}

type DepartmentQueryGetRequestBody struct{}

func (r *DepartmentQueryGetRequest) RequestBody() *DepartmentQueryGetRequestBody {
	return &r.requestBody
}

func (r *DepartmentQueryGetRequest) SetRequestBody(body DepartmentQueryGetRequestBody) {
	r.requestBody = body
}

func (r *DepartmentQueryGetRequest) NewResponseBody() *DepartmentQueryGetResponseBody {
	return &DepartmentQueryGetResponseBody{}
}

type DepartmentQueryGetResponseBody struct {
	FullResultSize int         `json:"fullResultSize"`
	From           int         `json:"from"`
	Count          int         `json:"count"`
	VersionDigest  string      `json:"versionDigest"`
	Values         Departments `json:"values"`
}

func (r *DepartmentQueryGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/department/query", r.PathParams())
}

func (r *DepartmentQueryGetRequest) Do() (DepartmentQueryGetResponseBody, error) {
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

func (r *DepartmentQueryGetRequest) All() (DepartmentQueryGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := DepartmentQueryGetResponseBody{}
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
