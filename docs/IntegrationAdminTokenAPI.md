# \IntegrationAdminTokenAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1IntegrationAdminToken**](IntegrationAdminTokenAPI.md#PostV1IntegrationAdminToken) | **Post** /V1/integration/admin/token | integration/admin/token



## PostV1IntegrationAdminToken

> string PostV1IntegrationAdminToken(ctx).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()

integration/admin/token



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
	postV1IntegrationAdminTokenRequest := *openapiclient.NewPostV1IntegrationAdminTokenRequest("Username_example", "Password_example") // PostV1IntegrationAdminTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAdminTokenAPI.PostV1IntegrationAdminToken(context.Background()).PostV1IntegrationAdminTokenRequest(postV1IntegrationAdminTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAdminTokenAPI.PostV1IntegrationAdminToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1IntegrationAdminToken`: string
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAdminTokenAPI.PostV1IntegrationAdminToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1IntegrationAdminTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1IntegrationAdminTokenRequest** | [**PostV1IntegrationAdminTokenRequest**](PostV1IntegrationAdminTokenRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

