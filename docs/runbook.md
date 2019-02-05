### Change OFAC refresh frequency

`OFAC_DATA_REFRESH=1h0m0s` can be set to refresh OFAC data more or less often. The value should match Go's `time.ParseDuration` syntax.

### Force OFAC data refresh

Make a request to `/ofac/refresh` on the **admin** HTTP interface (`:9094` by default).

```
$ curl http://localhost:9094/ofac/refresh
{"SDNs":7414,"altNames":9729,"addresses":11747}
```

### Change OFAC download URL

By default OFAC downloads [various files from treasury.gov](https://www.treasury.gov/resource-center/sanctions/SDN-List/Pages/default.aspx) on startup and will periodically download them to keep the data updated.

This URL change be changed and follows a template as several files are downloaded (example: `add.csv` or `sdn.csv`). To change where OFAC files are downloaded set:

`OFAC_DOWNLOAD_TEMPLATE='https://www.treasury.gov/ofac/downloads/%s'`

You should make the following files available at the new endpoint: `add.csv`, `alt.csv`, `sdn.csv`, `sdn_comments.csv`.

### Change SQLite storage location

To change where the SQLite database is stored on disk set `SQLITE_DB_PATH` as an environmental variable.

### Change name and/or address matching threshold

To adjust how sensitive search results are to their queries the following values can be modified, lower values indicate more results are included. By default all values are `0.90`.

- `NAME_SIMILARITY`: Similarity required for SDN name searches
- `ALT_SIMILARITY`: Similarity required for alternate name searches
- `ADDRESS_SIMILARITY`: Similarity required for address searches

### Webhook batch processing size

The size of each batch of watches to be processed (and their webhook called) can be adjusted with `WEBHOOK_BATCH_SIZE=100`. This is intended for performance improvements by using a larger batch size.
