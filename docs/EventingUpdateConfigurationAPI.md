# \EventingUpdateConfigurationAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1EventingUpdateConfiguration**](EventingUpdateConfigurationAPI.md#PutV1EventingUpdateConfiguration) | **Put** /V1/eventing/updateConfiguration | eventing/updateConfiguration



## PutV1EventingUpdateConfiguration

> bool PutV1EventingUpdateConfiguration(ctx).PutV1EventingUpdateConfigurationRequest(putV1EventingUpdateConfigurationRequest).Execute()

eventing/updateConfiguration



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
	putV1EventingUpdateConfigurationRequest := *openapiclient.NewPutV1EventingUpdateConfigurationRequest(*openapiclient.NewAdobeCommerceEventsClientDataConfigurationInterface(false, "MerchantId_example", "EnvironmentId_example", "ProviderId_example", "InstanceId_example", "WorkspaceConfiguration_example")) // PutV1EventingUpdateConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventingUpdateConfigurationAPI.PutV1EventingUpdateConfiguration(context.Background()).PutV1EventingUpdateConfigurationRequest(putV1EventingUpdateConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventingUpdateConfigurationAPI.PutV1EventingUpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1EventingUpdateConfiguration`: bool
	fmt.Fprintf(os.Stdout, "Response from `EventingUpdateConfigurationAPI.PutV1EventingUpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1EventingUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1EventingUpdateConfigurationRequest** | [**PutV1EventingUpdateConfigurationRequest**](PutV1EventingUpdateConfigurationRequest.md) |  | 

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

