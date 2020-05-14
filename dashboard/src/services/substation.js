import api from "./api";

const create = payload => {
  return api.post("substations", payload);
};

const list = () => {
  return api.get("substation/list");
};

export default { list, create };
