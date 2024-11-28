# \CustomerGroupsDefaultIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CustomerGroupsDefaultId**](CustomerGroupsDefaultIdAPI.md#PutV1CustomerGroupsDefaultId) | **Put** /V1/customerGroups/default/{id} | customerGroups/default/{id}



## PutV1CustomerGroupsDefaultId

> int32 PutV1CustomerGroupsDefaultId(ctx, id).Execute()

customerGroups/default/{id}



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
	resp, r, err := apiClient.CustomerGroupsDefaultIdAPI.PutV1CustomerGroupsDefaultId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsDefaultIdAPI.PutV1CustomerGroupsDefaultId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CustomerGroupsDefaultId`: int32
	fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsDefaultIdAPI.PutV1CustomerGroupsDefaultId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CustomerGroupsDefaultIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

