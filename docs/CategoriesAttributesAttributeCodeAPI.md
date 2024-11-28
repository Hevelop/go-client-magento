# \CategoriesAttributesAttributeCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CategoriesAttributesAttributeCode**](CategoriesAttributesAttributeCodeAPI.md#GetV1CategoriesAttributesAttributeCode) | **Get** /V1/categories/attributes/{attributeCode} | categories/attributes/{attributeCode}



## GetV1CategoriesAttributesAttributeCode

> CatalogDataCategoryAttributeInterface GetV1CategoriesAttributesAttributeCode(ctx, attributeCode).Execute()

categories/attributes/{attributeCode}



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
	resp, r, err := apiClient.CategoriesAttributesAttributeCodeAPI.GetV1CategoriesAttributesAttributeCode(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAttributesAttributeCodeAPI.GetV1CategoriesAttributesAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CategoriesAttributesAttributeCode`: CatalogDataCategoryAttributeInterface
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAttributesAttributeCodeAPI.GetV1CategoriesAttributesAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CategoriesAttributesAttributeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CatalogDataCategoryAttributeInterface**](CatalogDataCategoryAttributeInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

