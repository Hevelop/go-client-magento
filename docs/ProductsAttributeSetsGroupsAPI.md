# \ProductsAttributeSetsGroupsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ProductsAttributesetsGroups**](ProductsAttributeSetsGroupsAPI.md#PostV1ProductsAttributesetsGroups) | **Post** /V1/products/attribute-sets/groups | products/attribute-sets/groups



## PostV1ProductsAttributesetsGroups

> EavDataAttributeGroupInterface PostV1ProductsAttributesetsGroups(ctx).PostV1ProductsAttributesetsGroupsRequest(postV1ProductsAttributesetsGroupsRequest).Execute()

products/attribute-sets/groups



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
	postV1ProductsAttributesetsGroupsRequest := *openapiclient.NewPostV1ProductsAttributesetsGroupsRequest(*openapiclient.NewEavDataAttributeGroupInterface()) // PostV1ProductsAttributesetsGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsGroupsAPI.PostV1ProductsAttributesetsGroups(context.Background()).PostV1ProductsAttributesetsGroupsRequest(postV1ProductsAttributesetsGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsGroupsAPI.PostV1ProductsAttributesetsGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ProductsAttributesetsGroups`: EavDataAttributeGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsGroupsAPI.PostV1ProductsAttributesetsGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ProductsAttributesetsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ProductsAttributesetsGroupsRequest** | [**PostV1ProductsAttributesetsGroupsRequest**](PostV1ProductsAttributesetsGroupsRequest.md) |  | 

### Return type

[**EavDataAttributeGroupInterface**](EavDataAttributeGroupInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

