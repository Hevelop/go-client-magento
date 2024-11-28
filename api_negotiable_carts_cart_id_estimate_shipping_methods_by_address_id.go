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

type NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPI interface {

	/*
		PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressid negotiable-carts/{cartId}/estimate-shipping-methods-by-address-id

		Estimate shipping

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The shopping cart ID.
		@return ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest
	*/
	PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressid(ctx context.Context, cartId int32) ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest

	// PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidExecute executes the request
	//  @return []QuoteDataShippingMethodInterface
	PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidExecute(r ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest) ([]QuoteDataShippingMethodInterface, *http.Response, error)
}

// NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPIService NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPI service
type NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPIService service

type ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest struct {
	ctx                                                      context.Context
	ApiService                                               NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPI
	cartId                                                   int32
	postV1CartsMineEstimateshippingmethodsbyaddressidRequest *PostV1CartsMineEstimateshippingmethodsbyaddressidRequest
}

func (r ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest) PostV1CartsMineEstimateshippingmethodsbyaddressidRequest(postV1CartsMineEstimateshippingmethodsbyaddressidRequest PostV1CartsMineEstimateshippingmethodsbyaddressidRequest) ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest {
	r.postV1CartsMineEstimateshippingmethodsbyaddressidRequest = &postV1CartsMineEstimateshippingmethodsbyaddressidRequest
	return r
}

func (r ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest) Execute() ([]QuoteDataShippingMethodInterface, *http.Response, error) {
	return r.ApiService.PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidExecute(r)
}

/*
PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressid negotiable-carts/{cartId}/estimate-shipping-methods-by-address-id

Estimate shipping

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The shopping cart ID.
	@return ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest
*/
func (a *NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPIService) PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressid(ctx context.Context, cartId int32) ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest {
	return ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return []QuoteDataShippingMethodInterface
func (a *NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPIService) PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidExecute(r ApiPostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressidRequest) ([]QuoteDataShippingMethodInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []QuoteDataShippingMethodInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegotiableCartsCartIdEstimateShippingMethodsByAddressIdAPIService.PostV1NegotiablecartsCartIdEstimateshippingmethodsbyaddressid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/negotiable-carts/{cartId}/estimate-shipping-methods-by-address-id"
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
	localVarPostBody = r.postV1CartsMineEstimateshippingmethodsbyaddressidRequest
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