import UserService from "@/core/services/user";

import { USERS_SET, USER_ADD, USER_REMOVE } from "./mutations";

export const USER_FETCH = "fetchUser";
export const USERS_FETCH = "fetchUsers";
export const USER_CREATE = "createUser";
export const USER_DELETE = "deleteUser";
export const USER_UPDATE = "updateUser";

const actions = {
  [USER_DELETE]: async (context, slug) => {
    const { data } = await UserService.del(slug);
    context.commit(USER_REMOVE, slug);
    return data;
  },
  [USER_CREATE]: async (context, user) => {
    const { data } = await UserService.create(user);
    context.commit(USER_ADD, data);
    return data;
  },
  [USERS_FETCH]: async (context, page = 0) => {
    const { data } = await UserService.list(page);
    context.commit(USERS_SET, data);
    return data;
  },
  [USER_FETCH]: async (context, id) => {
    const { data } = await UserService.get(id);
    context.commit(USER_SET, id);
    return data;
  }
};

export default actions;
