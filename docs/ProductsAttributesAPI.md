# \ProductsAttributesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsAttributes**](ProductsAttributesAPI.md#GetV1ProductsAttributes) | **Get** /V1/products/attributes | products/attributes
[**PostV1ProductsAttributes**](ProductsAttributesAPI.md#PostV1ProductsAttributes) | **Post** /V1/products/attributes | products/attributes



## GetV1ProductsAttributes

> CatalogDataProductAttributeSearchResultsInterface GetV1ProductsAttributes(ctx).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()

products/attributes



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
	searchCriteriaFilterGroups0Filters0Field := "searchCriteriaFilterGroups0Filters0Field_example" // string | Field (optional)
	searchCriteriaFilterGroups0Filters0Value := "searchCriteriaFilterGroups0Filters0Value_example" // string | Value (optional)
	searchCriteriaFilterGroups0Filters0ConditionType := "searchCriteriaFilterGroups0Filters0ConditionType_example" // string | Condition type (optional)
	searchCriteriaSortOrders0Field := "searchCriteriaSortOrders0Field_example" // string | Sorting field. (optional)
	searchCriteriaSortOrders0Direction := "searchCriteriaSortOrders0Direction_example" // string | Sorting direction. (optional)
	searchCriteriaPageSize := int32(56) // int32 | Page size. (optional)
	searchCriteriaCurrentPage := int32(56) // int32 | Current page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAPI.GetV1ProductsAttributes(context.Background()).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAPI.GetV1ProductsAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsAttributes`: CatalogDataProductAttributeSearchResultsInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAPI.GetV1ProductsAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchCriteriaFilterGroups0Filters0Field** | **string** | Field | 
 **searchCriteriaFilterGroups0Filters0Value** | **string** | Value | 
 **searchCriteriaFilterGroups0Filters0ConditionType** | **string** | Condition type | 
 **searchCriteriaSortOrders0Field** | **string** | Sorting field. | 
 **searchCriteriaSortOrders0Direction** | **string** | Sorting direction. | 
 **searchCriteriaPageSize** | **int32** | Page size. | 
 **searchCriteriaCurrentPage** | **int32** | Current page. | 

### Return type

[**CatalogDataProductAttributeSearchResultsInterface**](CatalogDataProductAttributeSearchResultsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ProductsAttributes

> CatalogDataProductAttributeInterface PostV1ProductsAttributes(ctx).PostV1ProductsAttributesRequest(postV1ProductsAttributesRequest).Execute()

products/attributes



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
	postV1ProductsAttributesRequest := *openapiclient.NewPostV1ProductsAttributesRequest(*openapiclient.NewCatalogDataProductAttributeInterface("AttributeCode_example", "FrontendInput_example", "EntityTypeId_example", false, []openapiclient.EavDataAttributeFrontendLabelInterface{*openapiclient.NewEavDataAttributeFrontendLabelInterface()})) // PostV1ProductsAttributesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAPI.PostV1ProductsAttributes(context.Background()).PostV1ProductsAttributesRequest(postV1ProductsAttributesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAPI.PostV1ProductsAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsAttributes`: CatalogDataProductAttributeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAPI.PostV1ProductsAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsAttributesRequest** | [**PostV1ProductsAttributesRequest**](PostV1ProductsAttributesRequest.md) |  | 

### Return type

[**CatalogDataProductAttributeInterface**](CatalogDataProductAttributeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

