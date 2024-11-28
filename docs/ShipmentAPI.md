# \ShipmentAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1Shipment**](ShipmentAPI.md#PostV1Shipment) | **Post** /V1/shipment/ | shipment/



## PostV1Shipment

> SalesDataShipmentInterface PostV1Shipment(ctx).PostV1ShipmentRequest(postV1ShipmentRequest).Execute()

shipment/



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
	postV1ShipmentRequest := *openapiclient.NewPostV1ShipmentRequest(*openapiclient.NewSalesDataShipmentInterface(int32(123), []openapiclient.SalesDataShipmentItemInterface{*openapiclient.NewSalesDataShipmentItemInterface(int32(123), float32(123))}, []openapiclient.SalesDataShipmentTrackInterface{*openapiclient.NewSalesDataShipmentTrackInterface(int32(123), int32(123), float32(123), float32(123), "Description_example", "TrackNumber_example", "Title_example", "CarrierCode_example")}, []openapiclient.SalesDataShipmentCommentInterface{*openapiclient.NewSalesDataShipmentCommentInterface(int32(123), int32(123), "Comment_example", int32(123))})) // PostV1ShipmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentAPI.PostV1Shipment(context.Background()).PostV1ShipmentRequest(postV1ShipmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentAPI.PostV1Shipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Shipment`: SalesDataShipmentInterface
	fmt.Fprintf(os.Stdout, "Response from `ShipmentAPI.PostV1Shipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1ShipmentRequest** | [**PostV1ShipmentRequest**](PostV1ShipmentRequest.md) |  | 

### Return type

[**SalesDataShipmentInterface**](SalesDataShipmentInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

