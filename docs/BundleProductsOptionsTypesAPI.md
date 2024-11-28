# \BundleProductsOptionsTypesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1BundleproductsOptionsTypes**](BundleProductsOptionsTypesAPI.md#GetV1BundleproductsOptionsTypes) | **Get** /V1/bundle-products/options/types | bundle-products/options/types



## GetV1BundleproductsOptionsTypes

> []BundleDataOptionTypeInterface GetV1BundleproductsOptionsTypes(ctx).Execute()

bundle-products/options/types



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
	resp, r, err := apiClient.BundleProductsOptionsTypesAPI.GetV1BundleproductsOptionsTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleProductsOptionsTypesAPI.GetV1BundleproductsOptionsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1BundleproductsOptionsTypes`: []BundleDataOptionTypeInterface
	fmt.Fprintf(os.Stdout, "Response from `BundleProductsOptionsTypesAPI.GetV1BundleproductsOptionsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BundleproductsOptionsTypesRequest struct via the builder pattern


### Return type

[**[]BundleDataOptionTypeInterface**](BundleDataOptionTypeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

