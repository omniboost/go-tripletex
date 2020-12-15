package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewLedgerVoucherNonPostedGetRequest() LedgerVoucherNonPostedGetRequest {
	r := LedgerVoucherNonPostedGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewLedgerVoucherNonPostedGetQueryParams()
	r.pathParams = r.NewLedgerVoucherNonPostedGetPathParams()
	r.requestBody = r.NewLedgerVoucherNonPostedGetRequestBody()
	return r
}

type LedgerVoucherNonPostedGetRequest struct {
	client      *Client
	queryParams *LedgerVoucherNonPostedGetQueryParams
	pathParams  *LedgerVoucherNonPostedGetPathParams
	method      string
	headers     http.Header
	requestBody LedgerVoucherNonPostedGetRequestBody
}

func (r LedgerVoucherNonPostedGetRequest) NewLedgerVoucherNonPostedGetQueryParams() *LedgerVoucherNonPostedGetQueryParams {
	return &LedgerVoucherNonPostedGetQueryParams{
		From:  0,
		Count: 100,
	}
}

type LedgerVoucherNonPostedGetQueryParams struct {
	From     int  `schema:"from"`
	Count    int  `schema:"count"`
	DateFrom Date `schema:"dateFrom,omitempty"`
	DateTo   Date `schema:"dateTo,omitempty"`
}

func (p LedgerVoucherNonPostedGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerVoucherNonPostedGetRequest) QueryParams() *LedgerVoucherNonPostedGetQueryParams {
	return r.queryParams
}

func (r LedgerVoucherNonPostedGetRequest) NewLedgerVoucherNonPostedGetPathParams() *LedgerVoucherNonPostedGetPathParams {
	return &LedgerVoucherNonPostedGetPathParams{}
}

type LedgerVoucherNonPostedGetPathParams struct {
}

func (p *LedgerVoucherNonPostedGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerVoucherNonPostedGetRequest) PathParams() *LedgerVoucherNonPostedGetPathParams {
	return r.pathParams
}

func (r *LedgerVoucherNonPostedGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerVoucherNonPostedGetRequest) Method() string {
	return r.method
}

func (r LedgerVoucherNonPostedGetRequest) NewLedgerVoucherNonPostedGetRequestBody() LedgerVoucherNonPostedGetRequestBody {
	return LedgerVoucherNonPostedGetRequestBody{}
}

type LedgerVoucherNonPostedGetRequestBody struct{}

func (r *LedgerVoucherNonPostedGetRequest) RequestBody() *LedgerVoucherNonPostedGetRequestBody {
	return &r.requestBody
}

func (r *LedgerVoucherNonPostedGetRequest) SetRequestBody(body LedgerVoucherNonPostedGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerVoucherNonPostedGetRequest) NewResponseBody() *LedgerVoucherNonPostedGetResponseBody {
	return &LedgerVoucherNonPostedGetResponseBody{}
}

type LedgerVoucherNonPostedGetResponseBody struct {
	FullResultSize int    `json:"fullResultSize"`
	From           int    `json:"from"`
	Count          int    `json:"count"`
	VersionDigest  string `json:"versionDigest"`
	Values         []struct {
		ID          int    `json:"id"`
		Version     int    `json:"version"`
		URL         string `json:"url"`
		Date        string `json:"date"`
		Number      int    `json:"number"`
		Year        int    `json:"year"`
		Description string `json:"description"`
		VoucherType struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"voucherType"`
		ReverseVoucher interface{} `json:"reverseVoucher"`
		Postings       []struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"postings"`
		Document struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"document"`
		Attachment struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"attachment"`
		EdiDocument struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"ediDocument"`
	} `json:"values"`
	TotalNumberOfPostings int `json:"totalNumberOfPostings"`
}

func (r *LedgerVoucherNonPostedGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/ledger/voucher/>nonPosted", r.PathParams())
}

func (r *LedgerVoucherNonPostedGetRequest) Do() (LedgerVoucherNonPostedGetResponseBody, error) {
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

func (r *LedgerVoucherNonPostedGetRequest) All() (LedgerVoucherNonPostedGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := LedgerVoucherNonPostedGetResponseBody{}
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
