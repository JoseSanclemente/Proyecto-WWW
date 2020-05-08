import api from "./api";

const list = () => {
  return api.get("Transformer");
};

export default { list };
