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

type ProductsMediaTypesAttributeSetNameAPI interface {

	/*
		GetV1ProductsMediaTypesAttributeSetName products/media/types/{attributeSetName}

		Retrieve the list of media attributes (fronted input type is media_image) assigned to the given attribute set.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param attributeSetName
		@return ApiGetV1ProductsMediaTypesAttributeSetNameRequest
	*/
	GetV1ProductsMediaTypesAttributeSetName(ctx context.Context, attributeSetName string) ApiGetV1ProductsMediaTypesAttributeSetNameRequest

	// GetV1ProductsMediaTypesAttributeSetNameExecute executes the request
	//  @return []CatalogDataProductAttributeInterface
	GetV1ProductsMediaTypesAttributeSetNameExecute(r ApiGetV1ProductsMediaTypesAttributeSetNameRequest) ([]CatalogDataProductAttributeInterface, *http.Response, error)
}

// ProductsMediaTypesAttributeSetNameAPIService ProductsMediaTypesAttributeSetNameAPI service
type ProductsMediaTypesAttributeSetNameAPIService service

type ApiGetV1ProductsMediaTypesAttributeSetNameRequest struct {
	ctx              context.Context
	ApiService       ProductsMediaTypesAttributeSetNameAPI
	attributeSetName string
}

func (r ApiGetV1ProductsMediaTypesAttributeSetNameRequest) Execute() ([]CatalogDataProductAttributeInterface, *http.Response, error) {
	return r.ApiService.GetV1ProductsMediaTypesAttributeSetNameExecute(r)
}

/*
GetV1ProductsMediaTypesAttributeSetName products/media/types/{attributeSetName}

Retrieve the list of media attributes (fronted input type is media_image) assigned to the given attribute set.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param attributeSetName
	@return ApiGetV1ProductsMediaTypesAttributeSetNameRequest
*/
func (a *ProductsMediaTypesAttributeSetNameAPIService) GetV1ProductsMediaTypesAttributeSetName(ctx context.Context, attributeSetName string) ApiGetV1ProductsMediaTypesAttributeSetNameRequest {
	return ApiGetV1ProductsMediaTypesAttributeSetNameRequest{
		ApiService:       a,
		ctx:              ctx,
		attributeSetName: attributeSetName,
	}
}

// Execute executes the request
//
//	@return []CatalogDataProductAttributeInterface
func (a *ProductsMediaTypesAttributeSetNameAPIService) GetV1ProductsMediaTypesAttributeSetNameExecute(r ApiGetV1ProductsMediaTypesAttributeSetNameRequest) ([]CatalogDataProductAttributeInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []CatalogDataProductAttributeInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsMediaTypesAttributeSetNameAPIService.GetV1ProductsMediaTypesAttributeSetName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/products/media/types/{attributeSetName}"
	localVarPath = strings.Replace(localVarPath, "{"+"attributeSetName"+"}", url.PathEscape(parameterValueToString(r.attributeSetName, "attributeSetName")), -1)

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
