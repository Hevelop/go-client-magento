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

type StockItemsProductSkuAPI interface {

	/*
		GetV1StockItemsProductSku stockItems/{productSku}



		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param productSku
		@return ApiGetV1StockItemsProductSkuRequest
	*/
	GetV1StockItemsProductSku(ctx context.Context, productSku string) ApiGetV1StockItemsProductSkuRequest

	// GetV1StockItemsProductSkuExecute executes the request
	//  @return CatalogInventoryDataStockItemInterface
	GetV1StockItemsProductSkuExecute(r ApiGetV1StockItemsProductSkuRequest) (*CatalogInventoryDataStockItemInterface, *http.Response, error)
}

// StockItemsProductSkuAPIService StockItemsProductSkuAPI service
type StockItemsProductSkuAPIService service

type ApiGetV1StockItemsProductSkuRequest struct {
	ctx        context.Context
	ApiService StockItemsProductSkuAPI
	productSku string
	scopeId    *int32
}

func (r ApiGetV1StockItemsProductSkuRequest) ScopeId(scopeId int32) ApiGetV1StockItemsProductSkuRequest {
	r.scopeId = &scopeId
	return r
}

func (r ApiGetV1StockItemsProductSkuRequest) Execute() (*CatalogInventoryDataStockItemInterface, *http.Response, error) {
	return r.ApiService.GetV1StockItemsProductSkuExecute(r)
}

/*
GetV1StockItemsProductSku stockItems/{productSku}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param productSku
	@return ApiGetV1StockItemsProductSkuRequest
*/
func (a *StockItemsProductSkuAPIService) GetV1StockItemsProductSku(ctx context.Context, productSku string) ApiGetV1StockItemsProductSkuRequest {
	return ApiGetV1StockItemsProductSkuRequest{
		ApiService: a,
		ctx:        ctx,
		productSku: productSku,
	}
}

// Execute executes the request
//
//	@return CatalogInventoryDataStockItemInterface
func (a *StockItemsProductSkuAPIService) GetV1StockItemsProductSkuExecute(r ApiGetV1StockItemsProductSkuRequest) (*CatalogInventoryDataStockItemInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CatalogInventoryDataStockItemInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsProductSkuAPIService.GetV1StockItemsProductSku")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/stockItems/{productSku}"
	localVarPath = strings.Replace(localVarPath, "{"+"productSku"+"}", url.PathEscape(parameterValueToString(r.productSku, "productSku")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.scopeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scopeId", r.scopeId, "form", "")
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
