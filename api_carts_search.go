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

type CartsSearchAPI interface {

	/*
		GetV1CartsSearch carts/search

		Enables administrative users to list carts that match specified search criteria. This call returns an array of objects, but detailed information about each object’s attributes might not be included.  See https://developer.adobe.com/commerce/webapi/rest/attributes#CartRepositoryInterface to determine which call to use to get detailed information about all attributes for an object.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetV1CartsSearchRequest
	*/
	GetV1CartsSearch(ctx context.Context) ApiGetV1CartsSearchRequest

	// GetV1CartsSearchExecute executes the request
	//  @return QuoteDataCartSearchResultsInterface
	GetV1CartsSearchExecute(r ApiGetV1CartsSearchRequest) (*QuoteDataCartSearchResultsInterface, *http.Response, error)
}

// CartsSearchAPIService CartsSearchAPI service
type CartsSearchAPIService service

type ApiGetV1CartsSearchRequest struct {
	ctx                                              context.Context
	ApiService                                       CartsSearchAPI
	searchCriteriaFilterGroups0Filters0Field         *string
	searchCriteriaFilterGroups0Filters0Value         *string
	searchCriteriaFilterGroups0Filters0ConditionType *string
	searchCriteriaSortOrders0Field                   *string
	searchCriteriaSortOrders0Direction               *string
	searchCriteriaPageSize                           *int32
	searchCriteriaCurrentPage                        *int32
}

// Field
func (r ApiGetV1CartsSearchRequest) SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field string) ApiGetV1CartsSearchRequest {
	r.searchCriteriaFilterGroups0Filters0Field = &searchCriteriaFilterGroups0Filters0Field
	return r
}

// Value
func (r ApiGetV1CartsSearchRequest) SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value string) ApiGetV1CartsSearchRequest {
	r.searchCriteriaFilterGroups0Filters0Value = &searchCriteriaFilterGroups0Filters0Value
	return r
}

// Condition type
func (r ApiGetV1CartsSearchRequest) SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType string) ApiGetV1CartsSearchRequest {
	r.searchCriteriaFilterGroups0Filters0ConditionType = &searchCriteriaFilterGroups0Filters0ConditionType
	return r
}

// Sorting field.
func (r ApiGetV1CartsSearchRequest) SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field string) ApiGetV1CartsSearchRequest {
	r.searchCriteriaSortOrders0Field = &searchCriteriaSortOrders0Field
	return r
}

// Sorting direction.
func (r ApiGetV1CartsSearchRequest) SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction string) ApiGetV1CartsSearchRequest {
	r.searchCriteriaSortOrders0Direction = &searchCriteriaSortOrders0Direction
	return r
}

// Page size.
func (r ApiGetV1CartsSearchRequest) SearchCriteriaPageSize(searchCriteriaPageSize int32) ApiGetV1CartsSearchRequest {
	r.searchCriteriaPageSize = &searchCriteriaPageSize
	return r
}

// Current page.
func (r ApiGetV1CartsSearchRequest) SearchCriteriaCurrentPage(searchCriteriaCurrentPage int32) ApiGetV1CartsSearchRequest {
	r.searchCriteriaCurrentPage = &searchCriteriaCurrentPage
	return r
}

func (r ApiGetV1CartsSearchRequest) Execute() (*QuoteDataCartSearchResultsInterface, *http.Response, error) {
	return r.ApiService.GetV1CartsSearchExecute(r)
}

/*
GetV1CartsSearch carts/search

Enables administrative users to list carts that match specified search criteria. This call returns an array of objects, but detailed information about each object’s attributes might not be included.  See https://developer.adobe.com/commerce/webapi/rest/attributes#CartRepositoryInterface to determine which call to use to get detailed information about all attributes for an object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetV1CartsSearchRequest
*/
func (a *CartsSearchAPIService) GetV1CartsSearch(ctx context.Context) ApiGetV1CartsSearchRequest {
	return ApiGetV1CartsSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuoteDataCartSearchResultsInterface
func (a *CartsSearchAPIService) GetV1CartsSearchExecute(r ApiGetV1CartsSearchRequest) (*QuoteDataCartSearchResultsInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *QuoteDataCartSearchResultsInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsSearchAPIService.GetV1CartsSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/carts/search"

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
