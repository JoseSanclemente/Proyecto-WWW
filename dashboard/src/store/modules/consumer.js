import consumer from "@/services/consumer.js";

// initial state
const state = {
  consumers: [],
  dummy: false
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
    }).catch(error => {
      return error
    });
  },

  createConsumer({ commit }, payload) {
    commit("dummyMutation");
    return consumer.create(payload)
  },

  listContracts({ commit }, consumerID) {
    consumer.listConsumerContracts(consumerID).then(response => {
      commit("setContracts", response.data);
    }).catch(error => {
      return error
    });
  },

  createContract({ commit }, payload) {
    commit("dummyMutation")
    return consumer.createContract(payload)
  },

  payBill({ commit }, payload) {
    console.log(payload)
    commit("dummyMutation")
    return consumer.payConsumerBill(payload)
  },

  getPDF({ commit }, payload) {
    consumer.getConsumerPDF(payload).then(response => {
      const linkSource = `data:application/pdf;base64,${response.data}`;
      const downloadLink = document.createElement("a");
      const fileName = `bill_${payload}.pdf`;

      downloadLink.href = linkSource;
      downloadLink.download = fileName;
      downloadLink.click();

      commit("dummyMutation")
    }).catch(error => {
      return error
    });
  },

  loadCaptcha({ commit }){
    consumer.captcha().then(response => {
        commit("setCaptcha", response.data)
    });
  },
  
  sendLoginData({ commit }, payload){
    commit("setValidatedLogin");
    return consumer.login(payload);
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

  dummyMutation(state) {
    state.dummy = true
  }
  setConsumerSaved(state) {
    state.userSaved = true;
  },

  setCaptcha(state, captcha){
    state.captcha = captcha;
  },

  setValidatedLogin(state){
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
