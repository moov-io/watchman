---
layout: page
title: Search options
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Search

Moov Watchman offers numerous search options for inspecting the OFAC, SDN, CSL, and other related data. Certain endpoints don't inspect all supported lists.

## Supported combinations

- All fields, all lists
   - `?q=<string>`
- Name search, Searches OFAC, SSI, DPs, and BIS Entities
   - `?name=<string>`
   - An Address can be included, Only Searches the OFAC list
      - `address=<string>&city=<string>&state=<string>&providence=<string>&zip=<string>&country=<string>`
- ID search, Only searches the OFAC list
   - `?id=<string>`
- Alt Name search, Only searches the OFAC list
   - `?altName=<string>`
- Address search, Only searches the OFAC list
   - `&address=<string>&city=<string>&state=<string>&providence=<string>&zip=<string>&country=<string>`

## All in one

The most common endpoint for searching across all data Watchman has indexed. To perform this search make an HTTP query like the following:

See the [API documentation](https://moov-io.github.io/watchman/api/#get-/search) for full request/response data.

```
curl 'http://localhost:8084/search?q=nicolas+maduro&limit=1'
```
```
{
  "SDNs": [
    {
      "entityID": "22790",
      "sdnName": "MADURO MOROS, Nicolas",
      "sdnType": "individual",
      "program": "VENEZUELA",
      "title": "President of the Bolivarian Republic of Venezuela",
      "callSign": "",
      "vesselType": "",
      "tonnage": "",
      "grossRegisteredTonnage": "",
      "vesselFlag": "",
      "vesselOwner": "",
      "remarks": "DOB 23 Nov 1962; POB Caracas, Venezuela; citizen Venezuela; Gender Male; Cedula No. 5892464 (Venezuela); President of the Bolivarian Republic of Venezuela.",
      "match": 0.9444444444444444
    }
  ],
  "altNames": [
    {
      "entityID": "11318",
      "alternateID": "11868",
      "alternateType": "aka",
      "alternateName": "NICO",
      "alternateRemarks": "",
      "match": 0.8615384615384615
    }
  ],
  "addresses": [
    {
      "entityID": "16182",
      "addressID": "24386",
      "address": "Colonia Moderna",
      "cityStateProvincePostalCode": "San Pedro Sula, Cortes",
      "country": "Honduras",
      "addressRemarks": "",
      "match": 0.7804695304695305
    }
  ],
  "deniedPersons": [
    {
      "name": "LISONG MA",
      "streetAddress": "INMATE NUMBER - 80644-053, MOSHANNON VALLEY, CORRECTIONAL INSTITUTION, 555 GEO DRIVE",
      "city": "PHILIPSBURG",
      "state": "PA",
      "country": "US",
      "postalCode": "16866",
      "effectiveDate": "10/31/2014",
      "expirationDate": "05/27/2024",
      "standardOrder": "Y",
      "lastUpdate": "2014-11-07",
      "action": "FR NOTICE ADDED",
      "frCitation": "79 F.R. 66354 11/7/14",
      "match": 0.7673992673992674
    }
  ]
}
```

## SDN names

This search operation will only return results matching SDN names from your query:

```
curl 'http://localhost:8084/search?name=nicolas+maduro&limit=1'
```
```
{
  "SDNs": [
    {
      "entityID": "22790",
      "sdnName": "MADURO MOROS, Nicolas",
      "sdnType": "individual",
      "program": "VENEZUELA",
      "title": "President of the Bolivarian Republic of Venezuela",
      "callSign": "",
      "vesselType": "",
      "tonnage": "",
      "grossRegisteredTonnage": "",
      "vesselFlag": "",
      "vesselOwner": "",
      "remarks": "DOB 23 Nov 1962; POB Caracas, Venezuela; citizen Venezuela; Gender Male; Cedula No. 5892464 (Venezuela); President of the Bolivarian Republic of Venezuela.",
      "match": 0.9444444444444444
    }
  ],
  "altNames": null,
  "addresses": null,
  "deniedPersons": null
}
```

## SDN remark IDs

SDN Remarks contain semi-structured data which Watchman attempts to parse. One common element of this data is a National or Governmental ID which uniquely identifies an entity.

```
curl 'http://localhost:8084/search?id=5892464&limit=1'
```
```
{
  "SDNs": [
    {
      "entityID": "22790",
      "sdnName": "MADURO MOROS, Nicolas",
      "sdnType": "individual",
      "program": "VENEZUELA",
      "title": "President of the Bolivarian Republic of Venezuela",
      "callSign": "",
      "vesselType": "",
      "tonnage": "",
      "grossRegisteredTonnage": "",
      "vesselFlag": "",
      "vesselOwner": "",
      "remarks": "DOB 23 Nov 1962; POB Caracas, Venezuela; citizen Venezuela; Gender Male; Cedula No. 5892464 (Venezuela); President of the Bolivarian Republic of Venezuela.",
      "match": 1
    }
  ],
  "altNames": null,
  "addresses": null,
  "deniedPersons": null
}
```

## SDN alternate names

Often an entity will have multiple names which are in the OFAC dataset:

```
curl 'http://localhost:8084/search?altName=NATIONAL+BANK+OF+CUBA&limit=1'
```
```
{
  "SDNs": null,
  "altNames": [
    {
      "entityID": "306",
      "alternateID": "220",
      "alternateType": "aka",
      "alternateName": "NATIONAL BANK OF CUBA",
      "alternateRemarks": "",
      "match": 1
    }
  ],
  "addresses": null,
  "deniedPersons": null
}
```

Note - The SDN has an alternate name (in this case its primary name is its regional name):

```
curl 'http://localhost:8084/sdn/306'
```
```
{
  "entityID": "306",
  "sdnName": "BANCO NACIONAL DE CUBA",
  "sdnType": "",
  "program": "CUBA",
  "title": "",
  "callSign": "",
  "vesselType": "",
  "tonnage": "",
  "grossRegisteredTonnage": "",
  "vesselFlag": "",
  "vesselOwner": "",
  "remarks": "a.k.a. 'BNC'."
}
```

## SDN addresses

An address can also be a query against the OFAC data. There are multiple query parameters available here to further refine results:

- address
- city
- state
- providence
- zip
- country

```
curl 'http://localhost:8084/search?address=first+st&province=harare&country=zimbabew&limit=1'
```
```
{
  "SDNs": null,
  "altNames": null,
  "addresses": [
    {
      "entityID": "8178",
      "addressID": "7437",
      "address": "First Floor, Victory House, 88 Robert Mugabe Road",
      "cityStateProvincePostalCode": "Harare",
      "country": "Zimbabwe",
      "addressRemarks": "",
      "match": 0.8261904761904761
    }
  ],
  "deniedPersons": null
}
```

## Filtering

Moov Watchman offers filters to further refine search results. The supported query parameters are:

- `sdnType`: This is commonly `individual`, `aicraft`, or `vessel`.
- `program`: The specific U.S. sanctions program which added the entity. (Example: `SDGT`)

```
curl 'http://localhost:8084/search?name=EP&sdnType=aircraft&limit=1&program=sdgt'
```
```
{
  "SDNs": [
    {
      "entityID": "15431",
      "sdnName": "EP-GOM",
      "sdnType": "aircraft",
      "program": "SDGT",
      "title": "",
      "callSign": "",
      "vesselType": "",
      "tonnage": "",
      "grossRegisteredTonnage": "",
      "vesselFlag": "",
      "vesselOwner": "",
      "remarks": "Aircraft Construction Number (also called L/N or S/N or F/N) 8401; Aircraft Manufacture Date 1992; Aircraft Model IL76-TD; Aircraft Operator YAS AIR; Aircraft Manufacturer's Serial Number (MSN) 1023409321; Linked To: POUYA AIR.",
      "match": 0.84
    }
  ],
  "altNames": null,
  "addresses": null,
  "deniedPersons": null
}
```

## US Consolidated Screening List (CSL)

Moov Watchman offers searching the entire US CSL list. The supported query parameters are:

- `name`: Legal name of entity on list
- `limit`: Maximum number of results to return

Refer to the [API docs for searching US CSL](https://moov-io.github.io/watchman/api/#get-/search/us-csl) for more details.

```
curl "http://localhost:8084/search/us-csl?name=Al&limit=10
```
```
{
  "SDNs": null,
  "altNames": null,
  "addresses": null,
  "deniedPersons": null,
  "bisEntities": [ ... ],
  "militaryEndUsers": [ ... ],
  "sectoralSanctions": [ ... ],
  "unverifiedCSL": [ ... ],
  "nonproliferationSanctions": [ ... ],
  "foreignSanctionsEvaders": [ ... ],
  "palestinianLegislativeCouncil": [ ... ],
  "captaList": [ ... ],
  "itarDebarred": [ ... ],
  "nonSDNChineseMilitaryIndustrialComplex": [ ... ],
  "nonSDNMenuBasedSanctionsList": [ ... ],
  "refreshedAt": "2022-09-07T20:35:35.773313Z"
}
```
