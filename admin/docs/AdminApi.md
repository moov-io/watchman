# \AdminApi

All URIs are relative to *http://localhost:9094*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DebugSDN**](AdminApi.md#DebugSDN) | **Get** /debug/sdn/{sdnId} | Get an SDN and index debug information
[**GetVersion**](AdminApi.md#GetVersion) | **Get** /version | Show the current version
[**RefreshData**](AdminApi.md#RefreshData) | **Post** /data/refresh | Download and reindex all data sources



## DebugSDN

> DebugSdn DebugSDN(ctx, sdnId)

Get an SDN and index debug information

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

Show the current version

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

