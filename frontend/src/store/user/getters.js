const getters = {
  users(state) {
    let users = Object.values(state.users);
    if (state.filter == "") {
      return users;
    }

    let filteredUsers = [];
    let filter = state.filter.toLowerCase();
    for (let user of users) {
      for (let prop of Object.keys(user)) {
        if (
          typeof user[prop] === "string" &&
          user[prop].toLowerCase().indexOf(filter) !== -1
        ) {
          filteredUsers.push(user);
          break;
        }
      }
    }

    return filteredUsers;
  }
};

export default getters;
