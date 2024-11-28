# \BundleProductsProductSkuChildrenAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1BundleproductsProductSkuChildren**](BundleProductsProductSkuChildrenAPI.md#GetV1BundleproductsProductSkuChildren) | **Get** /V1/bundle-products/{productSku}/children | bundle-products/{productSku}/children



## GetV1BundleproductsProductSkuChildren

> []BundleDataLinkInterface GetV1BundleproductsProductSkuChildren(ctx, productSku).OptionId(optionId).Execute()

bundle-products/{productSku}/children



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
	productSku := "productSku_example" // string | 
	optionId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleProductsProductSkuChildrenAPI.GetV1BundleproductsProductSkuChildren(context.Background(), productSku).OptionId(optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsProductSkuChildrenAPI.GetV1BundleproductsProductSkuChildren``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1BundleproductsProductSkuChildren`: []BundleDataLinkInterface
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsProductSkuChildrenAPI.GetV1BundleproductsProductSkuChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productSku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BundleproductsProductSkuChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **optionId** | **int32** |  | 

### Return type

[**[]BundleDataLinkInterface**](BundleDataLinkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

