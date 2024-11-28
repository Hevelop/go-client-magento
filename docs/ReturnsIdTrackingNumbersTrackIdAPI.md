# \ReturnsIdTrackingNumbersTrackIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ReturnsIdTrackingnumbersTrackId**](ReturnsIdTrackingNumbersTrackIdAPI.md#DeleteV1ReturnsIdTrackingnumbersTrackId) | **Delete** /V1/returns/{id}/tracking-numbers/{trackId} | returns/{id}/tracking-numbers/{trackId}



## DeleteV1ReturnsIdTrackingnumbersTrackId

> bool DeleteV1ReturnsIdTrackingnumbersTrackId(ctx, id, trackId).Execute()

returns/{id}/tracking-numbers/{trackId}



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
	trackId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsIdTrackingNumbersTrackIdAPI.DeleteV1ReturnsIdTrackingnumbersTrackId(context.Background(), id, trackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdTrackingNumbersTrackIdAPI.DeleteV1ReturnsIdTrackingnumbersTrackId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ReturnsIdTrackingnumbersTrackId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdTrackingNumbersTrackIdAPI.DeleteV1ReturnsIdTrackingnumbersTrackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 
**trackId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ReturnsIdTrackingnumbersTrackIdRequest struct via the builder pattern


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

