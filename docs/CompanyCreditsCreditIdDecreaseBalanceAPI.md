# \CompanyCreditsCreditIdDecreaseBalanceAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CompanyCreditsCreditIdDecreaseBalance**](CompanyCreditsCreditIdDecreaseBalanceAPI.md#PostV1CompanyCreditsCreditIdDecreaseBalance) | **Post** /V1/companyCredits/{creditId}/decreaseBalance | companyCredits/{creditId}/decreaseBalance



## PostV1CompanyCreditsCreditIdDecreaseBalance

> bool PostV1CompanyCreditsCreditIdDecreaseBalance(ctx, creditId).PostV1CompanyCreditsCreditIdDecreaseBalanceRequest(postV1CompanyCreditsCreditIdDecreaseBalanceRequest).Execute()

companyCredits/{creditId}/decreaseBalance



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
	creditId := int32(56) // int32 | 
	postV1CompanyCreditsCreditIdDecreaseBalanceRequest := *openapiclient.NewPostV1CompanyCreditsCreditIdDecreaseBalanceRequest(float32(123), "Currency_example", int32(123)) // PostV1CompanyCreditsCreditIdDecreaseBalanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCreditsCreditIdDecreaseBalanceAPI.PostV1CompanyCreditsCreditIdDecreaseBalance(context.Background(), creditId).PostV1CompanyCreditsCreditIdDecreaseBalanceRequest(postV1CompanyCreditsCreditIdDecreaseBalanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCreditsCreditIdDecreaseBalanceAPI.PostV1CompanyCreditsCreditIdDecreaseBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CompanyCreditsCreditIdDecreaseBalance`: bool
	fmt.Fprintf(os.Stdout, "Response from `CompanyCreditsCreditIdDecreaseBalanceAPI.PostV1CompanyCreditsCreditIdDecreaseBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**creditId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CompanyCreditsCreditIdDecreaseBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CompanyCreditsCreditIdDecreaseBalanceRequest** | [**PostV1CompanyCreditsCreditIdDecreaseBalanceRequest**](PostV1CompanyCreditsCreditIdDecreaseBalanceRequest.md) |  | 

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

