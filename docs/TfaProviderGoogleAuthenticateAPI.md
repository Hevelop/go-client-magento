# \TfaProviderGoogleAuthenticateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderGoogleAuthenticate**](TfaProviderGoogleAuthenticateAPI.md#PostV1TfaProviderGoogleAuthenticate) | **Post** /V1/tfa/provider/google/authenticate | tfa/provider/google/authenticate



## PostV1TfaProviderGoogleAuthenticate

> string PostV1TfaProviderGoogleAuthenticate(ctx).PostV1TfaProviderAuthyAuthenticateRequest(postV1TfaProviderAuthyAuthenticateRequest).Execute()

tfa/provider/google/authenticate



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
	postV1TfaProviderAuthyAuthenticateRequest := *openapiclient.NewPostV1TfaProviderAuthyAuthenticateRequest("Username_example", "Password_example", "Otp_example") // PostV1TfaProviderAuthyAuthenticateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderGoogleAuthenticateAPI.PostV1TfaProviderGoogleAuthenticate(context.Background()).PostV1TfaProviderAuthyAuthenticateRequest(postV1TfaProviderAuthyAuthenticateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderGoogleAuthenticateAPI.PostV1TfaProviderGoogleAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderGoogleAuthenticate`: string
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderGoogleAuthenticateAPI.PostV1TfaProviderGoogleAuthenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderGoogleAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderAuthyAuthenticateRequest** | [**PostV1TfaProviderAuthyAuthenticateRequest**](PostV1TfaProviderAuthyAuthenticateRequest.md) |  | 

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

