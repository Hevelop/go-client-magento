# \StockItemsLowStockAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1StockItemsLowStock**](StockItemsLowStockAPI.md#GetV1StockItemsLowStock) | **Get** /V1/stockItems/lowStock/ | stockItems/lowStock/



## GetV1StockItemsLowStock

> CatalogInventoryDataStockItemCollectionInterface GetV1StockItemsLowStock(ctx).ScopeId(scopeId).Qty(qty).CurrentPage(currentPage).PageSize(pageSize).Execute()

stockItems/lowStock/



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
	scopeId := int32(56) // int32 | 
	qty := float32(8.14) // float32 | 
	currentPage := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StockItemsLowStockAPI.GetV1StockItemsLowStock(context.Background()).ScopeId(scopeId).Qty(qty).CurrentPage(currentPage).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StockItemsLowStockAPI.GetV1StockItemsLowStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1StockItemsLowStock`: CatalogInventoryDataStockItemCollectionInterface
	fmt.Fprintf(os.Stdout, "Response from `StockItemsLowStockAPI.GetV1StockItemsLowStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1StockItemsLowStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeId** | **int32** |  | 
 **qty** | **float32** |  | 
 **currentPage** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**CatalogInventoryDataStockItemCollectionInterface**](CatalogInventoryDataStockItemCollectionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

