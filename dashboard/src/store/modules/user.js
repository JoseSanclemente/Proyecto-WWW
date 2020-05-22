import user from "@/services/user.js";

// initial state
const state = {
    users: [],
    userSaved: false, //This is not used
    searched: [],
};

// getters
const getters = {};

// actions
const actions = {
    listUsers({ commit }) {
        user.list().then(response => {
            commit("setUsers", response.data);
            commit("setSearched");
        });
    },

    createUser({ commit, dispatch }, payload) {
        commit("setUserSaved");
        user.create(payload).then(() => {
            dispatch("listUsers")
        }).catch(error => {
            return error
        })

    },

    updateUser({ commit, dispatch }, payload) {
        commit("setUserSaved");
        payload.deleted = payload.deleted.toString()
        user.update(payload).then(() => {
            dispatch("listUsers")
        }).catch(error => {
            return error
        })
        return
    },
    sendLoginData({ commit }, payload) {
        commit("setValidatedLogin");
        let us = user.login(payload);
        console.log(us)
        return us;
    }
};

// mutations
const mutations = {
    setUsers(state, users) {
        state.users = users;
    },
    setUserSaved(state) {
        state.userSaved = true;
    },
    resetUserState(state) {
        state.userSaved = false;
    },
    setSearched(state) {
        state.searched = state.users
    },
    searchUser(state, email) {
        if (email) {
            state.searched = state.users.filter(item => item.email.toLowerCase().includes(email.toLowerCase()))
        } else {
            state.searched = state.users
        }
    },
    setValidatedLogin(state) {
        state.validated = true;
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
};
