import transformer from "@/services/transformer.js";

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
    listTransformers({ commit }) {
        transformer.list().then(response => {
            commit("setTransformers", response.data);
        });
    }
};

// mutations
const mutations = {
    setTransformers(state, transformers) {
        state.transformers = transformers;
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
