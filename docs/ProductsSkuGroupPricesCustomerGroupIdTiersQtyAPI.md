# \ProductsSkuGroupPricesCustomerGroupIdTiersQtyAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty**](ProductsSkuGroupPricesCustomerGroupIdTiersQtyAPI.md#DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty) | **Delete** /V1/products/{sku}/group-prices/{customerGroupId}/tiers/{qty} | products/{sku}/group-prices/{customerGroupId}/tiers/{qty}



## DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty

> bool DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty(ctx, sku, customerGroupId, qty).Execute()

products/{sku}/group-prices/{customerGroupId}/tiers/{qty}



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
	customerGroupId := "customerGroupId_example" // string | 'all' can be used to specify 'ALL GROUPS'
	qty := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuGroupPricesCustomerGroupIdTiersQtyAPI.DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty(context.Background(), sku, customerGroupId, qty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuGroupPricesCustomerGroupIdTiersQtyAPI.DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuGroupPricesCustomerGroupIdTiersQtyAPI.DeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**customerGroupId** | **string** | &#39;all&#39; can be used to specify &#39;ALL GROUPS&#39; | 
**qty** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsSkuGrouppricesCustomerGroupIdTiersQtyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

