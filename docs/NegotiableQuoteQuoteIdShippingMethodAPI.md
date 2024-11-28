# \NegotiableQuoteQuoteIdShippingMethodAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1NegotiableQuoteQuoteIdShippingMethod**](NegotiableQuoteQuoteIdShippingMethodAPI.md#PutV1NegotiableQuoteQuoteIdShippingMethod) | **Put** /V1/negotiableQuote/{quoteId}/shippingMethod | negotiableQuote/{quoteId}/shippingMethod



## PutV1NegotiableQuoteQuoteIdShippingMethod

> bool PutV1NegotiableQuoteQuoteIdShippingMethod(ctx, quoteId).PutV1NegotiableQuoteQuoteIdShippingMethodRequest(putV1NegotiableQuoteQuoteIdShippingMethodRequest).Execute()

negotiableQuote/{quoteId}/shippingMethod



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
	quoteId := int32(56) // int32 | Negotiable Quote id
	putV1NegotiableQuoteQuoteIdShippingMethodRequest := *openapiclient.NewPutV1NegotiableQuoteQuoteIdShippingMethodRequest("ShippingMethod_example") // PutV1NegotiableQuoteQuoteIdShippingMethodRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteQuoteIdShippingMethodAPI.PutV1NegotiableQuoteQuoteIdShippingMethod(context.Background(), quoteId).PutV1NegotiableQuoteQuoteIdShippingMethodRequest(putV1NegotiableQuoteQuoteIdShippingMethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteQuoteIdShippingMethodAPI.PutV1NegotiableQuoteQuoteIdShippingMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1NegotiableQuoteQuoteIdShippingMethod`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteQuoteIdShippingMethodAPI.PutV1NegotiableQuoteQuoteIdShippingMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **int32** | Negotiable Quote id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1NegotiableQuoteQuoteIdShippingMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1NegotiableQuoteQuoteIdShippingMethodRequest** | [**PutV1NegotiableQuoteQuoteIdShippingMethodRequest**](PutV1NegotiableQuoteQuoteIdShippingMethodRequest.md) |  | 

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

