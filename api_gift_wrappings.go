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

type GiftWrappingsAPI interface {

	/*
		GetV1Giftwrappings gift-wrappings

		Return list of gift wrapping data objects based on search criteria

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetV1GiftwrappingsRequest
	*/
	GetV1Giftwrappings(ctx context.Context) ApiGetV1GiftwrappingsRequest

	// GetV1GiftwrappingsExecute executes the request
	//  @return GiftWrappingDataWrappingSearchResultsInterface
	GetV1GiftwrappingsExecute(r ApiGetV1GiftwrappingsRequest) (*GiftWrappingDataWrappingSearchResultsInterface, *http.Response, error)

	/*
		PostV1Giftwrappings gift-wrappings

		Create/Update new gift wrapping with data object values

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPostV1GiftwrappingsRequest
	*/
	PostV1Giftwrappings(ctx context.Context) ApiPostV1GiftwrappingsRequest

	// PostV1GiftwrappingsExecute executes the request
	//  @return GiftWrappingDataWrappingInterface
	PostV1GiftwrappingsExecute(r ApiPostV1GiftwrappingsRequest) (*GiftWrappingDataWrappingInterface, *http.Response, error)
}

// GiftWrappingsAPIService GiftWrappingsAPI service
type GiftWrappingsAPIService service

type ApiGetV1GiftwrappingsRequest struct {
	ctx                                              context.Context
	ApiService                                       GiftWrappingsAPI
	searchCriteriaFilterGroups0Filters0Field         *string
	searchCriteriaFilterGroups0Filters0Value         *string
	searchCriteriaFilterGroups0Filters0ConditionType *string
	searchCriteriaSortOrders0Field                   *string
	searchCriteriaSortOrders0Direction               *string
	searchCriteriaPageSize                           *int32
	searchCriteriaCurrentPage                        *int32
}

// Field
func (r ApiGetV1GiftwrappingsRequest) SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field string) ApiGetV1GiftwrappingsRequest {
	r.searchCriteriaFilterGroups0Filters0Field = &searchCriteriaFilterGroups0Filters0Field
	return r
}

// Value
func (r ApiGetV1GiftwrappingsRequest) SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value string) ApiGetV1GiftwrappingsRequest {
	r.searchCriteriaFilterGroups0Filters0Value = &searchCriteriaFilterGroups0Filters0Value
	return r
}

// Condition type
func (r ApiGetV1GiftwrappingsRequest) SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType string) ApiGetV1GiftwrappingsRequest {
	r.searchCriteriaFilterGroups0Filters0ConditionType = &searchCriteriaFilterGroups0Filters0ConditionType
	return r
}

// Sorting field.
func (r ApiGetV1GiftwrappingsRequest) SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field string) ApiGetV1GiftwrappingsRequest {
	r.searchCriteriaSortOrders0Field = &searchCriteriaSortOrders0Field
	return r
}

// Sorting direction.
func (r ApiGetV1GiftwrappingsRequest) SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction string) ApiGetV1GiftwrappingsRequest {
	r.searchCriteriaSortOrders0Direction = &searchCriteriaSortOrders0Direction
	return r
}

// Page size.
func (r ApiGetV1GiftwrappingsRequest) SearchCriteriaPageSize(searchCriteriaPageSize int32) ApiGetV1GiftwrappingsRequest {
	r.searchCriteriaPageSize = &searchCriteriaPageSize
	return r
}

// Current page.
func (r ApiGetV1GiftwrappingsRequest) SearchCriteriaCurrentPage(searchCriteriaCurrentPage int32) ApiGetV1GiftwrappingsRequest {
	r.searchCriteriaCurrentPage = &searchCriteriaCurrentPage
	return r
}

func (r ApiGetV1GiftwrappingsRequest) Execute() (*GiftWrappingDataWrappingSearchResultsInterface, *http.Response, error) {
	return r.ApiService.GetV1GiftwrappingsExecute(r)
}

/*
GetV1Giftwrappings gift-wrappings

Return list of gift wrapping data objects based on search criteria

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetV1GiftwrappingsRequest
*/
func (a *GiftWrappingsAPIService) GetV1Giftwrappings(ctx context.Context) ApiGetV1GiftwrappingsRequest {
	return ApiGetV1GiftwrappingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GiftWrappingDataWrappingSearchResultsInterface
func (a *GiftWrappingsAPIService) GetV1GiftwrappingsExecute(r ApiGetV1GiftwrappingsRequest) (*GiftWrappingDataWrappingSearchResultsInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftWrappingDataWrappingSearchResultsInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftWrappingsAPIService.GetV1Giftwrappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/gift-wrappings"

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

type ApiPostV1GiftwrappingsRequest struct {
	ctx                        context.Context
	ApiService                 GiftWrappingsAPI
	postV1GiftwrappingsRequest *PostV1GiftwrappingsRequest
}

func (r ApiPostV1GiftwrappingsRequest) PostV1GiftwrappingsRequest(postV1GiftwrappingsRequest PostV1GiftwrappingsRequest) ApiPostV1GiftwrappingsRequest {
	r.postV1GiftwrappingsRequest = &postV1GiftwrappingsRequest
	return r
}

func (r ApiPostV1GiftwrappingsRequest) Execute() (*GiftWrappingDataWrappingInterface, *http.Response, error) {
	return r.ApiService.PostV1GiftwrappingsExecute(r)
}

/*
PostV1Giftwrappings gift-wrappings

Create/Update new gift wrapping with data object values

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostV1GiftwrappingsRequest
*/
func (a *GiftWrappingsAPIService) PostV1Giftwrappings(ctx context.Context) ApiPostV1GiftwrappingsRequest {
	return ApiPostV1GiftwrappingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GiftWrappingDataWrappingInterface
func (a *GiftWrappingsAPIService) PostV1GiftwrappingsExecute(r ApiPostV1GiftwrappingsRequest) (*GiftWrappingDataWrappingInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftWrappingDataWrappingInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftWrappingsAPIService.PostV1Giftwrappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/gift-wrappings"

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
	localVarPostBody = r.postV1GiftwrappingsRequest
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
