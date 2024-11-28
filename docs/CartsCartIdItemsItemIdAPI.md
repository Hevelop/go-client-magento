# \CartsCartIdItemsItemIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CartsCartIdItemsItemId**](CartsCartIdItemsItemIdAPI.md#DeleteV1CartsCartIdItemsItemId) | **Delete** /V1/carts/{cartId}/items/{itemId} | carts/{cartId}/items/{itemId}
[**PutV1CartsCartIdItemsItemId**](CartsCartIdItemsItemIdAPI.md#PutV1CartsCartIdItemsItemId) | **Put** /V1/carts/{cartId}/items/{itemId} | carts/{cartId}/items/{itemId}



## DeleteV1CartsCartIdItemsItemId

> bool DeleteV1CartsCartIdItemsItemId(ctx, cartId, itemId).Execute()

carts/{cartId}/items/{itemId}



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
	itemId := int32(56) // int32 | The item ID of the item to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdItemsItemIdAPI.DeleteV1CartsCartIdItemsItemId(context.Background(), cartId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdItemsItemIdAPI.DeleteV1CartsCartIdItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CartsCartIdItemsItemId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdItemsItemIdAPI.DeleteV1CartsCartIdItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 
**itemId** | **int32** | The item ID of the item to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CartsCartIdItemsItemIdRequest struct via the builder pattern


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


## PutV1CartsCartIdItemsItemId

> QuoteDataCartItemInterface PutV1CartsCartIdItemsItemId(ctx, cartId, itemId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()

carts/{cartId}/items/{itemId}



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
	resp, r, err := apiClient.CartsCartIdItemsItemIdAPI.PutV1CartsCartIdItemsItemId(context.Background(), cartId, itemId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdItemsItemIdAPI.PutV1CartsCartIdItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsCartIdItemsItemId`: QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdItemsItemIdAPI.PutV1CartsCartIdItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsCartIdItemsItemIdRequest struct via the builder pattern


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

