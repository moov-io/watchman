const path = require("path");
const express = require("express");
const winston = require("winston");
const expressWinston = require("express-winston");
const proxy = require("http-proxy-middleware");

const HTTP_BIND_ADDRESS = process.env.HTTP_BIND_ADDRESS || 3000;
const app = express();

app.use(
  expressWinston.logger({
    transports: [new winston.transports.Console()],
    format: winston.format.json()
  })
);
app.use(
  "/api",
  proxy({
    target: process.env.OFAC_ENDPOINT || "http://localhost:8084",
    pathRewrite: { "^/api": "" },
    changeOrigin: true
  })
);
app.use("/", express.static(path.join(__dirname, "/../build")));
app.get("/*lb-status", (req, res) => res.send("ok"));
app.get("/*", (req, res) => res.sendFile(path.join(__dirname, "/../build/index.html")));

app.listen(HTTP_BIND_ADDRESS, () => console.log(`App listening on ${HTTP_BIND_ADDRESS}`));
