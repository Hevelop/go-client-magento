# \NegotiableQuoteRequestAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1NegotiableQuoteRequest**](NegotiableQuoteRequestAPI.md#PostV1NegotiableQuoteRequest) | **Post** /V1/negotiableQuote/request | negotiableQuote/request



## PostV1NegotiableQuoteRequest

> bool PostV1NegotiableQuoteRequest(ctx).PostV1NegotiableQuoteRequestRequest(postV1NegotiableQuoteRequestRequest).Execute()

negotiableQuote/request



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
	postV1NegotiableQuoteRequestRequest := *openapiclient.NewPostV1NegotiableQuoteRequestRequest(int32(123), "QuoteName_example") // PostV1NegotiableQuoteRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteRequestAPI.PostV1NegotiableQuoteRequest(context.Background()).PostV1NegotiableQuoteRequestRequest(postV1NegotiableQuoteRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteRequestAPI.PostV1NegotiableQuoteRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiableQuoteRequest`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteRequestAPI.PostV1NegotiableQuoteRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiableQuoteRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1NegotiableQuoteRequestRequest** | [**PostV1NegotiableQuoteRequestRequest**](PostV1NegotiableQuoteRequestRequest.md) |  | 

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

