import api from "./api";

const create = payload => {
  return api.post("substations", payload);
};

const list = () => {
  return api.get("substation/list");
};

const update = substation => {
  api.put("substation", substation)
}

export default { list, create, update };
