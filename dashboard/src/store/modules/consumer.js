import consumer from "@/services/consumer.js";
import contracts from "@/services/contracts.js";

// initial state
const state = {
  consumers: [],
  contracts: []
};

// getters
const getters = {};

// actions
const actions = {
  listConsumers({ commit }) {
    consumer.list().then(response => {
      commit("setClients", response.data);
    });
  },

  createConsumer({ commit }, payload) {
    commit("setConsumerSaved");
    return consumer.create(payload)
  },

  loadAllContracts({ commit }) {
    contracts.list().then(response => {
      commit("setContracts", response.data);
    });
  }
};

// mutations
const mutations = {
  setClients(state, consumers) {
    state.consumers = consumers;
  },

  setContracts(state, contracts) {
    state.contracts = contracts;
  },

  setConsumerSaved(state) {
    state.userSaved = true;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
