# \TfaProviderU2fkeyVerifyAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderU2fkeyVerify**](TfaProviderU2fkeyVerifyAPI.md#PostV1TfaProviderU2fkeyVerify) | **Post** /V1/tfa/provider/u2fkey/verify | tfa/provider/u2fkey/verify



## PostV1TfaProviderU2fkeyVerify

> string PostV1TfaProviderU2fkeyVerify(ctx).PostV1TfaProviderU2fkeyVerifyRequest(postV1TfaProviderU2fkeyVerifyRequest).Execute()

tfa/provider/u2fkey/verify



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
	postV1TfaProviderU2fkeyVerifyRequest := *openapiclient.NewPostV1TfaProviderU2fkeyVerifyRequest("Username_example", "Password_example", "PublicKeyCredentialJson_example") // PostV1TfaProviderU2fkeyVerifyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderU2fkeyVerifyAPI.PostV1TfaProviderU2fkeyVerify(context.Background()).PostV1TfaProviderU2fkeyVerifyRequest(postV1TfaProviderU2fkeyVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderU2fkeyVerifyAPI.PostV1TfaProviderU2fkeyVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderU2fkeyVerify`: string
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderU2fkeyVerifyAPI.PostV1TfaProviderU2fkeyVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderU2fkeyVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderU2fkeyVerifyRequest** | [**PostV1TfaProviderU2fkeyVerifyRequest**](PostV1TfaProviderU2fkeyVerifyRequest.md) |  | 

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

