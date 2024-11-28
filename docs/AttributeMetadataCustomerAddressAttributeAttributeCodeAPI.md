# \AttributeMetadataCustomerAddressAttributeAttributeCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1AttributeMetadataCustomerAddressAttributeAttributeCode**](AttributeMetadataCustomerAddressAttributeAttributeCodeAPI.md#GetV1AttributeMetadataCustomerAddressAttributeAttributeCode) | **Get** /V1/attributeMetadata/customerAddress/attribute/{attributeCode} | attributeMetadata/customerAddress/attribute/{attributeCode}



## GetV1AttributeMetadataCustomerAddressAttributeAttributeCode

> CustomerDataAttributeMetadataInterface GetV1AttributeMetadataCustomerAddressAttributeAttributeCode(ctx, attributeCode).Execute()

attributeMetadata/customerAddress/attribute/{attributeCode}



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
	resp, r, err := apiClient.AttributeMetadataCustomerAddressAttributeAttributeCodeAPI.GetV1AttributeMetadataCustomerAddressAttributeAttributeCode(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeMetadataCustomerAddressAttributeAttributeCodeAPI.GetV1AttributeMetadataCustomerAddressAttributeAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AttributeMetadataCustomerAddressAttributeAttributeCode`: CustomerDataAttributeMetadataInterface
	fmt.Fprintf(os.Stdout, "Response from `AttributeMetadataCustomerAddressAttributeAttributeCodeAPI.GetV1AttributeMetadataCustomerAddressAttributeAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AttributeMetadataCustomerAddressAttributeAttributeCodeRequest struct via the builder pattern


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

