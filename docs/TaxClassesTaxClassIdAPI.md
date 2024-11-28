# \TaxClassesTaxClassIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1TaxClassesTaxClassId**](TaxClassesTaxClassIdAPI.md#DeleteV1TaxClassesTaxClassId) | **Delete** /V1/taxClasses/{taxClassId} | taxClasses/{taxClassId}
[**GetV1TaxClassesTaxClassId**](TaxClassesTaxClassIdAPI.md#GetV1TaxClassesTaxClassId) | **Get** /V1/taxClasses/{taxClassId} | taxClasses/{taxClassId}



## DeleteV1TaxClassesTaxClassId

> bool DeleteV1TaxClassesTaxClassId(ctx, taxClassId).Execute()

taxClasses/{taxClassId}



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
	taxClassId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxClassesTaxClassIdAPI.DeleteV1TaxClassesTaxClassId(context.Background(), taxClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxClassesTaxClassIdAPI.DeleteV1TaxClassesTaxClassId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1TaxClassesTaxClassId`: bool
	fmt.Fprintf(os.Stdout, "Response from `TaxClassesTaxClassIdAPI.DeleteV1TaxClassesTaxClassId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxClassId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1TaxClassesTaxClassIdRequest struct via the builder pattern


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


## GetV1TaxClassesTaxClassId

> TaxDataTaxClassInterface GetV1TaxClassesTaxClassId(ctx, taxClassId).Execute()

taxClasses/{taxClassId}



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
	taxClassId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxClassesTaxClassIdAPI.GetV1TaxClassesTaxClassId(context.Background(), taxClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxClassesTaxClassIdAPI.GetV1TaxClassesTaxClassId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TaxClassesTaxClassId`: TaxDataTaxClassInterface
	fmt.Fprintf(os.Stdout, "Response from `TaxClassesTaxClassIdAPI.GetV1TaxClassesTaxClassId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxClassId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TaxClassesTaxClassIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxDataTaxClassInterface**](TaxDataTaxClassInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

