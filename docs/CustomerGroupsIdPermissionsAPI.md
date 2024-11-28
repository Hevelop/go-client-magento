# \CustomerGroupsIdPermissionsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomerGroupsIdPermissions**](CustomerGroupsIdPermissionsAPI.md#GetV1CustomerGroupsIdPermissions) | **Get** /V1/customerGroups/{id}/permissions | customerGroups/{id}/permissions



## GetV1CustomerGroupsIdPermissions

> bool GetV1CustomerGroupsIdPermissions(ctx, id).Execute()

customerGroups/{id}/permissions



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerGroupsIdPermissionsAPI.GetV1CustomerGroupsIdPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsIdPermissionsAPI.GetV1CustomerGroupsIdPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomerGroupsIdPermissions`: bool
	fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsIdPermissionsAPI.GetV1CustomerGroupsIdPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomerGroupsIdPermissionsRequest struct via the builder pattern


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

