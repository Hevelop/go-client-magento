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

type GuestCartsCartIdCollectTotalsAPI interface {

	/*
		PutV1GuestcartsCartIdCollecttotals guest-carts/{cartId}/collect-totals

		Set shipping/billing methods and additional data for cart and collect totals for guest.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param cartId The cart ID.
		@return ApiPutV1GuestcartsCartIdCollecttotalsRequest
	*/
	PutV1GuestcartsCartIdCollecttotals(ctx context.Context, cartId string) ApiPutV1GuestcartsCartIdCollecttotalsRequest

	// PutV1GuestcartsCartIdCollecttotalsExecute executes the request
	//  @return QuoteDataTotalsInterface
	PutV1GuestcartsCartIdCollecttotalsExecute(r ApiPutV1GuestcartsCartIdCollecttotalsRequest) (*QuoteDataTotalsInterface, *http.Response, error)
}

// GuestCartsCartIdCollectTotalsAPIService GuestCartsCartIdCollectTotalsAPI service
type GuestCartsCartIdCollectTotalsAPIService service

type ApiPutV1GuestcartsCartIdCollecttotalsRequest struct {
	ctx                                context.Context
	ApiService                         GuestCartsCartIdCollectTotalsAPI
	cartId                             string
	putV1CartsMineCollecttotalsRequest *PutV1CartsMineCollecttotalsRequest
}

func (r ApiPutV1GuestcartsCartIdCollecttotalsRequest) PutV1CartsMineCollecttotalsRequest(putV1CartsMineCollecttotalsRequest PutV1CartsMineCollecttotalsRequest) ApiPutV1GuestcartsCartIdCollecttotalsRequest {
	r.putV1CartsMineCollecttotalsRequest = &putV1CartsMineCollecttotalsRequest
	return r
}

func (r ApiPutV1GuestcartsCartIdCollecttotalsRequest) Execute() (*QuoteDataTotalsInterface, *http.Response, error) {
	return r.ApiService.PutV1GuestcartsCartIdCollecttotalsExecute(r)
}

/*
PutV1GuestcartsCartIdCollecttotals guest-carts/{cartId}/collect-totals

Set shipping/billing methods and additional data for cart and collect totals for guest.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cartId The cart ID.
	@return ApiPutV1GuestcartsCartIdCollecttotalsRequest
*/
func (a *GuestCartsCartIdCollectTotalsAPIService) PutV1GuestcartsCartIdCollecttotals(ctx context.Context, cartId string) ApiPutV1GuestcartsCartIdCollecttotalsRequest {
	return ApiPutV1GuestcartsCartIdCollecttotalsRequest{
		ApiService: a,
		ctx:        ctx,
		cartId:     cartId,
	}
}

// Execute executes the request
//
//	@return QuoteDataTotalsInterface
func (a *GuestCartsCartIdCollectTotalsAPIService) PutV1GuestcartsCartIdCollecttotalsExecute(r ApiPutV1GuestcartsCartIdCollecttotalsRequest) (*QuoteDataTotalsInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *QuoteDataTotalsInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GuestCartsCartIdCollectTotalsAPIService.PutV1GuestcartsCartIdCollecttotals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/guest-carts/{cartId}/collect-totals"
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
	localVarPostBody = r.putV1CartsMineCollecttotalsRequest
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
