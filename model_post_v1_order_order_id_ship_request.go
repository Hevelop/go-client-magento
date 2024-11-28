/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
)

// checks if the PostV1OrderOrderIdShipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1OrderOrderIdShipRequest{}

// PostV1OrderOrderIdShipRequest struct for PostV1OrderOrderIdShipRequest
type PostV1OrderOrderIdShipRequest struct {
	Items                []SalesDataShipmentItemCreationInterface     `json:"items,omitempty"`
	Notify               *bool                                        `json:"notify,omitempty"`
	AppendComment        *bool                                        `json:"appendComment,omitempty"`
	Comment              *SalesDataShipmentCommentCreationInterface   `json:"comment,omitempty"`
	Tracks               []SalesDataShipmentTrackCreationInterface    `json:"tracks,omitempty"`
	Packages             []SalesDataShipmentPackageCreationInterface  `json:"packages,omitempty"`
	Arguments            *SalesDataShipmentCreationArgumentsInterface `json:"arguments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostV1OrderOrderIdShipRequest PostV1OrderOrderIdShipRequest

// NewPostV1OrderOrderIdShipRequest instantiates a new PostV1OrderOrderIdShipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1OrderOrderIdShipRequest() *PostV1OrderOrderIdShipRequest {
	this := PostV1OrderOrderIdShipRequest{}
	return &this
}

// NewPostV1OrderOrderIdShipRequestWithDefaults instantiates a new PostV1OrderOrderIdShipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1OrderOrderIdShipRequestWithDefaults() *PostV1OrderOrderIdShipRequest {
	this := PostV1OrderOrderIdShipRequest{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdShipRequest) GetItems() []SalesDataShipmentItemCreationInterface {
	if o == nil || IsNil(o.Items) {
		var ret []SalesDataShipmentItemCreationInterface
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdShipRequest) GetItemsOk() ([]SalesDataShipmentItemCreationInterface, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdShipRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SalesDataShipmentItemCreationInterface and assigns it to the Items field.
func (o *PostV1OrderOrderIdShipRequest) SetItems(v []SalesDataShipmentItemCreationInterface) {
	o.Items = v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdShipRequest) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdShipRequest) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdShipRequest) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *PostV1OrderOrderIdShipRequest) SetNotify(v bool) {
	o.Notify = &v
}

// GetAppendComment returns the AppendComment field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdShipRequest) GetAppendComment() bool {
	if o == nil || IsNil(o.AppendComment) {
		var ret bool
		return ret
	}
	return *o.AppendComment
}

// GetAppendCommentOk returns a tuple with the AppendComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdShipRequest) GetAppendCommentOk() (*bool, bool) {
	if o == nil || IsNil(o.AppendComment) {
		return nil, false
	}
	return o.AppendComment, true
}

// HasAppendComment returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdShipRequest) HasAppendComment() bool {
	if o != nil && !IsNil(o.AppendComment) {
		return true
	}

	return false
}

// SetAppendComment gets a reference to the given bool and assigns it to the AppendComment field.
func (o *PostV1OrderOrderIdShipRequest) SetAppendComment(v bool) {
	o.AppendComment = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdShipRequest) GetComment() SalesDataShipmentCommentCreationInterface {
	if o == nil || IsNil(o.Comment) {
		var ret SalesDataShipmentCommentCreationInterface
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdShipRequest) GetCommentOk() (*SalesDataShipmentCommentCreationInterface, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdShipRequest) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given SalesDataShipmentCommentCreationInterface and assigns it to the Comment field.
func (o *PostV1OrderOrderIdShipRequest) SetComment(v SalesDataShipmentCommentCreationInterface) {
	o.Comment = &v
}

// GetTracks returns the Tracks field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdShipRequest) GetTracks() []SalesDataShipmentTrackCreationInterface {
	if o == nil || IsNil(o.Tracks) {
		var ret []SalesDataShipmentTrackCreationInterface
		return ret
	}
	return o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdShipRequest) GetTracksOk() ([]SalesDataShipmentTrackCreationInterface, bool) {
	if o == nil || IsNil(o.Tracks) {
		return nil, false
	}
	return o.Tracks, true
}

// HasTracks returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdShipRequest) HasTracks() bool {
	if o != nil && !IsNil(o.Tracks) {
		return true
	}

	return false
}

// SetTracks gets a reference to the given []SalesDataShipmentTrackCreationInterface and assigns it to the Tracks field.
func (o *PostV1OrderOrderIdShipRequest) SetTracks(v []SalesDataShipmentTrackCreationInterface) {
	o.Tracks = v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdShipRequest) GetPackages() []SalesDataShipmentPackageCreationInterface {
	if o == nil || IsNil(o.Packages) {
		var ret []SalesDataShipmentPackageCreationInterface
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdShipRequest) GetPackagesOk() ([]SalesDataShipmentPackageCreationInterface, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdShipRequest) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []SalesDataShipmentPackageCreationInterface and assigns it to the Packages field.
func (o *PostV1OrderOrderIdShipRequest) SetPackages(v []SalesDataShipmentPackageCreationInterface) {
	o.Packages = v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdShipRequest) GetArguments() SalesDataShipmentCreationArgumentsInterface {
	if o == nil || IsNil(o.Arguments) {
		var ret SalesDataShipmentCreationArgumentsInterface
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdShipRequest) GetArgumentsOk() (*SalesDataShipmentCreationArgumentsInterface, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdShipRequest) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given SalesDataShipmentCreationArgumentsInterface and assigns it to the Arguments field.
func (o *PostV1OrderOrderIdShipRequest) SetArguments(v SalesDataShipmentCreationArgumentsInterface) {
	o.Arguments = &v
}

func (o PostV1OrderOrderIdShipRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1OrderOrderIdShipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.AppendComment) {
		toSerialize["appendComment"] = o.AppendComment
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Tracks) {
		toSerialize["tracks"] = o.Tracks
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1OrderOrderIdShipRequest) UnmarshalJSON(data []byte) (err error) {
	varPostV1OrderOrderIdShipRequest := _PostV1OrderOrderIdShipRequest{}

	err = json.Unmarshal(data, &varPostV1OrderOrderIdShipRequest)

	if err != nil {
		return err
	}

	*o = PostV1OrderOrderIdShipRequest(varPostV1OrderOrderIdShipRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "notify")
		delete(additionalProperties, "appendComment")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "tracks")
		delete(additionalProperties, "packages")
		delete(additionalProperties, "arguments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1OrderOrderIdShipRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1OrderOrderIdShipRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1OrderOrderIdShipRequest struct {
	value *PostV1OrderOrderIdShipRequest
	isSet bool
}

func (v NullablePostV1OrderOrderIdShipRequest) Get() *PostV1OrderOrderIdShipRequest {
	return v.value
}

func (v *NullablePostV1OrderOrderIdShipRequest) Set(val *PostV1OrderOrderIdShipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1OrderOrderIdShipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1OrderOrderIdShipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1OrderOrderIdShipRequest(val *PostV1OrderOrderIdShipRequest) *NullablePostV1OrderOrderIdShipRequest {
	return &NullablePostV1OrderOrderIdShipRequest{value: val, isSet: true}
}

func (v NullablePostV1OrderOrderIdShipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1OrderOrderIdShipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}