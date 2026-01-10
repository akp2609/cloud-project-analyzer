import client from "prom-client";
import onHeaders from "on-headers";

client.collectDefaultMetrics({timeout: 5000});

const requestDuration = new client.Histogram({
    name: "upload_http_request_pduration_ms",
    help: "Duration of HTTP requests in ms - upload service",
    labelNames: ["method","route","status","tenant"],
    buckets: [50,100,200,500,1000,2000]
});

const requestCount = new client.Counter({
  name: "upload_http_requests_total",
  help: "Total HTTP requests - upload service",
  labelNames: ["method","route","status","tenant"]
});

const metricsMiddleware = (req,res,next)=>{
    const start = Date.now();
    const route = req.route?.path ?? req.path ?? "unknown";
    const tenant = req.headers["x-tenant-id"] ?? "unknown";

    onHeaders(res,()=>{
        const duration = Date.now() - start;
        const status = String(res.statusCode);
        requestDuration.labels(req.method,route,status,tenant).observe(duration);
        requestCount.labels(req.method,route,status,tenant).inc();
    });

    return next();
}

const metricsHandler = async (_req, res) => {
  try {
    res.set("Content-Type", client.register.contentType);
    const metrics = await client.register.metrics();  
    res.end(metrics);
  } catch (err) {
    console.error("metrics error", err);
    res.status(500).send("Error generating metrics");
  }
};


export {requestDuration,requestCount,metricsHandler,metricsMiddleware};
