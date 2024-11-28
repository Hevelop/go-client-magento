# \AttributeMetadataCustomerAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1AttributeMetadataCustomerAddress**](AttributeMetadataCustomerAddressAPI.md#GetV1AttributeMetadataCustomerAddress) | **Get** /V1/attributeMetadata/customerAddress | attributeMetadata/customerAddress



## GetV1AttributeMetadataCustomerAddress

> []CustomerDataAttributeMetadataInterface GetV1AttributeMetadataCustomerAddress(ctx).Execute()

attributeMetadata/customerAddress



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
	resp, r, err := apiClient.AttributeMetadataCustomerAddressAPI.GetV1AttributeMetadataCustomerAddress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeMetadataCustomerAddressAPI.GetV1AttributeMetadataCustomerAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AttributeMetadataCustomerAddress`: []CustomerDataAttributeMetadataInterface
	fmt.Fprintf(os.Stdout, "Response from `AttributeMetadataCustomerAddressAPI.GetV1AttributeMetadataCustomerAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AttributeMetadataCustomerAddressRequest struct via the builder pattern


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

