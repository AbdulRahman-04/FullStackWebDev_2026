"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const config_1 = __importDefault(require("config"));
// db connect
require("./utils/dbConnect");
const app = (0, express_1.default)();
app.use(express_1.default.json());
const PORT = config_1.default.get("PORT") || "4404";
app.get("/", (req, res) => {
    try {
        res.status(200).json({ msg: "hello world" });
    }
    catch (error) {
        res.status(404).json(error);
    }
});
app.listen(Number(PORT), () => {
    console.log(`server running at port ${PORT}`);
});
