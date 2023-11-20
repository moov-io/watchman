# \WatchmanApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestDownloads**](WatchmanApi.md#GetLatestDownloads) | **Get** /downloads | Get latest downloads
[**GetSDNAddresses**](WatchmanApi.md#GetSDNAddresses) | **Get** /ofac/sdn/{sdnID}/addresses | Get SDN addresses
[**GetSDNAltNames**](WatchmanApi.md#GetSDNAltNames) | **Get** /ofac/sdn/{sdnID}/alts | Get SDN alt names
[**GetUIValues**](WatchmanApi.md#GetUIValues) | **Get** /ui/values/{key} | Get UI values
[**Ping**](WatchmanApi.md#Ping) | **Get** /ping | Ping Watchman service
[**Search**](WatchmanApi.md#Search) | **Get** /search | Search
[**SearchUSCSL**](WatchmanApi.md#SearchUSCSL) | **Get** /search/us-csl | Search US CSL



## GetLatestDownloads

> []Download GetLatestDownloads(ctx, optional)

Get latest downloads

Return list of recent downloads of list data.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestDownloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestDownloadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 
 **limit** | **optional.Int32**| Maximum number of downloads to return sorted by their timestamp in decending order. | 

### Return type

[**[]Download**](Download.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSDNAddresses

> []OfacEntityAddress GetSDNAddresses(ctx, sdnID, optional)

Get SDN addresses

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnID** | **string**| SDN ID | 
 **optional** | ***GetSDNAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNAddressesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 

### Return type

[**[]OfacEntityAddress**](OfacEntityAddress.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSDNAltNames

> []OfacAlt GetSDNAltNames(ctx, sdnID, optional)

Get SDN alt names

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnID** | **string**| SDN ID | 
 **optional** | ***GetSDNAltNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNAltNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 

### Return type

[**[]OfacAlt**](OfacAlt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUIValues

> []SdnType GetUIValues(ctx, key, optional)

Get UI values

Return an ordered distinct list of keys for an SDN property.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**SdnType**](.md)| SDN property to lookup. Values are sdnType, ofacProgram | 
 **optional** | ***GetUIValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUIValuesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of UI keys returned | 

### Return type

[**[]SdnType**](SdnType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> Ping(ctx, )

Ping Watchman service

Check if the Watchman service is running.

### Required Parameters

This endpoint does not need any parameter.

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


## Search

> Search Search(ctx, optional)

Search

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 
 **q** | **optional.String**| Search across Name, Alt Names, and SDN Address fields for all available sanctions lists. Entries may be returned in all response sub-objects. | 
 **name** | **optional.String**| Name which could correspond to an entry on the SDN, Denied Persons, Sectoral Sanctions Identifications, or BIS Entity List sanctions lists. Alt names are also searched. | 
 **address** | **optional.String**| Physical address which could correspond to a human on the SDN list. Only Address results will be returned. | 
 **city** | **optional.String**| City name as desginated by SDN guidelines. Only Address results will be returned. | 
 **state** | **optional.String**| State name as desginated by SDN guidelines. Only Address results will be returned. | 
 **providence** | **optional.String**| Providence name as desginated by SDN guidelines. Only Address results will be returned. | 
 **zip** | **optional.String**| Zip code as desginated by SDN guidelines. Only Address results will be returned. | 
 **country** | **optional.String**| Country name as desginated by SDN guidelines. Only Address results will be returned. | 
 **altName** | **optional.String**| Alternate name which could correspond to a human on the SDN list. Only Alt name results will be returned. | 
 **id** | **optional.String**| ID value often found in remarks property of an SDN. Takes the form of &#39;No. NNNNN&#39; as an alphanumeric value. | 
 **minMatch** | **optional.Float32**| Match percentage that search query must obtain for results to be returned. | 
 **limit** | **optional.Int32**| Maximum results returned by a search. Results are sorted by their match percentage in decending order. | 
 **sdnType** | [**optional.Interface of SdnType**](.md)| Optional filter to only return SDNs whose type case-insensitively matches. | 
 **program** | **optional.String**| Optional filter to only return SDNs whose program case-insensitively matches. | 

### Return type

[**Search**](Search.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUSCSL

> Search SearchUSCSL(ctx, optional)

Search US CSL

Search the US Consolidated Screening List

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchUSCSLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUSCSLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the system&#39;s logs | 
 **name** | **optional.String**| Name which could correspond to an entry on the CSL | 
 **limit** | **optional.Int32**| Maximum number of downloads to return sorted by their timestamp in decending order. | 

### Return type

[**Search**](Search.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

