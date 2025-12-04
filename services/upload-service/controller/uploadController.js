import {Storage} from "@google-cloud/storage";
import {PubSub} from "@google-cloud/pubsub";

const storage = new Storage({
  keyFilename: process.env.GOOGLE_APPLICATION_CREDENTIALS,
  projectId: "cloud-cost-resource-analyzer"
});
const pubsub = new PubSub({
  keyFilename: process.env.GOOGLE_APPLICATION_CREDENTIALS,
  projectId: "cloud-cost-resource-analyzer"
});

const RAW_BUCKET = process.env.RAW_BUCKET;
const CSV_TOPIC = process.env.CSV_TOPIC || "";
const port = process.env.PORT || 8080;

const uploadFile = async (req, res) => {
    try{
        if(!req.files || !req.files.file){
            return res.status(400).json({error: "No file uploaded"});
        }

        console.log("Bucket name", RAW_BUCKET);
        console.log("CSV_TOPIC:", process.env.CSV_TOPIC);
        console.log("Using credentials:", process.env.GOOGLE_APPLICATION_CREDENTIALS);


        const tenant = String(req.headers["x-tenant-id"] ?? "unknown").replace(/[^a-zA-Z0-9-_]/g,"_");

        const file = req.files.file;
        const timestamp = Date.now();
        const sanitizeName = file.name.replace(/\s+/g,"_");
        const gcsPath = `raw/${tenant}/${timestamp}-${sanitizeName}`;

        const bucket = storage.bucket(RAW_BUCKET);
        await bucket.file(gcsPath).save(file.data,{
            contentType: file.mimetype || "text/csv",
            resumable: false,
        })

        let messageId = null;
        if(CSV_TOPIC){
            const topic = pubsub.topic(CSV_TOPIC);
            const msg = {bucket: RAW_BUCKET, path: gcsPath, tenant, uploadAt: timestamp};
            messageId = await topic.publishJSON(msg);
        }

        return res.status(201).json({message: "uploaded",gcsPath,messageId});

    }
    catch(err){
        console.error("upload error",err);
        return res.status(500).json({error: "Internal server error during upload",detail: String(err.message || err)});
    }
}

export default uploadFile;