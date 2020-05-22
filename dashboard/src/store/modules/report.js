import report from "@/services/report.js"

// initial state
const state = {
    dummy: false,
    topConsumersList: []
};

const actions = {
    topConsumers({ commit }, payload) {
        commit("dummy")
        report.topConsumers(payload).then(response => {
            commit("setTopConsumers", response.data)
        }).catch(error => {
            console.log(error)
            return error
        })
    }
}

const mutations = {
    dummy(state) {
        state.dummy = true
    },

    setTopConsumers(state, payload) {
        state.topConsumersList = payload
    }
}

export default {
    namespaced: true,
    state,
    actions,
    mutations
};
