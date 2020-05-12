import Vue from "vue";
import Vuex from "vuex";
import substation from "./modules/substation.js";
import consumer from "./modules/consumer.js";
import user from "./modules/user.js";
import transformer from "./modules/transformer.js"

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    substation,
    transformer,
    consumer,
    user
  }
});
