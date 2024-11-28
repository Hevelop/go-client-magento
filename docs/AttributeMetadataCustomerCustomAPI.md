# \AttributeMetadataCustomerCustomAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1AttributeMetadataCustomerCustom**](AttributeMetadataCustomerCustomAPI.md#GetV1AttributeMetadataCustomerCustom) | **Get** /V1/attributeMetadata/customer/custom | attributeMetadata/customer/custom



## GetV1AttributeMetadataCustomerCustom

> []CustomerDataAttributeMetadataInterface GetV1AttributeMetadataCustomerCustom(ctx).DataInterfaceName(dataInterfaceName).Execute()

attributeMetadata/customer/custom



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
	dataInterfaceName := "dataInterfaceName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttributeMetadataCustomerCustomAPI.GetV1AttributeMetadataCustomerCustom(context.Background()).DataInterfaceName(dataInterfaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttributeMetadataCustomerCustomAPI.GetV1AttributeMetadataCustomerCustom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AttributeMetadataCustomerCustom`: []CustomerDataAttributeMetadataInterface
	fmt.Fprintf(os.Stdout, "Response from `AttributeMetadataCustomerCustomAPI.GetV1AttributeMetadataCustomerCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AttributeMetadataCustomerCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataInterfaceName** | **string** |  | 

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

