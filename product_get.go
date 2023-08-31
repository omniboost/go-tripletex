package tripletex

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-tripletex/utils"
)

func (c *Client) NewProductGetRequest() ProductGetRequest {
	r := ProductGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewProductGetQueryParams()
	r.pathParams = r.NewProductGetPathParams()
	r.requestBody = r.NewProductGetRequestBody()
	return r
}

type ProductGetRequest struct {
	client      *Client
	queryParams *ProductGetQueryParams
	pathParams  *ProductGetPathParams
	method      string
	headers     http.Header
	requestBody ProductGetRequestBody
}

func (r ProductGetRequest) NewProductGetQueryParams() *ProductGetQueryParams {
	return &ProductGetQueryParams{}
}

type ProductGetQueryParams struct {
	Number                        string `schema:"number,omitempty"`
	IDs                           string `schema:"ids,omitempty"`
	Name                          string `schema:"name,omitempty"`
	EAN                           string `schema:"ean,omitempty"`
	IsInactive                    bool   `schema:"isInactive,omitempty"`
	IsStockItem                   bool   `schema:"isStockItem,omitempty"`
	IsSupplierProduct             string `schema:"isSupplierProduct,omitempty"`
	SupplierID                    string `schema:"supplierId,omitempty"`
	CurrencyID                    string `schema:"currencyId,omitempty"`
	VATTypeID                     string `schema:"vatTypeId,omitempty"`
	ProductUnitID                 string `schema:"productUnitId,omitempty"`
	DepartmentID                  string `schema:"departmentId,omitempty"`
	AccountID                     string `schema:"accountId,omitempty"`
	CostExcludingVATCurrencyFrom  string `schema:"costExcludingVatCurrencyFrom,omitempty"`
	CostExcludingVATCurrencyTo    string `schema:"costExcludingVatCurrencyTo,omitempty"`
	PriceExcludingVatCurrencyFrom string `schema:"priceExcludingVatCurrencyFrom,omitempty"`
	PriceExcludingVatCurrencyTo   string `schema:"priceExcludingVatCurrencyTo,omitempty"`
	PriceIncludingVatCurrencyFrom string `schema:"priceIncludingVatCurrencyFrom,omitempty"`
	PriceIncludingVatCurrencyTo   string `schema:"priceIncludingVatCurrencyTo,omitempty"`
	From                          int    `schema:"from,omitempty"`
	Count                         int    `schema:"count,omitempty"`
}

func (p ProductGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ProductGetRequest) QueryParams() *ProductGetQueryParams {
	return r.queryParams
}

func (r ProductGetRequest) NewProductGetPathParams() *ProductGetPathParams {
	return &ProductGetPathParams{}
}

type ProductGetPathParams struct {
}

func (p *ProductGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProductGetRequest) PathParams() *ProductGetPathParams {
	return r.pathParams
}

func (r *ProductGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProductGetRequest) Method() string {
	return r.method
}

func (r ProductGetRequest) NewProductGetRequestBody() ProductGetRequestBody {
	return ProductGetRequestBody{}
}

type ProductGetRequestBody struct{}

func (r *ProductGetRequest) RequestBody() *ProductGetRequestBody {
	return &r.requestBody
}

func (r *ProductGetRequest) SetRequestBody(body ProductGetRequestBody) {
	r.requestBody = body
}

func (r *ProductGetRequest) NewResponseBody() *ProductGetResponseBody {
	return &ProductGetResponseBody{}
}

type ProductGetResponseBody struct {
	FullResultSize int      `json:"fullResultSize"`
	From           int      `json:"from"`
	Count          int      `json:"count"`
	VersionDigest  string   `json:"versionDigest"`
	Values         Products `json:"values"`
}

func (r *ProductGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("/product", r.PathParams())
}

func (r *ProductGetRequest) Do() (ProductGetResponseBody, error) {
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
