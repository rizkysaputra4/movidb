const routes = [
  { path: "/login", component: () => import("pages/Login.vue") },
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      { path: "/indexc", component: () => import("pages/Index.vue") },
      { path: "/dashboard", component: () => import("pages/Dashboard.vue") },
      { path: "/admin_list", component: () => import("pages/Admin-list.vue") },
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
