# \TfaProviderU2fkeyActivateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderU2fkeyActivate**](TfaProviderU2fkeyActivateAPI.md#PostV1TfaProviderU2fkeyActivate) | **Post** /V1/tfa/provider/u2fkey/activate | tfa/provider/u2fkey/activate



## PostV1TfaProviderU2fkeyActivate

> ErrorResponse PostV1TfaProviderU2fkeyActivate(ctx).PostV1TfaProviderU2fkeyActivateRequest(postV1TfaProviderU2fkeyActivateRequest).Execute()

tfa/provider/u2fkey/activate



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
	postV1TfaProviderU2fkeyActivateRequest := *openapiclient.NewPostV1TfaProviderU2fkeyActivateRequest("TfaToken_example", "PublicKeyCredentialJson_example") // PostV1TfaProviderU2fkeyActivateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderU2fkeyActivateAPI.PostV1TfaProviderU2fkeyActivate(context.Background()).PostV1TfaProviderU2fkeyActivateRequest(postV1TfaProviderU2fkeyActivateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderU2fkeyActivateAPI.PostV1TfaProviderU2fkeyActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderU2fkeyActivate`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderU2fkeyActivateAPI.PostV1TfaProviderU2fkeyActivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderU2fkeyActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderU2fkeyActivateRequest** | [**PostV1TfaProviderU2fkeyActivateRequest**](PostV1TfaProviderU2fkeyActivateRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

