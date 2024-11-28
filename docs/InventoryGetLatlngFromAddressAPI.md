# \InventoryGetLatlngFromAddressAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryGetlatlngfromaddress**](InventoryGetLatlngFromAddressAPI.md#GetV1InventoryGetlatlngfromaddress) | **Get** /V1/inventory/get-latlng-from-address | inventory/get-latlng-from-address



## GetV1InventoryGetlatlngfromaddress

> InventoryDistanceBasedSourceSelectionApiDataLatLngInterface GetV1InventoryGetlatlngfromaddress(ctx).AddressCountry(addressCountry).AddressPostcode(addressPostcode).AddressStreet(addressStreet).AddressRegion(addressRegion).AddressCity(addressCity).Execute()

inventory/get-latlng-from-address



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
	addressCountry := "addressCountry_example" // string | Shipping country (optional)
	addressPostcode := "addressPostcode_example" // string | Shipping postcode (optional)
	addressStreet := "addressStreet_example" // string | Shipping street address (optional)
	addressRegion := "addressRegion_example" // string | Shipping region (optional)
	addressCity := "addressCity_example" // string | Shipping city (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryGetLatlngFromAddressAPI.GetV1InventoryGetlatlngfromaddress(context.Background()).AddressCountry(addressCountry).AddressPostcode(addressPostcode).AddressStreet(addressStreet).AddressRegion(addressRegion).AddressCity(addressCity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGetLatlngFromAddressAPI.GetV1InventoryGetlatlngfromaddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryGetlatlngfromaddress`: InventoryDistanceBasedSourceSelectionApiDataLatLngInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryGetLatlngFromAddressAPI.GetV1InventoryGetlatlngfromaddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryGetlatlngfromaddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressCountry** | **string** | Shipping country | 
 **addressPostcode** | **string** | Shipping postcode | 
 **addressStreet** | **string** | Shipping street address | 
 **addressRegion** | **string** | Shipping region | 
 **addressCity** | **string** | Shipping city | 

### Return type

[**InventoryDistanceBasedSourceSelectionApiDataLatLngInterface**](InventoryDistanceBasedSourceSelectionApiDataLatLngInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

