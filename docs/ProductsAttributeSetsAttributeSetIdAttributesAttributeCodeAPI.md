# \ProductsAttributeSetsAttributeSetIdAttributesAttributeCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode**](ProductsAttributeSetsAttributeSetIdAttributesAttributeCodeAPI.md#DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode) | **Delete** /V1/products/attribute-sets/{attributeSetId}/attributes/{attributeCode} | products/attribute-sets/{attributeSetId}/attributes/{attributeCode}



## DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode

> bool DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode(ctx, attributeSetId, attributeCode).Execute()

products/attribute-sets/{attributeSetId}/attributes/{attributeCode}



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
	attributeCode := "attributeCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsAttributeSetIdAttributesAttributeCodeAPI.DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode(context.Background(), attributeSetId, attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsAttributeSetIdAttributesAttributeCodeAPI.DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsAttributeSetIdAttributesAttributeCodeAPI.DeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeSetId** | **string** |  | 
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsAttributesetsAttributeSetIdAttributesAttributeCodeRequest struct via the builder pattern


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

