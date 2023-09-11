package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewCountryGetRequest() CountryGetRequest {
	r := CountryGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewCountryGetQueryParams()
	r.pathParams = r.NewCountryGetPathParams()
	r.requestBody = r.NewCountryGetRequestBody()
	return r
}

type CountryGetRequest struct {
	client      *Client
	queryParams *CountryGetQueryParams
	pathParams  *CountryGetPathParams
	method      string
	headers     http.Header
	requestBody CountryGetRequestBody
}

func (r CountryGetRequest) NewCountryGetQueryParams() *CountryGetQueryParams {
	return &CountryGetQueryParams{}
}

type CountryGetQueryParams struct {
	ID         string `schema:"id,omitempty"`
	Code       string `schema:"code,omitempty"`
	IsDisabled bool   `schema:"isDisabled,omitempty"`
	From       int    `schema:"from,omitempty"`
	Count      int    `schema:"count,omitempty"`
	Sorting    string `schema:"sorting,omitempty"`
	Fields     string `schema:"fields,omitempty"`
}

func (p CountryGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CountryGetRequest) QueryParams() *CountryGetQueryParams {
	return r.queryParams
}

func (r CountryGetRequest) NewCountryGetPathParams() *CountryGetPathParams {
	return &CountryGetPathParams{}
}

type CountryGetPathParams struct {
}

func (p *CountryGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CountryGetRequest) PathParams() *CountryGetPathParams {
	return r.pathParams
}

func (r *CountryGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CountryGetRequest) Method() string {
	return r.method
}

func (r CountryGetRequest) NewCountryGetRequestBody() CountryGetRequestBody {
	return CountryGetRequestBody{}
}

type CountryGetRequestBody struct{}

func (r *CountryGetRequest) RequestBody() *CountryGetRequestBody {
	return &r.requestBody
}

func (r *CountryGetRequest) SetRequestBody(body CountryGetRequestBody) {
	r.requestBody = body
}

func (r *CountryGetRequest) NewResponseBody() *CountryGetResponseBody {
	return &CountryGetResponseBody{}
}

type CountryGetResponseBody struct {
	FullResultSize int       `json:"fullResultSize"`
	From           int       `json:"from"`
	Count          int       `json:"count"`
	VersionDigest  string    `json:"versionDigest"`
	Values         Countries `json:"values"`
}

func (r *CountryGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/country", r.PathParams())
}

func (r *CountryGetRequest) Do() (CountryGetResponseBody, error) {
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
