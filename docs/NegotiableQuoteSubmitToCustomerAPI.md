# \NegotiableQuoteSubmitToCustomerAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1NegotiableQuoteSubmitToCustomer**](NegotiableQuoteSubmitToCustomerAPI.md#PostV1NegotiableQuoteSubmitToCustomer) | **Post** /V1/negotiableQuote/submitToCustomer | negotiableQuote/submitToCustomer



## PostV1NegotiableQuoteSubmitToCustomer

> bool PostV1NegotiableQuoteSubmitToCustomer(ctx).PostV1NegotiableQuoteSubmitToCustomerRequest(postV1NegotiableQuoteSubmitToCustomerRequest).Execute()

negotiableQuote/submitToCustomer



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
	postV1NegotiableQuoteSubmitToCustomerRequest := *openapiclient.NewPostV1NegotiableQuoteSubmitToCustomerRequest(int32(123)) // PostV1NegotiableQuoteSubmitToCustomerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteSubmitToCustomerAPI.PostV1NegotiableQuoteSubmitToCustomer(context.Background()).PostV1NegotiableQuoteSubmitToCustomerRequest(postV1NegotiableQuoteSubmitToCustomerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteSubmitToCustomerAPI.PostV1NegotiableQuoteSubmitToCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiableQuoteSubmitToCustomer`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteSubmitToCustomerAPI.PostV1NegotiableQuoteSubmitToCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiableQuoteSubmitToCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1NegotiableQuoteSubmitToCustomerRequest** | [**PostV1NegotiableQuoteSubmitToCustomerRequest**](PostV1NegotiableQuoteSubmitToCustomerRequest.md) |  | 

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

