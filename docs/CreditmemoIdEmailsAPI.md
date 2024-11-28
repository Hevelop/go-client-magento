# \CreditmemoIdEmailsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CreditmemoIdEmails**](CreditmemoIdEmailsAPI.md#PostV1CreditmemoIdEmails) | **Post** /V1/creditmemo/{id}/emails | creditmemo/{id}/emails



## PostV1CreditmemoIdEmails

> bool PostV1CreditmemoIdEmails(ctx, id).Execute()

creditmemo/{id}/emails



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
	id := int32(56) // int32 | The credit memo ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditmemoIdEmailsAPI.PostV1CreditmemoIdEmails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditmemoIdEmailsAPI.PostV1CreditmemoIdEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CreditmemoIdEmails`: bool
	fmt.Fprintf(os.Stdout, "Response from `CreditmemoIdEmailsAPI.PostV1CreditmemoIdEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The credit memo ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CreditmemoIdEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

