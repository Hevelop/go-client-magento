# \ProductsLinksTypeAttributesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsLinksTypeAttributes**](ProductsLinksTypeAttributesAPI.md#GetV1ProductsLinksTypeAttributes) | **Get** /V1/products/links/{type}/attributes | products/links/{type}/attributes



## GetV1ProductsLinksTypeAttributes

> []CatalogDataProductLinkAttributeInterface GetV1ProductsLinksTypeAttributes(ctx, type_).Execute()

products/links/{type}/attributes



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
	type_ := "type__example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsLinksTypeAttributesAPI.GetV1ProductsLinksTypeAttributes(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsLinksTypeAttributesAPI.GetV1ProductsLinksTypeAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsLinksTypeAttributes`: []CatalogDataProductLinkAttributeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsLinksTypeAttributesAPI.GetV1ProductsLinksTypeAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsLinksTypeAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogDataProductLinkAttributeInterface**](CatalogDataProductLinkAttributeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

