import { applyMiddleware, combineReducers, createStore } from "redux";
import thunk from "redux-thunk";
import { UserReducer } from "../users/reducers";

export default function createInitStore() {
  return createStore(
    combineReducers({
      user: UserReducer
    }),
    applyMiddleware(thunk)
  )
}