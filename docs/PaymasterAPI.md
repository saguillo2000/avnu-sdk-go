# \PaymasterAPI

All URIs are relative to *https://starknet.api.avnu.fi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildTypedData1**](PaymasterAPI.md#BuildTypedData1) | **Post** /paymaster/v1/build-typed-data | Build the typed data
[**Execute1**](PaymasterAPI.md#Execute1) | **Post** /paymaster/v1/execute | Execute the calls
[**GetGasTokenPrices**](PaymasterAPI.md#GetGasTokenPrices) | **Get** /paymaster/v1/gas-token-prices | Get the supported gas tokens and their prices in ETH and USD
[**GetRewards1**](PaymasterAPI.md#GetRewards1) | **Get** /paymaster/v1/accounts/{address}/rewards | Retrieve account&#39;s rewards
[**GetSponsor**](PaymasterAPI.md#GetSponsor) | **Get** /paymaster/v1/sponsor-activity | 
[**IsCompatible1**](PaymasterAPI.md#IsCompatible1) | **Get** /paymaster/v1/accounts/{address}/compatible | Check if the account is compatible with the gasless service
[**RegisterReward1**](PaymasterAPI.md#RegisterReward1) | **Post** /paymaster/v1/accounts/{address}/rewards | Add a reward to an account
[**Status**](PaymasterAPI.md#Status) | **Get** /paymaster/v1/status | Get the current gasless service status



## BuildTypedData1

> TypedData BuildTypedData1(ctx).BuildTypedDataRequest(buildTypedDataRequest).ApiKey(apiKey).AskSignature(askSignature).Execute()

Build the typed data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	buildTypedDataRequest := *openapiclient.NewBuildTypedDataRequest("UserAddress_example", []openapiclient.Call{*openapiclient.NewCall("ContractAddress_example", "approve", []string{"Calldata_example"})}) // BuildTypedDataRequest | 
	apiKey := "apiKey_example" // string |  (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.BuildTypedData1(context.Background()).BuildTypedDataRequest(buildTypedDataRequest).ApiKey(apiKey).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.BuildTypedData1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildTypedData1`: TypedData
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.BuildTypedData1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildTypedData1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildTypedDataRequest** | [**BuildTypedDataRequest**](BuildTypedDataRequest.md) |  | 
 **apiKey** | **string** |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**TypedData**](TypedData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Execute1

> ExecuteResponse Execute1(ctx).ExecuteRequest(executeRequest).ApiKey(apiKey).AskSignature(askSignature).Execute()

Execute the calls



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	executeRequest := *openapiclient.NewExecuteRequest("UserAddress_example", "TypedData_example", []string{"Signature_example"}) // ExecuteRequest | 
	apiKey := "apiKey_example" // string |  (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.Execute1(context.Background()).ExecuteRequest(executeRequest).ApiKey(apiKey).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.Execute1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Execute1`: ExecuteResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.Execute1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecute1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeRequest** | [**ExecuteRequest**](ExecuteRequest.md) |  | 
 **apiKey** | **string** |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**ExecuteResponse**](ExecuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGasTokenPrices

> GaslessStatus GetGasTokenPrices(ctx).AskSignature(askSignature).Execute()

Get the supported gas tokens and their prices in ETH and USD

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.GetGasTokenPrices(context.Background()).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.GetGasTokenPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGasTokenPrices`: GaslessStatus
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.GetGasTokenPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGasTokenPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**GaslessStatus**](GaslessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewards1

> []Reward GetRewards1(ctx, address).Sponsor(sponsor).Campaign(campaign).Protocol(protocol).AskSignature(askSignature).Execute()

Retrieve account's rewards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := "address_example" // string | 
	sponsor := "sponsor_example" // string |  (optional)
	campaign := "campaign_example" // string |  (optional)
	protocol := "protocol_example" // string |  (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.GetRewards1(context.Background(), address).Sponsor(sponsor).Campaign(campaign).Protocol(protocol).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.GetRewards1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRewards1`: []Reward
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.GetRewards1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewards1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sponsor** | **string** |  | 
 **campaign** | **string** |  | 
 **protocol** | **string** |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**[]Reward**](Reward.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSponsor

> SponsorActivity GetSponsor(ctx).ApiKey(apiKey).StartDate(startDate).EndDate(endDate).AskSignature(askSignature).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiKey := "apiKey_example" // string | 
	startDate := time.Now() // time.Time | Default value 7 days ago (optional)
	endDate := time.Now() // time.Time | Default value is now (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.GetSponsor(context.Background()).ApiKey(apiKey).StartDate(startDate).EndDate(endDate).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.GetSponsor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSponsor`: SponsorActivity
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.GetSponsor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | 
 **startDate** | **time.Time** | Default value 7 days ago | 
 **endDate** | **time.Time** | Default value is now | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**SponsorActivity**](SponsorActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsCompatible1

> GaslessStatus IsCompatible1(ctx, address).AskSignature(askSignature).Execute()

Check if the account is compatible with the gasless service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7" // string | The account's address
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.IsCompatible1(context.Background(), address).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.IsCompatible1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsCompatible1`: GaslessStatus
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.IsCompatible1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | The account&#39;s address | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsCompatible1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**GaslessStatus**](GaslessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterReward1

> Reward RegisterReward1(ctx, address).ApiKey(apiKey).RegisterReward(registerReward).AskSignature(askSignature).Execute()

Add a reward to an account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := "address_example" // string | 
	apiKey := "apiKey_example" // string | 
	registerReward := *openapiclient.NewRegisterReward("0x039321741034d079C573bAd24dB5F012ed9614554301a2B08bDcb34E01d9C1BF", "Onboarding", int64(10), []openapiclient.WhitelistedCall{*openapiclient.NewWhitelistedCall("*", "approve")}) // RegisterReward | 
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.RegisterReward1(context.Background(), address).ApiKey(apiKey).RegisterReward(registerReward).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.RegisterReward1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterReward1`: Reward
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.RegisterReward1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterReward1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **string** |  | 
 **registerReward** | [**RegisterReward**](RegisterReward.md) |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**Reward**](Reward.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> GaslessStatus Status(ctx).AskSignature(askSignature).Execute()

Get the current gasless service status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymasterAPI.Status(context.Background()).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymasterAPI.Status``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Status`: GaslessStatus
	fmt.Fprintf(os.Stdout, "Response from `PaymasterAPI.Status`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**GaslessStatus**](GaslessStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

