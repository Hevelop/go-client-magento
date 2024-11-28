# \ReturnsAttributeMetadataCustomAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ReturnsAttributeMetadataCustom**](ReturnsAttributeMetadataCustomAPI.md#GetV1ReturnsAttributeMetadataCustom) | **Get** /V1/returnsAttributeMetadata/custom | returnsAttributeMetadata/custom



## GetV1ReturnsAttributeMetadataCustom

> []FrameworkMetadataObjectInterface GetV1ReturnsAttributeMetadataCustom(ctx).DataObjectClassName(dataObjectClassName).Execute()

returnsAttributeMetadata/custom



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
	dataObjectClassName := "dataObjectClassName_example" // string | Data object class name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAttributeMetadataCustomAPI.GetV1ReturnsAttributeMetadataCustom(context.Background()).DataObjectClassName(dataObjectClassName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAttributeMetadataCustomAPI.GetV1ReturnsAttributeMetadataCustom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ReturnsAttributeMetadataCustom`: []FrameworkMetadataObjectInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAttributeMetadataCustomAPI.GetV1ReturnsAttributeMetadataCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ReturnsAttributeMetadataCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataObjectClassName** | **string** | Data object class name | 

### Return type

[**[]FrameworkMetadataObjectInterface**](FrameworkMetadataObjectInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

