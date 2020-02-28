import axios from 'axios'
import * as Auth from './Authentication'
// todo: add authentication information
const intercept = payload => {
    payload = {
        ...payload,
        headers : {
            ...payload.headers,
            'Bearer-Token': Auth.GetAuthentication()
        }
    }
    return axios(payload)
}

export const  get = url => {
    return intercept({url: url, method:'get'})
}

export const del = url => {
    return intercept({method:"delete", url:url})
}

export const post = (url, payload)  => {
    return intercept({url: url, data: payload, method:'post'})
}
