# \CategoriesAttributesAttributeCodeOptionsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CategoriesAttributesAttributeCodeOptions**](CategoriesAttributesAttributeCodeOptionsAPI.md#GetV1CategoriesAttributesAttributeCodeOptions) | **Get** /V1/categories/attributes/{attributeCode}/options | categories/attributes/{attributeCode}/options



## GetV1CategoriesAttributesAttributeCodeOptions

> []EavDataAttributeOptionInterface GetV1CategoriesAttributesAttributeCodeOptions(ctx, attributeCode).Execute()

categories/attributes/{attributeCode}/options



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAttributesAttributeCodeOptionsAPI.GetV1CategoriesAttributesAttributeCodeOptions(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAttributesAttributeCodeOptionsAPI.GetV1CategoriesAttributesAttributeCodeOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CategoriesAttributesAttributeCodeOptions`: []EavDataAttributeOptionInterface
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAttributesAttributeCodeOptionsAPI.GetV1CategoriesAttributesAttributeCodeOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CategoriesAttributesAttributeCodeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EavDataAttributeOptionInterface**](EavDataAttributeOptionInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

