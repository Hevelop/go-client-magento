# \ProductsAttributeSetsGroupsGroupIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ProductsAttributesetsGroupsGroupId**](ProductsAttributeSetsGroupsGroupIdAPI.md#DeleteV1ProductsAttributesetsGroupsGroupId) | **Delete** /V1/products/attribute-sets/groups/{groupId} | products/attribute-sets/groups/{groupId}



## DeleteV1ProductsAttributesetsGroupsGroupId

> bool DeleteV1ProductsAttributesetsGroupsGroupId(ctx, groupId).Execute()

products/attribute-sets/groups/{groupId}



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
	groupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAttributeSetsGroupsGroupIdAPI.DeleteV1ProductsAttributesetsGroupsGroupId(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAttributeSetsGroupsGroupIdAPI.DeleteV1ProductsAttributesetsGroupsGroupId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ProductsAttributesetsGroupsGroupId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ProductsAttributeSetsGroupsGroupIdAPI.DeleteV1ProductsAttributesetsGroupsGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ProductsAttributesetsGroupsGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

