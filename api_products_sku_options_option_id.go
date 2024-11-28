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

type ProductsSkuOptionsOptionIdAPI interface {

	/*
		DeleteV1ProductsSkuOptionsOptionId products/{sku}/options/{optionId}



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sku
		@param optionId
		@return ApiDeleteV1ProductsSkuOptionsOptionIdRequest
	*/
	DeleteV1ProductsSkuOptionsOptionId(ctx context.Context, sku string, optionId int32) ApiDeleteV1ProductsSkuOptionsOptionIdRequest

	// DeleteV1ProductsSkuOptionsOptionIdExecute executes the request
	//  @return bool
	DeleteV1ProductsSkuOptionsOptionIdExecute(r ApiDeleteV1ProductsSkuOptionsOptionIdRequest) (bool, *http.Response, error)

	/*
		GetV1ProductsSkuOptionsOptionId products/{sku}/options/{optionId}

		Get custom option for a specific product

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sku
		@param optionId
		@return ApiGetV1ProductsSkuOptionsOptionIdRequest
	*/
	GetV1ProductsSkuOptionsOptionId(ctx context.Context, sku string, optionId int32) ApiGetV1ProductsSkuOptionsOptionIdRequest

	// GetV1ProductsSkuOptionsOptionIdExecute executes the request
	//  @return CatalogDataProductCustomOptionInterface
	GetV1ProductsSkuOptionsOptionIdExecute(r ApiGetV1ProductsSkuOptionsOptionIdRequest) (*CatalogDataProductCustomOptionInterface, *http.Response, error)
}

// ProductsSkuOptionsOptionIdAPIService ProductsSkuOptionsOptionIdAPI service
type ProductsSkuOptionsOptionIdAPIService service

type ApiDeleteV1ProductsSkuOptionsOptionIdRequest struct {
	ctx        context.Context
	ApiService ProductsSkuOptionsOptionIdAPI
	sku        string
	optionId   int32
}

func (r ApiDeleteV1ProductsSkuOptionsOptionIdRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.DeleteV1ProductsSkuOptionsOptionIdExecute(r)
}

/*
DeleteV1ProductsSkuOptionsOptionId products/{sku}/options/{optionId}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sku
	@param optionId
	@return ApiDeleteV1ProductsSkuOptionsOptionIdRequest
*/
func (a *ProductsSkuOptionsOptionIdAPIService) DeleteV1ProductsSkuOptionsOptionId(ctx context.Context, sku string, optionId int32) ApiDeleteV1ProductsSkuOptionsOptionIdRequest {
	return ApiDeleteV1ProductsSkuOptionsOptionIdRequest{
		ApiService: a,
		ctx:        ctx,
		sku:        sku,
		optionId:   optionId,
	}
}

// Execute executes the request
//
//	@return bool
func (a *ProductsSkuOptionsOptionIdAPIService) DeleteV1ProductsSkuOptionsOptionIdExecute(r ApiDeleteV1ProductsSkuOptionsOptionIdRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsSkuOptionsOptionIdAPIService.DeleteV1ProductsSkuOptionsOptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/products/{sku}/options/{optionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optionId"+"}", url.PathEscape(parameterValueToString(r.optionId, "optionId")), -1)

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

type ApiGetV1ProductsSkuOptionsOptionIdRequest struct {
	ctx        context.Context
	ApiService ProductsSkuOptionsOptionIdAPI
	sku        string
	optionId   int32
}

func (r ApiGetV1ProductsSkuOptionsOptionIdRequest) Execute() (*CatalogDataProductCustomOptionInterface, *http.Response, error) {
	return r.ApiService.GetV1ProductsSkuOptionsOptionIdExecute(r)
}

/*
GetV1ProductsSkuOptionsOptionId products/{sku}/options/{optionId}

Get custom option for a specific product

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sku
	@param optionId
	@return ApiGetV1ProductsSkuOptionsOptionIdRequest
*/
func (a *ProductsSkuOptionsOptionIdAPIService) GetV1ProductsSkuOptionsOptionId(ctx context.Context, sku string, optionId int32) ApiGetV1ProductsSkuOptionsOptionIdRequest {
	return ApiGetV1ProductsSkuOptionsOptionIdRequest{
		ApiService: a,
		ctx:        ctx,
		sku:        sku,
		optionId:   optionId,
	}
}

// Execute executes the request
//
//	@return CatalogDataProductCustomOptionInterface
func (a *ProductsSkuOptionsOptionIdAPIService) GetV1ProductsSkuOptionsOptionIdExecute(r ApiGetV1ProductsSkuOptionsOptionIdRequest) (*CatalogDataProductCustomOptionInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CatalogDataProductCustomOptionInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsSkuOptionsOptionIdAPIService.GetV1ProductsSkuOptionsOptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/products/{sku}/options/{optionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optionId"+"}", url.PathEscape(parameterValueToString(r.optionId, "optionId")), -1)

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
