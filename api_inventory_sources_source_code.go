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

type InventorySourcesSourceCodeAPI interface {

	/*
		GetV1InventorySourcesSourceCode inventory/sources/{sourceCode}

		Get Source data by given code. If you want to create plugin on get method, also you need to create separate plugin on getList method, because entity loading way is different for these methods

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sourceCode
		@return ApiGetV1InventorySourcesSourceCodeRequest
	*/
	GetV1InventorySourcesSourceCode(ctx context.Context, sourceCode string) ApiGetV1InventorySourcesSourceCodeRequest

	// GetV1InventorySourcesSourceCodeExecute executes the request
	//  @return InventoryApiDataSourceInterface
	GetV1InventorySourcesSourceCodeExecute(r ApiGetV1InventorySourcesSourceCodeRequest) (*InventoryApiDataSourceInterface, *http.Response, error)

	/*
		PutV1InventorySourcesSourceCode inventory/sources/{sourceCode}

		Save Source data

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sourceCode
		@return ApiPutV1InventorySourcesSourceCodeRequest
	*/
	PutV1InventorySourcesSourceCode(ctx context.Context, sourceCode string) ApiPutV1InventorySourcesSourceCodeRequest

	// PutV1InventorySourcesSourceCodeExecute executes the request
	//  @return ErrorResponse
	PutV1InventorySourcesSourceCodeExecute(r ApiPutV1InventorySourcesSourceCodeRequest) (*ErrorResponse, *http.Response, error)
}

// InventorySourcesSourceCodeAPIService InventorySourcesSourceCodeAPI service
type InventorySourcesSourceCodeAPIService service

type ApiGetV1InventorySourcesSourceCodeRequest struct {
	ctx        context.Context
	ApiService InventorySourcesSourceCodeAPI
	sourceCode string
}

func (r ApiGetV1InventorySourcesSourceCodeRequest) Execute() (*InventoryApiDataSourceInterface, *http.Response, error) {
	return r.ApiService.GetV1InventorySourcesSourceCodeExecute(r)
}

/*
GetV1InventorySourcesSourceCode inventory/sources/{sourceCode}

Get Source data by given code. If you want to create plugin on get method, also you need to create separate plugin on getList method, because entity loading way is different for these methods

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sourceCode
	@return ApiGetV1InventorySourcesSourceCodeRequest
*/
func (a *InventorySourcesSourceCodeAPIService) GetV1InventorySourcesSourceCode(ctx context.Context, sourceCode string) ApiGetV1InventorySourcesSourceCodeRequest {
	return ApiGetV1InventorySourcesSourceCodeRequest{
		ApiService: a,
		ctx:        ctx,
		sourceCode: sourceCode,
	}
}

// Execute executes the request
//
//	@return InventoryApiDataSourceInterface
func (a *InventorySourcesSourceCodeAPIService) GetV1InventorySourcesSourceCodeExecute(r ApiGetV1InventorySourcesSourceCodeRequest) (*InventoryApiDataSourceInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InventoryApiDataSourceInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventorySourcesSourceCodeAPIService.GetV1InventorySourcesSourceCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/inventory/sources/{sourceCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"sourceCode"+"}", url.PathEscape(parameterValueToString(r.sourceCode, "sourceCode")), -1)

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

type ApiPutV1InventorySourcesSourceCodeRequest struct {
	ctx                           context.Context
	ApiService                    InventorySourcesSourceCodeAPI
	sourceCode                    string
	postV1InventorySourcesRequest *PostV1InventorySourcesRequest
}

func (r ApiPutV1InventorySourcesSourceCodeRequest) PostV1InventorySourcesRequest(postV1InventorySourcesRequest PostV1InventorySourcesRequest) ApiPutV1InventorySourcesSourceCodeRequest {
	r.postV1InventorySourcesRequest = &postV1InventorySourcesRequest
	return r
}

func (r ApiPutV1InventorySourcesSourceCodeRequest) Execute() (*ErrorResponse, *http.Response, error) {
	return r.ApiService.PutV1InventorySourcesSourceCodeExecute(r)
}

/*
PutV1InventorySourcesSourceCode inventory/sources/{sourceCode}

Save Source data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sourceCode
	@return ApiPutV1InventorySourcesSourceCodeRequest
*/
func (a *InventorySourcesSourceCodeAPIService) PutV1InventorySourcesSourceCode(ctx context.Context, sourceCode string) ApiPutV1InventorySourcesSourceCodeRequest {
	return ApiPutV1InventorySourcesSourceCodeRequest{
		ApiService: a,
		ctx:        ctx,
		sourceCode: sourceCode,
	}
}

// Execute executes the request
//
//	@return ErrorResponse
func (a *InventorySourcesSourceCodeAPIService) PutV1InventorySourcesSourceCodeExecute(r ApiPutV1InventorySourcesSourceCodeRequest) (*ErrorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventorySourcesSourceCodeAPIService.PutV1InventorySourcesSourceCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/inventory/sources/{sourceCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"sourceCode"+"}", url.PathEscape(parameterValueToString(r.sourceCode, "sourceCode")), -1)

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
	localVarPostBody = r.postV1InventorySourcesRequest
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
