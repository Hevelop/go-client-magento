# \ProductsTypesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsTypes**](ProductsTypesAPI.md#GetV1ProductsTypes) | **Get** /V1/products/types | products/types



## GetV1ProductsTypes

> []CatalogDataProductTypeInterface GetV1ProductsTypes(ctx).Execute()

products/types



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
	resp, r, err := apiClient.ProductsTypesAPI.GetV1ProductsTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsTypesAPI.GetV1ProductsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsTypes`: []CatalogDataProductTypeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsTypesAPI.GetV1ProductsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsTypesRequest struct via the builder pattern


### Return type

[**[]CatalogDataProductTypeInterface**](CatalogDataProductTypeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

