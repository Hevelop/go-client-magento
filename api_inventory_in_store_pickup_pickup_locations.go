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

type InventoryInStorePickupPickupLocationsAPI interface {

	/*
		GetV1InventoryInstorepickupPickuplocations inventory/in-store-pickup/pickup-locations/

		Get Pickup Locations according to the results of filtration by Search Request.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetV1InventoryInstorepickupPickuplocationsRequest
	*/
	GetV1InventoryInstorepickupPickuplocations(ctx context.Context) ApiGetV1InventoryInstorepickupPickuplocationsRequest

	// GetV1InventoryInstorepickupPickuplocationsExecute executes the request
	//  @return InventoryInStorePickupApiDataSearchResultInterface
	GetV1InventoryInstorepickupPickuplocationsExecute(r ApiGetV1InventoryInstorepickupPickuplocationsRequest) (*InventoryInStorePickupApiDataSearchResultInterface, *http.Response, error)
}

// InventoryInStorePickupPickupLocationsAPIService InventoryInStorePickupPickupLocationsAPI service
type InventoryInStorePickupPickupLocationsAPIService service

type ApiGetV1InventoryInstorepickupPickuplocationsRequest struct {
	ctx                                                 context.Context
	ApiService                                          InventoryInStorePickupPickupLocationsAPI
	searchRequestAreaRadius                             *int32
	searchRequestAreaSearchTerm                         *string
	searchRequestFiltersCountryValue                    *string
	searchRequestFiltersCountryConditionType            *string
	searchRequestFiltersPostcodeValue                   *string
	searchRequestFiltersPostcodeConditionType           *string
	searchRequestFiltersRegionValue                     *string
	searchRequestFiltersRegionConditionType             *string
	searchRequestFiltersRegionIdValue                   *string
	searchRequestFiltersRegionIdConditionType           *string
	searchRequestFiltersCityValue                       *string
	searchRequestFiltersCityConditionType               *string
	searchRequestFiltersStreetValue                     *string
	searchRequestFiltersStreetConditionType             *string
	searchRequestFiltersNameValue                       *string
	searchRequestFiltersNameConditionType               *string
	searchRequestFiltersPickupLocationCodeValue         *string
	searchRequestFiltersPickupLocationCodeConditionType *string
	searchRequestPageSize                               *int32
	searchRequestCurrentPage                            *int32
	searchRequestScopeType                              *string
	searchRequestScopeCode                              *string
	searchRequestSort0Field                             *string
	searchRequestSort0Direction                         *string
	searchRequestExtensionAttributesProductsInfo0Sku    *string
}

// Search radius in KM.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestAreaRadius(searchRequestAreaRadius int32) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestAreaRadius = &searchRequestAreaRadius
	return r
}

