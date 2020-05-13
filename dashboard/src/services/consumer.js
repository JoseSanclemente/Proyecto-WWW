import api from "./api";

const create = payload => {
  return api.post("consumer", payload);
};

const list = () => {
  return api.get("consumer/list");
};

export default { list, create };
