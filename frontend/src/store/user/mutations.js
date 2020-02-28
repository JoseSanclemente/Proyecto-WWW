import Vue from "vue";

export const USER_SET = "setUser";
export const USER_REMOVE = "deleteUser";
export const USERS_SET = "setUsers";
export const USER_ADD = "addUser";

const mutations = {
  [USER_ADD]: (state, user) => {
    console.log(user);
    Vue.set(state.users, user.user_id, user);
  },
  [USER_REMOVE]: (state, id) => {
    Vue.delete(state.users, id);
  },
  [USER_SET]: (state, user) => {
    Vue.set(state.users, user.id, user);
  },
  [USERS_SET]: (state, users) => {
    state.users = {};
    for (let user of users) {
      Vue.set(state.users, user.user_id, user);
    }
  }
};

export default mutations;
