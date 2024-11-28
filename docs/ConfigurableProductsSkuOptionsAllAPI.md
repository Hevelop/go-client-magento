# \ConfigurableProductsSkuOptionsAllAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ConfigurableproductsSkuOptionsAll**](ConfigurableProductsSkuOptionsAllAPI.md#GetV1ConfigurableproductsSkuOptionsAll) | **Get** /V1/configurable-products/{sku}/options/all | configurable-products/{sku}/options/all



## GetV1ConfigurableproductsSkuOptionsAll

> []ConfigurableProductDataOptionInterface GetV1ConfigurableproductsSkuOptionsAll(ctx, sku).Execute()

configurable-products/{sku}/options/all



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
	resp, r, err := apiClient.ConfigurableProductsSkuOptionsAllAPI.GetV1ConfigurableproductsSkuOptionsAll(context.Background(), sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuOptionsAllAPI.GetV1ConfigurableproductsSkuOptionsAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ConfigurableproductsSkuOptionsAll`: []ConfigurableProductDataOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuOptionsAllAPI.GetV1ConfigurableproductsSkuOptionsAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ConfigurableproductsSkuOptionsAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigurableProductDataOptionInterface**](ConfigurableProductDataOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

