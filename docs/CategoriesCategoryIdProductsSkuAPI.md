# \CategoriesCategoryIdProductsSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CategoriesCategoryIdProductsSku**](CategoriesCategoryIdProductsSkuAPI.md#DeleteV1CategoriesCategoryIdProductsSku) | **Delete** /V1/categories/{categoryId}/products/{sku} | categories/{categoryId}/products/{sku}



## DeleteV1CategoriesCategoryIdProductsSku

> bool DeleteV1CategoriesCategoryIdProductsSku(ctx, categoryId, sku).Execute()

categories/{categoryId}/products/{sku}



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
	categoryId := int32(56) // int32 | 
	sku := "sku_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesCategoryIdProductsSkuAPI.DeleteV1CategoriesCategoryIdProductsSku(context.Background(), categoryId, sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesCategoryIdProductsSkuAPI.DeleteV1CategoriesCategoryIdProductsSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CategoriesCategoryIdProductsSku`: bool
	fmt.Fprintf(os.Stdout, "Response from `CategoriesCategoryIdProductsSkuAPI.DeleteV1CategoriesCategoryIdProductsSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **int32** |  | 
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CategoriesCategoryIdProductsSkuRequest struct via the builder pattern


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

