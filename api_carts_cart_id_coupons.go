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

type CartsCartIdCouponsAPI interface {

	/*
		DeleteV1CartsCartIdCoupons carts/{cartId}/coupons

		Deletes a coupon from a specified cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The cart ID.
		@return ApiDeleteV1CartsCartIdCouponsRequest
	*/
	DeleteV1CartsCartIdCoupons(ctx context.Context, cartId int32) ApiDeleteV1CartsCartIdCouponsRequest

	// DeleteV1CartsCartIdCouponsExecute executes the request
	//  @return bool
	DeleteV1CartsCartIdCouponsExecute(r ApiDeleteV1CartsCartIdCouponsRequest) (bool, *http.Response, error)

	/*
		GetV1CartsCartIdCoupons carts/{cartId}/coupons

		Returns information for a coupon in a specified cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The cart ID.
		@return ApiGetV1CartsCartIdCouponsRequest
	*/
	GetV1CartsCartIdCoupons(ctx context.Context, cartId int32) ApiGetV1CartsCartIdCouponsRequest

	// GetV1CartsCartIdCouponsExecute executes the request
	//  @return string
	GetV1CartsCartIdCouponsExecute(r ApiGetV1CartsCartIdCouponsRequest) (string, *http.Response, error)

	/*
		GetV2CartsCartIdCoupons carts/{cartId}/coupons

		Returns information for all coupons in a specified cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The cart ID.
		@return ApiGetV2CartsCartIdCouponsRequest
	*/
	GetV2CartsCartIdCoupons(ctx context.Context, cartId int32) ApiGetV2CartsCartIdCouponsRequest

	// GetV2CartsCartIdCouponsExecute executes the request
	//  @return []string
	GetV2CartsCartIdCouponsExecute(r ApiGetV2CartsCartIdCouponsRequest) ([]string, *http.Response, error)

	/*
		PostV2CartsCartIdCoupons carts/{cartId}/coupons

		Append the coupon code(s) to cart

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId
		@return ApiPostV2CartsCartIdCouponsRequest
	*/
	PostV2CartsCartIdCoupons(ctx context.Context, cartId int32) ApiPostV2CartsCartIdCouponsRequest

	// PostV2CartsCartIdCouponsExecute executes the request
	//  @return ErrorResponse
	PostV2CartsCartIdCouponsExecute(r ApiPostV2CartsCartIdCouponsRequest) (*ErrorResponse, *http.Response, error)

	/*
		PutV2CartsCartIdCoupons carts/{cartId}/coupons

		Replace the coupon code(s) in cart with the new code(s)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId
		@return ApiPutV2CartsCartIdCouponsRequest
	*/
	PutV2CartsCartIdCoupons(ctx context.Context, cartId int32) ApiPutV2CartsCartIdCouponsRequest

	// PutV2CartsCartIdCouponsExecute executes the request
	//  @return ErrorResponse
	PutV2CartsCartIdCouponsExecute(r ApiPutV2CartsCartIdCouponsRequest) (*ErrorResponse, *http.Response, error)
}

// CartsCartIdCouponsAPIService CartsCartIdCouponsAPI service
type CartsCartIdCouponsAPIService service

type ApiDeleteV1CartsCartIdCouponsRequest struct {
	ctx        context.Context
	ApiService CartsCartIdCouponsAPI
	cartId     int32
}

func (r ApiDeleteV1CartsCartIdCouponsRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.DeleteV1CartsCartIdCouponsExecute(r)
}

