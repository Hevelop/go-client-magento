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

type GuestCartsCartIdItemsAPI interface {

	/*
		GetV1GuestcartsCartIdItems guest-carts/{cartId}/items

		List items that are assigned to a specified cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The cart ID.
		@return ApiGetV1GuestcartsCartIdItemsRequest
	*/
	GetV1GuestcartsCartIdItems(ctx context.Context, cartId string) ApiGetV1GuestcartsCartIdItemsRequest

	// GetV1GuestcartsCartIdItemsExecute executes the request
	//  @return []QuoteDataCartItemInterface
	GetV1GuestcartsCartIdItemsExecute(r ApiGetV1GuestcartsCartIdItemsRequest) ([]QuoteDataCartItemInterface, *http.Response, error)

	/*
		PostV1GuestcartsCartIdItems guest-carts/{cartId}/items

		Add/update the specified cart item.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId
		@return ApiPostV1GuestcartsCartIdItemsRequest
	*/
	PostV1GuestcartsCartIdItems(ctx context.Context, cartId string) ApiPostV1GuestcartsCartIdItemsRequest

	// PostV1GuestcartsCartIdItemsExecute executes the request
	//  @return QuoteDataCartItemInterface
	PostV1GuestcartsCartIdItemsExecute(r ApiPostV1GuestcartsCartIdItemsRequest) (*QuoteDataCartItemInterface, *http.Response, error)
}

// GuestCartsCartIdItemsAPIService GuestCartsCartIdItemsAPI service
type GuestCartsCartIdItemsAPIService service

type ApiGetV1GuestcartsCartIdItemsRequest struct {
	ctx        context.Context
	ApiService GuestCartsCartIdItemsAPI
	cartId     string
}

func (r ApiGetV1GuestcartsCartIdItemsRequest) Execute() ([]QuoteDataCartItemInterface, *http.Response, error) {
	return r.ApiService.GetV1GuestcartsCartIdItemsExecute(r)
}

/*
GetV1GuestcartsCartIdItems guest-carts/{cartId}/items

List items that are assigned to a specified cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The cart ID.
	@return ApiGetV1GuestcartsCartIdItemsRequest
*/
func (a *GuestCartsCartIdItemsAPIService) GetV1GuestcartsCartIdItems(ctx context.Context, cartId string) ApiGetV1GuestcartsCartIdItemsRequest {
	return ApiGetV1GuestcartsCartIdItemsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return []QuoteDataCartItemInterface
func (a *GuestCartsCartIdItemsAPIService) GetV1GuestcartsCartIdItemsExecute(r ApiGetV1GuestcartsCartIdItemsRequest) ([]QuoteDataCartItemInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []QuoteDataCartItemInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestCartsCartIdItemsAPIService.GetV1GuestcartsCartIdItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/guest-carts/{cartId}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"cartId"+"}", url.PathEscape(parameterValueToString(r.cartId, "cartId")), -1)

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

type ApiPostV1GuestcartsCartIdItemsRequest struct {
	ctx                         context.Context
	ApiService                  GuestCartsCartIdItemsAPI
	cartId                      string
	postV1CartsMineItemsRequest *PostV1CartsMineItemsRequest
}

func (r ApiPostV1GuestcartsCartIdItemsRequest) PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest PostV1CartsMineItemsRequest) ApiPostV1GuestcartsCartIdItemsRequest {
	r.postV1CartsMineItemsRequest = &postV1CartsMineItemsRequest
	return r
}

func (r ApiPostV1GuestcartsCartIdItemsRequest) Execute() (*QuoteDataCartItemInterface, *http.Response, error) {
	return r.ApiService.PostV1GuestcartsCartIdItemsExecute(r)
}

/*
PostV1GuestcartsCartIdItems guest-carts/{cartId}/items

Add/update the specified cart item.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId
	@return ApiPostV1GuestcartsCartIdItemsRequest
*/
func (a *GuestCartsCartIdItemsAPIService) PostV1GuestcartsCartIdItems(ctx context.Context, cartId string) ApiPostV1GuestcartsCartIdItemsRequest {
	return ApiPostV1GuestcartsCartIdItemsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return QuoteDataCartItemInterface
func (a *GuestCartsCartIdItemsAPIService) PostV1GuestcartsCartIdItemsExecute(r ApiPostV1GuestcartsCartIdItemsRequest) (*QuoteDataCartItemInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *QuoteDataCartItemInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestCartsCartIdItemsAPIService.PostV1GuestcartsCartIdItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/guest-carts/{cartId}/items"
	localVarPath = strings.Replace(localVarPath, "{"+"cartId"+"}", url.PathEscape(parameterValueToString(r.cartId, "cartId")), -1)

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
	localVarPostBody = r.postV1CartsMineItemsRequest
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
