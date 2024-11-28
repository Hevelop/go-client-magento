# \ProductsOptionsTypesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsOptionsTypes**](ProductsOptionsTypesAPI.md#GetV1ProductsOptionsTypes) | **Get** /V1/products/options/types | products/options/types



## GetV1ProductsOptionsTypes

> []CatalogDataProductCustomOptionTypeInterface GetV1ProductsOptionsTypes(ctx).Execute()

products/options/types



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
	resp, r, err := apiClient.ProductsOptionsTypesAPI.GetV1ProductsOptionsTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsOptionsTypesAPI.GetV1ProductsOptionsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsOptionsTypes`: []CatalogDataProductCustomOptionTypeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsOptionsTypesAPI.GetV1ProductsOptionsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsOptionsTypesRequest struct via the builder pattern


### Return type

[**[]CatalogDataProductCustomOptionTypeInterface**](CatalogDataProductCustomOptionTypeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

