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

type CartsCartIdGiftMessageItemIdAPI interface {

	/*
		GetV1CartsCartIdGiftmessageItemId carts/{cartId}/gift-message/{itemId}

		Return the gift message for a specified item in a specified shopping cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The shopping cart ID.
		@param itemId The item ID.
		@return ApiGetV1CartsCartIdGiftmessageItemIdRequest
	*/
	GetV1CartsCartIdGiftmessageItemId(ctx context.Context, cartId int32, itemId int32) ApiGetV1CartsCartIdGiftmessageItemIdRequest

	// GetV1CartsCartIdGiftmessageItemIdExecute executes the request
	//  @return GiftMessageDataMessageInterface
	GetV1CartsCartIdGiftmessageItemIdExecute(r ApiGetV1CartsCartIdGiftmessageItemIdRequest) (*GiftMessageDataMessageInterface, *http.Response, error)

	/*
		PostV1CartsCartIdGiftmessageItemId carts/{cartId}/gift-message/{itemId}

		Set the gift message for a specified item in a specified shopping cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The cart ID.
		@param itemId The item ID.
		@return ApiPostV1CartsCartIdGiftmessageItemIdRequest
	*/
	PostV1CartsCartIdGiftmessageItemId(ctx context.Context, cartId int32, itemId int32) ApiPostV1CartsCartIdGiftmessageItemIdRequest

	// PostV1CartsCartIdGiftmessageItemIdExecute executes the request
	//  @return bool
	PostV1CartsCartIdGiftmessageItemIdExecute(r ApiPostV1CartsCartIdGiftmessageItemIdRequest) (bool, *http.Response, error)
}

// CartsCartIdGiftMessageItemIdAPIService CartsCartIdGiftMessageItemIdAPI service
type CartsCartIdGiftMessageItemIdAPIService service

type ApiGetV1CartsCartIdGiftmessageItemIdRequest struct {
	ctx        context.Context
	ApiService CartsCartIdGiftMessageItemIdAPI
	cartId     int32
	itemId     int32
}

func (r ApiGetV1CartsCartIdGiftmessageItemIdRequest) Execute() (*GiftMessageDataMessageInterface, *http.Response, error) {
	return r.ApiService.GetV1CartsCartIdGiftmessageItemIdExecute(r)
}

/*
GetV1CartsCartIdGiftmessageItemId carts/{cartId}/gift-message/{itemId}

Return the gift message for a specified item in a specified shopping cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The shopping cart ID.
	@param itemId The item ID.
	@return ApiGetV1CartsCartIdGiftmessageItemIdRequest
*/
func (a *CartsCartIdGiftMessageItemIdAPIService) GetV1CartsCartIdGiftmessageItemId(ctx context.Context, cartId int32, itemId int32) ApiGetV1CartsCartIdGiftmessageItemIdRequest {
	return ApiGetV1CartsCartIdGiftmessageItemIdRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return GiftMessageDataMessageInterface
func (a *CartsCartIdGiftMessageItemIdAPIService) GetV1CartsCartIdGiftmessageItemIdExecute(r ApiGetV1CartsCartIdGiftmessageItemIdRequest) (*GiftMessageDataMessageInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftMessageDataMessageInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdGiftMessageItemIdAPIService.GetV1CartsCartIdGiftmessageItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/carts/{cartId}/gift-message/{itemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"cartId"+"}", url.PathEscape(parameterValueToString(r.cartId, "cartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

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

type ApiPostV1CartsCartIdGiftmessageItemIdRequest struct {
	ctx                               context.Context
	ApiService                        CartsCartIdGiftMessageItemIdAPI
	cartId                            int32
	itemId                            int32
	postV1CartsMineGiftmessageRequest *PostV1CartsMineGiftmessageRequest
}

func (r ApiPostV1CartsCartIdGiftmessageItemIdRequest) PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest PostV1CartsMineGiftmessageRequest) ApiPostV1CartsCartIdGiftmessageItemIdRequest {
	r.postV1CartsMineGiftmessageRequest = &postV1CartsMineGiftmessageRequest
	return r
}

func (r ApiPostV1CartsCartIdGiftmessageItemIdRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.PostV1CartsCartIdGiftmessageItemIdExecute(r)
}

/*
PostV1CartsCartIdGiftmessageItemId carts/{cartId}/gift-message/{itemId}

Set the gift message for a specified item in a specified shopping cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The cart ID.
	@param itemId The item ID.
	@return ApiPostV1CartsCartIdGiftmessageItemIdRequest
*/
func (a *CartsCartIdGiftMessageItemIdAPIService) PostV1CartsCartIdGiftmessageItemId(ctx context.Context, cartId int32, itemId int32) ApiPostV1CartsCartIdGiftmessageItemIdRequest {
	return ApiPostV1CartsCartIdGiftmessageItemIdRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return bool
func (a *CartsCartIdGiftMessageItemIdAPIService) PostV1CartsCartIdGiftmessageItemIdExecute(r ApiPostV1CartsCartIdGiftmessageItemIdRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdGiftMessageItemIdAPIService.PostV1CartsCartIdGiftmessageItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/carts/{cartId}/gift-message/{itemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"cartId"+"}", url.PathEscape(parameterValueToString(r.cartId, "cartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

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
	localVarPostBody = r.postV1CartsMineGiftmessageRequest
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
