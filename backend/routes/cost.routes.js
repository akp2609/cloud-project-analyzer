import { Router } from "express";
import {uploadCostController} from "../controller/costController.js";

const router = Router();

router.post("/upload-cost",uploadCostController);

export default router;
