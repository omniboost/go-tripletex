package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewContactGetRequest() ContactGetRequest {
	r := ContactGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewContactGetQueryParams()
	r.pathParams = r.NewContactGetPathParams()
	r.requestBody = r.NewContactGetRequestBody()
	return r
}

type ContactGetRequest struct {
	client      *Client
	queryParams *ContactGetQueryParams
	pathParams  *ContactGetPathParams
	method      string
	headers     http.Header
	requestBody ContactGetRequestBody
}

func (r ContactGetRequest) NewContactGetQueryParams() *ContactGetQueryParams {
	return &ContactGetQueryParams{
		From:  0,
		Count: 1000,
	}
}

type ContactGetQueryParams struct {
	From  int `schema:"from"`
	Count int `schema:"count"`
}

func (p ContactGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ContactGetRequest) QueryParams() *ContactGetQueryParams {
	return r.queryParams
}

func (r ContactGetRequest) NewContactGetPathParams() *ContactGetPathParams {
	return &ContactGetPathParams{}
}

type ContactGetPathParams struct {
}

func (p *ContactGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ContactGetRequest) PathParams() *ContactGetPathParams {
	return r.pathParams
}

func (r *ContactGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ContactGetRequest) Method() string {
	return r.method
}

func (r ContactGetRequest) NewContactGetRequestBody() ContactGetRequestBody {
	return ContactGetRequestBody{}
}

type ContactGetRequestBody struct{}

func (r *ContactGetRequest) RequestBody() *ContactGetRequestBody {
	return &r.requestBody
}

func (r *ContactGetRequest) SetRequestBody(body ContactGetRequestBody) {
	r.requestBody = body
}

func (r *ContactGetRequest) NewResponseBody() *ContactGetResponseBody {
	return &ContactGetResponseBody{}
}

type ContactGetResponseBody struct {
	FullResultSize int      `json:"fullResultSize"`
	From           int      `json:"from"`
	Count          int      `json:"count"`
	VersionDigest  string   `json:"versionDigest"`
	Values         Contacts `json:"values"`
}

func (r *ContactGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/contact", r.PathParams())
}

func (r *ContactGetRequest) Do() (ContactGetResponseBody, error) {
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

func (r *ContactGetRequest) All() (ContactGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := ContactGetResponseBody{}
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
