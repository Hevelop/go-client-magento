# # InventoryApiDataSourceInterface
Represents physical storage, i.e. brick and mortar store or warehouse Used fully qualified namespaces in annotations for proper work of WebApi request parser

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceCode**| **string** | Source code  | [optional]
**Name**| **string** | Source name  | [optional]
**Email**| **string** | Source email  | [optional]
**ContactName**| **string** | Source contact name  | [optional]
**Enabled**| **bool** | If source is enabled. For new entity can be null  | [optional]
**Description**| **string** | Source description  | [optional]
**Latitude**| **float32** | Source latitude  | [optional]
**Longitude**| **float32** | Source longitude  | [optional]
**CountryId**| **string** | Source country id  | [optional]
**RegionId**| **int32** | Region id if source has registered region.  | [optional]
**Region**| **string** | Region title if source has custom region  | [optional]
**City**| **string** | Source city  | [optional]
**Street**| **string** | Source street name  | [optional]
**Postcode**| **string** | Source post code  | [optional]
**Phone**| **string** | Source phone number  | [optional]
**Fax**| **string** | Source fax  | [optional]
**UseDefaultCarrierConfig**| **bool** | Is need to use default config  | [optional]
**CarrierLinks**| [**[]InventoryApiDataSourceCarrierLinkInterface**](InventoryApiDataSourceCarrierLinkInterface.md) |   | [optional]
**ExtensionAttributes**| [**InventoryApiDataSourceExtensionInterface**](InventoryApiDataSourceExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

