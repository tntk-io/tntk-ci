import { Auth } from "../types/auth";
import { history } from "../../helpers/api";
import axios from "axios";
import { Buffer } from "buffer";

export const userActions = {
  login,
  register,
  logout,
  error,
  setError,
};

function login(name, password) {
    let auth = 'Basic ' + Buffer.from(name + ':' + password).toString('base64');
    const requestOptions = {
        headers: {
            'Authorization': auth
        }
    }

  return async (dispatch) => {
    try {
      dispatch({ type: Auth.FETCH });
      const response = await axios.post(
        "api/v1/auth/sign-in",
        {},
        requestOptions
      );

      localStorage.setItem("user", JSON.stringify(response.data));

      dispatch({
        type: Auth.FETCH_SUCCESS,
        payload: response.data,
      });
      history.push("/");
    } catch (e) {
      dispatch({ type: Auth.FETCH_ERROR, payload: e.message });
    }
  };
}

function register(name, password) {
    
  return async (dispatch) => {
    try {
      dispatch({ type: Auth.FETCH });
      const response = await axios.post(
        "api/v1/auth/sign-up",
        {name, password}
      );

      dispatch({
        type: Auth.FETCH_REGISTER,
        payload: response.data,
      });
      history.push("/login");
    } catch (e) {
      dispatch({ type: Auth.FETCH_ERROR, payload: e.response.data });
    }
  };
}

function logout() {
  localStorage.removeItem("user");
  return { type: Auth.LOGOUT };
}

function error(payload) {
  return { type: Auth.FETCH_ERROR, payload };
}

function setError(payload) {
  return { type: Auth.SET_ERROR, payload };
}
