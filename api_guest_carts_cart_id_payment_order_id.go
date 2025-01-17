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

type GuestCartsCartIdPaymentOrderIdAPI interface {

	/*
		GetV1GuestcartsCartIdPaymentorderId guest-carts/{cartId}/payment-order/{id}

		Get payment order for guest customer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId
		@param id
		@return ApiGetV1GuestcartsCartIdPaymentorderIdRequest
	*/
	GetV1GuestcartsCartIdPaymentorderId(ctx context.Context, cartId string, id string) ApiGetV1GuestcartsCartIdPaymentorderIdRequest

	// GetV1GuestcartsCartIdPaymentorderIdExecute executes the request
	//  @return PaymentServicesPaypalDataPaymentOrderDetailsInterface
	GetV1GuestcartsCartIdPaymentorderIdExecute(r ApiGetV1GuestcartsCartIdPaymentorderIdRequest) (*PaymentServicesPaypalDataPaymentOrderDetailsInterface, *http.Response, error)

	/*
		PostV1GuestcartsCartIdPaymentorderId guest-carts/{cartId}/payment-order/{id}

		Sync payment order for guest customer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId
		@param id
		@return ApiPostV1GuestcartsCartIdPaymentorderIdRequest
	*/
	PostV1GuestcartsCartIdPaymentorderId(ctx context.Context, cartId string, id string) ApiPostV1GuestcartsCartIdPaymentorderIdRequest

	// PostV1GuestcartsCartIdPaymentorderIdExecute executes the request
	//  @return bool
	PostV1GuestcartsCartIdPaymentorderIdExecute(r ApiPostV1GuestcartsCartIdPaymentorderIdRequest) (bool, *http.Response, error)
}

// GuestCartsCartIdPaymentOrderIdAPIService GuestCartsCartIdPaymentOrderIdAPI service
type GuestCartsCartIdPaymentOrderIdAPIService service

type ApiGetV1GuestcartsCartIdPaymentorderIdRequest struct {
	ctx        context.Context
	ApiService GuestCartsCartIdPaymentOrderIdAPI
	cartId     string
	id         string
}

func (r ApiGetV1GuestcartsCartIdPaymentorderIdRequest) Execute() (*PaymentServicesPaypalDataPaymentOrderDetailsInterface, *http.Response, error) {
	return r.ApiService.GetV1GuestcartsCartIdPaymentorderIdExecute(r)
}

/*
GetV1GuestcartsCartIdPaymentorderId guest-carts/{cartId}/payment-order/{id}

Get payment order for guest customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId
	@param id
	@return ApiGetV1GuestcartsCartIdPaymentorderIdRequest
*/
func (a *GuestCartsCartIdPaymentOrderIdAPIService) GetV1GuestcartsCartIdPaymentorderId(ctx context.Context, cartId string, id string) ApiGetV1GuestcartsCartIdPaymentorderIdRequest {
	return ApiGetV1GuestcartsCartIdPaymentorderIdRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
		id:         id,
	}
}

// Execute executes the request
//
//	@return PaymentServicesPaypalDataPaymentOrderDetailsInterface
func (a *GuestCartsCartIdPaymentOrderIdAPIService) GetV1GuestcartsCartIdPaymentorderIdExecute(r ApiGetV1GuestcartsCartIdPaymentorderIdRequest) (*PaymentServicesPaypalDataPaymentOrderDetailsInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaymentServicesPaypalDataPaymentOrderDetailsInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestCartsCartIdPaymentOrderIdAPIService.GetV1GuestcartsCartIdPaymentorderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/guest-carts/{cartId}/payment-order/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"cartId"+"}", url.PathEscape(parameterValueToString(r.cartId, "cartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiPostV1GuestcartsCartIdPaymentorderIdRequest struct {
	ctx        context.Context
	ApiService GuestCartsCartIdPaymentOrderIdAPI
	cartId     string
	id         string
}

func (r ApiPostV1GuestcartsCartIdPaymentorderIdRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.PostV1GuestcartsCartIdPaymentorderIdExecute(r)
}

/*
PostV1GuestcartsCartIdPaymentorderId guest-carts/{cartId}/payment-order/{id}

Sync payment order for guest customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId
	@param id
	@return ApiPostV1GuestcartsCartIdPaymentorderIdRequest
*/
func (a *GuestCartsCartIdPaymentOrderIdAPIService) PostV1GuestcartsCartIdPaymentorderId(ctx context.Context, cartId string, id string) ApiPostV1GuestcartsCartIdPaymentorderIdRequest {
	return ApiPostV1GuestcartsCartIdPaymentorderIdRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
		id:         id,
	}
}

// Execute executes the request
//
//	@return bool
func (a *GuestCartsCartIdPaymentOrderIdAPIService) PostV1GuestcartsCartIdPaymentorderIdExecute(r ApiPostV1GuestcartsCartIdPaymentorderIdRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestCartsCartIdPaymentOrderIdAPIService.PostV1GuestcartsCartIdPaymentorderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/guest-carts/{cartId}/payment-order/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"cartId"+"}", url.PathEscape(parameterValueToString(r.cartId, "cartId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
