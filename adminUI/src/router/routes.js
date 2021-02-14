const routes = [
  { path: "/login", component: () => import("pages/Login.vue") },
  {
    path: "/",
    component: () => import("layouts/MainLayout.vue"),
    children: [
      { path: "/indexc", component: () => import("pages/Index.vue") },
      { path: "/dashboard", component: () => import("pages/Dashboard.vue") },
      { path: "/admin-list", component: () => import("pages/Admin-list.vue") },
      {
        path: "/admin-promote",
        component: () => import("pages/Admin-promote.vue"),
      },
      { path: "/new-user", component: () => import("pages/AddUser.vue") },
      { path: "/profile", component: () => import("pages/Profile.vue") },
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