/*
DeleteV1CartsCartIdCoupons carts/{cartId}/coupons

Deletes a coupon from a specified cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The cart ID.
	@return ApiDeleteV1CartsCartIdCouponsRequest
*/
func (a *CartsCartIdCouponsAPIService) DeleteV1CartsCartIdCoupons(ctx context.Context, cartId int32) ApiDeleteV1CartsCartIdCouponsRequest {
	return ApiDeleteV1CartsCartIdCouponsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return bool
func (a *CartsCartIdCouponsAPIService) DeleteV1CartsCartIdCouponsExecute(r ApiDeleteV1CartsCartIdCouponsRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdCouponsAPIService.DeleteV1CartsCartIdCoupons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/carts/{cartId}/coupons"
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

type ApiGetV1CartsCartIdCouponsRequest struct {
	ctx        context.Context
	ApiService CartsCartIdCouponsAPI
	cartId     int32
}

func (r ApiGetV1CartsCartIdCouponsRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.GetV1CartsCartIdCouponsExecute(r)
}

/*
GetV1CartsCartIdCoupons carts/{cartId}/coupons

Returns information for a coupon in a specified cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The cart ID.
	@return ApiGetV1CartsCartIdCouponsRequest
*/
func (a *CartsCartIdCouponsAPIService) GetV1CartsCartIdCoupons(ctx context.Context, cartId int32) ApiGetV1CartsCartIdCouponsRequest {
	return ApiGetV1CartsCartIdCouponsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return string
func (a *CartsCartIdCouponsAPIService) GetV1CartsCartIdCouponsExecute(r ApiGetV1CartsCartIdCouponsRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdCouponsAPIService.GetV1CartsCartIdCoupons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/carts/{cartId}/coupons"
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

type ApiGetV2CartsCartIdCouponsRequest struct {
	ctx        context.Context
	ApiService CartsCartIdCouponsAPI
	cartId     int32
}

func (r ApiGetV2CartsCartIdCouponsRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetV2CartsCartIdCouponsExecute(r)
}

/*
GetV2CartsCartIdCoupons carts/{cartId}/coupons

Returns information for all coupons in a specified cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The cart ID.
	@return ApiGetV2CartsCartIdCouponsRequest
*/
func (a *CartsCartIdCouponsAPIService) GetV2CartsCartIdCoupons(ctx context.Context, cartId int32) ApiGetV2CartsCartIdCouponsRequest {
	return ApiGetV2CartsCartIdCouponsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return []string
func (a *CartsCartIdCouponsAPIService) GetV2CartsCartIdCouponsExecute(r ApiGetV2CartsCartIdCouponsRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdCouponsAPIService.GetV2CartsCartIdCoupons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V2/carts/{cartId}/coupons"
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

type ApiPostV2CartsCartIdCouponsRequest struct {
	ctx                          context.Context
	ApiService                   CartsCartIdCouponsAPI
	cartId                       int32
	putV2CartsMineCouponsRequest *PutV2CartsMineCouponsRequest
}

func (r ApiPostV2CartsCartIdCouponsRequest) PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest PutV2CartsMineCouponsRequest) ApiPostV2CartsCartIdCouponsRequest {
	r.putV2CartsMineCouponsRequest = &putV2CartsMineCouponsRequest
	return r
}

func (r ApiPostV2CartsCartIdCouponsRequest) Execute() (*ErrorResponse, *http.Response, error) {
	return r.ApiService.PostV2CartsCartIdCouponsExecute(r)
}

/*
PostV2CartsCartIdCoupons carts/{cartId}/coupons

Append the coupon code(s) to cart

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId
	@return ApiPostV2CartsCartIdCouponsRequest
*/
func (a *CartsCartIdCouponsAPIService) PostV2CartsCartIdCoupons(ctx context.Context, cartId int32) ApiPostV2CartsCartIdCouponsRequest {
	return ApiPostV2CartsCartIdCouponsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return ErrorResponse
func (a *CartsCartIdCouponsAPIService) PostV2CartsCartIdCouponsExecute(r ApiPostV2CartsCartIdCouponsRequest) (*ErrorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdCouponsAPIService.PostV2CartsCartIdCoupons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V2/carts/{cartId}/coupons"
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
	localVarPostBody = r.putV2CartsMineCouponsRequest
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

type ApiPutV2CartsCartIdCouponsRequest struct {
	ctx                          context.Context
	ApiService                   CartsCartIdCouponsAPI
	cartId                       int32
	putV2CartsMineCouponsRequest *PutV2CartsMineCouponsRequest
}

func (r ApiPutV2CartsCartIdCouponsRequest) PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest PutV2CartsMineCouponsRequest) ApiPutV2CartsCartIdCouponsRequest {
	r.putV2CartsMineCouponsRequest = &putV2CartsMineCouponsRequest
	return r
}

func (r ApiPutV2CartsCartIdCouponsRequest) Execute() (*ErrorResponse, *http.Response, error) {
	return r.ApiService.PutV2CartsCartIdCouponsExecute(r)
}

/*
PutV2CartsCartIdCoupons carts/{cartId}/coupons

Replace the coupon code(s) in cart with the new code(s)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId
	@return ApiPutV2CartsCartIdCouponsRequest
*/
func (a *CartsCartIdCouponsAPIService) PutV2CartsCartIdCoupons(ctx context.Context, cartId int32) ApiPutV2CartsCartIdCouponsRequest {
	return ApiPutV2CartsCartIdCouponsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return ErrorResponse
func (a *CartsCartIdCouponsAPIService) PutV2CartsCartIdCouponsExecute(r ApiPutV2CartsCartIdCouponsRequest) (*ErrorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdCouponsAPIService.PutV2CartsCartIdCoupons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V2/carts/{cartId}/coupons"
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
	localVarPostBody = r.putV2CartsMineCouponsRequest
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
