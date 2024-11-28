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

type SalesRulesRuleIdAPI interface {

	/*
		DeleteV1SalesRulesRuleId salesRules/{ruleId}

		Delete rule by ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ruleId
		@return ApiDeleteV1SalesRulesRuleIdRequest
	*/
	DeleteV1SalesRulesRuleId(ctx context.Context, ruleId int32) ApiDeleteV1SalesRulesRuleIdRequest

	// DeleteV1SalesRulesRuleIdExecute executes the request
	//  @return bool
	DeleteV1SalesRulesRuleIdExecute(r ApiDeleteV1SalesRulesRuleIdRequest) (bool, *http.Response, error)

	/*
		GetV1SalesRulesRuleId salesRules/{ruleId}

		Get rule by ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ruleId
		@return ApiGetV1SalesRulesRuleIdRequest
	*/
	GetV1SalesRulesRuleId(ctx context.Context, ruleId int32) ApiGetV1SalesRulesRuleIdRequest

	// GetV1SalesRulesRuleIdExecute executes the request
	//  @return SalesRuleDataRuleInterface
	GetV1SalesRulesRuleIdExecute(r ApiGetV1SalesRulesRuleIdRequest) (*SalesRuleDataRuleInterface, *http.Response, error)

	/*
		PutV1SalesRulesRuleId salesRules/{ruleId}

		Save sales rule.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ruleId
		@return ApiPutV1SalesRulesRuleIdRequest
	*/
	PutV1SalesRulesRuleId(ctx context.Context, ruleId string) ApiPutV1SalesRulesRuleIdRequest

	// PutV1SalesRulesRuleIdExecute executes the request
	//  @return SalesRuleDataRuleInterface
	PutV1SalesRulesRuleIdExecute(r ApiPutV1SalesRulesRuleIdRequest) (*SalesRuleDataRuleInterface, *http.Response, error)
}

// SalesRulesRuleIdAPIService SalesRulesRuleIdAPI service
type SalesRulesRuleIdAPIService service

type ApiDeleteV1SalesRulesRuleIdRequest struct {
	ctx        context.Context
	ApiService SalesRulesRuleIdAPI
	ruleId     int32
}

func (r ApiDeleteV1SalesRulesRuleIdRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.DeleteV1SalesRulesRuleIdExecute(r)
}

/*
DeleteV1SalesRulesRuleId salesRules/{ruleId}

Delete rule by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ruleId
	@return ApiDeleteV1SalesRulesRuleIdRequest
*/
func (a *SalesRulesRuleIdAPIService) DeleteV1SalesRulesRuleId(ctx context.Context, ruleId int32) ApiDeleteV1SalesRulesRuleIdRequest {
	return ApiDeleteV1SalesRulesRuleIdRequest{
		ApiService: a,
		ctx:        ctx,
		ruleId:     ruleId,
	}
}

// Execute executes the request
//
//	@return bool
func (a *SalesRulesRuleIdAPIService) DeleteV1SalesRulesRuleIdExecute(r ApiDeleteV1SalesRulesRuleIdRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesRulesRuleIdAPIService.DeleteV1SalesRulesRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/salesRules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiGetV1SalesRulesRuleIdRequest struct {
	ctx        context.Context
	ApiService SalesRulesRuleIdAPI
	ruleId     int32
}

func (r ApiGetV1SalesRulesRuleIdRequest) Execute() (*SalesRuleDataRuleInterface, *http.Response, error) {
	return r.ApiService.GetV1SalesRulesRuleIdExecute(r)
}

/*
GetV1SalesRulesRuleId salesRules/{ruleId}

Get rule by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ruleId
	@return ApiGetV1SalesRulesRuleIdRequest
*/
func (a *SalesRulesRuleIdAPIService) GetV1SalesRulesRuleId(ctx context.Context, ruleId int32) ApiGetV1SalesRulesRuleIdRequest {
	return ApiGetV1SalesRulesRuleIdRequest{
		ApiService: a,
		ctx:        ctx,
		ruleId:     ruleId,
	}
}

// Execute executes the request
//
//	@return SalesRuleDataRuleInterface
func (a *SalesRulesRuleIdAPIService) GetV1SalesRulesRuleIdExecute(r ApiGetV1SalesRulesRuleIdRequest) (*SalesRuleDataRuleInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SalesRuleDataRuleInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesRulesRuleIdAPIService.GetV1SalesRulesRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/salesRules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiPutV1SalesRulesRuleIdRequest struct {
	ctx                     context.Context
	ApiService              SalesRulesRuleIdAPI
	ruleId                  string
	postV1SalesRulesRequest *PostV1SalesRulesRequest
}

func (r ApiPutV1SalesRulesRuleIdRequest) PostV1SalesRulesRequest(postV1SalesRulesRequest PostV1SalesRulesRequest) ApiPutV1SalesRulesRuleIdRequest {
	r.postV1SalesRulesRequest = &postV1SalesRulesRequest
	return r
}

func (r ApiPutV1SalesRulesRuleIdRequest) Execute() (*SalesRuleDataRuleInterface, *http.Response, error) {
	return r.ApiService.PutV1SalesRulesRuleIdExecute(r)
}

/*
PutV1SalesRulesRuleId salesRules/{ruleId}

Save sales rule.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ruleId
	@return ApiPutV1SalesRulesRuleIdRequest
*/
func (a *SalesRulesRuleIdAPIService) PutV1SalesRulesRuleId(ctx context.Context, ruleId string) ApiPutV1SalesRulesRuleIdRequest {
	return ApiPutV1SalesRulesRuleIdRequest{
		ApiService: a,
		ctx:        ctx,
		ruleId:     ruleId,
	}
}

// Execute executes the request
//
//	@return SalesRuleDataRuleInterface
func (a *SalesRulesRuleIdAPIService) PutV1SalesRulesRuleIdExecute(r ApiPutV1SalesRulesRuleIdRequest) (*SalesRuleDataRuleInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SalesRuleDataRuleInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SalesRulesRuleIdAPIService.PutV1SalesRulesRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/salesRules/{ruleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ruleId"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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
	localVarPostBody = r.postV1SalesRulesRequest
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
		if localVarHTTPResponse.StatusCode == 500 {
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