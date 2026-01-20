import express , {Application , Request, Response} from "express"
import config from "config"
// db connect
import "./utils/dbConnect"

// user router
import userRouter from "./controllers/public/users"
//auth middleware
import authMiddleware from "./middleware/auth"
// todo router
import todoRouter from "./controllers/private/todos"

const app:Application = express()

app.use(express.json())

const PORT : string = config.get<string>("PORT") || "4404"


app.get("/", (req: Request, res:Response)=>{
    try {

        res.status(200).json({msg: "hello world"})
        
    } catch (error) {
        res.status(404).json(error)
    }
})


// public apis 
app.use("/api/public", userRouter)

// private apis 
app.use("/api/private", authMiddleware, todoRouter )



app.listen(Number(PORT), ()=>{
    console.log(`server running at port ${PORT}`);
    
})