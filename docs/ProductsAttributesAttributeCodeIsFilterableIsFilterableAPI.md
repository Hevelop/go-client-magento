# \ProductsAttributesAttributeCodeIsFilterableIsFilterableAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable**](ProductsAttributesAttributeCodeIsFilterableIsFilterableAPI.md#PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable) | **Put** /V1/products/attributes/{attributeCode}/is-filterable/{isFilterable} | products/attributes/{attributeCode}/is-filterable/{isFilterable}



## PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable

> bool PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable(ctx, attributeCode, isFilterable).Execute()

products/attributes/{attributeCode}/is-filterable/{isFilterable}



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
	attributeCode := "attributeCode_example" // string | 
	isFilterable := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributesAttributeCodeIsFilterableIsFilterableAPI.PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable(context.Background(), attributeCode, isFilterable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributesAttributeCodeIsFilterableIsFilterableAPI.PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributesAttributeCodeIsFilterableIsFilterableAPI.PutV1ProductsAttributesAttributeCodeIsfilterableIsFilterable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 
**isFilterable** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1ProductsAttributesAttributeCodeIsfilterableIsFilterableRequest struct via the builder pattern


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

