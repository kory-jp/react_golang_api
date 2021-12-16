import { User } from "./types"
import axios from 'axios';
import { showUserAction } from "./actions";
import { Dispatch } from "react";

export const showUser = (userId: number) => {
  return async(dispatch: Dispatch<{}>) => {
    axios
      .get(`http://localhost:8080/users/show/${userId}`)
      .then((response) => {
        const userDate = response.data
        dispatch(showUserAction(userDate))
      })
      .catch((error) => {
        console.log(error)
      })
  }
}