# \ProductsAttributeSetsAttributeSetIdAttributesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ProductsAttributesetsAttributeSetIdAttributes**](ProductsAttributeSetsAttributeSetIdAttributesAPI.md#GetV1ProductsAttributesetsAttributeSetIdAttributes) | **Get** /V1/products/attribute-sets/{attributeSetId}/attributes | products/attribute-sets/{attributeSetId}/attributes



## GetV1ProductsAttributesetsAttributeSetIdAttributes

> []CatalogDataProductAttributeInterface GetV1ProductsAttributesetsAttributeSetIdAttributes(ctx, attributeSetId).Execute()

products/attribute-sets/{attributeSetId}/attributes



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
	attributeSetId := "attributeSetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsAttributeSetIdAttributesAPI.GetV1ProductsAttributesetsAttributeSetIdAttributes(context.Background(), attributeSetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsAttributeSetIdAttributesAPI.GetV1ProductsAttributesetsAttributeSetIdAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ProductsAttributesetsAttributeSetIdAttributes`: []CatalogDataProductAttributeInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsAttributeSetIdAttributesAPI.GetV1ProductsAttributesetsAttributeSetIdAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ProductsAttributesetsAttributeSetIdAttributesRequest struct via the builder pattern


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

