# # CustomerDataAddressInterface
Customer address interface.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | ID  | [optional]
**CustomerId**| **int32** | Customer ID  | [optional]
**Region**| [**CustomerDataRegionInterface**](CustomerDataRegionInterface.md) |   | [optional]
**RegionId**| **int32** | Region ID  | [optional]
**CountryId**| **string** | Country code in ISO_3166-2 format  | [optional]
**Street**| **[]string** | Street  | [optional]
**Company**| **string** | Company  | [optional]
**Telephone**| **string** | Telephone number  | [optional]
**Fax**| **string** | Fax number  | [optional]
**Postcode**| **string** | Postcode  | [optional]
**City**| **string** | City name  | [optional]
**Firstname**| **string** | First name  | [optional]
**Lastname**| **string** | Last name  | [optional]
**Middlename**| **string** | Middle name  | [optional]
**Prefix**| **string** | Prefix  | [optional]
**Suffix**| **string** | Suffix  | [optional]
**VatId**| **string** | Vat id  | [optional]
**DefaultShipping**| **bool** | If this address is default shipping address.  | [optional]
**DefaultBilling**| **bool** | If this address is default billing address  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Customer\\Api\\Data\\AddressInterface  | [optional]
**CustomAttributes**| [**[]FrameworkAttributeInterface**](FrameworkAttributeInterface.md) | Custom attributes values.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

