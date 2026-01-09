import express from "express";
import dotenv from "dotenv";
import {metricsHandler, metricsMiddleware} from "./metrics.js";
import fileUpload from "express-fileupload";
import uploadFile from "./controller/uploadController.js";

const app = express();

dotenv.config();

app.use(
    fileUpload({
        limits: {fileSize: 50 * 1024 * 1024},
        abortOnLimit: true,
        useTempFiles: true,
    })
)

app.use(metricsMiddleware);

const RAW_BUCKET = process.env.RAW_BUCKET;
const CSV_TOPIC = process.env.CSV_TOPIC || "";
const port = process.env.PORT || 8080;

if(!RAW_BUCKET){
    console.log("FATAL: RAW_BUCKET is not set. Set RAW_BUCKET env var and retry");
}

app.get("/healthz", (_req,res)=> res.status(200).send("OK"));
app.get("/metrics", metricsHandler);

app.post("/upload",uploadFile);

app.listen(port, ()=>{
        console.log(`Upload service listening on port ${port}`);
        console.log(`RAW_BUCKET: ${RAW_BUCKET} CSV_TOPIC: ${CSV_TOPIC || "(not set"}`);
});


export default app;

