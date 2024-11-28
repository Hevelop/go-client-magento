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

type EavAttributeSetsAttributeSetIdAPI interface {

	/*
		DeleteV1EavAttributesetsAttributeSetId eav/attribute-sets/{attributeSetId}

		Remove attribute set by given ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param attributeSetId
		@return ApiDeleteV1EavAttributesetsAttributeSetIdRequest
	*/
	DeleteV1EavAttributesetsAttributeSetId(ctx context.Context, attributeSetId int32) ApiDeleteV1EavAttributesetsAttributeSetIdRequest

	// DeleteV1EavAttributesetsAttributeSetIdExecute executes the request
	//  @return bool
	DeleteV1EavAttributesetsAttributeSetIdExecute(r ApiDeleteV1EavAttributesetsAttributeSetIdRequest) (bool, *http.Response, error)

	/*
		GetV1EavAttributesetsAttributeSetId eav/attribute-sets/{attributeSetId}

		Retrieve attribute set information based on given ID

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param attributeSetId
		@return ApiGetV1EavAttributesetsAttributeSetIdRequest
	*/
	GetV1EavAttributesetsAttributeSetId(ctx context.Context, attributeSetId int32) ApiGetV1EavAttributesetsAttributeSetIdRequest

	// GetV1EavAttributesetsAttributeSetIdExecute executes the request
	//  @return EavDataAttributeSetInterface
	GetV1EavAttributesetsAttributeSetIdExecute(r ApiGetV1EavAttributesetsAttributeSetIdRequest) (*EavDataAttributeSetInterface, *http.Response, error)

	/*
		PutV1EavAttributesetsAttributeSetId eav/attribute-sets/{attributeSetId}

		Save attribute set data

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param attributeSetId
		@return ApiPutV1EavAttributesetsAttributeSetIdRequest
	*/
	PutV1EavAttributesetsAttributeSetId(ctx context.Context, attributeSetId string) ApiPutV1EavAttributesetsAttributeSetIdRequest

	// PutV1EavAttributesetsAttributeSetIdExecute executes the request
	//  @return EavDataAttributeSetInterface
	PutV1EavAttributesetsAttributeSetIdExecute(r ApiPutV1EavAttributesetsAttributeSetIdRequest) (*EavDataAttributeSetInterface, *http.Response, error)
}

// EavAttributeSetsAttributeSetIdAPIService EavAttributeSetsAttributeSetIdAPI service
type EavAttributeSetsAttributeSetIdAPIService service

type ApiDeleteV1EavAttributesetsAttributeSetIdRequest struct {
	ctx            context.Context
	ApiService     EavAttributeSetsAttributeSetIdAPI
	attributeSetId int32
}

func (r ApiDeleteV1EavAttributesetsAttributeSetIdRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.DeleteV1EavAttributesetsAttributeSetIdExecute(r)
}

/*
DeleteV1EavAttributesetsAttributeSetId eav/attribute-sets/{attributeSetId}

Remove attribute set by given ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param attributeSetId
	@return ApiDeleteV1EavAttributesetsAttributeSetIdRequest
*/
func (a *EavAttributeSetsAttributeSetIdAPIService) DeleteV1EavAttributesetsAttributeSetId(ctx context.Context, attributeSetId int32) ApiDeleteV1EavAttributesetsAttributeSetIdRequest {
	return ApiDeleteV1EavAttributesetsAttributeSetIdRequest{
		ApiService:     a,
		ctx:            ctx,
		attributeSetId: attributeSetId,
	}
}

