import transformer from "@/services/transformer.js";

// initial state
const state = {
    transformers: []
};

// getters
const getters = {};

// actions
const actions = {
    createTransformer(payload) {
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
    }
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
};
