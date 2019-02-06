# \OFACApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomerNameWatch**](OFACApi.md#AddCustomerNameWatch) | **Post** /customers/watch | Add customer watch by name
[**AddCustomerWatch**](OFACApi.md#AddCustomerWatch) | **Post** /customers/{customerId}/watch | Add OFAC watch on a Customer
[**GetCustomer**](OFACApi.md#GetCustomer) | **Get** /customers/{customerId} | Get information about a customer, addresses, alternate names, and their SDN metadata.
[**GetLatestDownloads**](OFACApi.md#GetLatestDownloads) | **Get** /downloads | Return list of recent re-downloads of OFAC data
[**GetSDN**](OFACApi.md#GetSDN) | **Get** /sdn/{sdnId} | Specially designated national
[**GetSDNAddresses**](OFACApi.md#GetSDNAddresses) | **Get** /sdn/{sdnId}/addresses | Get addresses for a given SDN
[**GetSDNAltNames**](OFACApi.md#GetSDNAltNames) | **Get** /sdn/{sdnId}/alts | Get alternate names for a given SDN
[**Ping**](OFACApi.md#Ping) | **Get** /ping | Ping the OFAC service to check if running
[**RemoveCustomerNameWatch**](OFACApi.md#RemoveCustomerNameWatch) | **Delete** /customers/watch/{watchId} | Remove a Customer name watch
[**RemoveCustomerWatch**](OFACApi.md#RemoveCustomerWatch) | **Delete** /customers/{customerId}/watch/{watchId} | Remove customer watch
[**SearchSDNs**](OFACApi.md#SearchSDNs) | **Get** /search | Search SDN names and metadata
[**UpdateCustomerStatus**](OFACApi.md#UpdateCustomerStatus) | **Put** /customers/{customerId} | Update a Customer&#39;s status to add or remove a manual block.


# **AddCustomerNameWatch**
> Watch AddCustomerNameWatch(ctx, name, watchRequest, optional)
Add customer watch by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Individual name used to match and send watch notifications | 
  **watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddCustomerNameWatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Watch**](Watch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddCustomerWatch**
> Watch AddCustomerWatch(ctx, customerId, watchRequest, optional)
Add OFAC watch on a Customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
  **watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddCustomerWatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Watch**](Watch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomer**
> OfacCustomer GetCustomer(ctx, customerId, optional)
Get information about a customer, addresses, alternate names, and their SDN metadata.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***GetCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacCustomer**](OFACCustomer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestDownloads**
> []Download GetLatestDownloads(ctx, optional)
Return list of recent re-downloads of OFAC data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestDownloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetLatestDownloadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Maximum results returned by a search | 

### Return type

[**[]Download**](Download.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSDN**
> Sdn GetSDN(ctx, sdnId, optional)
Specially designated national

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sdnId** | **string**| SDN ID | 
 **optional** | ***GetSDNOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSDNOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Sdn**](SDN.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSDNAddresses**
> []Address GetSDNAddresses(ctx, sdnId, optional)
Get addresses for a given SDN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sdnId** | **string**| SDN ID | 
 **optional** | ***GetSDNAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSDNAddressesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**[]Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSDNAltNames**
> []Alt GetSDNAltNames(ctx, sdnId, optional)
Get alternate names for a given SDN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sdnId** | **string**| SDN ID | 
 **optional** | ***GetSDNAltNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSDNAltNamesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**[]Alt**](Alt.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> Ping(ctx, )
Ping the OFAC service to check if running

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCustomerNameWatch**
> RemoveCustomerNameWatch(ctx, watchId, name, optional)
Remove a Customer name watch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **watchId** | **string**| Watch ID, used to identify a specific watch | 
  **name** | **string**| Customer or Company name watch | 
 **optional** | ***RemoveCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoveCustomerNameWatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCustomerWatch**
> RemoveCustomerWatch(ctx, customerId, watchId, optional)
Remove customer watch

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
  **watchId** | **string**| Watch ID, used to identify a specific watch | 
 **optional** | ***RemoveCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoveCustomerWatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **UpdateCustomerStatus**
> UpdateCustomerStatus(ctx, customerId, updateCustomerStatus, optional)
Update a Customer's status to add or remove a manual block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
  **updateCustomerStatus** | [**UpdateCustomerStatus**](UpdateCustomerStatus.md)|  | 
 **optional** | ***UpdateCustomerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateCustomerStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

