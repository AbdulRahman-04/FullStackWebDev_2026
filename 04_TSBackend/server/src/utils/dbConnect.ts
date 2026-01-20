import mongoose from "mongoose"
import config from "config"

const DB_URL: string = config.get("DB_URL")

async function DbConnect() {
    try {

        await mongoose.connect(DB_URL)
        console.log("db connected✅");
        
        
    } catch (error) {
        console.log(error);
        
    }
    
}

export default DbConnect()
