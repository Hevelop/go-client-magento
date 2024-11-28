# \InventoryLowQuantityNotificationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventoryLowquantitynotification**](InventoryLowQuantityNotificationAPI.md#PostV1InventoryLowquantitynotification) | **Post** /V1/inventory/low-quantity-notification | inventory/low-quantity-notification



## PostV1InventoryLowquantitynotification

> ErrorResponse PostV1InventoryLowquantitynotification(ctx).PostV1InventoryLowquantitynotificationRequest(postV1InventoryLowquantitynotificationRequest).Execute()

inventory/low-quantity-notification



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
	postV1InventoryLowquantitynotificationRequest := *openapiclient.NewPostV1InventoryLowquantitynotificationRequest([]openapiclient.InventoryLowQuantityNotificationApiDataSourceItemConfigurationInterface{*openapiclient.NewInventoryLowQuantityNotificationApiDataSourceItemConfigurationInterface()}) // PostV1InventoryLowquantitynotificationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLowQuantityNotificationAPI.PostV1InventoryLowquantitynotification(context.Background()).PostV1InventoryLowquantitynotificationRequest(postV1InventoryLowquantitynotificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLowQuantityNotificationAPI.PostV1InventoryLowquantitynotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventoryLowquantitynotification`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryLowQuantityNotificationAPI.PostV1InventoryLowquantitynotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventoryLowquantitynotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InventoryLowquantitynotificationRequest** | [**PostV1InventoryLowquantitynotificationRequest**](PostV1InventoryLowquantitynotificationRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

