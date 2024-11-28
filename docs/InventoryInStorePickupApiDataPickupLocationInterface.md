# # InventoryInStorePickupApiDataPickupLocationInterface
Represents sources projection on In-Store Pickup context. Realisation must follow immutable DTO concept. Partial immutability done according to restriction of current Extension Attributes implementation.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupLocationCode**| **string** | Source code of Pickup Location.  |
**Name**| **string** | Pickup Location name.  | [optional]
**Email**| **string** | Pickup Location contact email.  | [optional]
**Fax**| **string** | Fax contact info.  | [optional]
**ContactName**| **string** | Pickup Location contact name.  | [optional]
**Description**| **string** | Pickup Location description.  | [optional]
**Latitude**| **float32** | Pickup Location latitude.  | [optional]
**Longitude**| **float32** | Pickup Location longitude.  | [optional]
**CountryId**| **string** | Pickup Location country ID.  | [optional]
**RegionId**| **int32** | Pickup Location region ID.  | [optional]
**Region**| **string** | Pickup Location region.  | [optional]
**City**| **string** | Pickup Location city.  | [optional]
**Street**| **string** | Pickup Location street.  | [optional]
**Postcode**| **string** | Pickup Location postcode.  | [optional]
**Phone**| **string** | Pickup Location phone.  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\InventoryInStorePickupApi\\Api\\Data\\PickupLocationInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

