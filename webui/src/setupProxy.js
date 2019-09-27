const proxy = require("http-proxy-middleware");

module.exports = function(app) {
  app.use(
    "/api",
    proxy({
      target: process.env.OFAC_ENDPOINT || "http://localhost:8084",
      pathRewrite: { "^/api": "" },
      changeOrigin: true
    })
  );
};
