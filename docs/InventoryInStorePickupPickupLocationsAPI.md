# \InventoryInStorePickupPickupLocationsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryInstorepickupPickuplocations**](InventoryInStorePickupPickupLocationsAPI.md#GetV1InventoryInstorepickupPickuplocations) | **Get** /V1/inventory/in-store-pickup/pickup-locations/ | inventory/in-store-pickup/pickup-locations/



## GetV1InventoryInstorepickupPickuplocations

> InventoryInStorePickupApiDataSearchResultInterface GetV1InventoryInstorepickupPickuplocations(ctx).SearchRequestAreaRadius(searchRequestAreaRadius).SearchRequestAreaSearchTerm(searchRequestAreaSearchTerm).SearchRequestFiltersCountryValue(searchRequestFiltersCountryValue).SearchRequestFiltersCountryConditionType(searchRequestFiltersCountryConditionType).SearchRequestFiltersPostcodeValue(searchRequestFiltersPostcodeValue).SearchRequestFiltersPostcodeConditionType(searchRequestFiltersPostcodeConditionType).SearchRequestFiltersRegionValue(searchRequestFiltersRegionValue).SearchRequestFiltersRegionConditionType(searchRequestFiltersRegionConditionType).SearchRequestFiltersRegionIdValue(searchRequestFiltersRegionIdValue).SearchRequestFiltersRegionIdConditionType(searchRequestFiltersRegionIdConditionType).SearchRequestFiltersCityValue(searchRequestFiltersCityValue).SearchRequestFiltersCityConditionType(searchRequestFiltersCityConditionType).SearchRequestFiltersStreetValue(searchRequestFiltersStreetValue).SearchRequestFiltersStreetConditionType(searchRequestFiltersStreetConditionType).SearchRequestFiltersNameValue(searchRequestFiltersNameValue).SearchRequestFiltersNameConditionType(searchRequestFiltersNameConditionType).SearchRequestFiltersPickupLocationCodeValue(searchRequestFiltersPickupLocationCodeValue).SearchRequestFiltersPickupLocationCodeConditionType(searchRequestFiltersPickupLocationCodeConditionType).SearchRequestPageSize(searchRequestPageSize).SearchRequestCurrentPage(searchRequestCurrentPage).SearchRequestScopeType(searchRequestScopeType).SearchRequestScopeCode(searchRequestScopeCode).SearchRequestSort0Field(searchRequestSort0Field).SearchRequestSort0Direction(searchRequestSort0Direction).SearchRequestExtensionAttributesProductsInfo0Sku(searchRequestExtensionAttributesProductsInfo0Sku).Execute()

inventory/in-store-pickup/pickup-locations/



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-magento"
)

