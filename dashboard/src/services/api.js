import axios from "axios";
//import jwt from "./jwt";

const BASE_URL = "https://morning-brook-06940.herokuapp.com";
// todo: add authentication information
const intercept = payload => {
  payload = {
    ...payload,
    headers: {
      ...payload.headers
      //"Bearer-Token": jwt.getToken()
    },
    url: `${BASE_URL}/${payload.url}`
  };
  console.log("ℹ️ Axios payload", payload)
  return axios(payload);
};

const get = url => {
  return intercept({ url: url, method: "get" });
};

const del = url => {
  return intercept({ method: "delete", url: url });
};

const put = (url, payload) => {
  return intercept({ method: "put", data: payload, url: url })
}

const post = (url, payload) => {
  return intercept({ url: url, data: payload, method: "post" });
};

export default { get, del, post, put };
