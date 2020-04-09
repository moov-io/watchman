# \AdminApi

All URIs are relative to *http://localhost:9094*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DebugSDN**](AdminApi.md#DebugSDN) | **Get** /debug/sdn/{sdnId} | Debug SDN
[**GetVersion**](AdminApi.md#GetVersion) | **Get** /version | Get Version
[**RefreshData**](AdminApi.md#RefreshData) | **Post** /data/refresh | Download and reindex all data sources



## DebugSDN

> DebugSdn DebugSDN(ctx, sdnId)

Debug SDN

Get an SDN and search index debug information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnId** | **string**| SDN ID | 

### Return type

[**DebugSdn**](DebugSDN.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> string GetVersion(ctx, )

Get Version

Show the current version of Watchman

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshData

> DataRefresh RefreshData(ctx, )

Download and reindex all data sources

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DataRefresh**](DataRefresh.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

