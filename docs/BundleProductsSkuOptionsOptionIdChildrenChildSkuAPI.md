# \BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku**](BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI.md#DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku) | **Delete** /V1/bundle-products/{sku}/options/{optionId}/children/{childSku} | bundle-products/{sku}/options/{optionId}/children/{childSku}



## DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku

> bool DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku(ctx, sku, optionId, childSku).Execute()

bundle-products/{sku}/options/{optionId}/children/{childSku}



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
	optionId := int32(56) // int32 | 
	childSku := "childSku_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI.DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku(context.Background(), sku, optionId, childSku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI.DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku`: bool
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsSkuOptionsOptionIdChildrenChildSkuAPI.DeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**optionId** | **int32** |  | 
**childSku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1BundleproductsSkuOptionsOptionIdChildrenChildSkuRequest struct via the builder pattern


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

