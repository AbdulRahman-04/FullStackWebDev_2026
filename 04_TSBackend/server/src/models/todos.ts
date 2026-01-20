import mongoose, {Document, Schema, Model} from "mongoose";

interface ITodo extends Document {
    date: string,
    todono: number,
    todoTtitle: string,
    todoDescription: string,
    fileUpload: string
}

const toDoSchema = new Schema<ITodo>({
    date: {
        type: String,
        required: true
    },
    todono: {
        type: Number,
        required: true
    },
    todoTtitle: {
        type: String,
        required: true
    },
    todoDescription: {
        type: String,
        required: true
    },
    fileUpload: {
        type: String
    }
}, {
    timestamps: true
})

const todoModel:Model<ITodo> =  mongoose.model<ITodo>("todos", toDoSchema, "todos")

export {todoModel, ITodo}