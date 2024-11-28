# \ConfigurableProductsSkuChildrenChildSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ConfigurableproductsSkuChildrenChildSku**](ConfigurableProductsSkuChildrenChildSkuAPI.md#DeleteV1ConfigurableproductsSkuChildrenChildSku) | **Delete** /V1/configurable-products/{sku}/children/{childSku} | configurable-products/{sku}/children/{childSku}



## DeleteV1ConfigurableproductsSkuChildrenChildSku

> bool DeleteV1ConfigurableproductsSkuChildrenChildSku(ctx, sku, childSku).Execute()

configurable-products/{sku}/children/{childSku}



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
	childSku := "childSku_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurableProductsSkuChildrenChildSkuAPI.DeleteV1ConfigurableproductsSkuChildrenChildSku(context.Background(), sku, childSku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuChildrenChildSkuAPI.DeleteV1ConfigurableproductsSkuChildrenChildSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ConfigurableproductsSkuChildrenChildSku`: bool
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuChildrenChildSkuAPI.DeleteV1ConfigurableproductsSkuChildrenChildSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**childSku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ConfigurableproductsSkuChildrenChildSkuRequest struct via the builder pattern


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

