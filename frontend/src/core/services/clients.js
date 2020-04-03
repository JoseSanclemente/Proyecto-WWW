import api from "./api";

const create = payload => {
  return api.post("client", payload);
};

const list = (page = 0) => {
  return api.get("client");
};

export default { list, create };
