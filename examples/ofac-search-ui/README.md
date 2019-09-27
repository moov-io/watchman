## examples / ofac-search-ui

This project provides an interface to retrieve results from an OFAC service search endpoint.

### Available Scripts

#### `npm start`

Runs the app in the development mode. Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits. You will also see any lint errors in the console.

#### `npm run build`

Builds the application, producing bundles of the frontend assets, making the application ready to serve in a production environment.

#### `npm run server`

Serves the bundled application and proxies api requests to the OFAC search endpoint.

#### `npm run ofac-server`

Runs the latest docker image of the OFAC server as recommended in the main README. This is a convenience script and is not required if the server is already running.

### Configuration

| Environmental Variable | Description                                    | Default                 |
| ---------------------- | ---------------------------------------------- | ----------------------- |
| `OFAC_ENDPOINT`        | URL of the OFAC server to use for queries.     | `http://localhost:8084` |
| `HTTP_BIND_ADDRESS`    | Address for Search UI to bind its HTTP server. | `3000`                  |

### Learn More

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app) and the available scripts were modified. You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).
