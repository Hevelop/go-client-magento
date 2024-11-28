# \NegotiableCartsCartIdTotalsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1NegotiablecartsCartIdTotals**](NegotiableCartsCartIdTotalsAPI.md#GetV1NegotiablecartsCartIdTotals) | **Get** /V1/negotiable-carts/{cartId}/totals | negotiable-carts/{cartId}/totals



## GetV1NegotiablecartsCartIdTotals

> QuoteDataTotalsInterface GetV1NegotiablecartsCartIdTotals(ctx, cartId).Execute()

negotiable-carts/{cartId}/totals



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
	resp, r, err := apiClient.NegotiableCartsCartIdTotalsAPI.GetV1NegotiablecartsCartIdTotals(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdTotalsAPI.GetV1NegotiablecartsCartIdTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1NegotiablecartsCartIdTotals`: QuoteDataTotalsInterface
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdTotalsAPI.GetV1NegotiablecartsCartIdTotals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1NegotiablecartsCartIdTotalsRequest struct via the builder pattern


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

