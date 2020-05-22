import consumer from "@/services/consumer.js";

// initial state
const state = {
  consumers: [],
  dummy: false,
  contracts: [],
  searched: [],
  captcha: {},
  loggedConsumer: null
};

// getters
const getters = {};

// actions
const actions = {
  listConsumers({ commit }) {
    consumer.list().then(response => {
      commit("setConsumers", response.data);
      commit("setSearched")
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
  },

  loadCaptcha({ commit }) {
    consumer.captcha().then(response => {
      commit("setCaptcha", response.data)
    });
  },

  sendLoginData({ commit }, payload) {
    commit("setValidatedLogin");
    return consumer.login(payload);
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

  setSearched(state) {
    state.searched = state.consumers
  },

  searchConsumer(state, id) {
    if (id) {
      state.searched = state.consumers.filter(item => item.id.toLowerCase().includes(id.toLowerCase()))
    } else {
      state.searched = state.consumers
    }
  },

  dummyMutation(state) {
    state.dummy = true
  },
  setConsumerSaved(state) {
    state.userSaved = true;
  },

  setCaptcha(state, captcha) {
    state.captcha = captcha;
  },

  setValidatedLogin(state) {
    state.validated = true;
  },

  setLoggedConsumer(state, id) {
    state.loggedConsumer = state.consumers.filter(function (consumer) {
      return consumer.id == id
    })
  }

};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
