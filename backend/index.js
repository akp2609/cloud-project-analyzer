import express from "express";
import cors from "cors";
import dotenv from "dotenv";
import costRoutes from "../routes/cost.routes.js";


dotenv.config();
const app = express();
app.use(cors());
app.use(express.json());
app.use(fileUpload());
app.use("/cost",costRoutes);

app.get("/",(req,res)=>{
   res.json({message:"backend running"});
});

const PORT = process.env.PORT || 8080;
app.listen(PORT,()=>console.log(`Server running on port ${PORT}`));
