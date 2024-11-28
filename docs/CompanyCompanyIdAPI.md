# \CompanyCompanyIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CompanyCompanyId**](CompanyCompanyIdAPI.md#DeleteV1CompanyCompanyId) | **Delete** /V1/company/{companyId} | company/{companyId}
[**GetV1CompanyCompanyId**](CompanyCompanyIdAPI.md#GetV1CompanyCompanyId) | **Get** /V1/company/{companyId} | company/{companyId}
[**PutV1CompanyCompanyId**](CompanyCompanyIdAPI.md#PutV1CompanyCompanyId) | **Put** /V1/company/{companyId} | company/{companyId}



## DeleteV1CompanyCompanyId

> bool DeleteV1CompanyCompanyId(ctx, companyId).Execute()

company/{companyId}



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
	companyId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCompanyIdAPI.DeleteV1CompanyCompanyId(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyIdAPI.DeleteV1CompanyCompanyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CompanyCompanyId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyIdAPI.DeleteV1CompanyCompanyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CompanyCompanyIdRequest struct via the builder pattern


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


## GetV1CompanyCompanyId

> CompanyDataCompanyInterface GetV1CompanyCompanyId(ctx, companyId).Execute()

company/{companyId}



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
	companyId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCompanyIdAPI.GetV1CompanyCompanyId(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyIdAPI.GetV1CompanyCompanyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CompanyCompanyId`: CompanyDataCompanyInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyIdAPI.GetV1CompanyCompanyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CompanyCompanyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyDataCompanyInterface**](CompanyDataCompanyInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CompanyCompanyId

> CompanyDataCompanyInterface PutV1CompanyCompanyId(ctx, companyId).PostV1CompanyRequest(postV1CompanyRequest).Execute()

company/{companyId}



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
	companyId := "companyId_example" // string | 
	postV1CompanyRequest := *openapiclient.NewPostV1CompanyRequest(*openapiclient.NewCompanyDataCompanyInterface([]string{"Street_example"}, int32(123), int32(123), "RejectReason_example", "RejectedAt_example", int32(123))) // PostV1CompanyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCompanyIdAPI.PutV1CompanyCompanyId(context.Background(), companyId).PostV1CompanyRequest(postV1CompanyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyIdAPI.PutV1CompanyCompanyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CompanyCompanyId`: CompanyDataCompanyInterface
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyIdAPI.PutV1CompanyCompanyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CompanyCompanyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CompanyRequest** | [**PostV1CompanyRequest**](PostV1CompanyRequest.md) |  | 

### Return type

[**CompanyDataCompanyInterface**](CompanyDataCompanyInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

