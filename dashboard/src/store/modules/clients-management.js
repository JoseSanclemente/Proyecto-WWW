import clients from "@/services/clients.js";
import contracts from "@/services/contracts.js";

// initial state
const state = {
  clients: [],
  contracts: []
};

// getters
const getters = {};

// actions
const actions = {
  loadAllClients({ commit }) {
    clients.list().then(response => {
      commit("setClients", response.data);
    });
  },

  loadAllContracts({ commit }) {
    contracts.list().then(response => {
      commit("setContracts", response.data);
    });
  }
};

// mutations
const mutations = {
  setClients(state, clients) {
    state.clients = clients;
  },

  setContracts(state, contracts) {
    state.contracts = contracts;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
