const JWT_TOKEN_ID = "jwt";

const saveToken = token => {
  localStorage.setItem(JWT_TOKEN_ID, token);
};

const getToken = () => {
  return localStorage.getItem(JWT_TOKEN_ID);
};

const destroyToken = () => {
  localStorage.removeItem(JWT_TOKEN_ID);
};

export default { saveToken, getToken, destroyToken };
