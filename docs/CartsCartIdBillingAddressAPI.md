# \CartsCartIdBillingAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsCartIdBillingaddress**](CartsCartIdBillingAddressAPI.md#GetV1CartsCartIdBillingaddress) | **Get** /V1/carts/{cartId}/billing-address | carts/{cartId}/billing-address
[**PostV1CartsCartIdBillingaddress**](CartsCartIdBillingAddressAPI.md#PostV1CartsCartIdBillingaddress) | **Post** /V1/carts/{cartId}/billing-address | carts/{cartId}/billing-address



## GetV1CartsCartIdBillingaddress

> QuoteDataAddressInterface GetV1CartsCartIdBillingaddress(ctx, cartId).Execute()

carts/{cartId}/billing-address



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdBillingAddressAPI.GetV1CartsCartIdBillingaddress(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdBillingAddressAPI.GetV1CartsCartIdBillingaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartIdBillingaddress`: QuoteDataAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdBillingAddressAPI.GetV1CartsCartIdBillingaddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdBillingaddressRequest struct via the builder pattern


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


## PostV1CartsCartIdBillingaddress

> int32 PostV1CartsCartIdBillingaddress(ctx, cartId).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()

carts/{cartId}/billing-address



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
	resp, r, err := apiClient.CartsCartIdBillingAddressAPI.PostV1CartsCartIdBillingaddress(context.Background(), cartId).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdBillingAddressAPI.PostV1CartsCartIdBillingaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsCartIdBillingaddress`: int32
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdBillingAddressAPI.PostV1CartsCartIdBillingaddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsCartIdBillingaddressRequest struct via the builder pattern


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

