# \ProductsOptionsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsOptions**](ProductsOptionsAPI.md#PostV1ProductsOptions) | **Post** /V1/products/options | products/options



## PostV1ProductsOptions

> CatalogDataProductCustomOptionInterface PostV1ProductsOptions(ctx).PostV1ProductsOptionsRequest(postV1ProductsOptionsRequest).Execute()

products/options



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
	postV1ProductsOptionsRequest := *openapiclient.NewPostV1ProductsOptionsRequest(*openapiclient.NewCatalogDataProductCustomOptionInterface("ProductSku_example", "Title_example", "Type_example", int32(123), false)) // PostV1ProductsOptionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsOptionsAPI.PostV1ProductsOptions(context.Background()).PostV1ProductsOptionsRequest(postV1ProductsOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsOptionsAPI.PostV1ProductsOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsOptions`: CatalogDataProductCustomOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsOptionsAPI.PostV1ProductsOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsOptionsRequest** | [**PostV1ProductsOptionsRequest**](PostV1ProductsOptionsRequest.md) |  | 

### Return type

[**CatalogDataProductCustomOptionInterface**](CatalogDataProductCustomOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

