---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Docker

We publish a [public Docker image `moov/watchman`](https://hub.docker.com/r/moov/watchman/) from Docker Hub or use this repository. No configuration is required to serve on `:8084` and metrics at `:9094/metrics` in Prometheus format. We also have Docker images for [OpenShift](https://quay.io/repository/moov/watchman?tab=tags) published as `quay.io/moov/watchman`. Lastly, we offer a `moov/watchman:static` Docker image with files from 2019. This image can be useful for faster local testing or consistent results.

Pull & start the Docker image:
```
docker pull moov/watchman:latest
docker run -p 8084:8084 -p 9094:9094 moov/watchman:latest
```

Get information about a company using their entity ID:
```
curl "localhost:8084/ofac/companies/13374"
```
```
{
   "id":"13374",
   "sdn":{
      "entityID":"13374",
      "sdnName":"SYRONICS",
      "sdnType":"",
      "program":[
         "NPWMD"
      ],
      "title":"",
      "callSign":"",
      "vesselType":"",
      "tonnage":"",
      "grossRegisteredTonnage":"",
      "vesselFlag":"",
      "vesselOwner":"",
      "remarks":""
   },
   "addresses":[
      {
         "entityID":"13374",
         "addressID":"21360",
         "address":"Kaboon Street, PO Box 5966",
         "cityStateProvincePostalCode":"Damascus",
         "country":"Syria",
         "addressRemarks":""
      }
   ],
   "alts":[
      {
         "entityID":"13374",
         "alternateID":"15056",
         "alternateType":"aka",
         "alternateName":"SYRIAN ARAB CO. FOR ELECTRONIC INDUSTRIES",
         "alternateRemarks":""
      }
   ],
   "comments":null,
   "status":null
}
```