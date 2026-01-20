"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const mongoose_1 = __importDefault(require("mongoose"));
const config_1 = __importDefault(require("config"));
const DB_URL = config_1.default.get("DB_URL");
async function DbConnect() {
    try {
        await mongoose_1.default.connect(DB_URL);
        console.log("db connected✅");
    }
    catch (error) {
        console.log(error);
    }
}
exports.default = DbConnect();
