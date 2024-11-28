# \CustomerGroupsDefaultAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CustomerGroupsDefault**](CustomerGroupsDefaultAPI.md#GetV1CustomerGroupsDefault) | **Get** /V1/customerGroups/default | customerGroups/default



## GetV1CustomerGroupsDefault

> CustomerDataGroupInterface GetV1CustomerGroupsDefault(ctx).StoreId(storeId).Execute()

customerGroups/default



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
	storeId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerGroupsDefaultAPI.GetV1CustomerGroupsDefault(context.Background()).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsDefaultAPI.GetV1CustomerGroupsDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CustomerGroupsDefault`: CustomerDataGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsDefaultAPI.GetV1CustomerGroupsDefault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CustomerGroupsDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **int32** |  | 

### Return type

[**CustomerDataGroupInterface**](CustomerDataGroupInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

