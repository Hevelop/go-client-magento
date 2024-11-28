# \PurchaseOrderCartsCartIdEstimateShippingMethodsByAddressIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid**](PurchaseOrderCartsCartIdEstimateShippingMethodsByAddressIdAPI.md#PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid) | **Post** /V1/purchase-order-carts/{cartId}/estimate-shipping-methods-by-address-id | purchase-order-carts/{cartId}/estimate-shipping-methods-by-address-id



## PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid

> []QuoteDataShippingMethodInterface PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid(ctx, cartId).PostV1CartsMineEstimateshippingmethodsbyaddressidRequest(postV1CartsMineEstimateshippingmethodsbyaddressidRequest).Execute()

purchase-order-carts/{cartId}/estimate-shipping-methods-by-address-id



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
	cartId := int32(56) // int32 | The shopping cart ID.
	postV1CartsMineEstimateshippingmethodsbyaddressidRequest := *openapiclient.NewPostV1CartsMineEstimateshippingmethodsbyaddressidRequest(int32(123)) // PostV1CartsMineEstimateshippingmethodsbyaddressidRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderCartsCartIdEstimateShippingMethodsByAddressIdAPI.PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid(context.Background(), cartId).PostV1CartsMineEstimateshippingmethodsbyaddressidRequest(postV1CartsMineEstimateshippingmethodsbyaddressidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderCartsCartIdEstimateShippingMethodsByAddressIdAPI.PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid`: []QuoteDataShippingMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderCartsCartIdEstimateShippingMethodsByAddressIdAPI.PostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The shopping cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1PurchaseordercartsCartIdEstimateshippingmethodsbyaddressidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMineEstimateshippingmethodsbyaddressidRequest** | [**PostV1CartsMineEstimateshippingmethodsbyaddressidRequest**](PostV1CartsMineEstimateshippingmethodsbyaddressidRequest.md) |  | 

### Return type

[**[]QuoteDataShippingMethodInterface**](QuoteDataShippingMethodInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

