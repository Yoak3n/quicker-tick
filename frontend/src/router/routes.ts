import { RouteRecordRaw } from "vue-router";
import Modify from "../views/Modify.vue";

const routes: Array<RouteRecordRaw> = [
  { path: "/", component: () => import("../views/Home.vue") },
  { path: "/modify/:date", component: Modify ,name:"Modify"},
  { path: "/dashboard/:date",name:"Dashboard", component: () => import("../views/Dashboard.vue") },
  {path:"/calendar",name:"Calendar",component:()=>import("../views/Calendar.vue")}
]


export default routes;