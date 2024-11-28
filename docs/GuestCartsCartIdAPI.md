# \GuestCartsCartIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartId**](GuestCartsCartIdAPI.md#GetV1GuestcartsCartId) | **Get** /V1/guest-carts/{cartId} | guest-carts/{cartId}
[**PutV1GuestcartsCartId**](GuestCartsCartIdAPI.md#PutV1GuestcartsCartId) | **Put** /V1/guest-carts/{cartId} | guest-carts/{cartId}



## GetV1GuestcartsCartId

> QuoteDataCartInterface GetV1GuestcartsCartId(ctx, cartId).Execute()

guest-carts/{cartId}



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
	cartId := "cartId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdAPI.GetV1GuestcartsCartId(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdAPI.GetV1GuestcartsCartId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartId`: QuoteDataCartInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdAPI.GetV1GuestcartsCartId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteDataCartInterface**](QuoteDataCartInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1GuestcartsCartId

> bool PutV1GuestcartsCartId(ctx, cartId).PutV1CartsCartIdRequest(putV1CartsCartIdRequest).Execute()

guest-carts/{cartId}



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
	putV1CartsCartIdRequest := *openapiclient.NewPutV1CartsCartIdRequest(int32(123), int32(123)) // PutV1CartsCartIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdAPI.PutV1GuestcartsCartId(context.Background(), cartId).PutV1CartsCartIdRequest(putV1CartsCartIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdAPI.PutV1GuestcartsCartId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GuestcartsCartId`: bool
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdAPI.PutV1GuestcartsCartId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GuestcartsCartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CartsCartIdRequest** | [**PutV1CartsCartIdRequest**](PutV1CartsCartIdRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

