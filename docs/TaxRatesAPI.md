# \TaxRatesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TaxRates**](TaxRatesAPI.md#PostV1TaxRates) | **Post** /V1/taxRates | taxRates
[**PutV1TaxRates**](TaxRatesAPI.md#PutV1TaxRates) | **Put** /V1/taxRates | taxRates



## PostV1TaxRates

> TaxDataTaxRateInterface PostV1TaxRates(ctx).PutV1TaxRatesRequest(putV1TaxRatesRequest).Execute()

taxRates



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
	putV1TaxRatesRequest := *openapiclient.NewPutV1TaxRatesRequest(*openapiclient.NewTaxDataTaxRateInterface("TaxCountryId_example", float32(123), "Code_example")) // PutV1TaxRatesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesAPI.PostV1TaxRates(context.Background()).PutV1TaxRatesRequest(putV1TaxRatesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesAPI.PostV1TaxRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TaxRates`: TaxDataTaxRateInterface
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesAPI.PostV1TaxRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TaxRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1TaxRatesRequest** | [**PutV1TaxRatesRequest**](PutV1TaxRatesRequest.md) |  | 

### Return type

[**TaxDataTaxRateInterface**](TaxDataTaxRateInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1TaxRates

> TaxDataTaxRateInterface PutV1TaxRates(ctx).PutV1TaxRatesRequest(putV1TaxRatesRequest).Execute()

taxRates



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
	putV1TaxRatesRequest := *openapiclient.NewPutV1TaxRatesRequest(*openapiclient.NewTaxDataTaxRateInterface("TaxCountryId_example", float32(123), "Code_example")) // PutV1TaxRatesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesAPI.PutV1TaxRates(context.Background()).PutV1TaxRatesRequest(putV1TaxRatesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesAPI.PutV1TaxRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1TaxRates`: TaxDataTaxRateInterface
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesAPI.PutV1TaxRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1TaxRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1TaxRatesRequest** | [**PutV1TaxRatesRequest**](PutV1TaxRatesRequest.md) |  | 

### Return type

[**TaxDataTaxRateInterface**](TaxDataTaxRateInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

