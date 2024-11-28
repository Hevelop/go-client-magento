# \EavAttributeSetsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1EavAttributesets**](EavAttributeSetsAPI.md#PostV1EavAttributesets) | **Post** /V1/eav/attribute-sets | eav/attribute-sets



## PostV1EavAttributesets

> EavDataAttributeSetInterface PostV1EavAttributesets(ctx).PostV1EavAttributesetsRequest(postV1EavAttributesetsRequest).Execute()

eav/attribute-sets



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
	postV1EavAttributesetsRequest := *openapiclient.NewPostV1EavAttributesetsRequest("EntityTypeCode_example", *openapiclient.NewEavDataAttributeSetInterface("AttributeSetName_example", int32(123)), int32(123)) // PostV1EavAttributesetsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EavAttributeSetsAPI.PostV1EavAttributesets(context.Background()).PostV1EavAttributesetsRequest(postV1EavAttributesetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EavAttributeSetsAPI.PostV1EavAttributesets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1EavAttributesets`: EavDataAttributeSetInterface
	fmt.Fprintf(os.Stdout, "Response from `EavAttributeSetsAPI.PostV1EavAttributesets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1EavAttributesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1EavAttributesetsRequest** | [**PostV1EavAttributesetsRequest**](PostV1EavAttributesetsRequest.md) |  | 

### Return type

[**EavDataAttributeSetInterface**](EavDataAttributeSetInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

