import mongoose, {Schema, model, Document, Model} from "mongoose";

interface IUser extends Document {
    userName: string,
    email: string,
    password: string,
    phone: string,
    age: number,
    userVerified: {
        emailVerified: boolean | null,
        phoneVerified: boolean | null
    },
    userVerifyToken: {
        emailVerifyToken: string | null,
        phoneVerifyToken: string | null
    }

}

const UserSchema = new Schema<IUser>({

    userName: {
        type: String,
        required: true,
         maxlength: 70,
        minlength: 10
    },
    email: {
        type: String,
        required: true,
        unique: true
    },
     password: {
        type: String,
        required: true
    },
    age: {
        type: Number,
        required: true
    },
    phone:{
        type: String,
        required: true,
    },
    userVerified: {
        emailVerified: {
            type: Boolean,
            default: false
        },
        phoneVerified: {
            type: Boolean,
            default: false
        }
    },
    userVerifyToken: {
        emailVerifyToken: {
            type: String,
            default: null
        },
        phoneVerifyToken: {
            type: String,
            default: null
        }
    }

},{
    timestamps: true
})

const userModel: Model<IUser> = mongoose.model<IUser>("users", UserSchema, " users")

export {userModel, IUser}