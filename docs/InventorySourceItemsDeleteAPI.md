# \InventorySourceItemsDeleteAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventorySourceitemsdelete**](InventorySourceItemsDeleteAPI.md#PostV1InventorySourceitemsdelete) | **Post** /V1/inventory/source-items-delete | inventory/source-items-delete



## PostV1InventorySourceitemsdelete

> ErrorResponse PostV1InventorySourceitemsdelete(ctx).PostV1InventoryLowquantitynotificationsdeleteRequest(postV1InventoryLowquantitynotificationsdeleteRequest).Execute()

inventory/source-items-delete



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
	resp, r, err := apiClient.InventorySourceItemsDeleteAPI.PostV1InventorySourceitemsdelete(context.Background()).PostV1InventoryLowquantitynotificationsdeleteRequest(postV1InventoryLowquantitynotificationsdeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventorySourceItemsDeleteAPI.PostV1InventorySourceitemsdelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventorySourceitemsdelete`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `InventorySourceItemsDeleteAPI.PostV1InventorySourceitemsdelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventorySourceitemsdeleteRequest struct via the builder pattern


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

