import { ActionTypes } from "../store/actionTypes";
import { User, UsersActionTypes } from "./types";

const initialState: User = {
  id: 0,
  uuid: "",
  name: "",
  email: "",
  password: "",
  created_at: null
}

export const UserReducer = (state = initialState, action: UsersActionTypes): User => {
  switch(action.type) {
    case ActionTypes.showUser:
      return {...state, ...action.payload}
  }
  return state;
}