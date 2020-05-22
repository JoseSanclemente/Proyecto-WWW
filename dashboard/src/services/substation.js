import api from "./api";

const create = payload => {
  return api.post("substations", payload);
};

const list = () => {
  return api.get("substation/list");
};

const update = payload => {
  return api.put("substation", payload)
}

export default { list, create, update };
