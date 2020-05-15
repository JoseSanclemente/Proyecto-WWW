import consumer from "@/services/consumer.js";
import contracts from "@/services/contracts.js";

// initial state
const state = {
  consumers: [],
  contracts: [],
  captcha: {}
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
  },

  loadCaptcha({ commit }){
    consumer.captcha().then(response => {
        commit("setCaptcha", response.data)
    });
    console.log("AQUI ESTA EL CAPTCHA");
    console.log(state.captcha);
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

  setCaptcha(state, captcha){
    state.captcha = captcha;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
