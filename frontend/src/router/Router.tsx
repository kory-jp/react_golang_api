import { VFC } from "react";
import { Switch } from "react-router-dom";
import { Route } from "react-router";
import Login from "../components/pages/Login";
import Page404 from "../components/pages/Page404";
import Signup from "../components/pages/Signup";
import EditTodo from "../components/pages/todos/EditTodo";
import Index from "../components/pages/todos/Index";
import NewTodo from "../components/pages/todos/NewTodo";
import Top from "../components/pages/Top";
import DefaultTemplate from "../components/template/DefaultTemplate";

export const AppRouter: VFC = () => {
  return(
    <>
      <Switch>
        <DefaultTemplate>
          <Route exact path="/" component={Top}/>
          <Route exact path="/login" component={Login}/>
          <Route exact path="/signup" component={Signup}/>
          <Route exact path="/todo" component={Index}/>
          <Route exact path="/todo/new" component={NewTodo}/>
          <Route exact path="/todo/edit" component={EditTodo}/>
        </DefaultTemplate>
          <Route exact path="*" component={Page404}/>
      </Switch>
    </> 
  )
}