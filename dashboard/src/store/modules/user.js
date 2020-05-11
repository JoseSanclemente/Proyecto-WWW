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

    createUser({ commit }, payload) {
        commit("setUserSaved");
        return user.create(payload)
    },

    updateUser({ commit }, payload) {
        payload.deleted = payload.deleted.toString()
        commit("setUserSaved");
        return user.update(payload)
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
