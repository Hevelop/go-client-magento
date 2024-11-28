# \ProductsAttributeSetsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsAttributesets**](ProductsAttributeSetsAPI.md#PostV1ProductsAttributesets) | **Post** /V1/products/attribute-sets | products/attribute-sets



## PostV1ProductsAttributesets

> EavDataAttributeSetInterface PostV1ProductsAttributesets(ctx).PostV1ProductsAttributesetsRequest(postV1ProductsAttributesetsRequest).Execute()

products/attribute-sets



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
	postV1ProductsAttributesetsRequest := *openapiclient.NewPostV1ProductsAttributesetsRequest(*openapiclient.NewEavDataAttributeSetInterface("AttributeSetName_example", int32(123)), int32(123)) // PostV1ProductsAttributesetsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsAPI.PostV1ProductsAttributesets(context.Background()).PostV1ProductsAttributesetsRequest(postV1ProductsAttributesetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsAPI.PostV1ProductsAttributesets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsAttributesets`: EavDataAttributeSetInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsAPI.PostV1ProductsAttributesets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsAttributesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsAttributesetsRequest** | [**PostV1ProductsAttributesetsRequest**](PostV1ProductsAttributesetsRequest.md) |  | 

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

