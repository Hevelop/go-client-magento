# \InventoryLowQuantityNotificationsDeleteAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventoryLowquantitynotificationsdelete**](InventoryLowQuantityNotificationsDeleteAPI.md#PostV1InventoryLowquantitynotificationsdelete) | **Post** /V1/inventory/low-quantity-notifications-delete | inventory/low-quantity-notifications-delete



## PostV1InventoryLowquantitynotificationsdelete

> ErrorResponse PostV1InventoryLowquantitynotificationsdelete(ctx).PostV1InventoryLowquantitynotificationsdeleteRequest(postV1InventoryLowquantitynotificationsdeleteRequest).Execute()

inventory/low-quantity-notifications-delete



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
	postV1InventoryLowquantitynotificationsdeleteRequest := *openapiclient.NewPostV1InventoryLowquantitynotificationsdeleteRequest([]openapiclient.InventoryApiDataSourceItemInterface{*openapiclient.NewInventoryApiDataSourceItemInterface()}) // PostV1InventoryLowquantitynotificationsdeleteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLowQuantityNotificationsDeleteAPI.PostV1InventoryLowquantitynotificationsdelete(context.Background()).PostV1InventoryLowquantitynotificationsdeleteRequest(postV1InventoryLowquantitynotificationsdeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLowQuantityNotificationsDeleteAPI.PostV1InventoryLowquantitynotificationsdelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventoryLowquantitynotificationsdelete`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryLowQuantityNotificationsDeleteAPI.PostV1InventoryLowquantitynotificationsdelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventoryLowquantitynotificationsdeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InventoryLowquantitynotificationsdeleteRequest** | [**PostV1InventoryLowquantitynotificationsdeleteRequest**](PostV1InventoryLowquantitynotificationsdeleteRequest.md) |  | 

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

