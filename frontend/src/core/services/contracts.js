import api from "./api";

const list = (page = 0) => {
  return api.get("contract");
};

export default { list };
