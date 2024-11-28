# # DirectoryDataCurrencyInformationInterface
Currency Information interface.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCurrencyCode**| **string** | The base currency code for the store.  |
**BaseCurrencySymbol**| **string** | The currency symbol of the base currency for the store.  |
**DefaultDisplayCurrencyCode**| **string** | The default display currency code for the store.  |
**DefaultDisplayCurrencySymbol**| **string** | The currency symbol of the default display currency for the store.  |
**AvailableCurrencyCodes**| **[]string** | The list of allowed currency codes for the store.  |
**ExchangeRates**| [**[]DirectoryDataExchangeRateInterface**](DirectoryDataExchangeRateInterface.md) | The list of exchange rate information for the store.  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Directory\\Api\\Data\\CurrencyInformationInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

