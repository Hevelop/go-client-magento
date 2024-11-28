/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type BundleProductsProductSkuChildrenAPI interface {

	/*
		GetV1BundleproductsProductSkuChildren bundle-products/{productSku}/children

		Get all children for Bundle product

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param productSku
		@return ApiGetV1BundleproductsProductSkuChildrenRequest
	*/
	GetV1BundleproductsProductSkuChildren(ctx context.Context, productSku string) ApiGetV1BundleproductsProductSkuChildrenRequest

	// GetV1BundleproductsProductSkuChildrenExecute executes the request
	//  @return []BundleDataLinkInterface
	GetV1BundleproductsProductSkuChildrenExecute(r ApiGetV1BundleproductsProductSkuChildrenRequest) ([]BundleDataLinkInterface, *http.Response, error)
}

// BundleProductsProductSkuChildrenAPIService BundleProductsProductSkuChildrenAPI service
type BundleProductsProductSkuChildrenAPIService service

type ApiGetV1BundleproductsProductSkuChildrenRequest struct {
	ctx        context.Context
	ApiService BundleProductsProductSkuChildrenAPI
	productSku string
	optionId   *int32
}

func (r ApiGetV1BundleproductsProductSkuChildrenRequest) OptionId(optionId int32) ApiGetV1BundleproductsProductSkuChildrenRequest {
	r.optionId = &optionId
	return r
}

func (r ApiGetV1BundleproductsProductSkuChildrenRequest) Execute() ([]BundleDataLinkInterface, *http.Response, error) {
	return r.ApiService.GetV1BundleproductsProductSkuChildrenExecute(r)
}

/*
GetV1BundleproductsProductSkuChildren bundle-products/{productSku}/children

Get all children for Bundle product

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param productSku
	@return ApiGetV1BundleproductsProductSkuChildrenRequest
*/
func (a *BundleProductsProductSkuChildrenAPIService) GetV1BundleproductsProductSkuChildren(ctx context.Context, productSku string) ApiGetV1BundleproductsProductSkuChildrenRequest {
	return ApiGetV1BundleproductsProductSkuChildrenRequest{
		ApiService: a,
		ctx:        ctx,
		productSku: productSku,
	}
}

// Execute executes the request
//
//	@return []BundleDataLinkInterface
func (a *BundleProductsProductSkuChildrenAPIService) GetV1BundleproductsProductSkuChildrenExecute(r ApiGetV1BundleproductsProductSkuChildrenRequest) ([]BundleDataLinkInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BundleDataLinkInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundleProductsProductSkuChildrenAPIService.GetV1BundleproductsProductSkuChildren")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/bundle-products/{productSku}/children"
	localVarPath = strings.Replace(localVarPath, "{"+"productSku"+"}", url.PathEscape(parameterValueToString(r.productSku, "productSku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.optionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "optionId", r.optionId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
