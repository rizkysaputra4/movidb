import Index from "pages/Index.vue";
import Login from "pages/Login.vue";
import Dashboard from "pages/Dashboard.vue";

const routes = [
  { path: "/login", component: Login },
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      { path: "/indexc", component: Index },
      { path: "/dashboard", component: Dashboard },
    ],
    meta: { requiresAuth: true },
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: "*",
    component: () => import("pages/Error404.vue"),
  },
];

export default routes;
