# \InventoryIsProductSalableForRequestedQtySkuStockIdRequestedQtyAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty**](InventoryIsProductSalableForRequestedQtySkuStockIdRequestedQtyAPI.md#GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty) | **Get** /V1/inventory/is-product-salable-for-requested-qty/{sku}/{stockId}/{requestedQty} | inventory/is-product-salable-for-requested-qty/{sku}/{stockId}/{requestedQty}



## GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty

> InventorySalesApiDataProductSalableResultInterface GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty(ctx, sku, stockId, requestedQty).Execute()

inventory/is-product-salable-for-requested-qty/{sku}/{stockId}/{requestedQty}



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
	sku := "sku_example" // string | 
	stockId := int32(56) // int32 | 
	requestedQty := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryIsProductSalableForRequestedQtySkuStockIdRequestedQtyAPI.GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty(context.Background(), sku, stockId, requestedQty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryIsProductSalableForRequestedQtySkuStockIdRequestedQtyAPI.GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty`: InventorySalesApiDataProductSalableResultInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryIsProductSalableForRequestedQtySkuStockIdRequestedQtyAPI.GetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**stockId** | **int32** |  | 
**requestedQty** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryIsproductsalableforrequestedqtySkuStockIdRequestedQtyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InventorySalesApiDataProductSalableResultInterface**](InventorySalesApiDataProductSalableResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

