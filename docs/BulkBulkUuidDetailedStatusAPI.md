# \BulkBulkUuidDetailedStatusAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1BulkBulkUuidDetailedstatus**](BulkBulkUuidDetailedStatusAPI.md#GetV1BulkBulkUuidDetailedstatus) | **Get** /V1/bulk/{bulkUuid}/detailed-status | bulk/{bulkUuid}/detailed-status



## GetV1BulkBulkUuidDetailedstatus

> AsynchronousOperationsDataDetailedBulkOperationsStatusInterface GetV1BulkBulkUuidDetailedstatus(ctx, bulkUuid).Execute()

bulk/{bulkUuid}/detailed-status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkBulkUuidDetailedStatusAPI.GetV1BulkBulkUuidDetailedstatus(context.Background(), bulkUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkBulkUuidDetailedStatusAPI.GetV1BulkBulkUuidDetailedstatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1BulkBulkUuidDetailedstatus`: AsynchronousOperationsDataDetailedBulkOperationsStatusInterface
	fmt.Fprintf(os.Stdout, "Response from `BulkBulkUuidDetailedStatusAPI.GetV1BulkBulkUuidDetailedstatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bulkUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1BulkBulkUuidDetailedstatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsynchronousOperationsDataDetailedBulkOperationsStatusInterface**](AsynchronousOperationsDataDetailedBulkOperationsStatusInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

