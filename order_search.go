package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewOrderSearchRequest() OrderSearchRequest {
	r := OrderSearchRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewOrderSearchQueryParams()
	r.pathParams = r.NewOrderSearchPathParams()
	r.requestBody = r.NewOrderSearchRequestBody()
	return r
}

type OrderSearchRequest struct {
	client      *Client
	queryParams *OrderSearchQueryParams
	pathParams  *OrderSearchPathParams
	method      string
	headers     http.Header
	requestBody OrderSearchRequestBody
}

func (r OrderSearchRequest) NewOrderSearchQueryParams() *OrderSearchQueryParams {
	return &OrderSearchQueryParams{
		From:  0,
		Count: 100,
	}
}

type OrderSearchQueryParams struct {
	Number        string `schema:"number,omitempty"`
	From          int    `schema:"from,omitempty"`
	Count         int    `schema:"count,omitempty"`
	OrderDateFrom Date   `schema:"orderDateFrom,omitempty"`
	OrderDateTo   Date   `schema:"orderDateTo,omitempty"`
}

func (p OrderSearchQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *OrderSearchRequest) QueryParams() *OrderSearchQueryParams {
	return r.queryParams
}

func (r OrderSearchRequest) NewOrderSearchPathParams() *OrderSearchPathParams {
	return &OrderSearchPathParams{}
}

type OrderSearchPathParams struct {
}

func (p *OrderSearchPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OrderSearchRequest) PathParams() *OrderSearchPathParams {
	return r.pathParams
}

func (r *OrderSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *OrderSearchRequest) Method() string {
	return r.method
}

func (r OrderSearchRequest) NewOrderSearchRequestBody() OrderSearchRequestBody {
	return OrderSearchRequestBody{}
}

type OrderSearchRequestBody struct{}

func (r *OrderSearchRequest) RequestBody() *OrderSearchRequestBody {
	return &r.requestBody
}

func (r *OrderSearchRequest) SetRequestBody(body OrderSearchRequestBody) {
	r.requestBody = body
}

func (r *OrderSearchRequest) NewResponseBody() *OrderSearchResponseBody {
	return &OrderSearchResponseBody{}
}

type OrderSearchResponseBody struct {
	FullResultSize        int    `json:"fullResultSize"`
	From                  int    `json:"from"`
	Count                 int    `json:"count"`
	VersionDigest         string `json:"versionDigest"`
	Values                Orders `json:"values"`
	TotalNumberOfPostings int    `json:"totalNumberOfPostings"`
}

func (r *OrderSearchRequest) URL() url.URL {
	return r.client.GetEndpointURL("/order", r.PathParams())
}

func (r *OrderSearchRequest) Do() (OrderSearchResponseBody, error) {
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

func (r *OrderSearchRequest) All() (OrderSearchResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := OrderSearchResponseBody{}
	concat.Count = resp.Count
	concat.From = resp.From
	concat.FullResultSize = resp.FullResultSize
	concat.TotalNumberOfPostings = resp.TotalNumberOfPostings
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
		concat.TotalNumberOfPostings = resp.TotalNumberOfPostings
		concat.Values = append(concat.Values, resp.Values...)
		concat.VersionDigest = resp.VersionDigest
	}

	return concat, nil
}
