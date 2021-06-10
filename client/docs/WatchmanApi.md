# \WatchmanApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOfacCompanyNameWatch**](WatchmanApi.md#AddOfacCompanyNameWatch) | **Post** /ofac/companies/watch | Watch company
[**AddOfacCompanyWatch**](WatchmanApi.md#AddOfacCompanyWatch) | **Post** /ofac/companies/{companyID}/watch | Watch OFAC company
[**AddOfacCustomerNameWatch**](WatchmanApi.md#AddOfacCustomerNameWatch) | **Post** /ofac/customers/watch | Watch customer
[**AddOfacCustomerWatch**](WatchmanApi.md#AddOfacCustomerWatch) | **Post** /ofac/customers/{customerID}/watch | Watch OFAC customer
[**GetLatestDownloads**](WatchmanApi.md#GetLatestDownloads) | **Get** /downloads | Get latest downloads
[**GetOfacCompany**](WatchmanApi.md#GetOfacCompany) | **Get** /ofac/companies/{companyID} | Get company
[**GetOfacCustomer**](WatchmanApi.md#GetOfacCustomer) | **Get** /ofac/customers/{customerID} | Get Customer
[**GetSDN**](WatchmanApi.md#GetSDN) | **Get** /ofac/sdn/{sdnID} | Get SDN
[**GetSDNAddresses**](WatchmanApi.md#GetSDNAddresses) | **Get** /ofac/sdn/{sdnID}/addresses | Get SDN addresses
[**GetSDNAltNames**](WatchmanApi.md#GetSDNAltNames) | **Get** /ofac/sdn/{sdnID}/alts | Get SDN alt names
[**GetUIValues**](WatchmanApi.md#GetUIValues) | **Get** /ui/values/{key} | Get UI values
[**Ping**](WatchmanApi.md#Ping) | **Get** /ping | Ping Watchman
[**RemoveOfacCompanyNameWatch**](WatchmanApi.md#RemoveOfacCompanyNameWatch) | **Delete** /ofac/companies/watch/{watchID} | Remove company watch
[**RemoveOfacCompanyWatch**](WatchmanApi.md#RemoveOfacCompanyWatch) | **Delete** /ofac/companies/{companyID}/watch/{watchID} | Remove company watch
[**RemoveOfacCustomerNameWatch**](WatchmanApi.md#RemoveOfacCustomerNameWatch) | **Delete** /ofac/customers/watch/{watchID} | Remove customer watch
[**RemoveOfacCustomerWatch**](WatchmanApi.md#RemoveOfacCustomerWatch) | **Delete** /ofac/customers/{customerID}/watch/{watchID} | Remove customer watch
[**Search**](WatchmanApi.md#Search) | **Get** /search | Search SDNs
[**UpdateOfacCompanyStatus**](WatchmanApi.md#UpdateOfacCompanyStatus) | **Put** /ofac/companies/{companyID} | Update company
[**UpdateOfacCustomerStatus**](WatchmanApi.md#UpdateOfacCustomerStatus) | **Put** /ofac/customers/{customerID} | Update customer



## AddOfacCompanyNameWatch

> OfacWatch AddOfacCompanyNameWatch(ctx, name, ofacWatchRequest, optional)

Watch company

Watch a company by its name. The match percentage will be included in the webhook's JSON payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Company name used to match and send watch notifications | 
**ofacWatchRequest** | [**OfacWatchRequest**](OfacWatchRequest.md)|  | 
 **optional** | ***AddOfacCompanyNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOfacCompanyNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacWatch**](OfacWatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOfacCompanyWatch

> OfacWatch AddOfacCompanyWatch(ctx, companyID, ofacWatchRequest, optional)

Watch OFAC company

Add name watch on a OFAC Company

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
**ofacWatchRequest** | [**OfacWatchRequest**](OfacWatchRequest.md)|  | 
 **optional** | ***AddOfacCompanyWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOfacCompanyWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacWatch**](OfacWatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOfacCustomerNameWatch

> OfacWatch AddOfacCustomerNameWatch(ctx, name, ofacWatchRequest, optional)

Watch customer

Watch a customer by its name. The match percentage will be included in the webhook's JSON payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Individual name used to match and send watch notifications | 
**ofacWatchRequest** | [**OfacWatchRequest**](OfacWatchRequest.md)|  | 
 **optional** | ***AddOfacCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOfacCustomerNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacWatch**](OfacWatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddOfacCustomerWatch

> OfacWatch AddOfacCustomerWatch(ctx, customerID, ofacWatchRequest, optional)

Watch OFAC customer

Add name watch on a OFAC Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
**ofacWatchRequest** | [**OfacWatchRequest**](OfacWatchRequest.md)|  | 
 **optional** | ***AddOfacCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOfacCustomerWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacWatch**](OfacWatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestDownloads

> []Download GetLatestDownloads(ctx, optional)

Get latest downloads

Return list of recent downloads of list data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestDownloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestDownloadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
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


## GetOfacCompany

> OfacCompany GetOfacCompany(ctx, companyID, optional)

Get company

Get information about a company, trust or organization such as addresses, alternate names, and remarks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
 **optional** | ***GetOfacCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOfacCompanyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacCompany**](OfacCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfacCustomer

> OfacCustomer GetOfacCustomer(ctx, customerID, optional)

Get Customer

Get information about a customer, addresses, alternate names, and their SDN metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
 **optional** | ***GetOfacCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOfacCustomerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacCustomer**](OfacCustomer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSDN

> OfacSdn GetSDN(ctx, sdnID, optional)

Get SDN

Get SDN details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnID** | **string**| SDN ID | 
 **optional** | ***GetSDNOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacSdn**](OfacSDN.md)

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

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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
**key** | **string**| SDN property to lookup. Values are sdnType, ofacProgram | 
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

Ping Watchman

Check the Watchman service is running

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


## RemoveOfacCompanyNameWatch

> RemoveOfacCompanyNameWatch(ctx, watchID, name, optional)

Remove company watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchID** | **string**| Watch ID, used to identify a specific watch | 
**name** | **string**| Company name watch | 
 **optional** | ***RemoveOfacCompanyNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOfacCompanyNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## RemoveOfacCompanyWatch

> RemoveOfacCompanyWatch(ctx, companyID, watchID, optional)

Remove company watch

Delete a company name watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
**watchID** | **string**| Watch ID, used to identify a specific watch | 
 **optional** | ***RemoveOfacCompanyWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOfacCompanyWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## RemoveOfacCustomerNameWatch

> RemoveOfacCustomerNameWatch(ctx, watchID, name, optional)

Remove customer watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchID** | **string**| Watch ID, used to identify a specific watch | 
**name** | **string**| Customer or Company name watch | 
 **optional** | ***RemoveOfacCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOfacCustomerNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## RemoveOfacCustomerWatch

> RemoveOfacCustomerWatch(ctx, customerID, watchID, optional)

Remove customer watch

Delete a customer name watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
**watchID** | **string**| Watch ID, used to identify a specific watch | 
 **optional** | ***RemoveOfacCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOfacCustomerWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## Search

> Search Search(ctx, optional)

Search SDNs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **q** | **optional.String**| Search across Name, Alt Names, and SDN Address fields for all available sanctions lists. Entries may be returned in all response sub-objects. | 
 **name** | **optional.String**| Name which could correspond to an entry on the SDN, Denied Persons, Sectoral Sanctions Identifications, or BIS Entity List sanctions lists. Alt names are also searched. | 
 **address** | **optional.String**| Phsical address which could correspond to a human on the SDN list. Only Address results will be returned. | 
 **city** | **optional.String**| City name as desginated by SDN guidelines. Only Address results will be returned. | 
 **state** | **optional.String**| State name as desginated by SDN guidelines. Only Address results will be returned. | 
 **providence** | **optional.String**| Providence name as desginated by SDN guidelines. Only Address results will be returned. | 
 **zip** | **optional.String**| Zip code as desginated by SDN guidelines. Only Address results will be returned. | 
 **country** | **optional.String**| Country name as desginated by SDN guidelines. Only Address results will be returned. | 
 **altName** | **optional.String**| Alternate name which could correspond to a human on the SDN list. Only Alt name results will be returned. | 
 **id** | **optional.String**| ID value often found in remarks property of an SDN. Takes the form of &#39;No. NNNNN&#39; as an alphanumeric value. | 
 **minMatch** | **optional.Float32**| Match percentage that search query must obtain for results to be returned. | 
 **limit** | **optional.Int32**| Maximum results returned by a search. Results are sorted by their match percentage in decending order. | 
 **sdnType** | **optional.String**| Optional filter to only return SDNs whose type case-insensitively matches. | 
 **program** | **optional.String**| Optional filter to only return SDNs whose program case-insensitively matches | 

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


## UpdateOfacCompanyStatus

> UpdateOfacCompanyStatus(ctx, companyID, updateOfacCompanyStatus, optional)

Update company

Update a Companies sanction status to always block or always allow transactions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
**updateOfacCompanyStatus** | [**UpdateOfacCompanyStatus**](UpdateOfacCompanyStatus.md)|  | 
 **optional** | ***UpdateOfacCompanyStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOfacCompanyStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| User ID used to perform this search | 

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


## UpdateOfacCustomerStatus

> UpdateOfacCustomerStatus(ctx, customerID, updateOfacCustomerStatus, optional)

Update customer

Update a Customer sanction status to always block or always allow transactions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
**updateOfacCustomerStatus** | [**UpdateOfacCustomerStatus**](UpdateOfacCustomerStatus.md)|  | 
 **optional** | ***UpdateOfacCustomerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOfacCustomerStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| User ID used to perform this search | 

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

