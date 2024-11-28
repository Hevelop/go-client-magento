# \AdobeIoEventsCheckConfigurationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1AdobeIoEventsCheckConfiguration**](AdobeIoEventsCheckConfigurationAPI.md#GetV1AdobeIoEventsCheckConfiguration) | **Get** /V1/adobe_io_events/check_configuration | adobe_io_events/check_configuration



## GetV1AdobeIoEventsCheckConfiguration

> AdobeIoEventsClientConfigurationCheckResultInterface GetV1AdobeIoEventsCheckConfiguration(ctx).Execute()

adobe_io_events/check_configuration



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
	resp, r, err := apiClient.AdobeIoEventsCheckConfigurationAPI.GetV1AdobeIoEventsCheckConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdobeIoEventsCheckConfigurationAPI.GetV1AdobeIoEventsCheckConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1AdobeIoEventsCheckConfiguration`: AdobeIoEventsClientConfigurationCheckResultInterface
	fmt.Fprintf(os.Stdout, "Response from `AdobeIoEventsCheckConfigurationAPI.GetV1AdobeIoEventsCheckConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1AdobeIoEventsCheckConfigurationRequest struct via the builder pattern


### Return type

[**AdobeIoEventsClientConfigurationCheckResultInterface**](AdobeIoEventsClientConfigurationCheckResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

