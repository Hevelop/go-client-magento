# \AttributeMetadataCustomerAttributeAttributeCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1AttributeMetadataCustomerAttributeAttributeCode**](AttributeMetadataCustomerAttributeAttributeCodeAPI.md#GetV1AttributeMetadataCustomerAttributeAttributeCode) | **Get** /V1/attributeMetadata/customer/attribute/{attributeCode} | attributeMetadata/customer/attribute/{attributeCode}



## GetV1AttributeMetadataCustomerAttributeAttributeCode

> CustomerDataAttributeMetadataInterface GetV1AttributeMetadataCustomerAttributeAttributeCode(ctx, attributeCode).Execute()

attributeMetadata/customer/attribute/{attributeCode}



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
	resp, r, err := apiClient.AttributeMetadataCustomerAttributeAttributeCodeAPI.GetV1AttributeMetadataCustomerAttributeAttributeCode(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeMetadataCustomerAttributeAttributeCodeAPI.GetV1AttributeMetadataCustomerAttributeAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AttributeMetadataCustomerAttributeAttributeCode`: CustomerDataAttributeMetadataInterface
	fmt.Fprintf(os.Stdout, "Response from `AttributeMetadataCustomerAttributeAttributeCodeAPI.GetV1AttributeMetadataCustomerAttributeAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AttributeMetadataCustomerAttributeAttributeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDataAttributeMetadataInterface**](CustomerDataAttributeMetadataInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

