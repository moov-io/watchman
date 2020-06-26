## High Availability

Watchman doesn't currently support operating in a high availability (HA) mode. The major reason of doing so would be to distribute webhook notifications for high volume watches. We've not concerned ourselves with this as the notification rate is higher than anything we've needed at Moov. The complexity of implementing HA is estimated to be too much for our current needs.

Given these assumptions we've chosen to focus Watchman's vertical scaling (add CPUs and memory) instead of clustering. Currently two instances of Watchman will write over each other and send duplicate webhook notifications.

### Dependencies

#### Database

When Watchman is connected to a MySQL instance the database will need to be deployed in an acceptable manor for replication, failure recovery, and backups.

SQLite replication (possibly implemented via [rqlite](https://github.com/rqlite/rqlite)) has not been tested, but looks promising.
