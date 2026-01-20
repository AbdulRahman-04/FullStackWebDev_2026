import express, {Request,Response} from "express"
import { todoModel } from "../../models/todos"

const router = express.Router()

// create todo
router.post("/createtodo", async(req: Request, res: Response):Promise<void>=>{
    try {
 
        const {date, todono, todoTitle, todoDescription, fileUpload} = req.body

        if (!date || !todono || !todoTitle || !todoDescription ){
          res.status(400).json({msg: "fill all fields"})
          return
        }

        const newTodo = new todoModel({
            
            date, todono, todoTitle, todoDescription
        
        })

        await newTodo.save()
        res.status(200).json({msg: "todo created", newTodo})

        
    } catch (error) {
        res.status(500).json({msg:error})
    }
})


// get all todos 
router.get("/getall", async (req:Request, res:Response):Promise<void>=>{
    try {

        const todos = await todoModel.find({})
        if(!todos){
            res.status(500).json({msg:"no todos found"})
            return
        }

        res.status(200).json({msg: "todos are here✅", todos})
        
    } catch (error) {
        res.status(500).json({msg:error})
    }
})

// get one todo
router.get("/getone/:id", async(req:Request, res: Response):Promise<void>=>{
    try {
 
        const todo = await todoModel.findById(req.params.id)
        if(!todo){
            res.status(500).json({msg:"no todo found"})
            return
        }

        res.status(200).json({msg: "one todo is here", todo})

        
    } catch (error) {
        res.status(500).json({msg:error})
    }
})

// edit one todo
router.put("/edit/:id", async(req: Request, res: Response):Promise<void>=>{
    try {

        const {date, todono, todoTitle, todoDescription} = req.body

        const updateTodo = await todoModel.findByIdAndUpdate(req.params.id, {date, todono, todoTitle, todoDescription}, {new: true} )
         
        if(!updateTodo){
             res.status(500).json({msg:"no todo found"})
             return
        }

        res.status(200).json({msg: "todod updated✅", updateTodo})

    } catch (error) {
        res.status(500).json({msg:error})
    }
})

router.delete("/delete/:id", async(req: Request, res: Response):Promise<void>=>{
    try {
        const deleteTodo = await todoModel.findByIdAndDelete(req.params.id)
        if(!deleteTodo){
             res.status(500).json({msg:"no todo found"})
             return
        }

         res.status(200).json({msg:"todo deleted✅"})
        
    } catch (error) {
         res.status(500).json({msg:error})
    }
})

router.delete("/deleteall", async(req: Request, res: Response):Promise<void>=>{
    try {
        await todoModel.deleteMany({})
         res.status(200).json({msg:"all todos deleted"})
        
    } catch (error) {
         res.status(500).json({msg:error})
    }
})

export default router