# \NegotiableQuoteAttachmentContentAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1NegotiableQuoteAttachmentContent**](NegotiableQuoteAttachmentContentAPI.md#GetV1NegotiableQuoteAttachmentContent) | **Get** /V1/negotiableQuote/attachmentContent | negotiableQuote/attachmentContent



## GetV1NegotiableQuoteAttachmentContent

> []NegotiableQuoteDataAttachmentContentInterface GetV1NegotiableQuoteAttachmentContent(ctx).AttachmentIds(attachmentIds).Execute()

negotiableQuote/attachmentContent



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
	attachmentIds := []int32{int32(123)} // []int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteAttachmentContentAPI.GetV1NegotiableQuoteAttachmentContent(context.Background()).AttachmentIds(attachmentIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteAttachmentContentAPI.GetV1NegotiableQuoteAttachmentContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1NegotiableQuoteAttachmentContent`: []NegotiableQuoteDataAttachmentContentInterface
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteAttachmentContentAPI.GetV1NegotiableQuoteAttachmentContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1NegotiableQuoteAttachmentContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentIds** | **[]int32** |  | 

### Return type

[**[]NegotiableQuoteDataAttachmentContentInterface**](NegotiableQuoteDataAttachmentContentInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

