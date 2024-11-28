# \CustomerGroupsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CustomerGroups**](CustomerGroupsAPI.md#PostV1CustomerGroups) | **Post** /V1/customerGroups | customerGroups



## PostV1CustomerGroups

> CustomerDataGroupInterface PostV1CustomerGroups(ctx).PostV1CustomerGroupsRequest(postV1CustomerGroupsRequest).Execute()

customerGroups



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
	postV1CustomerGroupsRequest := *openapiclient.NewPostV1CustomerGroupsRequest(*openapiclient.NewCustomerDataGroupInterface("Code_example", int32(123))) // PostV1CustomerGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerGroupsAPI.PostV1CustomerGroups(context.Background()).PostV1CustomerGroupsRequest(postV1CustomerGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsAPI.PostV1CustomerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CustomerGroups`: CustomerDataGroupInterface
	fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsAPI.PostV1CustomerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CustomerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CustomerGroupsRequest** | [**PostV1CustomerGroupsRequest**](PostV1CustomerGroupsRequest.md) |  | 

### Return type

[**CustomerDataGroupInterface**](CustomerDataGroupInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

