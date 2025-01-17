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

type GuestCartsCartIdPaymentInformationAPI interface {

	/*
		GetV1GuestcartsCartIdPaymentinformation guest-carts/{cartId}/payment-information

		Get payment information

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId
		@return ApiGetV1GuestcartsCartIdPaymentinformationRequest
	*/
	GetV1GuestcartsCartIdPaymentinformation(ctx context.Context, cartId string) ApiGetV1GuestcartsCartIdPaymentinformationRequest

	// GetV1GuestcartsCartIdPaymentinformationExecute executes the request
	//  @return CheckoutDataPaymentDetailsInterface
	GetV1GuestcartsCartIdPaymentinformationExecute(r ApiGetV1GuestcartsCartIdPaymentinformationRequest) (*CheckoutDataPaymentDetailsInterface, *http.Response, error)

	/*
		PostV1GuestcartsCartIdPaymentinformation guest-carts/{cartId}/payment-information

		Set payment information and place order for a specified cart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId
		@return ApiPostV1GuestcartsCartIdPaymentinformationRequest
	*/
	PostV1GuestcartsCartIdPaymentinformation(ctx context.Context, cartId string) ApiPostV1GuestcartsCartIdPaymentinformationRequest

	// PostV1GuestcartsCartIdPaymentinformationExecute executes the request
	//  @return int32
	PostV1GuestcartsCartIdPaymentinformationExecute(r ApiPostV1GuestcartsCartIdPaymentinformationRequest) (int32, *http.Response, error)
}

// GuestCartsCartIdPaymentInformationAPIService GuestCartsCartIdPaymentInformationAPI service
type GuestCartsCartIdPaymentInformationAPIService service

type ApiGetV1GuestcartsCartIdPaymentinformationRequest struct {
	ctx        context.Context
	ApiService GuestCartsCartIdPaymentInformationAPI
	cartId     string
}

func (r ApiGetV1GuestcartsCartIdPaymentinformationRequest) Execute() (*CheckoutDataPaymentDetailsInterface, *http.Response, error) {
	return r.ApiService.GetV1GuestcartsCartIdPaymentinformationExecute(r)
}

/*
GetV1GuestcartsCartIdPaymentinformation guest-carts/{cartId}/payment-information

Get payment information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId
	@return ApiGetV1GuestcartsCartIdPaymentinformationRequest
*/
func (a *GuestCartsCartIdPaymentInformationAPIService) GetV1GuestcartsCartIdPaymentinformation(ctx context.Context, cartId string) ApiGetV1GuestcartsCartIdPaymentinformationRequest {
	return ApiGetV1GuestcartsCartIdPaymentinformationRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return CheckoutDataPaymentDetailsInterface
func (a *GuestCartsCartIdPaymentInformationAPIService) GetV1GuestcartsCartIdPaymentinformationExecute(r ApiGetV1GuestcartsCartIdPaymentinformationRequest) (*CheckoutDataPaymentDetailsInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CheckoutDataPaymentDetailsInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestCartsCartIdPaymentInformationAPIService.GetV1GuestcartsCartIdPaymentinformation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/guest-carts/{cartId}/payment-information"
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

type ApiPostV1GuestcartsCartIdPaymentinformationRequest struct {
	ctx                                             context.Context
	ApiService                                      GuestCartsCartIdPaymentInformationAPI
	cartId                                          string
	postV1GuestcartsCartIdPaymentinformationRequest *PostV1GuestcartsCartIdPaymentinformationRequest
}

func (r ApiPostV1GuestcartsCartIdPaymentinformationRequest) PostV1GuestcartsCartIdPaymentinformationRequest(postV1GuestcartsCartIdPaymentinformationRequest PostV1GuestcartsCartIdPaymentinformationRequest) ApiPostV1GuestcartsCartIdPaymentinformationRequest {
	r.postV1GuestcartsCartIdPaymentinformationRequest = &postV1GuestcartsCartIdPaymentinformationRequest
	return r
}

func (r ApiPostV1GuestcartsCartIdPaymentinformationRequest) Execute() (int32, *http.Response, error) {
	return r.ApiService.PostV1GuestcartsCartIdPaymentinformationExecute(r)
}

/*
PostV1GuestcartsCartIdPaymentinformation guest-carts/{cartId}/payment-information

Set payment information and place order for a specified cart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId
	@return ApiPostV1GuestcartsCartIdPaymentinformationRequest
*/
func (a *GuestCartsCartIdPaymentInformationAPIService) PostV1GuestcartsCartIdPaymentinformation(ctx context.Context, cartId string) ApiPostV1GuestcartsCartIdPaymentinformationRequest {
	return ApiPostV1GuestcartsCartIdPaymentinformationRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return int32
func (a *GuestCartsCartIdPaymentInformationAPIService) PostV1GuestcartsCartIdPaymentinformationExecute(r ApiPostV1GuestcartsCartIdPaymentinformationRequest) (int32, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue int32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestCartsCartIdPaymentInformationAPIService.PostV1GuestcartsCartIdPaymentinformation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/guest-carts/{cartId}/payment-information"
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
	localVarPostBody = r.postV1GuestcartsCartIdPaymentinformationRequest
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
