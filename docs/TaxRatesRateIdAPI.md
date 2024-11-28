# \TaxRatesRateIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1TaxRatesRateId**](TaxRatesRateIdAPI.md#DeleteV1TaxRatesRateId) | **Delete** /V1/taxRates/{rateId} | taxRates/{rateId}
[**GetV1TaxRatesRateId**](TaxRatesRateIdAPI.md#GetV1TaxRatesRateId) | **Get** /V1/taxRates/{rateId} | taxRates/{rateId}



## DeleteV1TaxRatesRateId

> bool DeleteV1TaxRatesRateId(ctx, rateId).Execute()

taxRates/{rateId}



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
	rateId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesRateIdAPI.DeleteV1TaxRatesRateId(context.Background(), rateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesRateIdAPI.DeleteV1TaxRatesRateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1TaxRatesRateId`: bool
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesRateIdAPI.DeleteV1TaxRatesRateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1TaxRatesRateIdRequest struct via the builder pattern


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


## GetV1TaxRatesRateId

> TaxDataTaxRateInterface GetV1TaxRatesRateId(ctx, rateId).Execute()

taxRates/{rateId}



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
	rateId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesRateIdAPI.GetV1TaxRatesRateId(context.Background(), rateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesRateIdAPI.GetV1TaxRatesRateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1TaxRatesRateId`: TaxDataTaxRateInterface
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesRateIdAPI.GetV1TaxRatesRateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rateId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1TaxRatesRateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxDataTaxRateInterface**](TaxDataTaxRateInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