// Search term string.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestAreaSearchTerm(searchRequestAreaSearchTerm string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestAreaSearchTerm = &searchRequestAreaSearchTerm
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersCountryValue(searchRequestFiltersCountryValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersCountryValue = &searchRequestFiltersCountryValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersCountryConditionType(searchRequestFiltersCountryConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersCountryConditionType = &searchRequestFiltersCountryConditionType
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersPostcodeValue(searchRequestFiltersPostcodeValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersPostcodeValue = &searchRequestFiltersPostcodeValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersPostcodeConditionType(searchRequestFiltersPostcodeConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersPostcodeConditionType = &searchRequestFiltersPostcodeConditionType
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersRegionValue(searchRequestFiltersRegionValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersRegionValue = &searchRequestFiltersRegionValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersRegionConditionType(searchRequestFiltersRegionConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersRegionConditionType = &searchRequestFiltersRegionConditionType
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersRegionIdValue(searchRequestFiltersRegionIdValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersRegionIdValue = &searchRequestFiltersRegionIdValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersRegionIdConditionType(searchRequestFiltersRegionIdConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersRegionIdConditionType = &searchRequestFiltersRegionIdConditionType
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersCityValue(searchRequestFiltersCityValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersCityValue = &searchRequestFiltersCityValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersCityConditionType(searchRequestFiltersCityConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersCityConditionType = &searchRequestFiltersCityConditionType
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersStreetValue(searchRequestFiltersStreetValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersStreetValue = &searchRequestFiltersStreetValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersStreetConditionType(searchRequestFiltersStreetConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersStreetConditionType = &searchRequestFiltersStreetConditionType
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersNameValue(searchRequestFiltersNameValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersNameValue = &searchRequestFiltersNameValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersNameConditionType(searchRequestFiltersNameConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersNameConditionType = &searchRequestFiltersNameConditionType
	return r
}

// Value.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersPickupLocationCodeValue(searchRequestFiltersPickupLocationCodeValue string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersPickupLocationCodeValue = &searchRequestFiltersPickupLocationCodeValue
	return r
}

// Condition Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestFiltersPickupLocationCodeConditionType(searchRequestFiltersPickupLocationCodeConditionType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestFiltersPickupLocationCodeConditionType = &searchRequestFiltersPickupLocationCodeConditionType
	return r
}

// Page size.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestPageSize(searchRequestPageSize int32) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestPageSize = &searchRequestPageSize
	return r
}

// Current page.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestCurrentPage(searchRequestCurrentPage int32) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestCurrentPage = &searchRequestCurrentPage
	return r
}

// Sales Channel Type.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestScopeType(searchRequestScopeType string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestScopeType = &searchRequestScopeType
	return r
}

// Sales Channel code.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestScopeCode(searchRequestScopeCode string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestScopeCode = &searchRequestScopeCode
	return r
}

// Sorting field.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestSort0Field(searchRequestSort0Field string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestSort0Field = &searchRequestSort0Field
	return r
}

// Sorting direction.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestSort0Direction(searchRequestSort0Direction string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestSort0Direction = &searchRequestSort0Direction
	return r
}

// Product SKU.
func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) SearchRequestExtensionAttributesProductsInfo0Sku(searchRequestExtensionAttributesProductsInfo0Sku string) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	r.searchRequestExtensionAttributesProductsInfo0Sku = &searchRequestExtensionAttributesProductsInfo0Sku
	return r
}

func (r ApiGetV1InventoryInstorepickupPickuplocationsRequest) Execute() (*InventoryInStorePickupApiDataSearchResultInterface, *http.Response, error) {
	return r.ApiService.GetV1InventoryInstorepickupPickuplocationsExecute(r)
}

/*
GetV1InventoryInstorepickupPickuplocations inventory/in-store-pickup/pickup-locations/

Get Pickup Locations according to the results of filtration by Search Request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetV1InventoryInstorepickupPickuplocationsRequest
*/
func (a *InventoryInStorePickupPickupLocationsAPIService) GetV1InventoryInstorepickupPickuplocations(ctx context.Context) ApiGetV1InventoryInstorepickupPickuplocationsRequest {
	return ApiGetV1InventoryInstorepickupPickuplocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InventoryInStorePickupApiDataSearchResultInterface
func (a *InventoryInStorePickupPickupLocationsAPIService) GetV1InventoryInstorepickupPickuplocationsExecute(r ApiGetV1InventoryInstorepickupPickuplocationsRequest) (*InventoryInStorePickupApiDataSearchResultInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InventoryInStorePickupApiDataSearchResultInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryInStorePickupPickupLocationsAPIService.GetV1InventoryInstorepickupPickuplocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/inventory/in-store-pickup/pickup-locations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.searchRequestAreaRadius != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[area][radius]", r.searchRequestAreaRadius, "form", "")
	}
	if r.searchRequestAreaSearchTerm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[area][searchTerm]", r.searchRequestAreaSearchTerm, "form", "")
	}
	if r.searchRequestFiltersCountryValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][country][value]", r.searchRequestFiltersCountryValue, "form", "")
	}
	if r.searchRequestFiltersCountryConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][country][conditionType]", r.searchRequestFiltersCountryConditionType, "form", "")
	}
	if r.searchRequestFiltersPostcodeValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][postcode][value]", r.searchRequestFiltersPostcodeValue, "form", "")
	}
	if r.searchRequestFiltersPostcodeConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][postcode][conditionType]", r.searchRequestFiltersPostcodeConditionType, "form", "")
	}
	if r.searchRequestFiltersRegionValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][region][value]", r.searchRequestFiltersRegionValue, "form", "")
	}
	if r.searchRequestFiltersRegionConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][region][conditionType]", r.searchRequestFiltersRegionConditionType, "form", "")
	}
	if r.searchRequestFiltersRegionIdValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][regionId][value]", r.searchRequestFiltersRegionIdValue, "form", "")
	}
	if r.searchRequestFiltersRegionIdConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][regionId][conditionType]", r.searchRequestFiltersRegionIdConditionType, "form", "")
	}
	if r.searchRequestFiltersCityValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][city][value]", r.searchRequestFiltersCityValue, "form", "")
	}
	if r.searchRequestFiltersCityConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][city][conditionType]", r.searchRequestFiltersCityConditionType, "form", "")
	}
	if r.searchRequestFiltersStreetValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][street][value]", r.searchRequestFiltersStreetValue, "form", "")
	}
	if r.searchRequestFiltersStreetConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][street][conditionType]", r.searchRequestFiltersStreetConditionType, "form", "")
	}
	if r.searchRequestFiltersNameValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][name][value]", r.searchRequestFiltersNameValue, "form", "")
	}
	if r.searchRequestFiltersNameConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][name][conditionType]", r.searchRequestFiltersNameConditionType, "form", "")
	}
	if r.searchRequestFiltersPickupLocationCodeValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][pickupLocationCode][value]", r.searchRequestFiltersPickupLocationCodeValue, "form", "")
	}
	if r.searchRequestFiltersPickupLocationCodeConditionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[filters][pickupLocationCode][conditionType]", r.searchRequestFiltersPickupLocationCodeConditionType, "form", "")
	}
	if r.searchRequestPageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[pageSize]", r.searchRequestPageSize, "form", "")
	}
	if r.searchRequestCurrentPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[currentPage]", r.searchRequestCurrentPage, "form", "")
	}
	if r.searchRequestScopeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[scopeType]", r.searchRequestScopeType, "form", "")
	}
	if r.searchRequestScopeCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[scopeCode]", r.searchRequestScopeCode, "form", "")
	}
	if r.searchRequestSort0Field != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[sort][0][field]", r.searchRequestSort0Field, "form", "")
	}
	if r.searchRequestSort0Direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[sort][0][direction]", r.searchRequestSort0Direction, "form", "")
	}
	if r.searchRequestExtensionAttributesProductsInfo0Sku != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchRequest[extensionAttributes][productsInfo][0][sku]", r.searchRequestExtensionAttributesProductsInfo0Sku, "form", "")
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
