import subestation from "@/services/substation.js";
import transformer from "@/services/transformer.js";
import { getTransformers } from "@/helpers/helpers.js"


// initial state
const state = {
  substations: [],
  transformers: [],
  subs: false,
};

// getters
const getters = {};

// actions
const actions = {
  createSubstation({ commit, dispatch }, payload) {
    commit("setSubstationsSaved")
    payload.latitude = payload.latitude.toString()
    payload.longitude = payload.longitude.toString()

    subestation.create(payload).then(() => {
      dispatch("listSubstations")
    }).catch(error => {
      return error
    })
    return
  },

  createTransformer({ commit }, payload) {
    payload.latitude = payload.latitude.toString()
    payload.longitude = payload.longitude.toString()

    transformer.create(payload).then(() => {
      commit("setTransformers")
    }).catch(error => {
      return error
    })
  },

  listSubstations({ commit }) {
    subestation.list().then(response => {
      commit("setSubstations", response.data);
      commit("setTransformers")
    });
  },
};

// mutations
const mutations = {
  setSubstations(state, subestations) {
    state.substations = subestations;
  },

  setTransformers(state) {
    let transformerList = getTransformers(state.substations)
    state.transformers = transformerList;
  },

  setSubstationsSaved(state) {
    state.subs = true;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
