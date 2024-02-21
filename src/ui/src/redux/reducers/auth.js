import { Auth } from "../types/auth";

let user = JSON.parse(localStorage.getItem("user"));
const initialState = user
    ? { loggedIn: true, user, loading: false, error: null }
    : { loggedIn: false, loading: false, error: null };

export function authReducer(state = initialState, action) {
    switch (action.type) {
        case Auth.FETCH:
            return {
                loading: true,
            };
        case Auth.FETCH_SUCCESS:
            return {
                loggedIn: true,
                user: action.payload,
                loading: false,
                error: null,
            };
        case Auth.FETCH_REGISTER:
            return {
                loggedIn: false
            };
        case Auth.FETCH_ERROR:
            return { loading: false, error: action.payload };
        case Auth.LOGOUT:
            return {};
        case Auth.SET_ERROR:
            return { ...state, error: action.payload };
        default:
            return state;
    }
}
