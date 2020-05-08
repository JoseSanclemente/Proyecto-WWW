import api from "./api";

const list = () => {
  return api.get("contract");
};

export default { list };
