---
layout: page
title: Model Context Protocol (MCP)
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Model Context Protocol (MCP) Server

> **Experimental Feature**: The MCP server is an experimental integration that may change in future releases.

Watchman provides an optional [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server implementation, allowing AI assistants and language models to directly interact with sanctions screening functionality.

## Overview

The MCP server enables seamless integration with MCP-compatible clients, providing AI agents with the ability to perform sanctions screening as part of their reasoning and decision-making processes.

## Configuration

To enable the MCP server, add the following to your Watchman configuration:

```yaml
mcp:
  enabled: true
```

When MCP is enabled, Watchman will serve MCP endpoints at `/mcp` in addition to the standard HTTP API.

## Available Tools

### search_entities

Searches for entities in sanctions lists using the same powerful matching algorithms as the HTTP API.

#### Parameters

- **request** (string, required): JSON string representing the search request, identical to the `/v2/search` HTTP endpoint
- **limit** (number, optional): Maximum number of results to return (default: 10)
- **minMatch** (number, optional): Minimum match score threshold (default: 0.0)

#### Request Format

The `request` parameter accepts the same JSON structure as the HTTP `/v2/search` endpoint:

```json
{
  "name": "Nicolas Maduro",
  "type": "person",
  "birthDate": "1962-11-23"
}
```

#### Response Format

Returns a JSON object with search results:

```json
{
  "query": {
    "name": "Nicolas Maduro",
    "type": "person",
    "birthDate": "1962-11-23"
  },
  "entities": [
    {
      "name": "Nicolas MADURO MOROS",
      "entityType": "person",
      "sourceList": "us_ofac",
      "sourceID": "OFAC-12345",
      "person": {
        "name": "Nicolas MADURO MOROS",
        "birthDate": "1962-11-23",
        "placeOfBirth": "Caracas, Venezuela"
      },
      "sanctionsInfo": {
        "programs": ["SDNT"],
        "description": "President of Venezuela"
      },
      "match": 0.95
    }
  ]
}
```

## Usage Examples

### Basic Person Search

```json
{
  "request": "{\"name\": \"Nicolas Maduro\", \"type\": \"person\"}",
  "limit": 5,
  "minMatch": 0.8
}
```

### Business Entity Search

```json
{
  "request": "{\"name\": \"Rosneft\", \"type\": \"business\"}",
  "limit": 3
}
```

### Advanced Search with Multiple Criteria

```json
{
  "request": "{\"name\": \"Vladimir Putin\", \"type\": \"person\", \"birthDate\": \"1952-10-07\", \"addresses\": [{\"country\": \"RU\"}]}",
  "minMatch": 0.9
}
```

## Supported Entity Types

The MCP server supports the same entity types as the HTTP API:

- **person**: Individual people
- **business**: Companies and organizations
- **organization**: Government and non-profit organizations
- **aircraft**: Aircraft with sanctions
- **vessel**: Ships with sanctions

## Search Parameters

All search parameters from the HTTP `/v2/search` endpoint are supported:

- `name`: Primary name to search for
- `altNames`: Alternative names
- `type`: Entity type (person, business, organization, aircraft, vessel)
- `gender`: Gender (male, female, unknown)
- `birthDate`: Birth date (YYYY-MM-DD format)
- `deathDate`: Death date
- `created`: Business/organization creation date
- `dissolved`: Business/organization dissolution date
- `addresses`: Array of address objects
- `phoneNumbers`: Phone numbers
- `emailAddresses`: Email addresses
- `cryptoAddresses`: Cryptocurrency addresses
- `governmentIDs`: Government-issued IDs (passports, tax IDs, etc.)

## Usage Examples

### Go Client Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/modelcontextprotocol/go-sdk/mcp"
    "github.com/moov-io/watchman/pkg/search"
)

func main() {
    ctx := context.Background()

    // Create MCP client
    client := mcp.NewClient(&mcp.Implementation{
        Name:    "watchman-client",
        Version: "1.0.0",
    }, nil)

    // Connect to Watchman MCP server
    session, err := client.Connect(ctx, &mcp.StreamableClientTransport{
        Endpoint: "http://localhost:8084/mcp",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer session.Close()

    // Search for an entity
    result, err := session.CallTool(ctx, &mcp.CallToolParams{
        Name: "search_entities",
        Arguments: map[string]any{
            "request": search.Entity[search.Value]{
                Name: "Nicolas Maduro",
                Type: search.EntityPerson,
            },
            "limit": 5,
            "minMatch": 0.8,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    // Print results
    for _, content := range result.Content {
        if textContent, ok := content.(*mcp.TextContent); ok {
            fmt.Println(textContent.Text)
        }
    }
}
```

### cURL Examples

Since the MCP server is configured in stateless mode, you can directly call tools without session initialization:

```bash
# Call the search_entities tool directly
curl -X POST "http://localhost:8084/mcp" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json,text/event-stream" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
      "name": "search_entities",
      "arguments": {
        "request": {
          "name": "Nicolas Maduro",
          "entityType": "person",
          "sourceList": "api-request",
          "sourceID": "test",
          "contact": {},
          "addresses": [],
          "cryptoAddresses": [],
          "affiliations": [],
          "historicalInfo": [],
          "sourceData": null
        },
        "limit": 5,
        "minMatch": 0.8
      }
    }
  }'
```

**Note**: The `request` object must include all required fields of the Entity struct. For production use, it's recommended to use an MCP client library that handles the protocol details automatically.

## MCP Client Integration

To use Watchman with an MCP client:

1. Configure Watchman with MCP enabled
2. Start Watchman with the standard HTTP server
3. Configure your MCP client to connect to `http://your-watchman-server/mcp`

## Limitations

- Currently only supports the `search_entities` tool
- Experimental feature that may change without notice

## Future Enhancements

Planned improvements include:

- Additional tools for list management and statistics
- Enhanced error handling and validation
- Support for streaming responses
- Integration with more MCP features

## Need Help?

If you encounter issues with the MCP server or have questions:

- Check the [GitHub Issues](https://github.com/moov-io/watchman/issues) for known problems
- Ask in the [#watchman channel](https://slack.moov.io/) on Slack
- Review the [HTTP API documentation](/watchman/search/) for search parameter details
