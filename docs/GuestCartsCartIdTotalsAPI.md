# \GuestCartsCartIdTotalsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1GuestcartsCartIdTotals**](GuestCartsCartIdTotalsAPI.md#GetV1GuestcartsCartIdTotals) | **Get** /V1/guest-carts/{cartId}/totals | guest-carts/{cartId}/totals



## GetV1GuestcartsCartIdTotals

> QuoteDataTotalsInterface GetV1GuestcartsCartIdTotals(ctx, cartId).Execute()

guest-carts/{cartId}/totals



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
	resp, r, err := apiClient.GuestCartsCartIdTotalsAPI.GetV1GuestcartsCartIdTotals(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdTotalsAPI.GetV1GuestcartsCartIdTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdTotals`: QuoteDataTotalsInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdTotalsAPI.GetV1GuestcartsCartIdTotals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdTotalsRequest struct via the builder pattern


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