// Execute executes the request
//
//	@return bool
func (a *EavAttributeSetsAttributeSetIdAPIService) DeleteV1EavAttributesetsAttributeSetIdExecute(r ApiDeleteV1EavAttributesetsAttributeSetIdRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EavAttributeSetsAttributeSetIdAPIService.DeleteV1EavAttributesetsAttributeSetId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/eav/attribute-sets/{attributeSetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"attributeSetId"+"}", url.PathEscape(parameterValueToString(r.attributeSetId, "attributeSetId")), -1)

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

type ApiGetV1EavAttributesetsAttributeSetIdRequest struct {
	ctx            context.Context
	ApiService     EavAttributeSetsAttributeSetIdAPI
	attributeSetId int32
}

func (r ApiGetV1EavAttributesetsAttributeSetIdRequest) Execute() (*EavDataAttributeSetInterface, *http.Response, error) {
	return r.ApiService.GetV1EavAttributesetsAttributeSetIdExecute(r)
}

/*
GetV1EavAttributesetsAttributeSetId eav/attribute-sets/{attributeSetId}

Retrieve attribute set information based on given ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param attributeSetId
	@return ApiGetV1EavAttributesetsAttributeSetIdRequest
*/
func (a *EavAttributeSetsAttributeSetIdAPIService) GetV1EavAttributesetsAttributeSetId(ctx context.Context, attributeSetId int32) ApiGetV1EavAttributesetsAttributeSetIdRequest {
	return ApiGetV1EavAttributesetsAttributeSetIdRequest{
		ApiService:     a,
		ctx:            ctx,
		attributeSetId: attributeSetId,
	}
}

// Execute executes the request
//
//	@return EavDataAttributeSetInterface
func (a *EavAttributeSetsAttributeSetIdAPIService) GetV1EavAttributesetsAttributeSetIdExecute(r ApiGetV1EavAttributesetsAttributeSetIdRequest) (*EavDataAttributeSetInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EavDataAttributeSetInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EavAttributeSetsAttributeSetIdAPIService.GetV1EavAttributesetsAttributeSetId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/eav/attribute-sets/{attributeSetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"attributeSetId"+"}", url.PathEscape(parameterValueToString(r.attributeSetId, "attributeSetId")), -1)

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

type ApiPutV1EavAttributesetsAttributeSetIdRequest struct {
	ctx                                        context.Context
	ApiService                                 EavAttributeSetsAttributeSetIdAPI
	attributeSetId                             string
	putV1EavAttributesetsAttributeSetIdRequest *PutV1EavAttributesetsAttributeSetIdRequest
}

func (r ApiPutV1EavAttributesetsAttributeSetIdRequest) PutV1EavAttributesetsAttributeSetIdRequest(putV1EavAttributesetsAttributeSetIdRequest PutV1EavAttributesetsAttributeSetIdRequest) ApiPutV1EavAttributesetsAttributeSetIdRequest {
	r.putV1EavAttributesetsAttributeSetIdRequest = &putV1EavAttributesetsAttributeSetIdRequest
	return r
}

func (r ApiPutV1EavAttributesetsAttributeSetIdRequest) Execute() (*EavDataAttributeSetInterface, *http.Response, error) {
	return r.ApiService.PutV1EavAttributesetsAttributeSetIdExecute(r)
}

/*
PutV1EavAttributesetsAttributeSetId eav/attribute-sets/{attributeSetId}

Save attribute set data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param attributeSetId
	@return ApiPutV1EavAttributesetsAttributeSetIdRequest
*/
func (a *EavAttributeSetsAttributeSetIdAPIService) PutV1EavAttributesetsAttributeSetId(ctx context.Context, attributeSetId string) ApiPutV1EavAttributesetsAttributeSetIdRequest {
	return ApiPutV1EavAttributesetsAttributeSetIdRequest{
		ApiService:     a,
		ctx:            ctx,
		attributeSetId: attributeSetId,
	}
}

// Execute executes the request
//
//	@return EavDataAttributeSetInterface
func (a *EavAttributeSetsAttributeSetIdAPIService) PutV1EavAttributesetsAttributeSetIdExecute(r ApiPutV1EavAttributesetsAttributeSetIdRequest) (*EavDataAttributeSetInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EavDataAttributeSetInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EavAttributeSetsAttributeSetIdAPIService.PutV1EavAttributesetsAttributeSetId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/eav/attribute-sets/{attributeSetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"attributeSetId"+"}", url.PathEscape(parameterValueToString(r.attributeSetId, "attributeSetId")), -1)

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
	localVarPostBody = r.putV1EavAttributesetsAttributeSetIdRequest
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
