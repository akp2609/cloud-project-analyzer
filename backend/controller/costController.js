import {processCostCSV} from "../services/costService.js";

const uploadCostController = async (req,res)=>{
    try{
        if(!req.files || !req.files.file){
            return res.status(400).json({error:"No CSV file uploaded"});
        }

        const csvData = req.files.file.data.toString();
        const processedData = await processCostCSV(csvData);
        return res.json(
            {message:"Cost data processed successfully",
            summary:processedData});
    }catch(err){
        console.error("Controller error",err.message);
        return res.status(500).json({error:err.message});
    }
}

export {uploadCostController};