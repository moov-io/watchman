---
layout: page
title: High availability
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# High availability

Watchman doesn't currently support operating in a high availability (HA) mode. The major reason of doing so would be to distribute webhook notifications for high volume watches. The complexity of implementing HA is estimated to be too much for the current needs of this project.

Given these assumptions we've chosen to focus on Watchman's vertical scaling (add CPUs and memory) instead of clustering. Currently two instances of Watchman will write over each other and send duplicate webhook notifications.

## Database dependency

When Watchman is connected to a MySQL instance the database will need to be deployed in an acceptable manner for replication, failure recovery, and backups.

SQLite replication (possibly implemented via [rqlite](https://github.com/rqlite/rqlite)) has not been tested, but looks promising.
