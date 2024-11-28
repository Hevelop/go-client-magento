# \ConfigurableProductsSkuOptionsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ConfigurableproductsSkuOptions**](ConfigurableProductsSkuOptionsAPI.md#PostV1ConfigurableproductsSkuOptions) | **Post** /V1/configurable-products/{sku}/options | configurable-products/{sku}/options



## PostV1ConfigurableproductsSkuOptions

> int32 PostV1ConfigurableproductsSkuOptions(ctx, sku).PostV1ConfigurableproductsSkuOptionsRequest(postV1ConfigurableproductsSkuOptionsRequest).Execute()

configurable-products/{sku}/options



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
	postV1ConfigurableproductsSkuOptionsRequest := *openapiclient.NewPostV1ConfigurableproductsSkuOptionsRequest(*openapiclient.NewConfigurableProductDataOptionInterface()) // PostV1ConfigurableproductsSkuOptionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurableProductsSkuOptionsAPI.PostV1ConfigurableproductsSkuOptions(context.Background(), sku).PostV1ConfigurableproductsSkuOptionsRequest(postV1ConfigurableproductsSkuOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuOptionsAPI.PostV1ConfigurableproductsSkuOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ConfigurableproductsSkuOptions`: int32
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuOptionsAPI.PostV1ConfigurableproductsSkuOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ConfigurableproductsSkuOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ConfigurableproductsSkuOptionsRequest** | [**PostV1ConfigurableproductsSkuOptionsRequest**](PostV1ConfigurableproductsSkuOptionsRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

