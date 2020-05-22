import consumer from "@/services/consumer.js";

// initial state
const state = {
  consumers: [],
  contracts: [],
  dummy: false
};

// getters
const getters = {};

// actions
const actions = {
  listConsumers({ commit }) {
    consumer.list().then(response => {
      commit("setConsumers", response.data);
    }).catch(error => {
      return error
    });
  },

  createConsumer({ commit, dispatch }, payload) {
    commit("dummyMutation");

    consumer.create(payload).then(() => {
      dispatch("listConsumers")
    }).catch(error => {
      return error
    })
  },

  updateConsumer({ commit, dispatch }, payload) {
    commit("dummyMutation");
    payload.deleted = payload.deleted.toString()

    consumer.update(payload).then(() => {
      dispatch("listConsumers")
    }).catch(error => {
      return error
    })
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
      console.log(error.response)
      if (error.response == null) {
        return "not_found"
      }
      return error
    });
  }
};

// mutations
const mutations = {
  setConsumers(state, consumers) {
    state.consumers = consumers;
  },

  setContracts(state, contracts) {
    state.contracts = contracts;
  },

  dummyMutation(state) {
    state.dummy = true
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
