# \DefaultAPI

All URIs are relative to *https://api.nango.dev*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionTriggerPost**](DefaultAPI.md#ActionTriggerPost) | **Post** /action/trigger | 
[**ConfigGet**](DefaultAPI.md#ConfigGet) | **Get** /config | 
[**ConfigPost**](DefaultAPI.md#ConfigPost) | **Post** /config | 
[**ConfigProviderConfigKeyDelete**](DefaultAPI.md#ConfigProviderConfigKeyDelete) | **Delete** /config/{providerConfigKey} | 
[**ConfigProviderConfigKeyGet**](DefaultAPI.md#ConfigProviderConfigKeyGet) | **Get** /config/{providerConfigKey} | 
[**ConfigPut**](DefaultAPI.md#ConfigPut) | **Put** /config | 
[**ConnectionConnectionIdDelete**](DefaultAPI.md#ConnectionConnectionIdDelete) | **Delete** /connection/{connectionId} | 
[**ConnectionConnectionIdGet**](DefaultAPI.md#ConnectionConnectionIdGet) | **Get** /connection/{connectionId} | 
[**ConnectionConnectionIdMetadataPatch**](DefaultAPI.md#ConnectionConnectionIdMetadataPatch) | **Patch** /connection/{connectionId}/metadata | 
[**ConnectionConnectionIdMetadataPost**](DefaultAPI.md#ConnectionConnectionIdMetadataPost) | **Post** /connection/{connectionId}/metadata | 
[**ConnectionGet**](DefaultAPI.md#ConnectionGet) | **Get** /connection | 
[**ConnectionMetadataPatch**](DefaultAPI.md#ConnectionMetadataPatch) | **Patch** /connection/metadata | 
[**ConnectionMetadataPost**](DefaultAPI.md#ConnectionMetadataPost) | **Post** /connection/metadata | 
[**ConnectionPost**](DefaultAPI.md#ConnectionPost) | **Post** /connection | 
[**EnvironmentVariablesGet**](DefaultAPI.md#EnvironmentVariablesGet) | **Get** /environment-variables | 
[**ProxyAnyPathDelete**](DefaultAPI.md#ProxyAnyPathDelete) | **Delete** /proxy/{anyPath} | 
[**ProxyAnyPathGet**](DefaultAPI.md#ProxyAnyPathGet) | **Get** /proxy/{anyPath} | 
[**ProxyAnyPathPatch**](DefaultAPI.md#ProxyAnyPathPatch) | **Patch** /proxy/{anyPath} | 
[**ProxyAnyPathPost**](DefaultAPI.md#ProxyAnyPathPost) | **Post** /proxy/{anyPath} | 
[**ProxyAnyPathPut**](DefaultAPI.md#ProxyAnyPathPut) | **Put** /proxy/{anyPath} | 
[**RecordsGet**](DefaultAPI.md#RecordsGet) | **Get** /records | 
[**ScriptsConfigGet**](DefaultAPI.md#ScriptsConfigGet) | **Get** /scripts/config | 
[**SyncPausePost**](DefaultAPI.md#SyncPausePost) | **Post** /sync/pause | 
[**SyncRecordsGet**](DefaultAPI.md#SyncRecordsGet) | **Get** /sync/records | 
[**SyncStartPost**](DefaultAPI.md#SyncStartPost) | **Post** /sync/start | 
[**SyncStatusGet**](DefaultAPI.md#SyncStatusGet) | **Get** /sync/status | 
[**SyncTriggerPost**](DefaultAPI.md#SyncTriggerPost) | **Post** /sync/trigger | 
[**SyncUpdateConnectionFrequencyPut**](DefaultAPI.md#SyncUpdateConnectionFrequencyPut) | **Put** /sync/update-connection-frequency | 



## ActionTriggerPost

> ActionTriggerPost200Response ActionTriggerPost(ctx).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).ActionTriggerPostRequest(actionTriggerPostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	actionTriggerPostRequest := *openapiclient.NewActionTriggerPostRequest("ActionName_example") // ActionTriggerPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ActionTriggerPost(context.Background()).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).ActionTriggerPostRequest(actionTriggerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ActionTriggerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionTriggerPost`: ActionTriggerPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ActionTriggerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionTriggerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **actionTriggerPostRequest** | [**ActionTriggerPostRequest**](ActionTriggerPostRequest.md) |  | 

### Return type

[**ActionTriggerPost200Response**](ActionTriggerPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigGet

> ConfigGet200Response ConfigGet(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigGet`: ConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConfigGetRequest struct via the builder pattern


### Return type

[**ConfigGet200Response**](ConfigGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigPost

> ConfigPut200Response ConfigPost(ctx).ConfigPostRequest(configPostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	configPostRequest := *openapiclient.NewConfigPostRequest("ProviderConfigKey_example", "Provider_example") // ConfigPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConfigPost(context.Background()).ConfigPostRequest(configPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigPost`: ConfigPut200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configPostRequest** | [**ConfigPostRequest**](ConfigPostRequest.md) |  | 

### Return type

[**ConfigPut200Response**](ConfigPut200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigProviderConfigKeyDelete

> ConfigProviderConfigKeyDelete(ctx, providerConfigKey).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	providerConfigKey := "providerConfigKey_example" // string | The integration ID that you created in Nango.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ConfigProviderConfigKeyDelete(context.Background(), providerConfigKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConfigProviderConfigKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerConfigKey** | **string** | The integration ID that you created in Nango. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigProviderConfigKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigProviderConfigKeyGet

> ConfigProviderConfigKeyGet200Response ConfigProviderConfigKeyGet(ctx, providerConfigKey).IncludeCreds(includeCreds).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	providerConfigKey := "providerConfigKey_example" // string | The integration ID that you created in Nango.
	includeCreds := true // bool | If true, the response will contain the client ID, secret, scopes, auth_mode and app link - if applicable. include_creds is false by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConfigProviderConfigKeyGet(context.Background(), providerConfigKey).IncludeCreds(includeCreds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConfigProviderConfigKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigProviderConfigKeyGet`: ConfigProviderConfigKeyGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConfigProviderConfigKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerConfigKey** | **string** | The integration ID that you created in Nango. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigProviderConfigKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCreds** | **bool** | If true, the response will contain the client ID, secret, scopes, auth_mode and app link - if applicable. include_creds is false by default. | 

### Return type

[**ConfigProviderConfigKeyGet200Response**](ConfigProviderConfigKeyGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigPut

> ConfigPut200Response ConfigPut(ctx).ConfigPutRequest(configPutRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	configPutRequest := *openapiclient.NewConfigPutRequest("ProviderConfigKey_example", "Provider_example", "OauthClientId_example", "OauthClientSecret_example") // ConfigPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConfigPut(context.Background()).ConfigPutRequest(configPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigPut`: ConfigPut200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConfigPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configPutRequest** | [**ConfigPutRequest**](ConfigPutRequest.md) |  | 

### Return type

[**ConfigPut200Response**](ConfigPut200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionConnectionIdDelete

> ConnectionConnectionIdDelete(ctx, connectionId).ProviderConfigKey(providerConfigKey).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ConnectionConnectionIdDelete(context.Background(), connectionId).ProviderConfigKey(providerConfigKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionConnectionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection ID used to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionConnectionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionConnectionIdGet

> ConnectionConnectionIdGet(ctx, connectionId).ProviderConfigKey(providerConfigKey).ForceRefresh(forceRefresh).RefreshToken(refreshToken).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	forceRefresh := true // bool | If true, Nango will attempt to refresh the access access token regardless of its expiration status (false by default). (optional)
	refreshToken := true // bool | If true, return the refresh token as part of the response (false by default). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ConnectionConnectionIdGet(context.Background(), connectionId).ProviderConfigKey(providerConfigKey).ForceRefresh(forceRefresh).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionConnectionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection ID used to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionConnectionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **forceRefresh** | **bool** | If true, Nango will attempt to refresh the access access token regardless of its expiration status (false by default). | 
 **refreshToken** | **bool** | If true, return the refresh token as part of the response (false by default). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionConnectionIdMetadataPatch

> map[string]interface{} ConnectionConnectionIdMetadataPatch(ctx, connectionId).ProviderConfigKey(providerConfigKey).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConnectionConnectionIdMetadataPatch(context.Background(), connectionId).ProviderConfigKey(providerConfigKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionConnectionIdMetadataPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionConnectionIdMetadataPatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConnectionConnectionIdMetadataPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection ID used to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionConnectionIdMetadataPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionConnectionIdMetadataPost

> ConnectionConnectionIdMetadataPost(ctx, connectionId).ProviderConfigKey(providerConfigKey).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ConnectionConnectionIdMetadataPost(context.Background(), connectionId).ProviderConfigKey(providerConfigKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionConnectionIdMetadataPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | The connection ID used to create the connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionConnectionIdMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionGet

> ConnectionGet200Response ConnectionGet(ctx).ConnectionId(connectionId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionId := "connectionId_example" // string | Filter the list of connections based on this connection ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConnectionGet(context.Background()).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionGet`: ConnectionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConnectionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionId** | **string** | Filter the list of connections based on this connection ID. | 

### Return type

[**ConnectionGet200Response**](ConnectionGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionMetadataPatch

> ConnectionMetadataPatch200Response ConnectionMetadataPatch(ctx).ConnectionMetadataPatchRequest(connectionMetadataPatchRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionMetadataPatchRequest := *openapiclient.NewConnectionMetadataPatchRequest(openapiclient._connection_metadata_patch_request_connection_id{ArrayOfString: new([]string)}, "ProviderConfigKey_example", map[string]interface{}(123)) // ConnectionMetadataPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConnectionMetadataPatch(context.Background()).ConnectionMetadataPatchRequest(connectionMetadataPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionMetadataPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionMetadataPatch`: ConnectionMetadataPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConnectionMetadataPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionMetadataPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionMetadataPatchRequest** | [**ConnectionMetadataPatchRequest**](ConnectionMetadataPatchRequest.md) |  | 

### Return type

[**ConnectionMetadataPatch200Response**](ConnectionMetadataPatch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionMetadataPost

> ConnectionMetadataPost201Response ConnectionMetadataPost(ctx).ConnectionMetadataPostRequest(connectionMetadataPostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionMetadataPostRequest := *openapiclient.NewConnectionMetadataPostRequest(openapiclient._connection_metadata_post_request_connection_id{ArrayOfString: new([]string)}, "ProviderConfigKey_example", map[string]interface{}(123)) // ConnectionMetadataPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConnectionMetadataPost(context.Background()).ConnectionMetadataPostRequest(connectionMetadataPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionMetadataPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectionMetadataPost`: ConnectionMetadataPost201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConnectionMetadataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionMetadataPostRequest** | [**ConnectionMetadataPostRequest**](ConnectionMetadataPostRequest.md) |  | 

### Return type

[**ConnectionMetadataPost201Response**](ConnectionMetadataPost201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionPost

> ConnectionPost(ctx).ConnectionPostRequest(connectionPostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	connectionPostRequest := *openapiclient.NewConnectionPostRequest("ConnectionId_example", "ProviderConfigKey_example") // ConnectionPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ConnectionPost(context.Background()).ConnectionPostRequest(connectionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConnectionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionPostRequest** | [**ConnectionPostRequest**](ConnectionPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentVariablesGet

> []EnvironmentVariablesGet200ResponseInner EnvironmentVariablesGet(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.EnvironmentVariablesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EnvironmentVariablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentVariablesGet`: []EnvironmentVariablesGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EnvironmentVariablesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentVariablesGetRequest struct via the builder pattern


### Return type

[**[]EnvironmentVariablesGet200ResponseInner**](EnvironmentVariablesGet200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyAnyPathDelete

> ProxyAnyPathDelete(ctx, anyPath).ANYQUERYPARAMS(aNYQUERYPARAMS).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	anyPath := "anyPath_example" // string | 
	aNYQUERYPARAMS := "aNYQUERYPARAMS_example" // string | 
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	retries := "retries_example" // string | The number of retries in case of failure (with exponential back-off). Optional, default 0. (optional)
	baseUrlOverride := "baseUrlOverride_example" // string | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional (optional)
	decompress := "decompress_example" // string | Override the decompress option when making requests. Optional, defaults to false (optional)
	nangoProxyANYHEADER := "nangoProxyANYHEADER_example" // string | Any other headers you send are passed on to the external API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ProxyAnyPathDelete(context.Background(), anyPath).ANYQUERYPARAMS(aNYQUERYPARAMS).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProxyAnyPathDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**anyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyAnyPathDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aNYQUERYPARAMS** | **string** |  | 
 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **retries** | **string** | The number of retries in case of failure (with exponential back-off). Optional, default 0. | 
 **baseUrlOverride** | **string** | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional | 
 **decompress** | **string** | Override the decompress option when making requests. Optional, defaults to false | 
 **nangoProxyANYHEADER** | **string** | Any other headers you send are passed on to the external API | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyAnyPathGet

> ProxyAnyPathGet(ctx, anyPath).ANYQUERYPARAMS(aNYQUERYPARAMS).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).RetryOn(retryOn).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	anyPath := "anyPath_example" // string | 
	aNYQUERYPARAMS := "aNYQUERYPARAMS_example" // string | 
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	retries := "retries_example" // string | The number of retries in case of failure (with exponential back-off). Optional, default 0. (optional)
	retryOn := "retryOn_example" // string | Comma separated status codes to explicitly retry on in addition to the default 5xx and 429. (optional)
	baseUrlOverride := "baseUrlOverride_example" // string | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional (optional)
	decompress := "decompress_example" // string | Override the decompress option when making requests. Optional, defaults to false (optional)
	nangoProxyANYHEADER := "nangoProxyANYHEADER_example" // string | Any other headers you send are passed on to the external API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ProxyAnyPathGet(context.Background(), anyPath).ANYQUERYPARAMS(aNYQUERYPARAMS).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).RetryOn(retryOn).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProxyAnyPathGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**anyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyAnyPathGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aNYQUERYPARAMS** | **string** |  | 
 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **retries** | **string** | The number of retries in case of failure (with exponential back-off). Optional, default 0. | 
 **retryOn** | **string** | Comma separated status codes to explicitly retry on in addition to the default 5xx and 429. | 
 **baseUrlOverride** | **string** | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional | 
 **decompress** | **string** | Override the decompress option when making requests. Optional, defaults to false | 
 **nangoProxyANYHEADER** | **string** | Any other headers you send are passed on to the external API | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyAnyPathPatch

> ProxyAnyPathPatch(ctx, anyPath).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).ProxyAnyPathPutRequest(proxyAnyPathPutRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	anyPath := "anyPath_example" // string | 
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	retries := "retries_example" // string | The number of retries in case of failure (with exponential back-off). Optional, default 0. (optional)
	baseUrlOverride := "baseUrlOverride_example" // string | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional (optional)
	decompress := "decompress_example" // string | Override the decompress option when making requests. Optional, defaults to false (optional)
	nangoProxyANYHEADER := "nangoProxyANYHEADER_example" // string | Any other headers you send are passed on to the external API (optional)
	proxyAnyPathPutRequest := *openapiclient.NewProxyAnyPathPutRequest() // ProxyAnyPathPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ProxyAnyPathPatch(context.Background(), anyPath).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).ProxyAnyPathPutRequest(proxyAnyPathPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProxyAnyPathPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**anyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyAnyPathPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **retries** | **string** | The number of retries in case of failure (with exponential back-off). Optional, default 0. | 
 **baseUrlOverride** | **string** | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional | 
 **decompress** | **string** | Override the decompress option when making requests. Optional, defaults to false | 
 **nangoProxyANYHEADER** | **string** | Any other headers you send are passed on to the external API | 
 **proxyAnyPathPutRequest** | [**ProxyAnyPathPutRequest**](ProxyAnyPathPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyAnyPathPost

> ProxyAnyPathPost(ctx, anyPath).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).ProxyAnyPathPutRequest(proxyAnyPathPutRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	anyPath := "anyPath_example" // string | 
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	retries := "retries_example" // string | The number of retries in case of failure (with exponential back-off). Optional, default 0. (optional)
	baseUrlOverride := "baseUrlOverride_example" // string | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional (optional)
	decompress := "decompress_example" // string | Override the decompress option when making requests. Optional, defaults to false (optional)
	nangoProxyANYHEADER := "nangoProxyANYHEADER_example" // string | Any other headers you send are passed on to the external API (optional)
	proxyAnyPathPutRequest := *openapiclient.NewProxyAnyPathPutRequest() // ProxyAnyPathPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ProxyAnyPathPost(context.Background(), anyPath).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).ProxyAnyPathPutRequest(proxyAnyPathPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProxyAnyPathPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**anyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyAnyPathPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **retries** | **string** | The number of retries in case of failure (with exponential back-off). Optional, default 0. | 
 **baseUrlOverride** | **string** | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional | 
 **decompress** | **string** | Override the decompress option when making requests. Optional, defaults to false | 
 **nangoProxyANYHEADER** | **string** | Any other headers you send are passed on to the external API | 
 **proxyAnyPathPutRequest** | [**ProxyAnyPathPutRequest**](ProxyAnyPathPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyAnyPathPut

> ProxyAnyPathPut(ctx, anyPath).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).ProxyAnyPathPutRequest(proxyAnyPathPutRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	anyPath := "anyPath_example" // string | 
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	retries := "retries_example" // string | The number of retries in case of failure (with exponential back-off). Optional, default 0. (optional)
	baseUrlOverride := "baseUrlOverride_example" // string | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional (optional)
	decompress := "decompress_example" // string | Override the decompress option when making requests. Optional, defaults to false (optional)
	nangoProxyANYHEADER := "nangoProxyANYHEADER_example" // string | Any other headers you send are passed on to the external API (optional)
	proxyAnyPathPutRequest := *openapiclient.NewProxyAnyPathPutRequest() // ProxyAnyPathPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ProxyAnyPathPut(context.Background(), anyPath).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Retries(retries).BaseUrlOverride(baseUrlOverride).Decompress(decompress).NangoProxyANYHEADER(nangoProxyANYHEADER).ProxyAnyPathPutRequest(proxyAnyPathPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProxyAnyPathPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**anyPath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyAnyPathPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **retries** | **string** | The number of retries in case of failure (with exponential back-off). Optional, default 0. | 
 **baseUrlOverride** | **string** | Provide an API base URL when the base API is not listed in the providers.yaml or it needs to be overridden. Optional | 
 **decompress** | **string** | Override the decompress option when making requests. Optional, defaults to false | 
 **nangoProxyANYHEADER** | **string** | Any other headers you send are passed on to the external API | 
 **proxyAnyPathPutRequest** | [**ProxyAnyPathPutRequest**](ProxyAnyPathPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordsGet

> RecordsGet200Response RecordsGet(ctx).Model(model).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Cursor(cursor).Limit(limit).Filter(filter).ModifiedAfter(modifiedAfter).DeltaDeprecated(deltaDeprecated).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	model := "model_example" // string | The data model to fetch
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	cursor := "cursor_example" // string | Each record from this method comes with a synchronization cursor in `_nango_metadata.cursor`. Save the last fetched record's cursor to track how far you've synced. By providing the cursor to this endpoint, you'll continue syncing from where you left off, receiving only post-cursor changes. This same cursor is used to paginate through the results of this endpoint.  (optional)
	limit := int32(56) // int32 | The maximum number of records to return. Defaults to 100. (optional)
	filter := "filter_example" // string | Filter to only show results that have been added or updated or deleted. (optional)
	modifiedAfter := "modifiedAfter_example" // string | Timestamp, e.g. 2023-05-31T11:46:13.390Z. If passed only records modified after this timestamp are returned, otherwise all records are returned. (optional)
	deltaDeprecated := "deltaDeprecated_example" // string | DEPRECATED (use modified_after) Timestamp, e.g. 2023-05-31T11:46:13.390Z. If passed, only records modified after this timestamp are returned, otherwise all records are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RecordsGet(context.Background()).Model(model).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Cursor(cursor).Limit(limit).Filter(filter).ModifiedAfter(modifiedAfter).DeltaDeprecated(deltaDeprecated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordsGet`: RecordsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RecordsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | The data model to fetch | 
 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **cursor** | **string** | Each record from this method comes with a synchronization cursor in &#x60;_nango_metadata.cursor&#x60;. Save the last fetched record&#39;s cursor to track how far you&#39;ve synced. By providing the cursor to this endpoint, you&#39;ll continue syncing from where you left off, receiving only post-cursor changes. This same cursor is used to paginate through the results of this endpoint.  | 
 **limit** | **int32** | The maximum number of records to return. Defaults to 100. | 
 **filter** | **string** | Filter to only show results that have been added or updated or deleted. | 
 **modifiedAfter** | **string** | Timestamp, e.g. 2023-05-31T11:46:13.390Z. If passed only records modified after this timestamp are returned, otherwise all records are returned. | 
 **deltaDeprecated** | **string** | DEPRECATED (use modified_after) Timestamp, e.g. 2023-05-31T11:46:13.390Z. If passed, only records modified after this timestamp are returned, otherwise all records are returned. | 

### Return type

[**RecordsGet200Response**](RecordsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScriptsConfigGet

> ScriptsConfigGet200Response ScriptsConfigGet(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ScriptsConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ScriptsConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScriptsConfigGet`: ScriptsConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ScriptsConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScriptsConfigGetRequest struct via the builder pattern


### Return type

[**ScriptsConfigGet200Response**](ScriptsConfigGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPausePost

> SyncPausePost(ctx).SyncPausePostRequest(syncPausePostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	syncPausePostRequest := *openapiclient.NewSyncPausePostRequest("ProviderConfigKey_example", []string{"Syncs_example"}) // SyncPausePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SyncPausePost(context.Background()).SyncPausePostRequest(syncPausePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SyncPausePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPausePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncPausePostRequest** | [**SyncPausePostRequest**](SyncPausePostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncRecordsGet

> []SyncRecordsGet200ResponseInner SyncRecordsGet(ctx).Model(model).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Delta(delta).Limit(limit).Offset(offset).SortBy(sortBy).Order(order).Filter(filter).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	model := "model_example" // string | The data model to fetch
	connectionId := "connectionId_example" // string | The connection ID used to create the connection.
	providerConfigKey := "providerConfigKey_example" // string | The integration ID used to create the connection (aka Unique Key).
	delta := "delta_example" // string | Timestamp, e.g. 2023-05-31T11:46:13.390Z. If passed only records added or updated after this timestamp are returned, otherwise all records are returned. (optional)
	limit := int32(56) // int32 | The maximum number of records to return. If not passed, all records are returned. (optional)
	offset := int32(56) // int32 | The number of records to skip. If not passed, no records are skipped. (optional)
	sortBy := "sortBy_example" // string | Set how the records are sorted. The default is id. The options are 'created_at', 'updated_at', 'id'. (optional)
	order := "order_example" // string | Set the order of results. The default is 'desc'. The options are 'desc' or 'asc'. (optional)
	filter := "filter_example" // string | Filter to only show results that have been added or updated or deleted. Helpful when used in conjuction with the delta parameter to retrieve a subset or records that were added, updated, or deleted with the latest sync. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SyncRecordsGet(context.Background()).Model(model).ConnectionId(connectionId).ProviderConfigKey(providerConfigKey).Delta(delta).Limit(limit).Offset(offset).SortBy(sortBy).Order(order).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SyncRecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncRecordsGet`: []SyncRecordsGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SyncRecordsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncRecordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | The data model to fetch | 
 **connectionId** | **string** | The connection ID used to create the connection. | 
 **providerConfigKey** | **string** | The integration ID used to create the connection (aka Unique Key). | 
 **delta** | **string** | Timestamp, e.g. 2023-05-31T11:46:13.390Z. If passed only records added or updated after this timestamp are returned, otherwise all records are returned. | 
 **limit** | **int32** | The maximum number of records to return. If not passed, all records are returned. | 
 **offset** | **int32** | The number of records to skip. If not passed, no records are skipped. | 
 **sortBy** | **string** | Set how the records are sorted. The default is id. The options are &#39;created_at&#39;, &#39;updated_at&#39;, &#39;id&#39;. | 
 **order** | **string** | Set the order of results. The default is &#39;desc&#39;. The options are &#39;desc&#39; or &#39;asc&#39;. | 
 **filter** | **string** | Filter to only show results that have been added or updated or deleted. Helpful when used in conjuction with the delta parameter to retrieve a subset or records that were added, updated, or deleted with the latest sync. | 

### Return type

[**[]SyncRecordsGet200ResponseInner**](SyncRecordsGet200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncStartPost

> SyncStartPost(ctx).SyncStartPostRequest(syncStartPostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	syncStartPostRequest := *openapiclient.NewSyncStartPostRequest("ProviderConfigKey_example", []string{"Syncs_example"}) // SyncStartPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SyncStartPost(context.Background()).SyncStartPostRequest(syncStartPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SyncStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncStartPostRequest** | [**SyncStartPostRequest**](SyncStartPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncStatusGet

> SyncStatusGet200Response SyncStatusGet(ctx).ProviderConfigKey(providerConfigKey).Syncs(syncs).ConnectionId(connectionId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	providerConfigKey := "providerConfigKey_example" // string | The ID of the integration you established within Nango
	syncs := "syncs_example" // string | The name of the syncs you want to fetch a status for. Pass in \"*\" to return all syncs per the integration
	connectionId := "connectionId_example" // string | The ID of the connection. If omitted, all connections will be surfaced. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SyncStatusGet(context.Background()).ProviderConfigKey(providerConfigKey).Syncs(syncs).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SyncStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncStatusGet`: SyncStatusGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SyncStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerConfigKey** | **string** | The ID of the integration you established within Nango | 
 **syncs** | **string** | The name of the syncs you want to fetch a status for. Pass in \&quot;*\&quot; to return all syncs per the integration | 
 **connectionId** | **string** | The ID of the connection. If omitted, all connections will be surfaced. | 

### Return type

[**SyncStatusGet200Response**](SyncStatusGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncTriggerPost

> SyncTriggerPost(ctx).SyncTriggerPostRequest(syncTriggerPostRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	syncTriggerPostRequest := *openapiclient.NewSyncTriggerPostRequest("ProviderConfigKey_example", []string{"Syncs_example"}) // SyncTriggerPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.SyncTriggerPost(context.Background()).SyncTriggerPostRequest(syncTriggerPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SyncTriggerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncTriggerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncTriggerPostRequest** | [**SyncTriggerPostRequest**](SyncTriggerPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncUpdateConnectionFrequencyPut

> SyncUpdateConnectionFrequencyPut200Response SyncUpdateConnectionFrequencyPut(ctx).SyncUpdateConnectionFrequencyPutRequest(syncUpdateConnectionFrequencyPutRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/capvc/nango-client"
)

func main() {
	syncUpdateConnectionFrequencyPutRequest := *openapiclient.NewSyncUpdateConnectionFrequencyPutRequest("ProviderConfigKey_example", "ConnectionId_example", "SyncName_example", "Frequency_example") // SyncUpdateConnectionFrequencyPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SyncUpdateConnectionFrequencyPut(context.Background()).SyncUpdateConnectionFrequencyPutRequest(syncUpdateConnectionFrequencyPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SyncUpdateConnectionFrequencyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncUpdateConnectionFrequencyPut`: SyncUpdateConnectionFrequencyPut200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SyncUpdateConnectionFrequencyPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncUpdateConnectionFrequencyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncUpdateConnectionFrequencyPutRequest** | [**SyncUpdateConnectionFrequencyPutRequest**](SyncUpdateConnectionFrequencyPutRequest.md) |  | 

### Return type

[**SyncUpdateConnectionFrequencyPut200Response**](SyncUpdateConnectionFrequencyPut200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

