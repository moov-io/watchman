---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Docker

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

We publish a [public Docker image `moov/watchman`](https://hub.docker.com/r/moov/watchman/) from Docker Hub or use this repository. No configuration is required to serve on `:8084`. We also have Docker images for [OpenShift](https://quay.io/repository/moov/watchman?tab=tags) published as `quay.io/moov/watchman`. Lastly, we offer a `moov/watchman:static` Docker image with files from 2019. This image can be useful for faster local testing or consistent results.

Pull & start the Docker image:
```
docker pull moov/watchman:latest
docker run -p 8084:8084 moov/watchman:latest
```

Get information about a company using their entity ID:

```
curl -s "http://localhost:8084/v2/search?name=Nicolas+Maduro&type=person&limit=1&minMatch=0.75" | jq .
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
