package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewDepartmentGetRequest() DepartmentGetRequest {
	r := DepartmentGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewDepartmentGetQueryParams()
	r.pathParams = r.NewDepartmentGetPathParams()
	r.requestBody = r.NewDepartmentGetRequestBody()
	return r
}

type DepartmentGetRequest struct {
	client      *Client
	queryParams *DepartmentGetQueryParams
	pathParams  *DepartmentGetPathParams
	method      string
	headers     http.Header
	requestBody DepartmentGetRequestBody
}

func (r DepartmentGetRequest) NewDepartmentGetQueryParams() *DepartmentGetQueryParams {
	return &DepartmentGetQueryParams{
		From:  0,
		Count: 1000,
	}
}

type DepartmentGetQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p DepartmentGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DepartmentGetRequest) QueryParams() *DepartmentGetQueryParams {
	return r.queryParams
}

func (r DepartmentGetRequest) NewDepartmentGetPathParams() *DepartmentGetPathParams {
	return &DepartmentGetPathParams{}
}

type DepartmentGetPathParams struct {
}

func (p *DepartmentGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DepartmentGetRequest) PathParams() *DepartmentGetPathParams {
	return r.pathParams
}

func (r *DepartmentGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DepartmentGetRequest) Method() string {
	return r.method
}

func (r DepartmentGetRequest) NewDepartmentGetRequestBody() DepartmentGetRequestBody {
	return DepartmentGetRequestBody{}
}

type DepartmentGetRequestBody struct{}

func (r *DepartmentGetRequest) RequestBody() *DepartmentGetRequestBody {
	return &r.requestBody
}

func (r *DepartmentGetRequest) SetRequestBody(body DepartmentGetRequestBody) {
	r.requestBody = body
}

func (r *DepartmentGetRequest) NewResponseBody() *DepartmentGetResponseBody {
	return &DepartmentGetResponseBody{}
}

type DepartmentGetResponseBody struct {
	FullResultSize int         `json:"fullResultSize"`
	From           int         `json:"from"`
	Count          int         `json:"count"`
	VersionDigest  string      `json:"versionDigest"`
	Values         Departments `json:"values"`
}

func (r *DepartmentGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/department", r.PathParams())
}

func (r *DepartmentGetRequest) Do() (DepartmentGetResponseBody, error) {
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

func (r *DepartmentGetRequest) All() (DepartmentGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := DepartmentGetResponseBody{}
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
