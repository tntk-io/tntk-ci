import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import { Provider } from 'react-redux';

import App from './App';
import reportWebVitals from "./reportWebVitals";
import { store } from './redux/store';

const host = window.location.hostname;
axios.defaults.baseURL = "https://" + host
// axios.defaults.baseURL = process.env.REACT_APP_API_URL;


ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById("root")
);

reportWebVitals();