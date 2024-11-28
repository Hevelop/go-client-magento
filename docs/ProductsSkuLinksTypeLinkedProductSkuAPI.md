# \ProductsSkuLinksTypeLinkedProductSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsSkuLinksTypeLinkedProductSku**](ProductsSkuLinksTypeLinkedProductSkuAPI.md#DeleteV1ProductsSkuLinksTypeLinkedProductSku) | **Delete** /V1/products/{sku}/links/{type}/{linkedProductSku} | products/{sku}/links/{type}/{linkedProductSku}



## DeleteV1ProductsSkuLinksTypeLinkedProductSku

> bool DeleteV1ProductsSkuLinksTypeLinkedProductSku(ctx, sku, type_, linkedProductSku).Execute()

products/{sku}/links/{type}/{linkedProductSku}



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
	type_ := "type__example" // string | 
	linkedProductSku := "linkedProductSku_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuLinksTypeLinkedProductSkuAPI.DeleteV1ProductsSkuLinksTypeLinkedProductSku(context.Background(), sku, type_, linkedProductSku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuLinksTypeLinkedProductSkuAPI.DeleteV1ProductsSkuLinksTypeLinkedProductSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsSkuLinksTypeLinkedProductSku`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuLinksTypeLinkedProductSkuAPI.DeleteV1ProductsSkuLinksTypeLinkedProductSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**type_** | **string** |  | 
**linkedProductSku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsSkuLinksTypeLinkedProductSkuRequest struct via the builder pattern


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

