import {Storage} from "@google-cloud/storage";

const storage = new Storage();
const bucketName = process.env.BUCKET_NAME;

const uploadJSON = async(path,jsonData)=>{
    const bucket = storage.bucket(bucketName);
    const file = bucket.file(path);

    await file.save(JSON.stringify(jsonData,null,2),{
        contentType:"application/json",
    });

    return true;
}

export {uploadJSON};