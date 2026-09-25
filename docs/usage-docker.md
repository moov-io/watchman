---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Quick Start with Docker

Images: [`moov/watchman`](https://hub.docker.com/r/moov/watchman/) on Docker Hub, [`quay.io/moov/watchman`](https://quay.io/repository/moov/watchman?tab=tags) for OpenShift. `moov/watchman:v2-static` ships frozen 2019 files for fast local tests. Older **v0.31** API docs: [v0.31.3](https://github.com/moov-io/watchman/tree/v0.31.3/docs).

Business API on **:8084**. Admin/metrics on **:9094**. Do not expose Watchman on the public internet. See [Network access](/watchman/network/).

```
docker run -p 8084:8084 -e INCLUDED_LISTS=us_ofac moov/watchman
```

WASM UI: [http://localhost:8084](http://localhost:8084). Full recipe: [Using Watchman](/watchman/using-watchman/). Env vars: [Configuration](/watchman/config/).

For an optional [deepparse](/watchman/config/#deepparse) sidecar used in tests and examples:

```
make setup-deepparse
```

Search a person (always include `type`; prefer `minMatch=0.80`):

```
curl -s "http://localhost:8084/v2/search?name=Nicolas+Maduro&type=person&limit=1&minMatch=0.80" | jq .
```
```json
{
  "entities": [
    {
      "name": "Nicolas MADURO MOROS",
      "entityType": "person",
      "sourceList": "us_ofac",
      "sourceID": "22790",
      "person": {
        "name": "Nicolas MADURO MOROS",
        "altNames": null,
        "gender": "male",
        "birthDate": "1962-11-23T00:00:00Z",
        "deathDate": null,
        "titles": [
          "President of the Bolivarian Republic of Venezuela"
        ],
        "governmentIDs": [
          {
            "type": "cedula",
            "country": "Venezuela",
            "identifier": "5892464"
          }
        ]
      },
      "business": null,
      "organization": null,
      "aircraft": null,
      "vessel": null,
      "contact": {
        "emailAddresses": null,
        "phoneNumbers": null,
        "faxNumbers": null,
        "websites": null
      },
      "addresses": null,
      "cryptoAddresses": null,
      "affiliations": null,
      "sanctionsInfo": null,
      "historicalInfo": null,
      "sourceData": {
        "entityID": "22790",
        "sdnName": "MADURO MOROS, Nicolas",
        "sdnType": "individual",
        "program": [
          "VENEZUELA",
          "IRAN-CON-ARMS-EO"
        ],
        "title": "President of the Bolivarian Republic of Venezuela",
        "callSign": "",
        "vesselType": "",
        "tonnage": "",
        "grossRegisteredTonnage": "",
        "vesselFlag": "",
        "vesselOwner": "",
        "remarks": "DOB 23 Nov 1962; POB Caracas, Venezuela; citizen Venezuela; Gender Male; Cedula No. 5892464 (Venezuela); President of the Bolivarian Republic of Venezuela."
      },
      "match": 0.7784062500000001
    }
  ]
}
```
