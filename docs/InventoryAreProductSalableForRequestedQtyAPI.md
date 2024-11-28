# \InventoryAreProductSalableForRequestedQtyAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryAreproductsalableforrequestedqty**](InventoryAreProductSalableForRequestedQtyAPI.md#GetV1InventoryAreproductsalableforrequestedqty) | **Get** /V1/inventory/are-product-salable-for-requested-qty/ | inventory/are-product-salable-for-requested-qty/



## GetV1InventoryAreproductsalableforrequestedqty

> []InventorySalesApiDataIsProductSalableForRequestedQtyResultInterface GetV1InventoryAreproductsalableforrequestedqty(ctx).StockId(stockId).SkuRequests0Sku(skuRequests0Sku).SkuRequests0Qty(skuRequests0Qty).Execute()

inventory/are-product-salable-for-requested-qty/



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
	skuRequests0Sku := "skuRequests0Sku_example" // string | Product sku. (optional)
	skuRequests0Qty := float32(8.14) // float32 | Product quantity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAreProductSalableForRequestedQtyAPI.GetV1InventoryAreproductsalableforrequestedqty(context.Background()).StockId(stockId).SkuRequests0Sku(skuRequests0Sku).SkuRequests0Qty(skuRequests0Qty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAreProductSalableForRequestedQtyAPI.GetV1InventoryAreproductsalableforrequestedqty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryAreproductsalableforrequestedqty`: []InventorySalesApiDataIsProductSalableForRequestedQtyResultInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryAreProductSalableForRequestedQtyAPI.GetV1InventoryAreproductsalableforrequestedqty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryAreproductsalableforrequestedqtyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockId** | **int32** |  | 
 **skuRequests0Sku** | **string** | Product sku. | 
 **skuRequests0Qty** | **float32** | Product quantity. | 

### Return type

[**[]InventorySalesApiDataIsProductSalableForRequestedQtyResultInterface**](InventorySalesApiDataIsProductSalableForRequestedQtyResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

