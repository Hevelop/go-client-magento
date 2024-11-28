# \GuestCartsCartIdBillingAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdBillingaddress**](GuestCartsCartIdBillingAddressAPI.md#GetV1GuestcartsCartIdBillingaddress) | **Get** /V1/guest-carts/{cartId}/billing-address | guest-carts/{cartId}/billing-address
[**PostV1GuestcartsCartIdBillingaddress**](GuestCartsCartIdBillingAddressAPI.md#PostV1GuestcartsCartIdBillingaddress) | **Post** /V1/guest-carts/{cartId}/billing-address | guest-carts/{cartId}/billing-address



## GetV1GuestcartsCartIdBillingaddress

> QuoteDataAddressInterface GetV1GuestcartsCartIdBillingaddress(ctx, cartId).Execute()

guest-carts/{cartId}/billing-address



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
	cartId := "cartId_example" // string | The cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdBillingAddressAPI.GetV1GuestcartsCartIdBillingaddress(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdBillingAddressAPI.GetV1GuestcartsCartIdBillingaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdBillingaddress`: QuoteDataAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdBillingAddressAPI.GetV1GuestcartsCartIdBillingaddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdBillingaddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteDataAddressInterface**](QuoteDataAddressInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1GuestcartsCartIdBillingaddress

> int32 PostV1GuestcartsCartIdBillingaddress(ctx, cartId).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()

guest-carts/{cartId}/billing-address



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
	cartId := "cartId_example" // string | The cart ID.
	postV1CartsMineBillingaddressRequest := *openapiclient.NewPostV1CartsMineBillingaddressRequest(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example")) // PostV1CartsMineBillingaddressRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdBillingAddressAPI.PostV1GuestcartsCartIdBillingaddress(context.Background(), cartId).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdBillingAddressAPI.PostV1GuestcartsCartIdBillingaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdBillingaddress`: int32
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdBillingAddressAPI.PostV1GuestcartsCartIdBillingaddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdBillingaddressRequest struct via the builder pattern


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

