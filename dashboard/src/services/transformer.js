import api from "./api";

const create = payload => {
  return api.post("transformer", payload);
};

const list = () => {
  return api.get("transformer/list");
};

export default { list, create };
