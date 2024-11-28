# \BundleProductsSkuOptionsAllAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1BundleproductsSkuOptionsAll**](BundleProductsSkuOptionsAllAPI.md#GetV1BundleproductsSkuOptionsAll) | **Get** /V1/bundle-products/{sku}/options/all | bundle-products/{sku}/options/all



## GetV1BundleproductsSkuOptionsAll

> []BundleDataOptionInterface GetV1BundleproductsSkuOptionsAll(ctx, sku).Execute()

bundle-products/{sku}/options/all



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsSkuOptionsAllAPI.GetV1BundleproductsSkuOptionsAll(context.Background(), sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsSkuOptionsAllAPI.GetV1BundleproductsSkuOptionsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1BundleproductsSkuOptionsAll`: []BundleDataOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsSkuOptionsAllAPI.GetV1BundleproductsSkuOptionsAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BundleproductsSkuOptionsAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BundleDataOptionInterface**](BundleDataOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

