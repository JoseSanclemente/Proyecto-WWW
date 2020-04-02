import subestation from "@/core/services/substation.js";

// initial state
const state = {
  substations: []
};

// getters
const getters = {};

// actions
const actions = {
  loadAllSubstations({ commit }) {
    subestation.list().then(response => {
      commit("setSubstations", response.data);
    });
  }
};

// mutations
const mutations = {
  setSubstations(state, subestations) {
    state.substations = subestations;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
