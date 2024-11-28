# \PurchaseOrderCartsCartIdBillingAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1PurchaseordercartsCartIdBillingaddress**](PurchaseOrderCartsCartIdBillingAddressAPI.md#PostV1PurchaseordercartsCartIdBillingaddress) | **Post** /V1/purchase-order-carts/{cartId}/billing-address | purchase-order-carts/{cartId}/billing-address



## PostV1PurchaseordercartsCartIdBillingaddress

> int32 PostV1PurchaseordercartsCartIdBillingaddress(ctx, cartId).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()

purchase-order-carts/{cartId}/billing-address



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
	cartId := int32(56) // int32 | The cart ID.
	postV1CartsMineBillingaddressRequest := *openapiclient.NewPostV1CartsMineBillingaddressRequest(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example")) // PostV1CartsMineBillingaddressRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderCartsCartIdBillingAddressAPI.PostV1PurchaseordercartsCartIdBillingaddress(context.Background(), cartId).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderCartsCartIdBillingAddressAPI.PostV1PurchaseordercartsCartIdBillingaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1PurchaseordercartsCartIdBillingaddress`: int32
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderCartsCartIdBillingAddressAPI.PostV1PurchaseordercartsCartIdBillingaddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1PurchaseordercartsCartIdBillingaddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMineBillingaddressRequest** | [**PostV1CartsMineBillingaddressRequest**](PostV1CartsMineBillingaddressRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

