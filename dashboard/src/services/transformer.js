import api from "./api";

const create = payload => {
  return api.post("transformer", payload);
};

const list = () => {
  return api.get("transformer");
};

export default { list, create };
