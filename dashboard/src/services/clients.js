import api from "./api";

const create = payload => {
  return api.post("client", payload);
};

const list = () => {
  return api.get("client");
};

export default { list, create };
