# \ProductsAttributeSetsAttributesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsAttributesetsAttributes**](ProductsAttributeSetsAttributesAPI.md#PostV1ProductsAttributesetsAttributes) | **Post** /V1/products/attribute-sets/attributes | products/attribute-sets/attributes



## PostV1ProductsAttributesetsAttributes

> int32 PostV1ProductsAttributesetsAttributes(ctx).PostV1ProductsAttributesetsAttributesRequest(postV1ProductsAttributesetsAttributesRequest).Execute()

products/attribute-sets/attributes



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
	postV1ProductsAttributesetsAttributesRequest := *openapiclient.NewPostV1ProductsAttributesetsAttributesRequest(int32(123), int32(123), "AttributeCode_example", int32(123)) // PostV1ProductsAttributesetsAttributesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsAttributesAPI.PostV1ProductsAttributesetsAttributes(context.Background()).PostV1ProductsAttributesetsAttributesRequest(postV1ProductsAttributesetsAttributesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsAttributesAPI.PostV1ProductsAttributesetsAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsAttributesetsAttributes`: int32
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsAttributesAPI.PostV1ProductsAttributesetsAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsAttributesetsAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsAttributesetsAttributesRequest** | [**PostV1ProductsAttributesetsAttributesRequest**](PostV1ProductsAttributesetsAttributesRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

