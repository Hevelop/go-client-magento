# \ProductsSkuLinksTypeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsSkuLinksType**](ProductsSkuLinksTypeAPI.md#GetV1ProductsSkuLinksType) | **Get** /V1/products/{sku}/links/{type} | products/{sku}/links/{type}



## GetV1ProductsSkuLinksType

> []CatalogDataProductLinkInterface GetV1ProductsSkuLinksType(ctx, sku, type_).Execute()

products/{sku}/links/{type}



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsSkuLinksTypeAPI.GetV1ProductsSkuLinksType(context.Background(), sku, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsSkuLinksTypeAPI.GetV1ProductsSkuLinksType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsSkuLinksType`: []CatalogDataProductLinkInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsSkuLinksTypeAPI.GetV1ProductsSkuLinksType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sku** | **string** |  | 
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsSkuLinksTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CatalogDataProductLinkInterface**](CatalogDataProductLinkInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

