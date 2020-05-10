import api from "./api";

const create = user => {
  return api.post("employee", user);
};

// todo: add pagination
const list = () => {
  return api.get("user");
};

const del = id => {
  return api.del(`user/${id}`);
};

const get = id => {
  return api.get(`user/${id}`);
};

export default { list, del, get, create };
