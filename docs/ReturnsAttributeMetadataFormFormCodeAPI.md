# \ReturnsAttributeMetadataFormFormCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ReturnsAttributeMetadataFormFormCode**](ReturnsAttributeMetadataFormFormCodeAPI.md#GetV1ReturnsAttributeMetadataFormFormCode) | **Get** /V1/returnsAttributeMetadata/form/{formCode} | returnsAttributeMetadata/form/{formCode}



## GetV1ReturnsAttributeMetadataFormFormCode

> []CustomerDataAttributeMetadataInterface GetV1ReturnsAttributeMetadataFormFormCode(ctx, formCode).Execute()

returnsAttributeMetadata/form/{formCode}



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
	formCode := "formCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAttributeMetadataFormFormCodeAPI.GetV1ReturnsAttributeMetadataFormFormCode(context.Background(), formCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAttributeMetadataFormFormCodeAPI.GetV1ReturnsAttributeMetadataFormFormCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ReturnsAttributeMetadataFormFormCode`: []CustomerDataAttributeMetadataInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAttributeMetadataFormFormCodeAPI.GetV1ReturnsAttributeMetadataFormFormCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ReturnsAttributeMetadataFormFormCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomerDataAttributeMetadataInterface**](CustomerDataAttributeMetadataInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

