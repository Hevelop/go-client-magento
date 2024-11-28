# \InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId**](InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI.md#GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId) | **Get** /V1/inventory/get-sources-assigned-to-stock-ordered-by-priority/{stockId} | inventory/get-sources-assigned-to-stock-ordered-by-priority/{stockId}



## GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId

> []InventoryApiDataSourceInterface GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId(ctx, stockId).Execute()

inventory/get-sources-assigned-to-stock-ordered-by-priority/{stockId}



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
	stockId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI.GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId(context.Background(), stockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI.GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId`: []InventoryApiDataSourceInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryGetSourcesAssignedToStockOrderedByPriorityStockIdAPI.GetV1InventoryGetsourcesassignedtostockorderedbypriorityStockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryGetsourcesassignedtostockorderedbypriorityStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InventoryApiDataSourceInterface**](InventoryApiDataSourceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

