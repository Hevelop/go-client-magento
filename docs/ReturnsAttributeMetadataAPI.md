# \ReturnsAttributeMetadataAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ReturnsAttributeMetadata**](ReturnsAttributeMetadataAPI.md#GetV1ReturnsAttributeMetadata) | **Get** /V1/returnsAttributeMetadata | returnsAttributeMetadata



## GetV1ReturnsAttributeMetadata

> []CustomerDataAttributeMetadataInterface GetV1ReturnsAttributeMetadata(ctx).Execute()

returnsAttributeMetadata



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsAttributeMetadataAPI.GetV1ReturnsAttributeMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsAttributeMetadataAPI.GetV1ReturnsAttributeMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ReturnsAttributeMetadata`: []CustomerDataAttributeMetadataInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsAttributeMetadataAPI.GetV1ReturnsAttributeMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ReturnsAttributeMetadataRequest struct via the builder pattern


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

