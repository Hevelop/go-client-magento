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

type InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI interface {

	/*
		GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId inventory/get-sources-assigned-to-stock-ordered-by-priority/{stockId}

		Get Sources assigned to Stock ordered by priority If Stock with given id doesn't exist then return an empty array

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param stockId
		@return ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest
	*/
	GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId(ctx context.Context, stockId int32) ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest

	// GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdExecute executes the request
	//  @return []InventoryApiDataSourceInterface
	GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdExecute(r ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest) ([]InventoryApiDataSourceInterface, *http.Response, error)
}

// InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPIService InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI service
type InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPIService service

type ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest struct {
	ctx        context.Context
	ApiService InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI
	stockId    int32
}

func (r ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest) Execute() ([]InventoryApiDataSourceInterface, *http.Response, error) {
	return r.ApiService.GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdExecute(r)
}

/*
GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId inventory/get-sources-assigned-to-stock-ordered-by-priority/{stockId}

Get Sources assigned to Stock ordered by priority If Stock with given id doesn't exist then return an empty array

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockId
	@return ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest
*/
func (a *InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPIService) GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId(ctx context.Context, stockId int32) ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest {
	return ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest{
		ApiService: a,
		ctx:        ctx,
		stockId:    stockId,
	}
}

// Execute executes the request
//
//	@return []InventoryApiDataSourceInterface
func (a *InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPIService) GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdExecute(r ApiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest) ([]InventoryApiDataSourceInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []InventoryApiDataSourceInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPIService.GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/inventory/get-sources-assigned-to-stock-ordered-by-priority/{stockId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockId"+"}", url.PathEscape(parameterValueToString(r.stockId, "stockId")), -1)

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