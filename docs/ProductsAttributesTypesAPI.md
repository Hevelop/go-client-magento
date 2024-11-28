# \ProductsAttributesTypesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsAttributesTypes**](ProductsAttributesTypesAPI.md#GetV1ProductsAttributesTypes) | **Get** /V1/products/attributes/types | products/attributes/types



## GetV1ProductsAttributesTypes

> []CatalogDataProductAttributeTypeInterface GetV1ProductsAttributesTypes(ctx).Execute()

products/attributes/types



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
	resp, r, err := apiClient.ProductsAttributesTypesAPI.GetV1ProductsAttributesTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesTypesAPI.GetV1ProductsAttributesTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsAttributesTypes`: []CatalogDataProductAttributeTypeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesTypesAPI.GetV1ProductsAttributesTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsAttributesTypesRequest struct via the builder pattern


### Return type

[**[]CatalogDataProductAttributeTypeInterface**](CatalogDataProductAttributeTypeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
