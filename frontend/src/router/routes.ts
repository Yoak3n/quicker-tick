import { RouteRecordRaw } from "vue-router";
import Modify from "../views/Modify.vue";

const routes: Array<RouteRecordRaw> = [
  { path: "/", component: () => import("../views/Home.vue") },
  { path: "/modify/:date", component: Modify ,name:"Modify"},
  { path: "/dashboard", component: () => import("../views/Dashboard.vue") },
]


export default routes;