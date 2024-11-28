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

type ProductsSkuLinksAPI interface {

	/*
		PostV1ProductsSkuLinks products/{sku}/links

		Assign a product link to another product

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sku
		@return ApiPostV1ProductsSkuLinksRequest
	*/
	PostV1ProductsSkuLinks(ctx context.Context, sku string) ApiPostV1ProductsSkuLinksRequest

	// PostV1ProductsSkuLinksExecute executes the request
	//  @return bool
	PostV1ProductsSkuLinksExecute(r ApiPostV1ProductsSkuLinksRequest) (bool, *http.Response, error)

	/*
		PutV1ProductsSkuLinks products/{sku}/links

		Save product link

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sku
		@return ApiPutV1ProductsSkuLinksRequest
	*/
	PutV1ProductsSkuLinks(ctx context.Context, sku string) ApiPutV1ProductsSkuLinksRequest

	// PutV1ProductsSkuLinksExecute executes the request
	//  @return bool
	PutV1ProductsSkuLinksExecute(r ApiPutV1ProductsSkuLinksRequest) (bool, *http.Response, error)
}

// ProductsSkuLinksAPIService ProductsSkuLinksAPI service
type ProductsSkuLinksAPIService service

type ApiPostV1ProductsSkuLinksRequest struct {
	ctx                           context.Context
	ApiService                    ProductsSkuLinksAPI
	sku                           string
	postV1ProductsSkuLinksRequest *PostV1ProductsSkuLinksRequest
}

func (r ApiPostV1ProductsSkuLinksRequest) PostV1ProductsSkuLinksRequest(postV1ProductsSkuLinksRequest PostV1ProductsSkuLinksRequest) ApiPostV1ProductsSkuLinksRequest {
	r.postV1ProductsSkuLinksRequest = &postV1ProductsSkuLinksRequest
	return r
}

func (r ApiPostV1ProductsSkuLinksRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.PostV1ProductsSkuLinksExecute(r)
}

/*
PostV1ProductsSkuLinks products/{sku}/links

Assign a product link to another product

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sku
	@return ApiPostV1ProductsSkuLinksRequest
*/
func (a *ProductsSkuLinksAPIService) PostV1ProductsSkuLinks(ctx context.Context, sku string) ApiPostV1ProductsSkuLinksRequest {
	return ApiPostV1ProductsSkuLinksRequest{
		ApiService: a,
		ctx:        ctx,
		sku:        sku,
	}
}

// Execute executes the request
//
//	@return bool
func (a *ProductsSkuLinksAPIService) PostV1ProductsSkuLinksExecute(r ApiPostV1ProductsSkuLinksRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsSkuLinksAPIService.PostV1ProductsSkuLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/products/{sku}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = r.postV1ProductsSkuLinksRequest
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

type ApiPutV1ProductsSkuLinksRequest struct {
	ctx                          context.Context
	ApiService                   ProductsSkuLinksAPI
	sku                          string
	putV1ProductsSkuLinksRequest *PutV1ProductsSkuLinksRequest
}

func (r ApiPutV1ProductsSkuLinksRequest) PutV1ProductsSkuLinksRequest(putV1ProductsSkuLinksRequest PutV1ProductsSkuLinksRequest) ApiPutV1ProductsSkuLinksRequest {
	r.putV1ProductsSkuLinksRequest = &putV1ProductsSkuLinksRequest
	return r
}

func (r ApiPutV1ProductsSkuLinksRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.PutV1ProductsSkuLinksExecute(r)
}

/*
PutV1ProductsSkuLinks products/{sku}/links

Save product link

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sku
	@return ApiPutV1ProductsSkuLinksRequest
*/
func (a *ProductsSkuLinksAPIService) PutV1ProductsSkuLinks(ctx context.Context, sku string) ApiPutV1ProductsSkuLinksRequest {
	return ApiPutV1ProductsSkuLinksRequest{
		ApiService: a,
		ctx:        ctx,
		sku:        sku,
	}
}

// Execute executes the request
//
//	@return bool
func (a *ProductsSkuLinksAPIService) PutV1ProductsSkuLinksExecute(r ApiPutV1ProductsSkuLinksRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsSkuLinksAPIService.PutV1ProductsSkuLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/products/{sku}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = r.putV1ProductsSkuLinksRequest
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