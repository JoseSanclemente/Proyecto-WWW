import Vue from "vue";
import Vuex from "vuex";
import assetsManagement from "./modules/assets-management.js";
import clientsManagement from "./modules/clients-management.js";
import user from "./modules/user.js";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    assetsManagement,
    clientsManagement,
    user
  }
});
