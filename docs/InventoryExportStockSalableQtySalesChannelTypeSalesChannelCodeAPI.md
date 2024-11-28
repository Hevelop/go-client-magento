# \InventoryExportStockSalableQtySalesChannelTypeSalesChannelCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode**](InventoryExportStockSalableQtySalesChannelTypeSalesChannelCodeAPI.md#GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode) | **Get** /V1/inventory/export-stock-salable-qty/{salesChannelType}/{salesChannelCode} | inventory/export-stock-salable-qty/{salesChannelType}/{salesChannelCode}



## GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode

> InventoryExportStockApiDataExportStockSalableQtySearchResultInterface GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode(ctx, salesChannelType, salesChannelCode).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()

inventory/export-stock-salable-qty/{salesChannelType}/{salesChannelCode}



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
	salesChannelType := "salesChannelType_example" // string | 
	salesChannelCode := "salesChannelCode_example" // string | 
	searchCriteriaFilterGroups0Filters0Field := "searchCriteriaFilterGroups0Filters0Field_example" // string | Field (optional)
	searchCriteriaFilterGroups0Filters0Value := "searchCriteriaFilterGroups0Filters0Value_example" // string | Value (optional)
	searchCriteriaFilterGroups0Filters0ConditionType := "searchCriteriaFilterGroups0Filters0ConditionType_example" // string | Condition type (optional)
	searchCriteriaSortOrders0Field := "searchCriteriaSortOrders0Field_example" // string | Sorting field. (optional)
	searchCriteriaSortOrders0Direction := "searchCriteriaSortOrders0Direction_example" // string | Sorting direction. (optional)
	searchCriteriaPageSize := int32(56) // int32 | Page size. (optional)
	searchCriteriaCurrentPage := int32(56) // int32 | Current page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryExportStockSalableQtySalesChannelTypeSalesChannelCodeAPI.GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode(context.Background(), salesChannelType, salesChannelCode).SearchCriteriaFilterGroups0Filters0Field(searchCriteriaFilterGroups0Filters0Field).SearchCriteriaFilterGroups0Filters0Value(searchCriteriaFilterGroups0Filters0Value).SearchCriteriaFilterGroups0Filters0ConditionType(searchCriteriaFilterGroups0Filters0ConditionType).SearchCriteriaSortOrders0Field(searchCriteriaSortOrders0Field).SearchCriteriaSortOrders0Direction(searchCriteriaSortOrders0Direction).SearchCriteriaPageSize(searchCriteriaPageSize).SearchCriteriaCurrentPage(searchCriteriaCurrentPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryExportStockSalableQtySalesChannelTypeSalesChannelCodeAPI.GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode`: InventoryExportStockApiDataExportStockSalableQtySearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryExportStockSalableQtySalesChannelTypeSalesChannelCodeAPI.GetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**salesChannelType** | **string** |  | 
**salesChannelCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryExportstocksalableqtySalesChannelTypeSalesChannelCodeRequest struct via the builder pattern


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

[**InventoryExportStockApiDataExportStockSalableQtySearchResultInterface**](InventoryExportStockApiDataExportStockSalableQtySearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
