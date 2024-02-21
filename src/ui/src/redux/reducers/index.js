import { combineReducers } from "redux";
import { authReducer } from "./auth.js";
import { modalReducer } from "./modal.js";

export const rootReducer = combineReducers({
    auth: authReducer,
    modal: modalReducer
})