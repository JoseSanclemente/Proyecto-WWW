import api from "./api";

const list = () => {
  return api.get("Subestation");
};

export default { list };
