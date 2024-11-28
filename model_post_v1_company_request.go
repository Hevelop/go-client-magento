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

// checks if the PostV1CompanyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1CompanyRequest{}

// PostV1CompanyRequest struct for PostV1CompanyRequest
type PostV1CompanyRequest struct {
	Company              CompanyDataCompanyInterface `json:"company"`
	AdditionalProperties map[string]interface{}
}

type _PostV1CompanyRequest PostV1CompanyRequest

// NewPostV1CompanyRequest instantiates a new PostV1CompanyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1CompanyRequest(company CompanyDataCompanyInterface) *PostV1CompanyRequest {
	this := PostV1CompanyRequest{}
	this.Company = company
	return &this
}

// NewPostV1CompanyRequestWithDefaults instantiates a new PostV1CompanyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1CompanyRequestWithDefaults() *PostV1CompanyRequest {
	this := PostV1CompanyRequest{}
	return &this
}

// GetCompany returns the Company field value
func (o *PostV1CompanyRequest) GetCompany() CompanyDataCompanyInterface {
	if o == nil {
		var ret CompanyDataCompanyInterface
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *PostV1CompanyRequest) GetCompanyOk() (*CompanyDataCompanyInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *PostV1CompanyRequest) SetCompany(v CompanyDataCompanyInterface) {
	o.Company = v
}

func (o PostV1CompanyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1CompanyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company"] = o.Company

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1CompanyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company",
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

	varPostV1CompanyRequest := _PostV1CompanyRequest{}

	err = json.Unmarshal(data, &varPostV1CompanyRequest)

	if err != nil {
		return err
	}

	*o = PostV1CompanyRequest(varPostV1CompanyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1CompanyRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1CompanyRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1CompanyRequest struct {
	value *PostV1CompanyRequest
	isSet bool
}

func (v NullablePostV1CompanyRequest) Get() *PostV1CompanyRequest {
	return v.value
}

func (v *NullablePostV1CompanyRequest) Set(val *PostV1CompanyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1CompanyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1CompanyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1CompanyRequest(val *PostV1CompanyRequest) *NullablePostV1CompanyRequest {
	return &NullablePostV1CompanyRequest{value: val, isSet: true}
}

func (v NullablePostV1CompanyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1CompanyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
