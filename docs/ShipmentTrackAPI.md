# \ShipmentTrackAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1ShipmentTrack**](ShipmentTrackAPI.md#PostV1ShipmentTrack) | **Post** /V1/shipment/track | shipment/track



## PostV1ShipmentTrack

> SalesDataShipmentTrackInterface PostV1ShipmentTrack(ctx).PostV1ShipmentTrackRequest(postV1ShipmentTrackRequest).Execute()

shipment/track



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
	postV1ShipmentTrackRequest := *openapiclient.NewPostV1ShipmentTrackRequest(*openapiclient.NewSalesDataShipmentTrackInterface(int32(123), int32(123), float32(123), float32(123), "Description_example", "TrackNumber_example", "Title_example", "CarrierCode_example")) // PostV1ShipmentTrackRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentTrackAPI.PostV1ShipmentTrack(context.Background()).PostV1ShipmentTrackRequest(postV1ShipmentTrackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentTrackAPI.PostV1ShipmentTrack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ShipmentTrack`: SalesDataShipmentTrackInterface
	fmt.Fprintf(os.Stdout, "Response from `ShipmentTrackAPI.PostV1ShipmentTrack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ShipmentTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ShipmentTrackRequest** | [**PostV1ShipmentTrackRequest**](PostV1ShipmentTrackRequest.md) |  | 

### Return type

[**SalesDataShipmentTrackInterface**](SalesDataShipmentTrackInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

