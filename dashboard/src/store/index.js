import Vue from "vue";
import Vuex from "vuex";
import substation from "./modules/substation.js";
import consumer from "./modules/consumer.js";
import user from "./modules/user.js";
import report from "./modules/report.js"

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    substation,
    consumer,
    user,
    report
  }
});
