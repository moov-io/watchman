---
layout: page
title: Performance
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Watchman Performance Characteristics

> For documentation on older releases of Watchman (v0.31.x series), please visit the [older docs website](https://github.com/moov-io/watchman/tree/v0.31.3/docs) in our GitHub repository.

Watchman is designed to deliver fast, reliable sanctions and watchlist screening for financial services, balancing compliance needs with performance demands. By leveraging several key optimizations,
Watchman ensures stable query times and efficient resource usage, even under heavy load. These performance traits make it a robust choice for production environments where speed and accuracy are critical.

One of Watchman’s core strengths is its **precomputation and normalization** of data. On startup, Watchman downloads and processes sanctions lists (like US OFAC and US/UK/EU Consolidated Screening)
and applies transformations—such as reordering names (e.g., "MADURO MOROS, Nicolas" to "Nicolas MADURO MOROS") — to provide standardized results. Wachman also uses the [libpostal](https://github.com/openvenues/libpostal)
library to parse and normalize postal addresses, improving match accuracy at the cost of higher memory usage. This upfront work ensures that searches are faster by reducing the need for on-the-fly
processing, though it does introduce some memory overhead due to `libpostal`’s requirements.

> The libpostal library can take 2GB of memory to run. Make sure Watchman has enough memory to support your load ontop of libpostal's requirements.

Watchman operates entirely with **in-memory lists**, storing all sanction data in memory without disk persistence. This eliminates I/O bottlenecks, enabling rapid search operations.
The trade-off is that data is reloaded on restart, but this ensures freshness and avoids stale data slowing down queries. Combined with a high-performance search implementation using the
Jaro-Winkler algorithm, Watchman delivers quick and accurate fuzzy matching for names and addresses, with scoring from 0.0 (no match) to 1.0 (exact match).
The in-memory approach, paired with precomputed indexes, allows Watchman to handle large query volumes without relying on an external database.

To manage high throughput, Watchman employs **dynamic concurrency with goroutines**. It uses a feedback loop to adjust the number of goroutines based on load, ensuring stable response
times even on shared hardware. As shown in the first graph below, which tracks search requests per second (req/s) over time, Watchman maintains consistent query performance despite fluctuating load.

![Graph 1: Stable query times with search req/s over time](../images/stable-response-times.png)

The second graph illustrates how Watchman dynamically adjusts goroutine group sizes, optimizing for overall time to score and keeping response timings steady (typically between 250-1000 ms).
This concurrency model was a significant improvement over earlier versions, where consistent load could cause slowdowns or crashes. The v0.5x series further refined this by consolidating to a single search model,
encouraging richer query data for better performance.

![Graph 2: Dynamic adjustment of goroutine group sizes for optimal scoring time](../images/dynamic-goroutines.png)
