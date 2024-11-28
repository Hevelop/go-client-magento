# \ConfigurableProductsSkuChildrenAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ConfigurableproductsSkuChildren**](ConfigurableProductsSkuChildrenAPI.md#GetV1ConfigurableproductsSkuChildren) | **Get** /V1/configurable-products/{sku}/children | configurable-products/{sku}/children



## GetV1ConfigurableproductsSkuChildren

> []CatalogDataProductInterface GetV1ConfigurableproductsSkuChildren(ctx, sku).Execute()

configurable-products/{sku}/children



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
	resp, r, err := apiClient.ConfigurableProductsSkuChildrenAPI.GetV1ConfigurableproductsSkuChildren(context.Background(), sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurableProductsSkuChildrenAPI.GetV1ConfigurableproductsSkuChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ConfigurableproductsSkuChildren`: []CatalogDataProductInterface
	fmt.Fprintf(os.Stdout, "Response from `ConfigurableProductsSkuChildrenAPI.GetV1ConfigurableproductsSkuChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ConfigurableproductsSkuChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogDataProductInterface**](CatalogDataProductInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

