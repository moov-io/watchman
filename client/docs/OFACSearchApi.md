# \OFACSearchApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSDNs**](OFACSearchApi.md#SearchSDNs) | **Get** /search | Search SDN names and metadata


# **SearchSDNs**
> Search SearchSDNs(ctx, optional)
Search SDN names and metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchSDNsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchSDNsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **name** | **optional.String**| Name which could correspond to a human on the SDN list | 
 **address** | **optional.String**| Phsical address which could correspond to a human on the SDN list | 
 **altName** | **optional.String**| Alternate name which could correspond to a human on the SDN list | 
 **limit** | **optional.Int32**| Maximum results returned by a search | 

### Return type

[**Search**](Search.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

