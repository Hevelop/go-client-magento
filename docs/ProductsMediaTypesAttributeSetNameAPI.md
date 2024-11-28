# \ProductsMediaTypesAttributeSetNameAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsMediaTypesAttributeSetName**](ProductsMediaTypesAttributeSetNameAPI.md#GetV1ProductsMediaTypesAttributeSetName) | **Get** /V1/products/media/types/{attributeSetName} | products/media/types/{attributeSetName}



## GetV1ProductsMediaTypesAttributeSetName

> []CatalogDataProductAttributeInterface GetV1ProductsMediaTypesAttributeSetName(ctx, attributeSetName).Execute()

products/media/types/{attributeSetName}



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
	attributeSetName := "attributeSetName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsMediaTypesAttributeSetNameAPI.GetV1ProductsMediaTypesAttributeSetName(context.Background(), attributeSetName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsMediaTypesAttributeSetNameAPI.GetV1ProductsMediaTypesAttributeSetName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsMediaTypesAttributeSetName`: []CatalogDataProductAttributeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsMediaTypesAttributeSetNameAPI.GetV1ProductsMediaTypesAttributeSetName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsMediaTypesAttributeSetNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CatalogDataProductAttributeInterface**](CatalogDataProductAttributeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

