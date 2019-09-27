## webui

This project provides an interface to retrieve results from an OFAC service search endpoint.

### Available Scripts

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app) and the available scripts were modified.

#### `npm start`

Runs the app in the development mode. Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits. You will also see any lint errors in the console.

#### `npm run build`

Builds the application, producing bundles of the frontend assets, making the application ready to serve in a production environment.

#### `npm run server`

Serves the bundled application and proxies api requests to the OFAC search endpoint.

By default, the server will start on port `3000`. Setting the `HTTP_BIND_ADDRESS` environment variable will override this value.

#### `npm run ofac-server`

Runs the latest docker image of the OFAC server as recommended in the main README. This is a convenience script and is not required if the server is already running.

### Connecting the UI to the OFAC search endpoint

By default, the api proxy will forward requests to [http://localhost:8084](http://localhost:8084). Setting the `OFAC_ENDPOINT` environment variable will override this value.

### Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).
