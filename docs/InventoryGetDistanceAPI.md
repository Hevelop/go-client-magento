# \InventoryGetDistanceAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryGetdistance**](InventoryGetDistanceAPI.md#GetV1InventoryGetdistance) | **Get** /V1/inventory/get-distance | inventory/get-distance



## GetV1InventoryGetdistance

> float32 GetV1InventoryGetdistance(ctx).SourceLat(sourceLat).SourceLng(sourceLng).DestinationLat(destinationLat).DestinationLng(destinationLng).Execute()

inventory/get-distance



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
	sourceLat := float32(8.14) // float32 | Latitude (optional)
	sourceLng := float32(8.14) // float32 | Longitude (optional)
	destinationLat := float32(8.14) // float32 | Latitude (optional)
	destinationLng := float32(8.14) // float32 | Longitude (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGetDistanceAPI.GetV1InventoryGetdistance(context.Background()).SourceLat(sourceLat).SourceLng(sourceLng).DestinationLat(destinationLat).DestinationLng(destinationLng).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGetDistanceAPI.GetV1InventoryGetdistance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryGetdistance`: float32
	fmt.Fprintf(os.Stdout, "Response from `InventoryGetDistanceAPI.GetV1InventoryGetdistance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryGetdistanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLat** | **float32** | Latitude | 
 **sourceLng** | **float32** | Longitude | 
 **destinationLat** | **float32** | Latitude | 
 **destinationLng** | **float32** | Longitude | 

### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

