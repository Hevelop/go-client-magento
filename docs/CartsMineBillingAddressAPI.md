# \CartsMineBillingAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMineBillingaddress**](CartsMineBillingAddressAPI.md#GetV1CartsMineBillingaddress) | **Get** /V1/carts/mine/billing-address | carts/mine/billing-address
[**PostV1CartsMineBillingaddress**](CartsMineBillingAddressAPI.md#PostV1CartsMineBillingaddress) | **Post** /V1/carts/mine/billing-address | carts/mine/billing-address



## GetV1CartsMineBillingaddress

> QuoteDataAddressInterface GetV1CartsMineBillingaddress(ctx).Execute()

carts/mine/billing-address



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineBillingAddressAPI.GetV1CartsMineBillingaddress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineBillingAddressAPI.GetV1CartsMineBillingaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMineBillingaddress`: QuoteDataAddressInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineBillingAddressAPI.GetV1CartsMineBillingaddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineBillingaddressRequest struct via the builder pattern


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


## PostV1CartsMineBillingaddress

> int32 PostV1CartsMineBillingaddress(ctx).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()

carts/mine/billing-address



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
	postV1CartsMineBillingaddressRequest := *openapiclient.NewPostV1CartsMineBillingaddressRequest(*openapiclient.NewQuoteDataAddressInterface("Region_example", int32(123), "RegionCode_example", "CountryId_example", []string{"Street_example"}, "Telephone_example", "Postcode_example", "City_example", "Firstname_example", "Lastname_example", "Email_example")) // PostV1CartsMineBillingaddressRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineBillingAddressAPI.PostV1CartsMineBillingaddress(context.Background()).PostV1CartsMineBillingaddressRequest(postV1CartsMineBillingaddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineBillingAddressAPI.PostV1CartsMineBillingaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineBillingaddress`: int32
	fmt.Fprintf(os.Stdout, "Response from `CartsMineBillingAddressAPI.PostV1CartsMineBillingaddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineBillingaddressRequest struct via the builder pattern


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

