import Vue from "vue";
import Vuex from "vuex";
import assetsManagement from "./modules/assets-management.js";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    assetsManagement
  }
});
