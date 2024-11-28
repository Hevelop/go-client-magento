# \PurchaseOrderCartsCartIdTotalsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1PurchaseordercartsCartIdTotals**](PurchaseOrderCartsCartIdTotalsAPI.md#GetV1PurchaseordercartsCartIdTotals) | **Get** /V1/purchase-order-carts/{cartId}/totals | purchase-order-carts/{cartId}/totals



## GetV1PurchaseordercartsCartIdTotals

> QuoteDataTotalsInterface GetV1PurchaseordercartsCartIdTotals(ctx, cartId).Execute()

purchase-order-carts/{cartId}/totals



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
	resp, r, err := apiClient.PurchaseOrderCartsCartIdTotalsAPI.GetV1PurchaseordercartsCartIdTotals(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderCartsCartIdTotalsAPI.GetV1PurchaseordercartsCartIdTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1PurchaseordercartsCartIdTotals`: QuoteDataTotalsInterface
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderCartsCartIdTotalsAPI.GetV1PurchaseordercartsCartIdTotals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1PurchaseordercartsCartIdTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteDataTotalsInterface**](QuoteDataTotalsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

