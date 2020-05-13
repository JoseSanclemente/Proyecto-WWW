import api from "./api";

const create = payload => {
  return api.post("substations", payload);
};

const list = () => {
  return api.get("substations");
};

export default { list, create };
