# \ConfigurableProductsSkuChildAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ConfigurableproductsSkuChild**](ConfigurableProductsSkuChildAPI.md#PostV1ConfigurableproductsSkuChild) | **Post** /V1/configurable-products/{sku}/child | configurable-products/{sku}/child



## PostV1ConfigurableproductsSkuChild

> bool PostV1ConfigurableproductsSkuChild(ctx, sku).PostV1ConfigurableproductsSkuChildRequest(postV1ConfigurableproductsSkuChildRequest).Execute()

configurable-products/{sku}/child



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
	postV1ConfigurableproductsSkuChildRequest := *openapiclient.NewPostV1ConfigurableproductsSkuChildRequest("ChildSku_example") // PostV1ConfigurableproductsSkuChildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurableProductsSkuChildAPI.PostV1ConfigurableproductsSkuChild(context.Background(), sku).PostV1ConfigurableproductsSkuChildRequest(postV1ConfigurableproductsSkuChildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuChildAPI.PostV1ConfigurableproductsSkuChild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ConfigurableproductsSkuChild`: bool
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuChildAPI.PostV1ConfigurableproductsSkuChild`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ConfigurableproductsSkuChildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ConfigurableproductsSkuChildRequest** | [**PostV1ConfigurableproductsSkuChildRequest**](PostV1ConfigurableproductsSkuChildRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

