# \TfaProviderDuoSecurityAuthenticateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderDuoSecurityAuthenticate**](TfaProviderDuoSecurityAuthenticateAPI.md#PostV1TfaProviderDuoSecurityAuthenticate) | **Post** /V1/tfa/provider/duo_security/authenticate | tfa/provider/duo_security/authenticate



## PostV1TfaProviderDuoSecurityAuthenticate

> string PostV1TfaProviderDuoSecurityAuthenticate(ctx).PostV1TfaProviderDuoSecurityAuthenticateRequest(postV1TfaProviderDuoSecurityAuthenticateRequest).Execute()

tfa/provider/duo_security/authenticate



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
	postV1TfaProviderDuoSecurityAuthenticateRequest := *openapiclient.NewPostV1TfaProviderDuoSecurityAuthenticateRequest("Username_example", "Password_example", "SignatureResponse_example") // PostV1TfaProviderDuoSecurityAuthenticateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderDuoSecurityAuthenticateAPI.PostV1TfaProviderDuoSecurityAuthenticate(context.Background()).PostV1TfaProviderDuoSecurityAuthenticateRequest(postV1TfaProviderDuoSecurityAuthenticateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderDuoSecurityAuthenticateAPI.PostV1TfaProviderDuoSecurityAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderDuoSecurityAuthenticate`: string
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderDuoSecurityAuthenticateAPI.PostV1TfaProviderDuoSecurityAuthenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderDuoSecurityAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderDuoSecurityAuthenticateRequest** | [**PostV1TfaProviderDuoSecurityAuthenticateRequest**](PostV1TfaProviderDuoSecurityAuthenticateRequest.md) |  | 

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

