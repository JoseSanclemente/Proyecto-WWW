import user from "@/services/user.js";

// initial state
const state = {
    users: [],
    userSaved: false, //This is not used
};

// getters
const getters = {};

// actions
const actions = {
    listUsers({ commit }) {
        user.list().then(response => {
            commit("setUsers", response.data);
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
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
};
