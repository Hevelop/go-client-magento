# \ShipmentIdLabelAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ShipmentIdLabel**](ShipmentIdLabelAPI.md#GetV1ShipmentIdLabel) | **Get** /V1/shipment/{id}/label | shipment/{id}/label



## GetV1ShipmentIdLabel

> string GetV1ShipmentIdLabel(ctx, id).Execute()

shipment/{id}/label



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
	id := int32(56) // int32 | The shipment label ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentIdLabelAPI.GetV1ShipmentIdLabel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentIdLabelAPI.GetV1ShipmentIdLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ShipmentIdLabel`: string
	fmt.Fprintf(os.Stdout, "Response from `ShipmentIdLabelAPI.GetV1ShipmentIdLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The shipment label ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ShipmentIdLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

