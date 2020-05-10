import users from "@/services/user.js";

// initial state
const state = {
    users: [],
    userSaved: false,
};

// getters
const getters = {};

// actions
const actions = {
    loadAllUsers({ commit }) {
        users.list().then(response => {
            commit("setUsers", response.data);
        });
    },

    createUser({ commit, payload }) {
        users.create(payload).then(() => {
            commit("setUserSaved");
        });
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
