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

type InventoryStockResolverTypeCodeAPI interface {

	/*
		GetV1InventoryStockresolverTypeCode inventory/stock-resolver/{type}/{code}

		Resolve Stock by Sales Channel type and code

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param type_
		@param code
		@return ApiGetV1InventoryStockresolverTypeCodeRequest
	*/
	GetV1InventoryStockresolverTypeCode(ctx context.Context, type_ string, code string) ApiGetV1InventoryStockresolverTypeCodeRequest

	// GetV1InventoryStockresolverTypeCodeExecute executes the request
	//  @return InventoryApiDataStockInterface
	GetV1InventoryStockresolverTypeCodeExecute(r ApiGetV1InventoryStockresolverTypeCodeRequest) (*InventoryApiDataStockInterface, *http.Response, error)
}

// InventoryStockResolverTypeCodeAPIService InventoryStockResolverTypeCodeAPI service
type InventoryStockResolverTypeCodeAPIService service

type ApiGetV1InventoryStockresolverTypeCodeRequest struct {
	ctx        context.Context
	ApiService InventoryStockResolverTypeCodeAPI
	type_      string
	code       string
}

func (r ApiGetV1InventoryStockresolverTypeCodeRequest) Execute() (*InventoryApiDataStockInterface, *http.Response, error) {
	return r.ApiService.GetV1InventoryStockresolverTypeCodeExecute(r)
}

/*
GetV1InventoryStockresolverTypeCode inventory/stock-resolver/{type}/{code}

Resolve Stock by Sales Channel type and code

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param type_
	@param code
	@return ApiGetV1InventoryStockresolverTypeCodeRequest
*/
func (a *InventoryStockResolverTypeCodeAPIService) GetV1InventoryStockresolverTypeCode(ctx context.Context, type_ string, code string) ApiGetV1InventoryStockresolverTypeCodeRequest {
	return ApiGetV1InventoryStockresolverTypeCodeRequest{
		ApiService: a,
		ctx:        ctx,
		type_:      type_,
		code:       code,
	}
}

// Execute executes the request
//
//	@return InventoryApiDataStockInterface
func (a *InventoryStockResolverTypeCodeAPIService) GetV1InventoryStockresolverTypeCodeExecute(r ApiGetV1InventoryStockresolverTypeCodeRequest) (*InventoryApiDataStockInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InventoryApiDataStockInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockResolverTypeCodeAPIService.GetV1InventoryStockresolverTypeCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/inventory/stock-resolver/{type}/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", url.PathEscape(parameterValueToString(r.type_, "type_")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
