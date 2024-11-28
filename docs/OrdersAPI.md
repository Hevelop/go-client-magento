# \OrdersAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1Orders**](OrdersAPI.md#GetV1Orders) | **Get** /V1/orders | orders
[**PostV1Orders**](OrdersAPI.md#PostV1Orders) | **Post** /V1/orders | orders



## GetV1Orders

> SalesDataOrderSearchResultInterface GetV1Orders(ctx).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()

orders



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
	resp, r, err := apiClient.OrdersAPI.GetV1Orders(context.Background()).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetV1Orders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1Orders`: SalesDataOrderSearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetV1Orders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1OrdersRequest struct via the builder pattern


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

[**SalesDataOrderSearchResultInterface**](SalesDataOrderSearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1Orders

> SalesDataOrderInterface PostV1Orders(ctx).PostV1OrdersRequest(postV1OrdersRequest).Execute()

orders



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
	postV1OrdersRequest := *openapiclient.NewPostV1OrdersRequest(*openapiclient.NewSalesDataOrderInterface(float32(123), "CustomerEmail_example", float32(123), []openapiclient.SalesDataOrderItemInterface{*openapiclient.NewSalesDataOrderItemInterface("Sku_example")})) // PostV1OrdersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.PostV1Orders(context.Background()).PostV1OrdersRequest(postV1OrdersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PostV1Orders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Orders`: SalesDataOrderInterface
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.PostV1Orders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1OrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1OrdersRequest** | [**PostV1OrdersRequest**](PostV1OrdersRequest.md) |  | 

### Return type

[**SalesDataOrderInterface**](SalesDataOrderInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

