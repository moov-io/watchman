---
layout: page
title: Model Context Protocol (MCP)
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Model Context Protocol (MCP) Server

> **Experimental Feature**: The MCP server is an experimental integration that may change in future releases.

Watchman provides an HTTP `/mcp` endpoint with full Model Context Protocol (MCP) support, allowing agents to search the loaded lists.
The endpoint is protected by **[MCPS](https://datatracker.ietf.org/doc/draft-sharif-mcps-secure-mcp/)**, which cryptographically verifies
agent chains and calls at the point of screening. Untrusted requests are automatically blocked, and every request is proven tamper-resistant
through cryptographic attestation.

Only agents that can prove their identity are permitted to screen entities, verify transactions, or initiate payments.

Check out the **[live demo](https://showcase.cybersecai.co.uk/demo.html#live-check)** to see seamless agent onboarding, verification, and screening in action.

## Overview

The MCP server enables seamless integration with MCP-compatible clients, providing AI agents with the ability to perform sanctions screening as part of their reasoning and decision-making processes.

## Configuration

To enable the MCP server, add the following to your Watchman configuration:

```yaml
Watchman:
  MCP:
    Enabled: true
    Signing:
  	  Enabled: true
	  KeyPath: "/path/to/private-key.pem"
	  PubPath: "/path/to/public-key.pem"
    AgentPass:
	  Enabled: true
	  TrustAnchorPath: "/path/to/ca.pem"
      # MinTrustLevel rejects agents whose trust level is below level.
      # Level must be in the range 0-4; values outside this range are
      # clamped to the nearest bound. Callers typically set this to 2
      # in production to reject unverified (L0) and developer-sandbox
      # (L1) agents.
	  MinTrustLevel: 0
      # RequiredScopes rejects agents that lack any of the listed
      # scopes. Useful for Watchman-style integrations that want to
      # demand e.g. WithRequiredScopes("sanctions:search") before
      # accepting a screening request.
	  RequiredScopes: ["sanctions:search"]
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
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
        "name": "search_entities",
        "arguments": {
            "request": {
                "name": "john",
                "entityType": "person"
            },
            "limit": 1,
            "minMatch": 0.25
        }
    }
}
```

#### Response Format

Returns a JSON object with search results:

```json
{
    "query": {
        "name": "John",
        "entityType": "person",
        "sourceList": "",
        "sourceID": "",
        "person": {
            "name": "John",
            "altNames": null,
            "gender": "",
            "birthDate": null,
            "placeOfBirth": "",
            "deathDate": null,
            "titles": null,
            "governmentIDs": null
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
        "sourceData": null
    },
    "entities": [
        {
            "name": "John NUMBI",
            "entityType": "person",
            "sourceList": "us_ofac",
            "sourceID": "20420",
            "person": {
                "name": "John NUMBI",
                "altNames": null,
                "gender": "male",
                "birthDate": "1957-01-01T00:00:00Z",
                "placeOfBirth": "",
                "deathDate": null,
                "titles": [
                    "General; Former National Inspector, Congolese National Police"
                ],
                "governmentIDs": null
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
                "entityID": "20420",
                "sdnName": "NUMBI, John",
                "sdnType": "individual",
                "program": [
                    "DRCONGO"
                ],
                "title": "General; Former National Inspector, Congolese National Police",
                "callSign": "",
                "vesselType": "",
                "tonnage": "",
                "grossRegisteredTonnage": "",
                "vesselFlag": "",
                "vesselOwner": "",
                "remarks": "DOB 1957; POB Kolwezi, Katanga Province, Democratic Republic of the Congo; Gender Male; General; Former National Inspector, Congolese National Police."
            },
            "match": 0.7445624999999998
        }
    ]
}
```

## Usage Examples

### Basic Person Search

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
        "name": "search_entities",
        "arguments": {
            "request": {
                "name": "Nicholas Maduro",
                "entityType": "person"
            },
            "limit": 1,
            "minMatch": 0.25
        }
    }
}
```

### Business Entity Search

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
        "name": "search_entities",
        "arguments": {
            "request": {
                "name": "Rosneft",
                "entityType": "business"
            },
            "limit": 1,
            "minMatch": 0.25
        }
    }
}
```

### Advanced Search with Multiple Criteria

```json
{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
        "name": "search_entities",
        "arguments": {
            "request": {
                "name": "Vladimir Putin",
                "entityType": "person",
                "person": {
                    "birthDate": "1952-10-07"
                },
                "addresses": [
                    {
                        "country": "RU"
                    }
                ]
            },
            "limit": 1,
            "minMatch": 0.25
        }
    }
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
				Name: "John",
				Type: search.EntityPerson,
			},
			"limit":    1,
			"minMatch": 0.25,
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
curl -s -X POST "http://localhost:8084/mcp" \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/call",
    "params": {
        "name": "search_entities",
        "arguments": {
            "request": {
                "name": "john",
                "entityType": "person"
            },
            "limit": 1,
            "minMatch": 0.25
        }
    }
}'
```

**Note**: The `request` object must include all required fields of the Entity struct. For production use, it's recommended to use an MCP client library that handles the protocol details automatically.

## Client Integration

To use Watchman with an MCP client:

1. Configure Watchman with MCP enabled
2. Start Watchman with the standard HTTP server
3. Configure your MCP client to connect to `http://your-watchman-server/mcp`

## Need Help?

If you encounter issues with the MCP server or have questions:

- Check the [GitHub Issues](https://github.com/moov-io/watchman/issues) for known problems
- Ask in the [#watchman channel](https://slack.moov.io/) on Slack
- Review the [HTTP API documentation](/watchman/search/) for search parameter details
