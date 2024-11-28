# \InventoryAreProductsSalableAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryAreproductssalable**](InventoryAreProductsSalableAPI.md#GetV1InventoryAreproductssalable) | **Get** /V1/inventory/are-products-salable | inventory/are-products-salable



## GetV1InventoryAreproductssalable

> []InventorySalesApiDataIsProductSalableResultInterface GetV1InventoryAreproductssalable(ctx).Skus(skus).StockId(stockId).Execute()

inventory/are-products-salable



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
	skus := []string{"Inner_example"} // []string | 
	stockId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAreProductsSalableAPI.GetV1InventoryAreproductssalable(context.Background()).Skus(skus).StockId(stockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAreProductsSalableAPI.GetV1InventoryAreproductssalable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryAreproductssalable`: []InventorySalesApiDataIsProductSalableResultInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryAreProductsSalableAPI.GetV1InventoryAreproductssalable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryAreproductssalableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skus** | **[]string** |  | 
 **stockId** | **int32** |  | 

### Return type

[**[]InventorySalesApiDataIsProductSalableResultInterface**](InventorySalesApiDataIsProductSalableResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

