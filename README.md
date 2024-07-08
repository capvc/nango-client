# Go API client for nango

Nango API specs used to authorize & sync data with external APIs.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.7.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import nango "github.com/capvc/nango-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `nango.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), nango.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `nango.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), nango.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `nango.ContextOperationServerIndices` and `nango.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), nango.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), nango.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.nango.dev*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**ActionTriggerPost**](docs/DefaultAPI.md#actiontriggerpost) | **Post** /action/trigger | 
*DefaultAPI* | [**ConfigGet**](docs/DefaultAPI.md#configget) | **Get** /config | 
*DefaultAPI* | [**ConfigPost**](docs/DefaultAPI.md#configpost) | **Post** /config | 
*DefaultAPI* | [**ConfigProviderConfigKeyDelete**](docs/DefaultAPI.md#configproviderconfigkeydelete) | **Delete** /config/{providerConfigKey} | 
*DefaultAPI* | [**ConfigProviderConfigKeyGet**](docs/DefaultAPI.md#configproviderconfigkeyget) | **Get** /config/{providerConfigKey} | 
*DefaultAPI* | [**ConfigPut**](docs/DefaultAPI.md#configput) | **Put** /config | 
*DefaultAPI* | [**ConnectionConnectionIdDelete**](docs/DefaultAPI.md#connectionconnectioniddelete) | **Delete** /connection/{connectionId} | 
*DefaultAPI* | [**ConnectionConnectionIdGet**](docs/DefaultAPI.md#connectionconnectionidget) | **Get** /connection/{connectionId} | 
*DefaultAPI* | [**ConnectionConnectionIdMetadataPatch**](docs/DefaultAPI.md#connectionconnectionidmetadatapatch) | **Patch** /connection/{connectionId}/metadata | 
*DefaultAPI* | [**ConnectionConnectionIdMetadataPost**](docs/DefaultAPI.md#connectionconnectionidmetadatapost) | **Post** /connection/{connectionId}/metadata | 
*DefaultAPI* | [**ConnectionGet**](docs/DefaultAPI.md#connectionget) | **Get** /connection | 
*DefaultAPI* | [**ConnectionMetadataPatch**](docs/DefaultAPI.md#connectionmetadatapatch) | **Patch** /connection/metadata | 
*DefaultAPI* | [**ConnectionMetadataPost**](docs/DefaultAPI.md#connectionmetadatapost) | **Post** /connection/metadata | 
*DefaultAPI* | [**ConnectionPost**](docs/DefaultAPI.md#connectionpost) | **Post** /connection | 
*DefaultAPI* | [**EnvironmentVariablesGet**](docs/DefaultAPI.md#environmentvariablesget) | **Get** /environment-variables | 
*DefaultAPI* | [**ProxyAnyPathDelete**](docs/DefaultAPI.md#proxyanypathdelete) | **Delete** /proxy/{anyPath} | 
*DefaultAPI* | [**ProxyAnyPathGet**](docs/DefaultAPI.md#proxyanypathget) | **Get** /proxy/{anyPath} | 
*DefaultAPI* | [**ProxyAnyPathPatch**](docs/DefaultAPI.md#proxyanypathpatch) | **Patch** /proxy/{anyPath} | 
*DefaultAPI* | [**ProxyAnyPathPost**](docs/DefaultAPI.md#proxyanypathpost) | **Post** /proxy/{anyPath} | 
*DefaultAPI* | [**ProxyAnyPathPut**](docs/DefaultAPI.md#proxyanypathput) | **Put** /proxy/{anyPath} | 
*DefaultAPI* | [**RecordsGet**](docs/DefaultAPI.md#recordsget) | **Get** /records | 
*DefaultAPI* | [**ScriptsConfigGet**](docs/DefaultAPI.md#scriptsconfigget) | **Get** /scripts/config | 
*DefaultAPI* | [**SyncPausePost**](docs/DefaultAPI.md#syncpausepost) | **Post** /sync/pause | 
*DefaultAPI* | [**SyncRecordsGet**](docs/DefaultAPI.md#syncrecordsget) | **Get** /sync/records | 
*DefaultAPI* | [**SyncStartPost**](docs/DefaultAPI.md#syncstartpost) | **Post** /sync/start | 
*DefaultAPI* | [**SyncStatusGet**](docs/DefaultAPI.md#syncstatusget) | **Get** /sync/status | 
*DefaultAPI* | [**SyncTriggerPost**](docs/DefaultAPI.md#synctriggerpost) | **Post** /sync/trigger | 
*DefaultAPI* | [**SyncUpdateConnectionFrequencyPut**](docs/DefaultAPI.md#syncupdateconnectionfrequencyput) | **Put** /sync/update-connection-frequency | 


## Documentation For Models

 - [ActionTriggerPost200Response](docs/ActionTriggerPost200Response.md)
 - [ActionTriggerPostRequest](docs/ActionTriggerPostRequest.md)
 - [ConfigGet200Response](docs/ConfigGet200Response.md)
 - [ConfigGet200ResponseConfigsInner](docs/ConfigGet200ResponseConfigsInner.md)
 - [ConfigPostRequest](docs/ConfigPostRequest.md)
 - [ConfigProviderConfigKeyGet200Response](docs/ConfigProviderConfigKeyGet200Response.md)
 - [ConfigProviderConfigKeyGet200ResponseConfig](docs/ConfigProviderConfigKeyGet200ResponseConfig.md)
 - [ConfigProviderConfigKeyGet200ResponseConfigActionsInner](docs/ConfigProviderConfigKeyGet200ResponseConfigActionsInner.md)
 - [ConfigProviderConfigKeyGet200ResponseConfigSyncsInner](docs/ConfigProviderConfigKeyGet200ResponseConfigSyncsInner.md)
 - [ConfigPut200Response](docs/ConfigPut200Response.md)
 - [ConfigPut400Response](docs/ConfigPut400Response.md)
 - [ConfigPutRequest](docs/ConfigPutRequest.md)
 - [ConnectionGet200Response](docs/ConnectionGet200Response.md)
 - [ConnectionGet200ResponseConfigsInner](docs/ConnectionGet200ResponseConfigsInner.md)
 - [ConnectionMetadataPatch200Response](docs/ConnectionMetadataPatch200Response.md)
 - [ConnectionMetadataPatch200ResponseConnectionId](docs/ConnectionMetadataPatch200ResponseConnectionId.md)
 - [ConnectionMetadataPatchRequest](docs/ConnectionMetadataPatchRequest.md)
 - [ConnectionMetadataPatchRequestConnectionId](docs/ConnectionMetadataPatchRequestConnectionId.md)
 - [ConnectionMetadataPost201Response](docs/ConnectionMetadataPost201Response.md)
 - [ConnectionMetadataPost201ResponseConnectionId](docs/ConnectionMetadataPost201ResponseConnectionId.md)
 - [ConnectionMetadataPostRequest](docs/ConnectionMetadataPostRequest.md)
 - [ConnectionMetadataPostRequestConnectionId](docs/ConnectionMetadataPostRequestConnectionId.md)
 - [ConnectionPostRequest](docs/ConnectionPostRequest.md)
 - [EnvironmentVariablesGet200ResponseInner](docs/EnvironmentVariablesGet200ResponseInner.md)
 - [ProxyAnyPathPutRequest](docs/ProxyAnyPathPutRequest.md)
 - [RecordsGet200Response](docs/RecordsGet200Response.md)
 - [RecordsGet200ResponseRecordsInner](docs/RecordsGet200ResponseRecordsInner.md)
 - [RecordsGet200ResponseRecordsInnerNangoMetadata](docs/RecordsGet200ResponseRecordsInnerNangoMetadata.md)
 - [ScriptsConfigGet200Response](docs/ScriptsConfigGet200Response.md)
 - [ScriptsConfigGet200ResponseActionsInner](docs/ScriptsConfigGet200ResponseActionsInner.md)
 - [ScriptsConfigGet200ResponseActionsInnerModelsInner](docs/ScriptsConfigGet200ResponseActionsInnerModelsInner.md)
 - [ScriptsConfigGet200ResponseActionsInnerModelsInnerFieldsInner](docs/ScriptsConfigGet200ResponseActionsInnerModelsInnerFieldsInner.md)
 - [ScriptsConfigGet200ResponseSyncsInner](docs/ScriptsConfigGet200ResponseSyncsInner.md)
 - [ScriptsConfigGet200ResponseSyncsInnerModelsInner](docs/ScriptsConfigGet200ResponseSyncsInnerModelsInner.md)
 - [ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner](docs/ScriptsConfigGet200ResponseSyncsInnerModelsInnerFieldsInner.md)
 - [SyncPausePostRequest](docs/SyncPausePostRequest.md)
 - [SyncRecordsGet200ResponseInner](docs/SyncRecordsGet200ResponseInner.md)
 - [SyncRecordsGet200ResponseInnerNangoMetadata](docs/SyncRecordsGet200ResponseInnerNangoMetadata.md)
 - [SyncStartPostRequest](docs/SyncStartPostRequest.md)
 - [SyncStatusGet200Response](docs/SyncStatusGet200Response.md)
 - [SyncStatusGet200ResponseSyncsInner](docs/SyncStatusGet200ResponseSyncsInner.md)
 - [SyncTriggerPostRequest](docs/SyncTriggerPostRequest.md)
 - [SyncUpdateConnectionFrequencyPut200Response](docs/SyncUpdateConnectionFrequencyPut200Response.md)
 - [SyncUpdateConnectionFrequencyPutRequest](docs/SyncUpdateConnectionFrequencyPutRequest.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



