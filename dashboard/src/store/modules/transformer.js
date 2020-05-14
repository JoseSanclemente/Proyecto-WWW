import transformer from "@/services/transformer.js";
import { getTransformers } from "@/helpers/helpers.js"

// initial state
const state = {
    transformers: [],
    transSaved: false,
};

// getters
const getters = {};

// actions
const actions = {
    createTransformer({ commit }, payload) {
        commit("setTransSaved")
        payload.latitude = payload.latitude.toString()
        payload.longitude = payload.longitude.toString()
        return transformer.create(payload)
    },
};

// mutations
const mutations = {
    setTransformers(state, substations) {
        let transformerList = getTransformers(substations)
        state.transformers = transformerList;
    },
    setTransSaved(state) {
        state.transSaved = true
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
};
