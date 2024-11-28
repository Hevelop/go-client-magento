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

type ProductsSkuWebsitesWebsiteIdAPI interface {

	/*
		DeleteV1ProductsSkuWebsitesWebsiteId products/{sku}/websites/{websiteId}

		Remove the website assignment from the product by product sku

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param sku
		@param websiteId
		@return ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest
	*/
	DeleteV1ProductsSkuWebsitesWebsiteId(ctx context.Context, sku string, websiteId int32) ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest

	// DeleteV1ProductsSkuWebsitesWebsiteIdExecute executes the request
	//  @return bool
	DeleteV1ProductsSkuWebsitesWebsiteIdExecute(r ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest) (bool, *http.Response, error)
}

// ProductsSkuWebsitesWebsiteIdAPIService ProductsSkuWebsitesWebsiteIdAPI service
type ProductsSkuWebsitesWebsiteIdAPIService service

type ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest struct {
	ctx        context.Context
	ApiService ProductsSkuWebsitesWebsiteIdAPI
	sku        string
	websiteId  int32
}

func (r ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.DeleteV1ProductsSkuWebsitesWebsiteIdExecute(r)
}

/*
DeleteV1ProductsSkuWebsitesWebsiteId products/{sku}/websites/{websiteId}

Remove the website assignment from the product by product sku

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param sku
	@param websiteId
	@return ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest
*/
func (a *ProductsSkuWebsitesWebsiteIdAPIService) DeleteV1ProductsSkuWebsitesWebsiteId(ctx context.Context, sku string, websiteId int32) ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest {
	return ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest{
		ApiService: a,
		ctx:        ctx,
		sku:        sku,
		websiteId:  websiteId,
	}
}

// Execute executes the request
//
//	@return bool
func (a *ProductsSkuWebsitesWebsiteIdAPIService) DeleteV1ProductsSkuWebsitesWebsiteIdExecute(r ApiDeleteV1ProductsSkuWebsitesWebsiteIdRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsSkuWebsitesWebsiteIdAPIService.DeleteV1ProductsSkuWebsitesWebsiteId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/products/{sku}/websites/{websiteId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sku"+"}", url.PathEscape(parameterValueToString(r.sku, "sku")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"websiteId"+"}", url.PathEscape(parameterValueToString(r.websiteId, "websiteId")), -1)

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