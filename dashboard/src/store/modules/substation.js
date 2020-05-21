import substation from "@/services/substation.js";

// initial state
const state = {
  substations: [],
  subs: false,
};

// getters
const getters = {};

// actions
const actions = {
  createSubstation({ commit }, payload) {
    commit("setSubstationsSaved")
    payload.latitude = payload.latitude.toString()
    payload.longitude = payload.longitude.toString()

    return substation.create(payload)
  },
  listSubstations({ commit }) {
    substation.list().then(response => {
      commit("setSubstations", response.data);
    });
  },
  deactivateSubstation({ commit }, payload) {
    payload.deleted = "true"
    substation.update(payload)
    commit("dummyMutation")
  }
};

// mutations
const mutations = {
  setSubstations(state, subestations) {
    state.substations = subestations;
  },
  setSubstationsSaved(state) {
    state.subs = true;
  },
  dummyMutation() {
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
