const JWT_TOKEN_ID = "jwt"

export const saveToken = token => {
    localStorage.setItem(JWT_TOKEN_ID, token)
}

export const getToken = () => {
    return localStorage.getItem(JWT_TOKEN_ID)
}

export const destroyToken = () => {
    localStorage.removeItem(JWT_TOKEN_ID)
}