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

type CartsCartIdCouponsDeleteByCodesAPI interface {

	/*
		PostV2CartsCartIdCouponsDeleteByCodes carts/{cartId}/coupons/deleteByCodes

		Deletes coupon(s) from a specified cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The cart ID.
		@return ApiPostV2CartsCartIdCouponsDeleteByCodesRequest
	*/
	PostV2CartsCartIdCouponsDeleteByCodes(ctx context.Context, cartId int32) ApiPostV2CartsCartIdCouponsDeleteByCodesRequest

	// PostV2CartsCartIdCouponsDeleteByCodesExecute executes the request
	//  @return ErrorResponse
	PostV2CartsCartIdCouponsDeleteByCodesExecute(r ApiPostV2CartsCartIdCouponsDeleteByCodesRequest) (*ErrorResponse, *http.Response, error)
}

// CartsCartIdCouponsDeleteByCodesAPIService CartsCartIdCouponsDeleteByCodesAPI service
type CartsCartIdCouponsDeleteByCodesAPIService service

type ApiPostV2CartsCartIdCouponsDeleteByCodesRequest struct {
	ctx                                        context.Context
	ApiService                                 CartsCartIdCouponsDeleteByCodesAPI
	cartId                                     int32
	postV2CartsMineCouponsDeleteByCodesRequest *PostV2CartsMineCouponsDeleteByCodesRequest
}

func (r ApiPostV2CartsCartIdCouponsDeleteByCodesRequest) PostV2CartsMineCouponsDeleteByCodesRequest(postV2CartsMineCouponsDeleteByCodesRequest PostV2CartsMineCouponsDeleteByCodesRequest) ApiPostV2CartsCartIdCouponsDeleteByCodesRequest {
	r.postV2CartsMineCouponsDeleteByCodesRequest = &postV2CartsMineCouponsDeleteByCodesRequest
	return r
}

func (r ApiPostV2CartsCartIdCouponsDeleteByCodesRequest) Execute() (*ErrorResponse, *http.Response, error) {
	return r.ApiService.PostV2CartsCartIdCouponsDeleteByCodesExecute(r)
}

/*
PostV2CartsCartIdCouponsDeleteByCodes carts/{cartId}/coupons/deleteByCodes

Deletes coupon(s) from a specified cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The cart ID.
	@return ApiPostV2CartsCartIdCouponsDeleteByCodesRequest
*/
func (a *CartsCartIdCouponsDeleteByCodesAPIService) PostV2CartsCartIdCouponsDeleteByCodes(ctx context.Context, cartId int32) ApiPostV2CartsCartIdCouponsDeleteByCodesRequest {
	return ApiPostV2CartsCartIdCouponsDeleteByCodesRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return ErrorResponse
func (a *CartsCartIdCouponsDeleteByCodesAPIService) PostV2CartsCartIdCouponsDeleteByCodesExecute(r ApiPostV2CartsCartIdCouponsDeleteByCodesRequest) (*ErrorResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ErrorResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartsCartIdCouponsDeleteByCodesAPIService.PostV2CartsCartIdCouponsDeleteByCodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V2/carts/{cartId}/coupons/deleteByCodes"
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
	localVarPostBody = r.postV2CartsMineCouponsDeleteByCodesRequest
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
