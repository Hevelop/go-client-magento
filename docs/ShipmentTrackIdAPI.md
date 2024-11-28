# \ShipmentTrackIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1ShipmentTrackId**](ShipmentTrackIdAPI.md#DeleteV1ShipmentTrackId) | **Delete** /V1/shipment/track/{id} | shipment/track/{id}



## DeleteV1ShipmentTrackId

> bool DeleteV1ShipmentTrackId(ctx, id).Execute()

shipment/track/{id}



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
	id := int32(56) // int32 | The shipment track ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentTrackIdAPI.DeleteV1ShipmentTrackId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentTrackIdAPI.DeleteV1ShipmentTrackId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1ShipmentTrackId`: bool
	fmt.Fprintf(os.Stdout, "Response from `ShipmentTrackIdAPI.DeleteV1ShipmentTrackId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The shipment track ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1ShipmentTrackIdRequest struct via the builder pattern


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

