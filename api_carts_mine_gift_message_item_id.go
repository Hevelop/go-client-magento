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

type CartsMineGiftMessageItemIdAPI interface {

	/*
		GetV1CartsMineGiftmessageItemId carts/mine/gift-message/{itemId}

		Return the gift message for a specified item in a specified shopping cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param itemId The item ID.
		@return ApiGetV1CartsMineGiftmessageItemIdRequest
	*/
	GetV1CartsMineGiftmessageItemId(ctx context.Context, itemId int32) ApiGetV1CartsMineGiftmessageItemIdRequest

	// GetV1CartsMineGiftmessageItemIdExecute executes the request
	//  @return GiftMessageDataMessageInterface
	GetV1CartsMineGiftmessageItemIdExecute(r ApiGetV1CartsMineGiftmessageItemIdRequest) (*GiftMessageDataMessageInterface, *http.Response, error)

	/*
		PostV1CartsMineGiftmessageItemId carts/mine/gift-message/{itemId}

		Set the gift message for a specified item in a specified shopping cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param itemId The item ID.
		@return ApiPostV1CartsMineGiftmessageItemIdRequest
	*/
	PostV1CartsMineGiftmessageItemId(ctx context.Context, itemId int32) ApiPostV1CartsMineGiftmessageItemIdRequest

	// PostV1CartsMineGiftmessageItemIdExecute executes the request
	//  @return bool
	PostV1CartsMineGiftmessageItemIdExecute(r ApiPostV1CartsMineGiftmessageItemIdRequest) (bool, *http.Response, error)
}

// CartsMineGiftMessageItemIdAPIService CartsMineGiftMessageItemIdAPI service
type CartsMineGiftMessageItemIdAPIService service

type ApiGetV1CartsMineGiftmessageItemIdRequest struct {
	ctx        context.Context
	ApiService CartsMineGiftMessageItemIdAPI
	itemId     int32
}

func (r ApiGetV1CartsMineGiftmessageItemIdRequest) Execute() (*GiftMessageDataMessageInterface, *http.Response, error) {
	return r.ApiService.GetV1CartsMineGiftmessageItemIdExecute(r)
}

/*
GetV1CartsMineGiftmessageItemId carts/mine/gift-message/{itemId}

Return the gift message for a specified item in a specified shopping cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item ID.
	@return ApiGetV1CartsMineGiftmessageItemIdRequest
*/
func (a *CartsMineGiftMessageItemIdAPIService) GetV1CartsMineGiftmessageItemId(ctx context.Context, itemId int32) ApiGetV1CartsMineGiftmessageItemIdRequest {
	return ApiGetV1CartsMineGiftmessageItemIdRequest{
		ApiService: a,
		ctx:        ctx,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return GiftMessageDataMessageInterface
func (a *CartsMineGiftMessageItemIdAPIService) GetV1CartsMineGiftmessageItemIdExecute(r ApiGetV1CartsMineGiftmessageItemIdRequest) (*GiftMessageDataMessageInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftMessageDataMessageInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsMineGiftMessageItemIdAPIService.GetV1CartsMineGiftmessageItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/carts/mine/gift-message/{itemId}"
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

type ApiPostV1CartsMineGiftmessageItemIdRequest struct {
	ctx                               context.Context
	ApiService                        CartsMineGiftMessageItemIdAPI
	itemId                            int32
	postV1CartsMineGiftmessageRequest *PostV1CartsMineGiftmessageRequest
}

func (r ApiPostV1CartsMineGiftmessageItemIdRequest) PostV1CartsMineGiftmessageRequest(postV1CartsMineGiftmessageRequest PostV1CartsMineGiftmessageRequest) ApiPostV1CartsMineGiftmessageItemIdRequest {
	r.postV1CartsMineGiftmessageRequest = &postV1CartsMineGiftmessageRequest
	return r
}

func (r ApiPostV1CartsMineGiftmessageItemIdRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.PostV1CartsMineGiftmessageItemIdExecute(r)
}

/*
PostV1CartsMineGiftmessageItemId carts/mine/gift-message/{itemId}

Set the gift message for a specified item in a specified shopping cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item ID.
	@return ApiPostV1CartsMineGiftmessageItemIdRequest
*/
func (a *CartsMineGiftMessageItemIdAPIService) PostV1CartsMineGiftmessageItemId(ctx context.Context, itemId int32) ApiPostV1CartsMineGiftmessageItemIdRequest {
	return ApiPostV1CartsMineGiftmessageItemIdRequest{
		ApiService: a,
		ctx:        ctx,
		itemId:     itemId,
	}
}

// Execute executes the request
//
//	@return bool
func (a *CartsMineGiftMessageItemIdAPIService) PostV1CartsMineGiftmessageItemIdExecute(r ApiPostV1CartsMineGiftmessageItemIdRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsMineGiftMessageItemIdAPIService.PostV1CartsMineGiftmessageItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/carts/mine/gift-message/{itemId}"
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
