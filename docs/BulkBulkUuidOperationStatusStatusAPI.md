# \BulkBulkUuidOperationStatusStatusAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1BulkBulkUuidOperationstatusStatus**](BulkBulkUuidOperationStatusStatusAPI.md#GetV1BulkBulkUuidOperationstatusStatus) | **Get** /V1/bulk/{bulkUuid}/operation-status/{status} | bulk/{bulkUuid}/operation-status/{status}



## GetV1BulkBulkUuidOperationstatusStatus

> int32 GetV1BulkBulkUuidOperationstatusStatus(ctx, bulkUuid, status).Execute()

bulk/{bulkUuid}/operation-status/{status}



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
	bulkUuid := "bulkUuid_example" // string | 
	status := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkBulkUuidOperationStatusStatusAPI.GetV1BulkBulkUuidOperationstatusStatus(context.Background(), bulkUuid, status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkBulkUuidOperationStatusStatusAPI.GetV1BulkBulkUuidOperationstatusStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1BulkBulkUuidOperationstatusStatus`: int32
	fmt.Fprintf(os.Stdout, "Response from `BulkBulkUuidOperationStatusStatusAPI.GetV1BulkBulkUuidOperationstatusStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bulkUuid** | **string** |  | 
**status** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BulkBulkUuidOperationstatusStatusRequest struct via the builder pattern


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

