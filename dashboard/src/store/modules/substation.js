import subestation from "@/services/substation.js";


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

    return subestation.create(payload)
  },
  listSubstations({ commit }) {
    subestation.list().then(response => {
      //console.log(response.data)
      commit("setSubstations", response.data);
    });
  },
};

// mutations
const mutations = {
  setSubstations(state, subestations) {
    state.substations = subestations;
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
