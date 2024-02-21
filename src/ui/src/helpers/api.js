import axios from "axios";
import { createBrowserHistory } from "history";

export const history = createBrowserHistory();

export const apiConfig = {
    getFiles: `api/v1/files`,
    createFile: `api/v1/request`,
    getFile: `api/v1/file`,
    deleteFile: `api/v1/file`,
    signIn: `api/v1/auth/sign-in`,
    signUp: `api/v1/auth/sign-up`
};

export const getPdfFiles = () => {
    return new Promise((res, rej) => {
        axios.get(`${apiConfig.getFiles}`)
    })
};

export const addPdfFile = (url, user) => {
    let data = { url: url };
    const requestOptions = {
        headers: {
            'Token': user.token,
            'Content-Type': 'application/json'
        }
    };
    return new Promise((resolve, reject) => {
        axios.post(`${apiConfig.createFile}`, data, requestOptions)
            .then((res) => {
                resolve(res);
            })
            .catch((e) => reject(e));
    })
};
