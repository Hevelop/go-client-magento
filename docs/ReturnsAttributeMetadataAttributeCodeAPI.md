# \ReturnsAttributeMetadataAttributeCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ReturnsAttributeMetadataAttributeCode**](ReturnsAttributeMetadataAttributeCodeAPI.md#GetV1ReturnsAttributeMetadataAttributeCode) | **Get** /V1/returnsAttributeMetadata/{attributeCode} | returnsAttributeMetadata/{attributeCode}



## GetV1ReturnsAttributeMetadataAttributeCode

> CustomerDataAttributeMetadataInterface GetV1ReturnsAttributeMetadataAttributeCode(ctx, attributeCode).Execute()

returnsAttributeMetadata/{attributeCode}



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
	resp, r, err := apiClient.ReturnsAttributeMetadataAttributeCodeAPI.GetV1ReturnsAttributeMetadataAttributeCode(context.Background(), attributeCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAttributeMetadataAttributeCodeAPI.GetV1ReturnsAttributeMetadataAttributeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ReturnsAttributeMetadataAttributeCode`: CustomerDataAttributeMetadataInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAttributeMetadataAttributeCodeAPI.GetV1ReturnsAttributeMetadataAttributeCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ReturnsAttributeMetadataAttributeCodeRequest struct via the builder pattern


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

