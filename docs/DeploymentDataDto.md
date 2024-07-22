# DeploymentDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassHash** | **string** |  | 
**Salt** | **string** |  | 
**Unique** | **string** |  | 
**Calldata** | **[]string** |  | 
**Sigdata** | **[]string** |  | 

## Methods

### NewDeploymentDataDto

`func NewDeploymentDataDto(classHash string, salt string, unique string, calldata []string, sigdata []string, ) *DeploymentDataDto`

NewDeploymentDataDto instantiates a new DeploymentDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDataDtoWithDefaults

`func NewDeploymentDataDtoWithDefaults() *DeploymentDataDto`

NewDeploymentDataDtoWithDefaults instantiates a new DeploymentDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassHash

`func (o *DeploymentDataDto) GetClassHash() string`

GetClassHash returns the ClassHash field if non-nil, zero value otherwise.

### GetClassHashOk

`func (o *DeploymentDataDto) GetClassHashOk() (*string, bool)`

GetClassHashOk returns a tuple with the ClassHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassHash

`func (o *DeploymentDataDto) SetClassHash(v string)`

SetClassHash sets ClassHash field to given value.


### GetSalt

`func (o *DeploymentDataDto) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *DeploymentDataDto) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *DeploymentDataDto) SetSalt(v string)`

SetSalt sets Salt field to given value.


### GetUnique

`func (o *DeploymentDataDto) GetUnique() string`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *DeploymentDataDto) GetUniqueOk() (*string, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *DeploymentDataDto) SetUnique(v string)`

SetUnique sets Unique field to given value.


### GetCalldata

`func (o *DeploymentDataDto) GetCalldata() []string`

GetCalldata returns the Calldata field if non-nil, zero value otherwise.

### GetCalldataOk

`func (o *DeploymentDataDto) GetCalldataOk() (*[]string, bool)`

GetCalldataOk returns a tuple with the Calldata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalldata

`func (o *DeploymentDataDto) SetCalldata(v []string)`

SetCalldata sets Calldata field to given value.


### GetSigdata

`func (o *DeploymentDataDto) GetSigdata() []string`

GetSigdata returns the Sigdata field if non-nil, zero value otherwise.

### GetSigdataOk

`func (o *DeploymentDataDto) GetSigdataOk() (*[]string, bool)`

GetSigdataOk returns a tuple with the Sigdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigdata

`func (o *DeploymentDataDto) SetSigdata(v []string)`

SetSigdata sets Sigdata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