func main() {
	searchRequestAreaRadius := int32(56) // int32 | Search radius in KM. (optional)
	searchRequestAreaSearchTerm := "searchRequestAreaSearchTerm_example" // string | Search term string. (optional)
	searchRequestFiltersCountryValue := "searchRequestFiltersCountryValue_example" // string | Value. (optional)
	searchRequestFiltersCountryConditionType := "searchRequestFiltersCountryConditionType_example" // string | Condition Type. (optional)
	searchRequestFiltersPostcodeValue := "searchRequestFiltersPostcodeValue_example" // string | Value. (optional)
	searchRequestFiltersPostcodeConditionType := "searchRequestFiltersPostcodeConditionType_example" // string | Condition Type. (optional)
	searchRequestFiltersRegionValue := "searchRequestFiltersRegionValue_example" // string | Value. (optional)
	searchRequestFiltersRegionConditionType := "searchRequestFiltersRegionConditionType_example" // string | Condition Type. (optional)
	searchRequestFiltersRegionIdValue := "searchRequestFiltersRegionIdValue_example" // string | Value. (optional)
	searchRequestFiltersRegionIdConditionType := "searchRequestFiltersRegionIdConditionType_example" // string | Condition Type. (optional)
	searchRequestFiltersCityValue := "searchRequestFiltersCityValue_example" // string | Value. (optional)
	searchRequestFiltersCityConditionType := "searchRequestFiltersCityConditionType_example" // string | Condition Type. (optional)
	searchRequestFiltersStreetValue := "searchRequestFiltersStreetValue_example" // string | Value. (optional)
	searchRequestFiltersStreetConditionType := "searchRequestFiltersStreetConditionType_example" // string | Condition Type. (optional)
	searchRequestFiltersNameValue := "searchRequestFiltersNameValue_example" // string | Value. (optional)
	searchRequestFiltersNameConditionType := "searchRequestFiltersNameConditionType_example" // string | Condition Type. (optional)
	searchRequestFiltersPickupLocationCodeValue := "searchRequestFiltersPickupLocationCodeValue_example" // string | Value. (optional)
	searchRequestFiltersPickupLocationCodeConditionType := "searchRequestFiltersPickupLocationCodeConditionType_example" // string | Condition Type. (optional)
	searchRequestPageSize := int32(56) // int32 | Page size. (optional)
	searchRequestCurrentPage := int32(56) // int32 | Current page. (optional)
	searchRequestScopeType := "searchRequestScopeType_example" // string | Sales Channel Type. (optional)
	searchRequestScopeCode := "searchRequestScopeCode_example" // string | Sales Channel code. (optional)
	searchRequestSort0Field := "searchRequestSort0Field_example" // string | Sorting field. (optional)
	searchRequestSort0Direction := "searchRequestSort0Direction_example" // string | Sorting direction. (optional)
	searchRequestExtensionAttributesProductsInfo0Sku := "searchRequestExtensionAttributesProductsInfo0Sku_example" // string | Product SKU. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryInStorePickupPickupLocationsAPI.GetV1InventoryInstorepickupPickuplocations(context.Background()).SearchRequestAreaRadius(searchRequestAreaRadius).SearchRequestAreaSearchTerm(searchRequestAreaSearchTerm).SearchRequestFiltersCountryValue(searchRequestFiltersCountryValue).SearchRequestFiltersCountryConditionType(searchRequestFiltersCountryConditionType).SearchRequestFiltersPostcodeValue(searchRequestFiltersPostcodeValue).SearchRequestFiltersPostcodeConditionType(searchRequestFiltersPostcodeConditionType).SearchRequestFiltersRegionValue(searchRequestFiltersRegionValue).SearchRequestFiltersRegionConditionType(searchRequestFiltersRegionConditionType).SearchRequestFiltersRegionIdValue(searchRequestFiltersRegionIdValue).SearchRequestFiltersRegionIdConditionType(searchRequestFiltersRegionIdConditionType).SearchRequestFiltersCityValue(searchRequestFiltersCityValue).SearchRequestFiltersCityConditionType(searchRequestFiltersCityConditionType).SearchRequestFiltersStreetValue(searchRequestFiltersStreetValue).SearchRequestFiltersStreetConditionType(searchRequestFiltersStreetConditionType).SearchRequestFiltersNameValue(searchRequestFiltersNameValue).SearchRequestFiltersNameConditionType(searchRequestFiltersNameConditionType).SearchRequestFiltersPickupLocationCodeValue(searchRequestFiltersPickupLocationCodeValue).SearchRequestFiltersPickupLocationCodeConditionType(searchRequestFiltersPickupLocationCodeConditionType).SearchRequestPageSize(searchRequestPageSize).SearchRequestCurrentPage(searchRequestCurrentPage).SearchRequestScopeType(searchRequestScopeType).SearchRequestScopeCode(searchRequestScopeCode).SearchRequestSort0Field(searchRequestSort0Field).SearchRequestSort0Direction(searchRequestSort0Direction).SearchRequestExtensionAttributesProductsInfo0Sku(searchRequestExtensionAttributesProductsInfo0Sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryInStorePickupPickupLocationsAPI.GetV1InventoryInstorepickupPickuplocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryInstorepickupPickuplocations`: InventoryInStorePickupApiDataSearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryInStorePickupPickupLocationsAPI.GetV1InventoryInstorepickupPickuplocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryInstorepickupPickuplocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequestAreaRadius** | **int32** | Search radius in KM. | 
 **searchRequestAreaSearchTerm** | **string** | Search term string. | 
 **searchRequestFiltersCountryValue** | **string** | Value. | 
 **searchRequestFiltersCountryConditionType** | **string** | Condition Type. | 
 **searchRequestFiltersPostcodeValue** | **string** | Value. | 
 **searchRequestFiltersPostcodeConditionType** | **string** | Condition Type. | 
 **searchRequestFiltersRegionValue** | **string** | Value. | 
 **searchRequestFiltersRegionConditionType** | **string** | Condition Type. | 
 **searchRequestFiltersRegionIdValue** | **string** | Value. | 
 **searchRequestFiltersRegionIdConditionType** | **string** | Condition Type. | 
 **searchRequestFiltersCityValue** | **string** | Value. | 
 **searchRequestFiltersCityConditionType** | **string** | Condition Type. | 
 **searchRequestFiltersStreetValue** | **string** | Value. | 
 **searchRequestFiltersStreetConditionType** | **string** | Condition Type. | 
 **searchRequestFiltersNameValue** | **string** | Value. | 
 **searchRequestFiltersNameConditionType** | **string** | Condition Type. | 
 **searchRequestFiltersPickupLocationCodeValue** | **string** | Value. | 
 **searchRequestFiltersPickupLocationCodeConditionType** | **string** | Condition Type. | 
 **searchRequestPageSize** | **int32** | Page size. | 
 **searchRequestCurrentPage** | **int32** | Current page. | 
 **searchRequestScopeType** | **string** | Sales Channel Type. | 
 **searchRequestScopeCode** | **string** | Sales Channel code. | 
 **searchRequestSort0Field** | **string** | Sorting field. | 
 **searchRequestSort0Direction** | **string** | Sorting direction. | 
 **searchRequestExtensionAttributesProductsInfo0Sku** | **string** | Product SKU. | 

### Return type

[**InventoryInStorePickupApiDataSearchResultInterface**](InventoryInStorePickupApiDataSearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

