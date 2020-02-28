import * as api from './api'

export const create = user => {
    return api.post('user', user)
}

// todo: add pagination
export const list = (page=0) => {
    return Api.get('user')
}

export const del = id =>  {
    return Api.del(`user/${id}`)
}

export const  get = id => {
    return api.get(`user/${id}`)
}