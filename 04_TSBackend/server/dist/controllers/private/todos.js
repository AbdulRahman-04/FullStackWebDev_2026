"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const todos_1 = require("../../models/todos");
const router = express_1.default.Router();
// create todo
router.post("/createtodo", async (req, res) => {
    try {
        const { date, todono, todoTitle, todoDescription, fileUpload } = req.body;
        if (!date || !todono || !todoTitle || !todoDescription) {
            res.status(400).json({ msg: "fill all fields" });
            return;
        }
        const newTodo = new todos_1.todoModel({
            date, todono, todoTitle, todoDescription
        });
        await newTodo.save();
        res.status(200).json({ msg: "todo created", newTodo });
    }
    catch (error) {
        res.status(500).json({ msg: error });
    }
});
// get all todos 
router.get("/getall", async (req, res) => {
    try {
        const todos = await todos_1.todoModel.find({});
        if (!todos) {
            res.status(500).json({ msg: "no todos found" });
            return;
        }
        res.status(200).json({ msg: "todos are here✅", todos });
    }
    catch (error) {
        res.status(500).json({ msg: error });
    }
});
// get one todo
router.get("/getone/:id", async (req, res) => {
    try {
        const todo = await todos_1.todoModel.findById(req.params.id);
        if (!todo) {
            res.status(500).json({ msg: "no todo found" });
            return;
        }
        res.status(200).json({ msg: "one todo is here", todo });
    }
    catch (error) {
        res.status(500).json({ msg: error });
    }
});
// edit one todo
router.put("/edit/:id", async (req, res) => {
    try {
        const { date, todono, todoTitle, todoDescription } = req.body;
        const updateTodo = await todos_1.todoModel.findByIdAndUpdate(req.params.id, { date, todono, todoTitle, todoDescription }, { new: true });
        if (!updateTodo) {
            res.status(500).json({ msg: "no todo found" });
            return;
        }
        res.status(200).json({ msg: "todod updated✅", updateTodo });
    }
    catch (error) {
        res.status(500).json({ msg: error });
    }
});
router.delete("/delete/:id", async (req, res) => {
    try {
        const deleteTodo = await todos_1.todoModel.findByIdAndDelete(req.params.id);
        if (!deleteTodo) {
            res.status(500).json({ msg: "no todo found" });
            return;
        }
        res.status(200).json({ msg: "todo deleted✅" });
    }
    catch (error) {
        res.status(500).json({ msg: error });
    }
});
router.delete("/deleteall", async (req, res) => {
    try {
        await todos_1.todoModel.deleteMany({});
        res.status(200).json({ msg: "all todos deleted" });
    }
    catch (error) {
        res.status(500).json({ msg: error });
    }
});
exports.default = router;
