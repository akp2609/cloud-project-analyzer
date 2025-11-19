import {uploadJSON} from "../repositories/gcsRepository.js";

const saveCostToStorage = async(data)=>{
    const filename = `cost-data/${Date.now()}-cost.json`;
    return await uploadJSON(filename,data);
}

export {saveCostToStorage} 