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

type InventoryStockSourceLinksDeleteAPI interface {

	/*
		PostV1InventoryStocksourcelinksdelete inventory/stock-source-links-delete

		Remove StockSourceLink list list

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPostV1InventoryStocksourcelinksdeleteRequest
	*/
	PostV1InventoryStocksourcelinksdelete(ctx context.Context) ApiPostV1InventoryStocksourcelinksdeleteRequest

	// PostV1InventoryStocksourcelinksdeleteExecute executes the request
	//  @return ErrorResponse
	PostV1InventoryStocksourcelinksdeleteExecute(r ApiPostV1InventoryStocksourcelinksdeleteRequest) (*ErrorResponse, *http.Response, error)
}

// InventoryStockSourceLinksDeleteAPIService InventoryStockSourceLinksDeleteAPI service
type InventoryStockSourceLinksDeleteAPIService service

type ApiPostV1InventoryStocksourcelinksdeleteRequest struct {
	ctx                                    context.Context
	ApiService                             InventoryStockSourceLinksDeleteAPI
	postV1InventoryStocksourcelinksRequest *PostV1InventoryStocksourcelinksRequest
}

func (r ApiPostV1InventoryStocksourcelinksdeleteRequest) PostV1InventoryStocksourcelinksRequest(postV1InventoryStocksourcelinksRequest PostV1InventoryStocksourcelinksRequest) ApiPostV1InventoryStocksourcelinksdeleteRequest {
	r.postV1InventoryStocksourcelinksRequest = &postV1InventoryStocksourcelinksRequest
	return r
}

func (r ApiPostV1InventoryStocksourcelinksdeleteRequest) Execute() (*ErrorResponse, *http.Response, error) {
	return r.ApiService.PostV1InventoryStocksourcelinksdeleteExecute(r)
}

/*
PostV1InventoryStocksourcelinksdelete inventory/stock-source-links-delete

Remove StockSourceLink list list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostV1InventoryStocksourcelinksdeleteRequest
*/
func (a *InventoryStockSourceLinksDeleteAPIService) PostV1InventoryStocksourcelinksdelete(ctx context.Context) ApiPostV1InventoryStocksourcelinksdeleteRequest {
	return ApiPostV1InventoryStocksourcelinksdeleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ErrorResponse
func (a *InventoryStockSourceLinksDeleteAPIService) PostV1InventoryStocksourcelinksdeleteExecute(r ApiPostV1InventoryStocksourcelinksdeleteRequest) (*ErrorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockSourceLinksDeleteAPIService.PostV1InventoryStocksourcelinksdelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/inventory/stock-source-links-delete"

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
	localVarPostBody = r.postV1InventoryStocksourcelinksRequest
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
