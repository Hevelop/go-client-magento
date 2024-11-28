# \OrderNotifyOrdersAreReadyForPickupAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1OrderNotifyordersarereadyforpickup**](OrderNotifyOrdersAreReadyForPickupAPI.md#PostV1OrderNotifyordersarereadyforpickup) | **Post** /V1/order/notify-orders-are-ready-for-pickup | order/notify-orders-are-ready-for-pickup



## PostV1OrderNotifyordersarereadyforpickup

> InventoryInStorePickupSalesApiDataResultInterface PostV1OrderNotifyordersarereadyforpickup(ctx).PostV1OrderNotifyordersarereadyforpickupRequest(postV1OrderNotifyordersarereadyforpickupRequest).Execute()

order/notify-orders-are-ready-for-pickup



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
	postV1OrderNotifyordersarereadyforpickupRequest := *openapiclient.NewPostV1OrderNotifyordersarereadyforpickupRequest([]int32{int32(123)}) // PostV1OrderNotifyordersarereadyforpickupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderNotifyOrdersAreReadyForPickupAPI.PostV1OrderNotifyordersarereadyforpickup(context.Background()).PostV1OrderNotifyordersarereadyforpickupRequest(postV1OrderNotifyordersarereadyforpickupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderNotifyOrdersAreReadyForPickupAPI.PostV1OrderNotifyordersarereadyforpickup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1OrderNotifyordersarereadyforpickup`: InventoryInStorePickupSalesApiDataResultInterface
	fmt.Fprintf(os.Stdout, "Response from `OrderNotifyOrdersAreReadyForPickupAPI.PostV1OrderNotifyordersarereadyforpickup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1OrderNotifyordersarereadyforpickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1OrderNotifyordersarereadyforpickupRequest** | [**PostV1OrderNotifyordersarereadyforpickupRequest**](PostV1OrderNotifyordersarereadyforpickupRequest.md) |  | 

### Return type

[**InventoryInStorePickupSalesApiDataResultInterface**](InventoryInStorePickupSalesApiDataResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

