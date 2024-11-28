# \ProductsLinksTypesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsLinksTypes**](ProductsLinksTypesAPI.md#GetV1ProductsLinksTypes) | **Get** /V1/products/links/types | products/links/types



## GetV1ProductsLinksTypes

> []CatalogDataProductLinkTypeInterface GetV1ProductsLinksTypes(ctx).Execute()

products/links/types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsLinksTypesAPI.GetV1ProductsLinksTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsLinksTypesAPI.GetV1ProductsLinksTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsLinksTypes`: []CatalogDataProductLinkTypeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsLinksTypesAPI.GetV1ProductsLinksTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsLinksTypesRequest struct via the builder pattern


### Return type

[**[]CatalogDataProductLinkTypeInterface**](CatalogDataProductLinkTypeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

