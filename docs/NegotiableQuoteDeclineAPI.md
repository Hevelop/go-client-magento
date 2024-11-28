# \NegotiableQuoteDeclineAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1NegotiableQuoteDecline**](NegotiableQuoteDeclineAPI.md#PostV1NegotiableQuoteDecline) | **Post** /V1/negotiableQuote/decline | negotiableQuote/decline



## PostV1NegotiableQuoteDecline

> bool PostV1NegotiableQuoteDecline(ctx).PostV1NegotiableQuoteDeclineRequest(postV1NegotiableQuoteDeclineRequest).Execute()

negotiableQuote/decline



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
	postV1NegotiableQuoteDeclineRequest := *openapiclient.NewPostV1NegotiableQuoteDeclineRequest(int32(123), "Reason_example") // PostV1NegotiableQuoteDeclineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteDeclineAPI.PostV1NegotiableQuoteDecline(context.Background()).PostV1NegotiableQuoteDeclineRequest(postV1NegotiableQuoteDeclineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteDeclineAPI.PostV1NegotiableQuoteDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiableQuoteDecline`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteDeclineAPI.PostV1NegotiableQuoteDecline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiableQuoteDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1NegotiableQuoteDeclineRequest** | [**PostV1NegotiableQuoteDeclineRequest**](PostV1NegotiableQuoteDeclineRequest.md) |  | 

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

