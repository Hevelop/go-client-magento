# \NegotiableQuoteDraftAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1NegotiableQuoteDraft**](NegotiableQuoteDraftAPI.md#PostV1NegotiableQuoteDraft) | **Post** /V1/negotiableQuote/draft | negotiableQuote/draft



## PostV1NegotiableQuoteDraft

> int32 PostV1NegotiableQuoteDraft(ctx).PostV1NegotiableQuoteDraftRequest(postV1NegotiableQuoteDraftRequest).Execute()

negotiableQuote/draft



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
	postV1NegotiableQuoteDraftRequest := *openapiclient.NewPostV1NegotiableQuoteDraftRequest(int32(123)) // PostV1NegotiableQuoteDraftRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteDraftAPI.PostV1NegotiableQuoteDraft(context.Background()).PostV1NegotiableQuoteDraftRequest(postV1NegotiableQuoteDraftRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteDraftAPI.PostV1NegotiableQuoteDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiableQuoteDraft`: int32
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteDraftAPI.PostV1NegotiableQuoteDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiableQuoteDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1NegotiableQuoteDraftRequest** | [**PostV1NegotiableQuoteDraftRequest**](PostV1NegotiableQuoteDraftRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

