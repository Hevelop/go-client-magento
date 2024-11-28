/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
	"fmt"
)

// checks if the SalesDataOrderAddressInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesDataOrderAddressInterface{}

// SalesDataOrderAddressInterface Order address interface. An order is a document that a web store issues to a customer. Magento generates a sales order that lists the product items, billing and shipping addresses, and shipping and payment methods. A corresponding external document, known as a purchase order, is emailed to the customer.
type SalesDataOrderAddressInterface struct {
	// Address type.
	AddressType string `json:"address_type"`
	// City.
	City string `json:"city"`
	// Company.
	Company *string `json:"company,omitempty"`
	// Country ID.
	CountryId string `json:"country_id"`
	// Country address ID.
	CustomerAddressId *int32 `json:"customer_address_id,omitempty"`
	// Customer ID.
	CustomerId *int32 `json:"customer_id,omitempty"`
	// Email address.
	Email *string `json:"email,omitempty"`
	// Order address ID.
	EntityId *int32 `json:"entity_id,omitempty"`
	// Fax number.
	Fax *string `json:"fax,omitempty"`
	// First name.
	Firstname string `json:"firstname"`
	// Last name.
	Lastname string `json:"lastname"`
	// Middle name.
	Middlename *string `json:"middlename,omitempty"`
	// Parent ID.
	ParentId *int32 `json:"parent_id,omitempty"`
	// Postal code.
	Postcode string `json:"postcode"`
	// Prefix.
	Prefix *string `json:"prefix,omitempty"`
	// Region.
	Region *string `json:"region,omitempty"`
	// Region code.
	RegionCode *string `json:"region_code,omitempty"`
	// Region ID.
	RegionId *int32 `json:"region_id,omitempty"`
	// Array of any street values. Otherwise, null.
	Street []string `json:"street,omitempty"`
	// Suffix.
	Suffix *string `json:"suffix,omitempty"`
	// Telephone number.
	Telephone string `json:"telephone"`
	// VAT ID.
	VatId *string `json:"vat_id,omitempty"`
	// VAT-is-valid flag value.
	VatIsValid *int32 `json:"vat_is_valid,omitempty"`
	// VAT request date.
	VatRequestDate *string `json:"vat_request_date,omitempty"`
	// VAT request ID.
	VatRequestId *string `json:"vat_request_id,omitempty"`
	// VAT-request-success flag value.
	VatRequestSuccess *int32 `json:"vat_request_success,omitempty"`
	// ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\OrderAddressInterface
	ExtensionAttributes  map[string]interface{} `json:"extension_attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesDataOrderAddressInterface SalesDataOrderAddressInterface

// NewSalesDataOrderAddressInterface instantiates a new SalesDataOrderAddressInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesDataOrderAddressInterface(addressType string, city string, countryId string, firstname string, lastname string, postcode string, telephone string) *SalesDataOrderAddressInterface {
	this := SalesDataOrderAddressInterface{}
	this.AddressType = addressType
	this.City = city
	this.CountryId = countryId
	this.Firstname = firstname
	this.Lastname = lastname
	this.Postcode = postcode
	this.Telephone = telephone
	return &this
}

// NewSalesDataOrderAddressInterfaceWithDefaults instantiates a new SalesDataOrderAddressInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesDataOrderAddressInterfaceWithDefaults() *SalesDataOrderAddressInterface {
	this := SalesDataOrderAddressInterface{}
	return &this
}

// GetAddressType returns the AddressType field value
func (o *SalesDataOrderAddressInterface) GetAddressType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressType
}

// GetAddressTypeOk returns a tuple with the AddressType field value
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetAddressTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressType, true
}

// SetAddressType sets field value
func (o *SalesDataOrderAddressInterface) SetAddressType(v string) {
	o.AddressType = v
}

// GetCity returns the City field value
func (o *SalesDataOrderAddressInterface) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *SalesDataOrderAddressInterface) SetCity(v string) {
	o.City = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *SalesDataOrderAddressInterface) SetCompany(v string) {
	o.Company = &v
}

// GetCountryId returns the CountryId field value
func (o *SalesDataOrderAddressInterface) GetCountryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetCountryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryId, true
}

// SetCountryId sets field value
func (o *SalesDataOrderAddressInterface) SetCountryId(v string) {
	o.CountryId = v
}

// GetCustomerAddressId returns the CustomerAddressId field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetCustomerAddressId() int32 {
	if o == nil || IsNil(o.CustomerAddressId) {
		var ret int32
		return ret
	}
	return *o.CustomerAddressId
}

// GetCustomerAddressIdOk returns a tuple with the CustomerAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetCustomerAddressIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerAddressId) {
		return nil, false
	}
	return o.CustomerAddressId, true
}

// HasCustomerAddressId returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasCustomerAddressId() bool {
	if o != nil && !IsNil(o.CustomerAddressId) {
		return true
	}

	return false
}

// SetCustomerAddressId gets a reference to the given int32 and assigns it to the CustomerAddressId field.
func (o *SalesDataOrderAddressInterface) SetCustomerAddressId(v int32) {
	o.CustomerAddressId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetCustomerId() int32 {
	if o == nil || IsNil(o.CustomerId) {
		var ret int32
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetCustomerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given int32 and assigns it to the CustomerId field.
func (o *SalesDataOrderAddressInterface) SetCustomerId(v int32) {
	o.CustomerId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SalesDataOrderAddressInterface) SetEmail(v string) {
	o.Email = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetEntityId() int32 {
	if o == nil || IsNil(o.EntityId) {
		var ret int32
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given int32 and assigns it to the EntityId field.
func (o *SalesDataOrderAddressInterface) SetEntityId(v int32) {
	o.EntityId = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetFax() string {
	if o == nil || IsNil(o.Fax) {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetFaxOk() (*string, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *SalesDataOrderAddressInterface) SetFax(v string) {
	o.Fax = &v
}

// GetFirstname returns the Firstname field value
func (o *SalesDataOrderAddressInterface) GetFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firstname, true
}

// SetFirstname sets field value
func (o *SalesDataOrderAddressInterface) SetFirstname(v string) {
	o.Firstname = v
}

// GetLastname returns the Lastname field value
func (o *SalesDataOrderAddressInterface) GetLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lastname, true
}

// SetLastname sets field value
func (o *SalesDataOrderAddressInterface) SetLastname(v string) {
	o.Lastname = v
}

// GetMiddlename returns the Middlename field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetMiddlename() string {
	if o == nil || IsNil(o.Middlename) {
		var ret string
		return ret
	}
	return *o.Middlename
}

// GetMiddlenameOk returns a tuple with the Middlename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetMiddlenameOk() (*string, bool) {
	if o == nil || IsNil(o.Middlename) {
		return nil, false
	}
	return o.Middlename, true
}

// HasMiddlename returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasMiddlename() bool {
	if o != nil && !IsNil(o.Middlename) {
		return true
	}

	return false
}

// SetMiddlename gets a reference to the given string and assigns it to the Middlename field.
func (o *SalesDataOrderAddressInterface) SetMiddlename(v string) {
	o.Middlename = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *SalesDataOrderAddressInterface) SetParentId(v int32) {
	o.ParentId = &v
}

// GetPostcode returns the Postcode field value
func (o *SalesDataOrderAddressInterface) GetPostcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Postcode
}

// GetPostcodeOk returns a tuple with the Postcode field value
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetPostcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Postcode, true
}

// SetPostcode sets field value
func (o *SalesDataOrderAddressInterface) SetPostcode(v string) {
	o.Postcode = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *SalesDataOrderAddressInterface) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *SalesDataOrderAddressInterface) SetRegion(v string) {
	o.Region = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *SalesDataOrderAddressInterface) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetRegionId() int32 {
	if o == nil || IsNil(o.RegionId) {
		var ret int32
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int32 and assigns it to the RegionId field.
func (o *SalesDataOrderAddressInterface) SetRegionId(v int32) {
	o.RegionId = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetStreet() []string {
	if o == nil || IsNil(o.Street) {
		var ret []string
		return ret
	}
	return o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetStreetOk() ([]string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given []string and assigns it to the Street field.
func (o *SalesDataOrderAddressInterface) SetStreet(v []string) {
	o.Street = v
}

// GetSuffix returns the Suffix field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetSuffix() string {
	if o == nil || IsNil(o.Suffix) {
		var ret string
		return ret
	}
	return *o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.Suffix) {
		return nil, false
	}
	return o.Suffix, true
}

// HasSuffix returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasSuffix() bool {
	if o != nil && !IsNil(o.Suffix) {
		return true
	}

	return false
}

// SetSuffix gets a reference to the given string and assigns it to the Suffix field.
func (o *SalesDataOrderAddressInterface) SetSuffix(v string) {
	o.Suffix = &v
}

// GetTelephone returns the Telephone field value
func (o *SalesDataOrderAddressInterface) GetTelephone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetTelephoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Telephone, true
}

// SetTelephone sets field value
func (o *SalesDataOrderAddressInterface) SetTelephone(v string) {
	o.Telephone = v
}

// GetVatId returns the VatId field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetVatId() string {
	if o == nil || IsNil(o.VatId) {
		var ret string
		return ret
	}
	return *o.VatId
}

// GetVatIdOk returns a tuple with the VatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetVatIdOk() (*string, bool) {
	if o == nil || IsNil(o.VatId) {
		return nil, false
	}
	return o.VatId, true
}

// HasVatId returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasVatId() bool {
	if o != nil && !IsNil(o.VatId) {
		return true
	}

	return false
}

// SetVatId gets a reference to the given string and assigns it to the VatId field.
func (o *SalesDataOrderAddressInterface) SetVatId(v string) {
	o.VatId = &v
}

// GetVatIsValid returns the VatIsValid field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetVatIsValid() int32 {
	if o == nil || IsNil(o.VatIsValid) {
		var ret int32
		return ret
	}
	return *o.VatIsValid
}

// GetVatIsValidOk returns a tuple with the VatIsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetVatIsValidOk() (*int32, bool) {
	if o == nil || IsNil(o.VatIsValid) {
		return nil, false
	}
	return o.VatIsValid, true
}

// HasVatIsValid returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasVatIsValid() bool {
	if o != nil && !IsNil(o.VatIsValid) {
		return true
	}

	return false
}

// SetVatIsValid gets a reference to the given int32 and assigns it to the VatIsValid field.
func (o *SalesDataOrderAddressInterface) SetVatIsValid(v int32) {
	o.VatIsValid = &v
}

// GetVatRequestDate returns the VatRequestDate field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetVatRequestDate() string {
	if o == nil || IsNil(o.VatRequestDate) {
		var ret string
		return ret
	}
	return *o.VatRequestDate
}

// GetVatRequestDateOk returns a tuple with the VatRequestDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetVatRequestDateOk() (*string, bool) {
	if o == nil || IsNil(o.VatRequestDate) {
		return nil, false
	}
	return o.VatRequestDate, true
}

// HasVatRequestDate returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasVatRequestDate() bool {
	if o != nil && !IsNil(o.VatRequestDate) {
		return true
	}

	return false
}

// SetVatRequestDate gets a reference to the given string and assigns it to the VatRequestDate field.
func (o *SalesDataOrderAddressInterface) SetVatRequestDate(v string) {
	o.VatRequestDate = &v
}

// GetVatRequestId returns the VatRequestId field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetVatRequestId() string {
	if o == nil || IsNil(o.VatRequestId) {
		var ret string
		return ret
	}
	return *o.VatRequestId
}

// GetVatRequestIdOk returns a tuple with the VatRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetVatRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.VatRequestId) {
		return nil, false
	}
	return o.VatRequestId, true
}

// HasVatRequestId returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasVatRequestId() bool {
	if o != nil && !IsNil(o.VatRequestId) {
		return true
	}

	return false
}

// SetVatRequestId gets a reference to the given string and assigns it to the VatRequestId field.
func (o *SalesDataOrderAddressInterface) SetVatRequestId(v string) {
	o.VatRequestId = &v
}

// GetVatRequestSuccess returns the VatRequestSuccess field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetVatRequestSuccess() int32 {
	if o == nil || IsNil(o.VatRequestSuccess) {
		var ret int32
		return ret
	}
	return *o.VatRequestSuccess
}

// GetVatRequestSuccessOk returns a tuple with the VatRequestSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetVatRequestSuccessOk() (*int32, bool) {
	if o == nil || IsNil(o.VatRequestSuccess) {
		return nil, false
	}
	return o.VatRequestSuccess, true
}

// HasVatRequestSuccess returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasVatRequestSuccess() bool {
	if o != nil && !IsNil(o.VatRequestSuccess) {
		return true
	}

	return false
}

// SetVatRequestSuccess gets a reference to the given int32 and assigns it to the VatRequestSuccess field.
func (o *SalesDataOrderAddressInterface) SetVatRequestSuccess(v int32) {
	o.VatRequestSuccess = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *SalesDataOrderAddressInterface) GetExtensionAttributes() map[string]interface{} {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesDataOrderAddressInterface) GetExtensionAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return map[string]interface{}{}, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *SalesDataOrderAddressInterface) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given map[string]interface{} and assigns it to the ExtensionAttributes field.
func (o *SalesDataOrderAddressInterface) SetExtensionAttributes(v map[string]interface{}) {
	o.ExtensionAttributes = v
}

func (o SalesDataOrderAddressInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesDataOrderAddressInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address_type"] = o.AddressType
	toSerialize["city"] = o.City
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["country_id"] = o.CountryId
	if !IsNil(o.CustomerAddressId) {
		toSerialize["customer_address_id"] = o.CustomerAddressId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	toSerialize["firstname"] = o.Firstname
	toSerialize["lastname"] = o.Lastname
	if !IsNil(o.Middlename) {
		toSerialize["middlename"] = o.Middlename
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	toSerialize["postcode"] = o.Postcode
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.RegionCode) {
		toSerialize["region_code"] = o.RegionCode
	}
	if !IsNil(o.RegionId) {
		toSerialize["region_id"] = o.RegionId
	}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !IsNil(o.Suffix) {
		toSerialize["suffix"] = o.Suffix
	}
	toSerialize["telephone"] = o.Telephone
	if !IsNil(o.VatId) {
		toSerialize["vat_id"] = o.VatId
	}
	if !IsNil(o.VatIsValid) {
		toSerialize["vat_is_valid"] = o.VatIsValid
	}
	if !IsNil(o.VatRequestDate) {
		toSerialize["vat_request_date"] = o.VatRequestDate
	}
	if !IsNil(o.VatRequestId) {
		toSerialize["vat_request_id"] = o.VatRequestId
	}
	if !IsNil(o.VatRequestSuccess) {
		toSerialize["vat_request_success"] = o.VatRequestSuccess
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extension_attributes"] = o.ExtensionAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SalesDataOrderAddressInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address_type",
		"city",
		"country_id",
		"firstname",
		"lastname",
		"postcode",
		"telephone",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSalesDataOrderAddressInterface := _SalesDataOrderAddressInterface{}

	err = json.Unmarshal(data, &varSalesDataOrderAddressInterface)

	if err != nil {
		return err
	}

	*o = SalesDataOrderAddressInterface(varSalesDataOrderAddressInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address_type")
		delete(additionalProperties, "city")
		delete(additionalProperties, "company")
		delete(additionalProperties, "country_id")
		delete(additionalProperties, "customer_address_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "email")
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "fax")
		delete(additionalProperties, "firstname")
		delete(additionalProperties, "lastname")
		delete(additionalProperties, "middlename")
		delete(additionalProperties, "parent_id")
		delete(additionalProperties, "postcode")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "region")
		delete(additionalProperties, "region_code")
		delete(additionalProperties, "region_id")
		delete(additionalProperties, "street")
		delete(additionalProperties, "suffix")
		delete(additionalProperties, "telephone")
		delete(additionalProperties, "vat_id")
		delete(additionalProperties, "vat_is_valid")
		delete(additionalProperties, "vat_request_date")
		delete(additionalProperties, "vat_request_id")
		delete(additionalProperties, "vat_request_success")
		delete(additionalProperties, "extension_attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SalesDataOrderAddressInterface) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SalesDataOrderAddressInterface) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSalesDataOrderAddressInterface struct {
	value *SalesDataOrderAddressInterface
	isSet bool
}

func (v NullableSalesDataOrderAddressInterface) Get() *SalesDataOrderAddressInterface {
	return v.value
}

func (v *NullableSalesDataOrderAddressInterface) Set(val *SalesDataOrderAddressInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesDataOrderAddressInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesDataOrderAddressInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesDataOrderAddressInterface(val *SalesDataOrderAddressInterface) *NullableSalesDataOrderAddressInterface {
	return &NullableSalesDataOrderAddressInterface{value: val, isSet: true}
}

func (v NullableSalesDataOrderAddressInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesDataOrderAddressInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
