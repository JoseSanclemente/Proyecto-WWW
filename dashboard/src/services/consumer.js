import api from "./api";

const create = payload => {
  return api.post("consumer", payload);
};

const list = () => {
  return api.get("consumer/list");
};

const captcha = () => {
  return api.get("consumer/captcha");
}

const login = payload => {
  return api.post("consumer", payload);
}

export default { list, create };
