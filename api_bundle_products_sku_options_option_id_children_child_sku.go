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

type BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI interface {

	/*
		DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku bundle-products/{sku}/options/{optionId}/children/{childSku}

		Remove product from Bundle product option

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sku
		@param optionId
		@param childSku
		@return ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest
	*/
	DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku(ctx context.Context, sku string, optionId int32, childSku string) ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest

	// DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuExecute executes the request
	//  @return bool
	DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuExecute(r ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest) (bool, *http.Response, error)
}

// BundleProductsSkuOptionsOptionIdChildrenChildSkuAPIService BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI service
type BundleProductsSkuOptionsOptionIdChildrenChildSkuAPIService service

type ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest struct {
	ctx        context.Context
	ApiService BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI
	sku        string
	optionId   int32
	childSku   string
}

func (r ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuExecute(r)
}

/*
DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku bundle-products/{sku}/options/{optionId}/children/{childSku}

Remove product from Bundle product option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sku
	@param optionId
	@param childSku
	@return ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest
*/
func (a *BundleProductsSkuOptionsOptionIdChildrenChildSkuAPIService) DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku(ctx context.Context, sku string, optionId int32, childSku string) ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest {
	return ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest{
		ApiService: a,
		ctx:        ctx,
		sku:        sku,
		optionId:   optionId,
		childSku:   childSku,
	}
}

// Execute executes the request
//
//	@return bool
func (a *BundleProductsSkuOptionsOptionIdChildrenChildSkuAPIService) DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuExecute(r ApiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundleProductsSkuOptionsOptionIdChildrenChildSkuAPIService.DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/bundle-products/{sku}/options/{optionId}/children/{childSku}"
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optionId"+"}", url.PathEscape(parameterValueToString(r.optionId, "optionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"childSku"+"}", url.PathEscape(parameterValueToString(r.childSku, "childSku")), -1)

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