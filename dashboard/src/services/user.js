import api from "./api";

const create = user => {
  return api.post("employee", user);
};

// todo: add pagination
const list = () => {
  return api.get("employee/list");
};

const update = user => {
  return api.put("employee", user);
};

const get = id => {
  return api.get(`user/${id}`);
};

const login = payload => {
  return api.post("employee/login", payload);
}

export default { list, update, get, create, login };
