# \CartsLicenceAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsLicence**](CartsLicenceAPI.md#GetV1CartsLicence) | **Get** /V1/carts/licence | carts/licence



## GetV1CartsLicence

> []CheckoutAgreementsDataAgreementInterface GetV1CartsLicence(ctx).Execute()

carts/licence



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsLicenceAPI.GetV1CartsLicence(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsLicenceAPI.GetV1CartsLicence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsLicence`: []CheckoutAgreementsDataAgreementInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsLicenceAPI.GetV1CartsLicence`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsLicenceRequest struct via the builder pattern


### Return type

[**[]CheckoutAgreementsDataAgreementInterface**](CheckoutAgreementsDataAgreementInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

