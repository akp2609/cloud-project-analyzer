import {parseCSV} from "../utils/csvParser.js";
import {saveCostToStorage} from "./gcsService.js";

const processCostCSV = async (csvData)=>{
    const records = parseCSV(csvData);
    
    let toatalCost = 0;
    const daily = {};

    for(const row of records){
        const cost = parseFloat(row["Cost"]);
        const date = row["Start Date"];

        toatalCost += cost;

        if(!daily[date]){
            daily[date] = 0;}
        daily[date] += cost;
    
    }

    const result = {toatalCost,daily};

    await saveCostToStorage(result);

    return result;
}

export {processCostCSV}