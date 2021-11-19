import EditTodo from "../components/pages/EditTodo";
import Index from "../components/pages/Index";
import NewTodo from "../components/pages/NewTodo";
import Page404 from "../components/pages/Page404";


export const todoRoutes = [
  {
    path: "/",
    exact: true,
    children: <Index/>
  },
  {
    path: "/new",
    exact: false,
    children: <NewTodo/>
  },
  {
    path: "/edit",
    exact: false,
    children: <EditTodo/>
  },
  {
    path: "*",
    exact: false,
    children: <Page404/>
  }
]