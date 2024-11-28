# # TwoFactorAuthProviderInterface
2FA provider interface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled**| **bool** | True if this provider has been enabled by admin  |
**Engine**| [**TwoFactorAuthEngineInterface**](TwoFactorAuthEngineInterface.md) |   |
**Code**| **string** | Provider code  |
**Name**| **string** | Provider name  |
**Icon**| **string** | Icon  |
**ResetAllowed**| **bool** | True if this provider configuration can be reset  |
**ConfigureAction**| **string** | Configure action  |
**AuthAction**| **string** | Auth action  |
**ExtraActions**| **[]string** | Allowed extra actions  |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

