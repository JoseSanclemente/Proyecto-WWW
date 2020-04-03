import clients from "@/core/services/clients.js";

// initial state
const state = {
  clients: []
};

// getters
const getters = {};

// actions
const actions = {
  loadAllClients({ commit }) {
    clients.list().then(response => {
      commit("setClients", response.data);
    });
  }
};

// mutations
const mutations = {
  setClients(state, clients) {
    state.clients = clients;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
