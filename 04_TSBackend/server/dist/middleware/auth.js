"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
const config_1 = __importDefault(require("config"));
const KEY = config_1.default.get("JWT_KEY");
const authMiddleware = (req, res, next) => {
    const authHeader = req.headers["authorization"];
    if (!authHeader) {
        res.status(401).json({ msg: "no token provided" });
        return;
    }
    const token = authHeader.split(" ")[1];
    try {
        const decoded = jsonwebtoken_1.default.verify(token, KEY);
        req.user = decoded;
        next();
    }
    catch (error) {
        console.error("JWT verification failed:", error);
        res.status(403).json({ msg: "Invalid or expired token ❌" });
    }
};
exports.default = authMiddleware;
