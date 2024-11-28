# \TfaProviderDuoSecurityActivateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TfaProviderDuoSecurityActivate**](TfaProviderDuoSecurityActivateAPI.md#PostV1TfaProviderDuoSecurityActivate) | **Post** /V1/tfa/provider/duo_security/activate | tfa/provider/duo_security/activate



## PostV1TfaProviderDuoSecurityActivate

> ErrorResponse PostV1TfaProviderDuoSecurityActivate(ctx).PostV1TfaProviderDuoSecurityActivateRequest(postV1TfaProviderDuoSecurityActivateRequest).Execute()

tfa/provider/duo_security/activate



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
	postV1TfaProviderDuoSecurityActivateRequest := *openapiclient.NewPostV1TfaProviderDuoSecurityActivateRequest("TfaToken_example", "SignatureResponse_example") // PostV1TfaProviderDuoSecurityActivateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaProviderDuoSecurityActivateAPI.PostV1TfaProviderDuoSecurityActivate(context.Background()).PostV1TfaProviderDuoSecurityActivateRequest(postV1TfaProviderDuoSecurityActivateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaProviderDuoSecurityActivateAPI.PostV1TfaProviderDuoSecurityActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TfaProviderDuoSecurityActivate`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `TfaProviderDuoSecurityActivateAPI.PostV1TfaProviderDuoSecurityActivate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TfaProviderDuoSecurityActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1TfaProviderDuoSecurityActivateRequest** | [**PostV1TfaProviderDuoSecurityActivateRequest**](PostV1TfaProviderDuoSecurityActivateRequest.md) |  | 

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

