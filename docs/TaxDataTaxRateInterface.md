# # TaxDataTaxRateInterface
Tax rate interface.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | Id  | [optional]
**TaxCountryId**| **string** | Country id  |
**TaxRegionId**| **int32** | Region id  | [optional]
**RegionName**| **string** | Region name  | [optional]
**TaxPostcode**| **string** | Postcode  | [optional]
**ZipIsRange**| **int32** | Zip is range  | [optional]
**ZipFrom**| **int32** | Zip range from  | [optional]
**ZipTo**| **int32** | Zip range to  | [optional]
**Rate**| **float32** | Tax rate in percentage  |
**Code**| **string** | Tax rate code  |
**Titles**| [**[]TaxDataTaxRateTitleInterface**](TaxDataTaxRateTitleInterface.md) | Tax rate titles  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Tax\\Api\\Data\\TaxRateInterface  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

