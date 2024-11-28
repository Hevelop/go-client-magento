# \GuestCartsCartIdItemsItemIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1GuestcartsCartIdItemsItemId**](GuestCartsCartIdItemsItemIdAPI.md#DeleteV1GuestcartsCartIdItemsItemId) | **Delete** /V1/guest-carts/{cartId}/items/{itemId} | guest-carts/{cartId}/items/{itemId}
[**PutV1GuestcartsCartIdItemsItemId**](GuestCartsCartIdItemsItemIdAPI.md#PutV1GuestcartsCartIdItemsItemId) | **Put** /V1/guest-carts/{cartId}/items/{itemId} | guest-carts/{cartId}/items/{itemId}



## DeleteV1GuestcartsCartIdItemsItemId

> bool DeleteV1GuestcartsCartIdItemsItemId(ctx, cartId, itemId).Execute()

guest-carts/{cartId}/items/{itemId}



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
	itemId := int32(56) // int32 | The item ID of the item to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdItemsItemIdAPI.DeleteV1GuestcartsCartIdItemsItemId(context.Background(), cartId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdItemsItemIdAPI.DeleteV1GuestcartsCartIdItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1GuestcartsCartIdItemsItemId`: bool
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdItemsItemIdAPI.DeleteV1GuestcartsCartIdItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 
**itemId** | **int32** | The item ID of the item to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1GuestcartsCartIdItemsItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1GuestcartsCartIdItemsItemId

> QuoteDataCartItemInterface PutV1GuestcartsCartIdItemsItemId(ctx, cartId, itemId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()

guest-carts/{cartId}/items/{itemId}



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
	itemId := "itemId_example" // string | 
	postV1CartsMineItemsRequest := *openapiclient.NewPostV1CartsMineItemsRequest(*openapiclient.NewQuoteDataCartItemInterface(float32(123), "QuoteId_example")) // PostV1CartsMineItemsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdItemsItemIdAPI.PutV1GuestcartsCartIdItemsItemId(context.Background(), cartId, itemId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdItemsItemIdAPI.PutV1GuestcartsCartIdItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GuestcartsCartIdItemsItemId`: QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdItemsItemIdAPI.PutV1GuestcartsCartIdItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GuestcartsCartIdItemsItemIdRequest struct via the builder pattern


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

