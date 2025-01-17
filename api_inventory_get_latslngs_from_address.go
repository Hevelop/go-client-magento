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

type InventoryGetLatslngsFromAddressAPI interface {

	/*
		GetV1InventoryGetlatslngsfromaddress inventory/get-latslngs-from-address

		Get all available latitude and longitude objects from address.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetV1InventoryGetlatslngsfromaddressRequest
	*/
	GetV1InventoryGetlatslngsfromaddress(ctx context.Context) ApiGetV1InventoryGetlatslngsfromaddressRequest

	// GetV1InventoryGetlatslngsfromaddressExecute executes the request
	//  @return []InventoryDistanceBasedSourceSelectionApiDataLatLngInterface
	GetV1InventoryGetlatslngsfromaddressExecute(r ApiGetV1InventoryGetlatslngsfromaddressRequest) ([]InventoryDistanceBasedSourceSelectionApiDataLatLngInterface, *http.Response, error)
}

// InventoryGetLatslngsFromAddressAPIService InventoryGetLatslngsFromAddressAPI service
type InventoryGetLatslngsFromAddressAPIService service

type ApiGetV1InventoryGetlatslngsfromaddressRequest struct {
	ctx             context.Context
	ApiService      InventoryGetLatslngsFromAddressAPI
	addressCountry  *string
	addressPostcode *string
	addressStreet   *string
	addressRegion   *string
	addressCity     *string
}

// Shipping country
func (r ApiGetV1InventoryGetlatslngsfromaddressRequest) AddressCountry(addressCountry string) ApiGetV1InventoryGetlatslngsfromaddressRequest {
	r.addressCountry = &addressCountry
	return r
}

// Shipping postcode
func (r ApiGetV1InventoryGetlatslngsfromaddressRequest) AddressPostcode(addressPostcode string) ApiGetV1InventoryGetlatslngsfromaddressRequest {
	r.addressPostcode = &addressPostcode
	return r
}

// Shipping street address
func (r ApiGetV1InventoryGetlatslngsfromaddressRequest) AddressStreet(addressStreet string) ApiGetV1InventoryGetlatslngsfromaddressRequest {
	r.addressStreet = &addressStreet
	return r
}

// Shipping region
func (r ApiGetV1InventoryGetlatslngsfromaddressRequest) AddressRegion(addressRegion string) ApiGetV1InventoryGetlatslngsfromaddressRequest {
	r.addressRegion = &addressRegion
	return r
}

// Shipping city
func (r ApiGetV1InventoryGetlatslngsfromaddressRequest) AddressCity(addressCity string) ApiGetV1InventoryGetlatslngsfromaddressRequest {
	r.addressCity = &addressCity
	return r
}

func (r ApiGetV1InventoryGetlatslngsfromaddressRequest) Execute() ([]InventoryDistanceBasedSourceSelectionApiDataLatLngInterface, *http.Response, error) {
	return r.ApiService.GetV1InventoryGetlatslngsfromaddressExecute(r)
}

/*
GetV1InventoryGetlatslngsfromaddress inventory/get-latslngs-from-address

Get all available latitude and longitude objects from address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetV1InventoryGetlatslngsfromaddressRequest
*/
func (a *InventoryGetLatslngsFromAddressAPIService) GetV1InventoryGetlatslngsfromaddress(ctx context.Context) ApiGetV1InventoryGetlatslngsfromaddressRequest {
	return ApiGetV1InventoryGetlatslngsfromaddressRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []InventoryDistanceBasedSourceSelectionApiDataLatLngInterface
func (a *InventoryGetLatslngsFromAddressAPIService) GetV1InventoryGetlatslngsfromaddressExecute(r ApiGetV1InventoryGetlatslngsfromaddressRequest) ([]InventoryDistanceBasedSourceSelectionApiDataLatLngInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []InventoryDistanceBasedSourceSelectionApiDataLatLngInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryGetLatslngsFromAddressAPIService.GetV1InventoryGetlatslngsfromaddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/inventory/get-latslngs-from-address"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.addressCountry != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address[country]", r.addressCountry, "form", "")
	}
	if r.addressPostcode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address[postcode]", r.addressPostcode, "form", "")
	}
	if r.addressStreet != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address[street]", r.addressStreet, "form", "")
	}
	if r.addressRegion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address[region]", r.addressRegion, "form", "")
	}
	if r.addressCity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address[city]", r.addressCity, "form", "")
	}
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
