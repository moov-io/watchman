# Ssi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityID** | **string** | The ID assigned to an entity by the Treasury Department | [optional] 
**Type** | [**SsiType**](SsiType.md) |  | [optional] 
**Programs** | **[]string** | Sanction programs for which the entity is flagged | [optional] 
**Name** | **string** | The name of the entity | [optional] 
**Addresses** | **[]string** | Addresses associated with the entity | [optional] 
**Remarks** | **[]string** | Additional details regarding the entity | [optional] 
**AlternateNames** | **[]string** | Known aliases associated with the entity | [optional] 
**Ids** | **[]string** | IDs on file for the entity | [optional] 
**SourceListURL** | **string** | The link to the official SSI list | [optional] 
**SourceInfoURL** | **string** | The link for information regarding the source | [optional] 
**Match** | **float32** | Match percentage of search query | [optional] 
**MatchedName** | **string** | The highest scoring term from the search query. This term is the precomputed indexed value used by the search algorithm. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


