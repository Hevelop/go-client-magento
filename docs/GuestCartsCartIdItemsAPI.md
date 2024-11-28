# \GuestCartsCartIdItemsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdItems**](GuestCartsCartIdItemsAPI.md#GetV1GuestcartsCartIdItems) | **Get** /V1/guest-carts/{cartId}/items | guest-carts/{cartId}/items
[**PostV1GuestcartsCartIdItems**](GuestCartsCartIdItemsAPI.md#PostV1GuestcartsCartIdItems) | **Post** /V1/guest-carts/{cartId}/items | guest-carts/{cartId}/items



## GetV1GuestcartsCartIdItems

> []QuoteDataCartItemInterface GetV1GuestcartsCartIdItems(ctx, cartId).Execute()

guest-carts/{cartId}/items



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
	resp, r, err := apiClient.GuestCartsCartIdItemsAPI.GetV1GuestcartsCartIdItems(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdItemsAPI.GetV1GuestcartsCartIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdItems`: []QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdItemsAPI.GetV1GuestcartsCartIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]QuoteDataCartItemInterface**](QuoteDataCartItemInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1GuestcartsCartIdItems

> QuoteDataCartItemInterface PostV1GuestcartsCartIdItems(ctx, cartId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()

guest-carts/{cartId}/items



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
	postV1CartsMineItemsRequest := *openapiclient.NewPostV1CartsMineItemsRequest(*openapiclient.NewQuoteDataCartItemInterface(float32(123), "QuoteId_example")) // PostV1CartsMineItemsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdItemsAPI.PostV1GuestcartsCartIdItems(context.Background(), cartId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdItemsAPI.PostV1GuestcartsCartIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1GuestcartsCartIdItems`: QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdItemsAPI.PostV1GuestcartsCartIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1GuestcartsCartIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMineItemsRequest** | [**PostV1CartsMineItemsRequest**](PostV1CartsMineItemsRequest.md) |  | 

### Return type

[**QuoteDataCartItemInterface**](QuoteDataCartItemInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

