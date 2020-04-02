import subestation from "@/core/services/substation.js";
import transformer from "@/core/services/transformer.js";

// initial state
const state = {
  substations: [],
  transformers: []
};

// getters
const getters = {};

// actions
const actions = {
  loadAllSubstations({ commit }) {
    subestation.list().then(response => {
      commit("setSubstations", response.data);
    });
  },

  loadAllTransformers({ commit }) {
    transformer.list().then(response => {
      commit("setTransformers", response.data);
    });
  }
};

// mutations
const mutations = {
  setSubstations(state, subestations) {
    state.substations = subestations;
  },

  setTransformers(state, transformers) {
    state.transformers = transformers;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
