import api from "./api";

const create = payload => {
  return api.post("transformer", payload);
};

const list = () => {
  return api.get("transformer/list");
};

const update = payload => {
  api.put("transformer", payload)
}

export default { list, create, update };
