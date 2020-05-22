import substation from "@/services/substation.js";
import transformer from "@/services/transformer.js";
import { getTransformers } from "@/helpers/helpers.js"


// initial state
const state = {
  substations: [],
  searchedSubstations: [],
  transformers: [],
  searchedTransformers: [],
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

    substation.create(payload).then(() => {
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
    substation.list().then(response => {
      commit("setSubstations", response.data);
      commit("setTransformers")
    });
  },

  updateTransformer({ commit }, payload) {
    commit("dummyMutation")
    payload.deleted = payload.deleted.toString()
    return transformer.update(payload)
  },

  updateSubstation({ commit }, payload) {
    commit("dummyMutation")
    payload.deleted = payload.deleted.toString()
    return substation.update(payload)
  },
};

// mutations
const mutations = {
  setSubstations(state, subestations) {
    state.substations = subestations;
    state.searchedSubstations = subestations
  },

  setTransformers(state) {
    let transformerList = getTransformers(state.substations)

    state.searchedTransformers = transformerList
    state.transformers = transformerList;
  },

  setSubstationsSaved(state) {
    state.subs = true;
  },

  searchTransformer(state, name) {
    if (name) {
      state.searchedTransformers = state.transformers.filter(item => item.name.toLowerCase().includes(name.toLowerCase()))
    } else {
      state.searchedTransformers = state.transformers
    }
  },

  searchSubstation(state, name) {
    if (name) {
      state.searchedSubstations = state.substations.filter(item => item.name.toLowerCase().includes(name.toLowerCase()))
    } else {
      state.searchedSubstations = state.substations
    }
  },

  dummyMutation(){}
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
