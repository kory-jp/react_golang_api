import { ActionTypes } from "../store/actionTypes"
import { User, UsersActionTypes } from "./types"

export const showUserAction = (userState: User): UsersActionTypes => {
  return {
    type: ActionTypes.showUser,
    payload: userState
  }
}