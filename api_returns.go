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
)

type ReturnsAPI interface {

	/*
		GetV1Returns returns

		Return list of rma data objects based on search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetV1ReturnsRequest
	*/
	GetV1Returns(ctx context.Context) ApiGetV1ReturnsRequest

	// GetV1ReturnsExecute executes the request
	//  @return RmaDataRmaSearchResultInterface
	GetV1ReturnsExecute(r ApiGetV1ReturnsRequest) (*RmaDataRmaSearchResultInterface, *http.Response, error)

	/*
		PostV1Returns returns

		Save RMA

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPostV1ReturnsRequest
	*/
	PostV1Returns(ctx context.Context) ApiPostV1ReturnsRequest

	// PostV1ReturnsExecute executes the request
	//  @return RmaDataRmaInterface
	PostV1ReturnsExecute(r ApiPostV1ReturnsRequest) (*RmaDataRmaInterface, *http.Response, error)
}

// ReturnsAPIService ReturnsAPI service
type ReturnsAPIService service

type ApiGetV1ReturnsRequest struct {
	ctx                                              context.Context
	ApiService                                       ReturnsAPI
	searchCriteriaFilterGroups0Filters0Field         *string
	searchCriteriaFilterGroups0Filters0Value         *string
	searchCriteriaFilterGroups0Filters0ConditionType *string
	searchCriteriaSortOrders0Field                   *string
	searchCriteriaSortOrders0Direction               *string
	searchCriteriaPageSize                           *int32
	searchCriteriaCurrentPage                        *int32
}

// Field
func (r ApiGetV1ReturnsRequest) SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field string) ApiGetV1ReturnsRequest {
	r.searchCriteriaFilterGroups0Filters0Field = &searchCriteriaFilterGroups0Filters0Field
	return r
}

// Value
func (r ApiGetV1ReturnsRequest) SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value string) ApiGetV1ReturnsRequest {
	r.searchCriteriaFilterGroups0Filters0Value = &searchCriteriaFilterGroups0Filters0Value
	return r
}

// Condition type
func (r ApiGetV1ReturnsRequest) SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType string) ApiGetV1ReturnsRequest {
	r.searchCriteriaFilterGroups0Filters0ConditionType = &searchCriteriaFilterGroups0Filters0ConditionType
	return r
}

// Sorting field.
func (r ApiGetV1ReturnsRequest) SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field string) ApiGetV1ReturnsRequest {
	r.searchCriteriaSortOrders0Field = &searchCriteriaSortOrders0Field
	return r
}

// Sorting direction.
func (r ApiGetV1ReturnsRequest) SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction string) ApiGetV1ReturnsRequest {
	r.searchCriteriaSortOrders0Direction = &searchCriteriaSortOrders0Direction
	return r
}

// Page size.
func (r ApiGetV1ReturnsRequest) SearchCriteriaPageSize(searchCriteriaPageSize int32) ApiGetV1ReturnsRequest {
	r.searchCriteriaPageSize = &searchCriteriaPageSize
	return r
}

// Current page.
func (r ApiGetV1ReturnsRequest) SearchCriteriaCurrentPage(searchCriteriaCurrentPage int32) ApiGetV1ReturnsRequest {
	r.searchCriteriaCurrentPage = &searchCriteriaCurrentPage
	return r
}

func (r ApiGetV1ReturnsRequest) Execute() (*RmaDataRmaSearchResultInterface, *http.Response, error) {
	return r.ApiService.GetV1ReturnsExecute(r)
}

/*
GetV1Returns returns

Return list of rma data objects based on search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetV1ReturnsRequest
*/
func (a *ReturnsAPIService) GetV1Returns(ctx context.Context) ApiGetV1ReturnsRequest {
	return ApiGetV1ReturnsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RmaDataRmaSearchResultInterface
func (a *ReturnsAPIService) GetV1ReturnsExecute(r ApiGetV1ReturnsRequest) (*RmaDataRmaSearchResultInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RmaDataRmaSearchResultInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsAPIService.GetV1Returns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/returns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.searchCriteriaFilterGroups0Filters0Field != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchCriteria[filterGroups][0][filters][0][field]", r.searchCriteriaFilterGroups0Filters0Field, "form", "")
	}
	if r.searchCriteriaFilterGroups0Filters0Value != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchCriteria[filterGroups][0][filters][0][value]", r.searchCriteriaFilterGroups0Filters0Value, "form", "")
	}
	if r.searchCriteriaFilterGroups0Filters0ConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchCriteria[filterGroups][0][filters][0][conditionType]", r.searchCriteriaFilterGroups0Filters0ConditionType, "form", "")
	}
	if r.searchCriteriaSortOrders0Field != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchCriteria[sortOrders][0][field]", r.searchCriteriaSortOrders0Field, "form", "")
	}
	if r.searchCriteriaSortOrders0Direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchCriteria[sortOrders][0][direction]", r.searchCriteriaSortOrders0Direction, "form", "")
	}
	if r.searchCriteriaPageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchCriteria[pageSize]", r.searchCriteriaPageSize, "form", "")
	}
	if r.searchCriteriaCurrentPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchCriteria[currentPage]", r.searchCriteriaCurrentPage, "form", "")
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

type ApiPostV1ReturnsRequest struct {
	ctx                  context.Context
	ApiService           ReturnsAPI
	postV1ReturnsRequest *PostV1ReturnsRequest
}

func (r ApiPostV1ReturnsRequest) PostV1ReturnsRequest(postV1ReturnsRequest PostV1ReturnsRequest) ApiPostV1ReturnsRequest {
	r.postV1ReturnsRequest = &postV1ReturnsRequest
	return r
}

func (r ApiPostV1ReturnsRequest) Execute() (*RmaDataRmaInterface, *http.Response, error) {
	return r.ApiService.PostV1ReturnsExecute(r)
}

/*
PostV1Returns returns

Save RMA

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostV1ReturnsRequest
*/
func (a *ReturnsAPIService) PostV1Returns(ctx context.Context) ApiPostV1ReturnsRequest {
	return ApiPostV1ReturnsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RmaDataRmaInterface
func (a *ReturnsAPIService) PostV1ReturnsExecute(r ApiPostV1ReturnsRequest) (*RmaDataRmaInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RmaDataRmaInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsAPIService.PostV1Returns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/returns"

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
	localVarPostBody = r.postV1ReturnsRequest
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
