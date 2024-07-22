# \SwapAPI

All URIs are relative to *https://starknet.api.avnu.fi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Build**](SwapAPI.md#Build) | **Post** /swap/v2/build | Build swap calldata
[**Build1**](SwapAPI.md#Build1) | **Post** /swap/v1/build | Build swap calldata
[**BuildTypedData2**](SwapAPI.md#BuildTypedData2) | **Post** /swap/v2/build-typed-data | 
[**Execute2**](SwapAPI.md#Execute2) | **Post** /swap/v2/execute | Execute swap
[**GetPrices**](SwapAPI.md#GetPrices) | **Get** /swap/v2/prices | Find the prices of DEX applications
[**GetPrices1**](SwapAPI.md#GetPrices1) | **Get** /swap/v1/prices | Find the prices of DEX applications
[**GetQuotes**](SwapAPI.md#GetQuotes) | **Get** /swap/v2/quotes | Find the best quotes
[**GetQuotes1**](SwapAPI.md#GetQuotes1) | **Get** /swap/v1/quotes | Find the best quotes
[**GetSources**](SwapAPI.md#GetSources) | **Get** /swap/v2/sources | Fetch the list of supported sources
[**GetSources1**](SwapAPI.md#GetSources1) | **Get** /swap/v1/sources | Fetch the list of supported sources
[**GetTokens**](SwapAPI.md#GetTokens) | **Get** /swap/v2/tokens | Fetch supported tokens
[**GetTokens1**](SwapAPI.md#GetTokens1) | **Get** /swap/v1/tokens | Fetch supported tokens



## Build

> BuildSwapResponse Build(ctx).BuildSwapRequestDtoV2(buildSwapRequestDtoV2).AskSignature(askSignature).Execute()

Build swap calldata



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
	buildSwapRequestDtoV2 := *openapiclient.NewBuildSwapRequestDtoV2("QuoteId_example", float32(0.05), false) // BuildSwapRequestDtoV2 | 
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.Build(context.Background()).BuildSwapRequestDtoV2(buildSwapRequestDtoV2).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.Build``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Build`: BuildSwapResponse
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.Build`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildSwapRequestDtoV2** | [**BuildSwapRequestDtoV2**](BuildSwapRequestDtoV2.md) |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**BuildSwapResponse**](BuildSwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Build1

> BuildSwapResponse Build1(ctx).BuildSwapRequest(buildSwapRequest).AskSignature(askSignature).Execute()

Build swap calldata



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
	buildSwapRequest := *openapiclient.NewBuildSwapRequest("QuoteId_example", float32(0.05)) // BuildSwapRequest | 
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.Build1(context.Background()).BuildSwapRequest(buildSwapRequest).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.Build1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Build1`: BuildSwapResponse
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.Build1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuild1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildSwapRequest** | [**BuildSwapRequest**](BuildSwapRequest.md) |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**BuildSwapResponse**](BuildSwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildTypedData2

> TypedData BuildTypedData2(ctx).BuildSwapTypedDataRequest(buildSwapTypedDataRequest).AskSignature(askSignature).Execute()



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
	buildSwapTypedDataRequest := *openapiclient.NewBuildSwapTypedDataRequest("QuoteId_example", float32(0.05), false, "0x053c91253bc9682c04929ca02ed00b3e423f6710d2ee7e0d5ebb06f3ecf368a8") // BuildSwapTypedDataRequest | 
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.BuildTypedData2(context.Background()).BuildSwapTypedDataRequest(buildSwapTypedDataRequest).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.BuildTypedData2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildTypedData2`: TypedData
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.BuildTypedData2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildTypedData2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buildSwapTypedDataRequest** | [**BuildSwapTypedDataRequest**](BuildSwapTypedDataRequest.md) |  | 
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


## Execute2

> ExecuteSwapResponse Execute2(ctx).ExecuteSwapRequest(executeSwapRequest).AskSignature(askSignature).Execute()

Execute swap



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
	executeSwapRequest := *openapiclient.NewExecuteSwapRequest("QuoteId_example", []string{"Signature_example"}) // ExecuteSwapRequest | 
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.Execute2(context.Background()).ExecuteSwapRequest(executeSwapRequest).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.Execute2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Execute2`: ExecuteSwapResponse
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.Execute2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecute2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeSwapRequest** | [**ExecuteSwapRequest**](ExecuteSwapRequest.md) |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**ExecuteSwapResponse**](ExecuteSwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrices

> Quote GetPrices(ctx).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).AskSignature(askSignature).Execute()

Find the prices of DEX applications



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
	sellTokenAddress := "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7" // string | The token address user wants to sell
	buyTokenAddress := "0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d" // string | The token address user wants to buy
	sellAmount := "0x038d7ea4c68000" // string | The Amount of token user wants to sell
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetPrices(context.Background()).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrices`: Quote
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sellTokenAddress** | **string** | The token address user wants to sell | 
 **buyTokenAddress** | **string** | The token address user wants to buy | 
 **sellAmount** | **string** | The Amount of token user wants to sell | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**Quote**](Quote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrices1

> Quote GetPrices1(ctx).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).AskSignature(askSignature).Execute()

Find the prices of DEX applications



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
	sellTokenAddress := "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7" // string | The token address user wants to sell
	buyTokenAddress := "0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d" // string | The token address user wants to buy
	sellAmount := "0x038d7ea4c68000" // string | The Amount of token user wants to sell
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetPrices1(context.Background()).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetPrices1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrices1`: Quote
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetPrices1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrices1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sellTokenAddress** | **string** | The token address user wants to sell | 
 **buyTokenAddress** | **string** | The token address user wants to buy | 
 **sellAmount** | **string** | The Amount of token user wants to sell | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**Quote**](Quote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes

> Quote GetQuotes(ctx).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).BuyAmount(buyAmount).TakerAddress(takerAddress).ExcludeSources(excludeSources).Size(size).IntegratorFees(integratorFees).IntegratorFeeRecipient(integratorFeeRecipient).IntegratorName(integratorName).Origin(origin).UserAgent(userAgent).XFORWARDEDFOR(xFORWARDEDFOR).AskSignature(askSignature).Execute()

Find the best quotes



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
	sellTokenAddress := "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7" // string | The token address user wants to sell
	buyTokenAddress := "0xda114221cb83fa859dbdb4c44beeaa0bb37c7537ad5ae66fe5e0efd20e6eb3" // string | The token address user wants to buy
	sellAmount := "0x38d7ea4c68000" // string | The amount of token user wants to sell. It must be defined if buyAmount is not defined. (optional)
	buyAmount := "0x38d7ea4c68000" // string | The exact amount of token user wants to buy. It must be defined if sellAmount is not defined. (optional)
	takerAddress := "0x052D8E9778D026588a51595E30B0F45609B4F771EecF0E335CdeFeD1d84a9D89" // string | The address which will fill the quote (optional)
	excludeSources := []string{"Inner_example"} // []string | The sources that the user wants to exclude (optional)
	size := int64(1) // int64 | The maximum number of returned quotes (optional)
	integratorFees := "0x3" // string | Fee amount in bps, 30 is 0.3% (optional)
	integratorFeeRecipient := "0x01238E9778D026588a51595E30B0F45609B4F771EecF0E335CdeFeD1d84a9D89" // string | Required when `integratorFees` is defined. You need to provide the address of your fee collector. (optional)
	integratorName := "Swagger UI" // string | The name of your application (optional)
	origin := "origin_example" // string |  (optional)
	userAgent := "userAgent_example" // string |  (optional)
	xFORWARDEDFOR := "xFORWARDEDFOR_example" // string |  (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetQuotes(context.Background()).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).BuyAmount(buyAmount).TakerAddress(takerAddress).ExcludeSources(excludeSources).Size(size).IntegratorFees(integratorFees).IntegratorFeeRecipient(integratorFeeRecipient).IntegratorName(integratorName).Origin(origin).UserAgent(userAgent).XFORWARDEDFOR(xFORWARDEDFOR).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotes`: Quote
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sellTokenAddress** | **string** | The token address user wants to sell | 
 **buyTokenAddress** | **string** | The token address user wants to buy | 
 **sellAmount** | **string** | The amount of token user wants to sell. It must be defined if buyAmount is not defined. | 
 **buyAmount** | **string** | The exact amount of token user wants to buy. It must be defined if sellAmount is not defined. | 
 **takerAddress** | **string** | The address which will fill the quote | 
 **excludeSources** | **[]string** | The sources that the user wants to exclude | 
 **size** | **int64** | The maximum number of returned quotes | 
 **integratorFees** | **string** | Fee amount in bps, 30 is 0.3% | 
 **integratorFeeRecipient** | **string** | Required when &#x60;integratorFees&#x60; is defined. You need to provide the address of your fee collector. | 
 **integratorName** | **string** | The name of your application | 
 **origin** | **string** |  | 
 **userAgent** | **string** |  | 
 **xFORWARDEDFOR** | **string** |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**Quote**](Quote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes1

> Quote GetQuotes1(ctx).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).BuyAmount(buyAmount).TakerAddress(takerAddress).ExcludeSources(excludeSources).Size(size).IntegratorFees(integratorFees).IntegratorFeeRecipient(integratorFeeRecipient).IntegratorName(integratorName).Origin(origin).UserAgent(userAgent).XFORWARDEDFOR(xFORWARDEDFOR).AskSignature(askSignature).Execute()

Find the best quotes



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
	sellTokenAddress := "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7" // string | The token address user wants to sell
	buyTokenAddress := "0xda114221cb83fa859dbdb4c44beeaa0bb37c7537ad5ae66fe5e0efd20e6eb3" // string | The token address user wants to buy
	sellAmount := "0x38d7ea4c68000" // string | The amount of token user wants to sell. It must be defined if buyAmount is not defined. (optional)
	buyAmount := "0x38d7ea4c68000" // string | The exact amount of token user wants to buy. It must be defined if sellAmount is not defined. (optional)
	takerAddress := "0x052D8E9778D026588a51595E30B0F45609B4F771EecF0E335CdeFeD1d84a9D89" // string | The address which will fill the quote (optional)
	excludeSources := []string{"Inner_example"} // []string | The sources that the user wants to exclude (optional)
	size := int64(1) // int64 | The maximum number of returned quotes (optional)
	integratorFees := "0x3" // string | Fee amount in bps, 30 is 0.3% (optional)
	integratorFeeRecipient := "0x01238E9778D026588a51595E30B0F45609B4F771EecF0E335CdeFeD1d84a9D89" // string | Required when `integratorFees` is defined. You need to provide the address of your fee collector. (optional)
	integratorName := "Swagger UI" // string | The name of your application (optional)
	origin := "origin_example" // string |  (optional)
	userAgent := "userAgent_example" // string |  (optional)
	xFORWARDEDFOR := "xFORWARDEDFOR_example" // string |  (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetQuotes1(context.Background()).SellTokenAddress(sellTokenAddress).BuyTokenAddress(buyTokenAddress).SellAmount(sellAmount).BuyAmount(buyAmount).TakerAddress(takerAddress).ExcludeSources(excludeSources).Size(size).IntegratorFees(integratorFees).IntegratorFeeRecipient(integratorFeeRecipient).IntegratorName(integratorName).Origin(origin).UserAgent(userAgent).XFORWARDEDFOR(xFORWARDEDFOR).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetQuotes1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotes1`: Quote
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetQuotes1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotes1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sellTokenAddress** | **string** | The token address user wants to sell | 
 **buyTokenAddress** | **string** | The token address user wants to buy | 
 **sellAmount** | **string** | The amount of token user wants to sell. It must be defined if buyAmount is not defined. | 
 **buyAmount** | **string** | The exact amount of token user wants to buy. It must be defined if sellAmount is not defined. | 
 **takerAddress** | **string** | The address which will fill the quote | 
 **excludeSources** | **[]string** | The sources that the user wants to exclude | 
 **size** | **int64** | The maximum number of returned quotes | 
 **integratorFees** | **string** | Fee amount in bps, 30 is 0.3% | 
 **integratorFeeRecipient** | **string** | Required when &#x60;integratorFees&#x60; is defined. You need to provide the address of your fee collector. | 
 **integratorName** | **string** | The name of your application | 
 **origin** | **string** |  | 
 **userAgent** | **string** |  | 
 **xFORWARDEDFOR** | **string** |  | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**Quote**](Quote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSources

> ExecuteSwapResponse GetSources(ctx).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()

Fetch the list of supported sources



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
	page := int32(56) // int32 | Zero-based page index (0..N) (optional) (default to 0)
	size := int32(56) // int32 | The size of the page to be returned (optional) (default to 20)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetSources(context.Background()).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSources`: ExecuteSwapResponse
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | [default to 0]
 **size** | **int32** | The size of the page to be returned | [default to 20]
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**ExecuteSwapResponse**](ExecuteSwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSources1

> ExecuteSwapResponse GetSources1(ctx).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()

Fetch the list of supported sources



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
	page := int32(56) // int32 | Zero-based page index (0..N) (optional) (default to 0)
	size := int32(56) // int32 | The size of the page to be returned (optional) (default to 20)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetSources1(context.Background()).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetSources1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSources1`: ExecuteSwapResponse
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetSources1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSources1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | [default to 0]
 **size** | **int32** | The size of the page to be returned | [default to 20]
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**ExecuteSwapResponse**](ExecuteSwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> ExecuteSwapResponse GetTokens(ctx).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()

Fetch supported tokens



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
	page := int32(56) // int32 | Zero-based page index (0..N) (optional) (default to 0)
	size := int32(56) // int32 | The size of the page to be returned (optional) (default to 20)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetTokens(context.Background()).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: ExecuteSwapResponse
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | [default to 0]
 **size** | **int32** | The size of the page to be returned | [default to 20]
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**ExecuteSwapResponse**](ExecuteSwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens1

> ExecuteSwapResponse GetTokens1(ctx).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()

Fetch supported tokens



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
	page := int32(56) // int32 | Zero-based page index (0..N) (optional) (default to 0)
	size := int32(56) // int32 | The size of the page to be returned (optional) (default to 20)
	sort := []string{"Inner_example"} // []string | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. (optional)
	askSignature := "askSignature_example" // string | When the given value is provided with the value 'true', the response header 'signature' will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.GetTokens1(context.Background()).Page(page).Size(size).Sort(sort).AskSignature(askSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.GetTokens1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens1`: ExecuteSwapResponse
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.GetTokens1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokens1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Zero-based page index (0..N) | [default to 0]
 **size** | **int32** | The size of the page to be returned | [default to 20]
 **sort** | **[]string** | Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 
 **askSignature** | **string** | When the given value is provided with the value &#39;true&#39;, the response header &#39;signature&#39; will be returned. | 

### Return type

[**ExecuteSwapResponse**](ExecuteSwapResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

