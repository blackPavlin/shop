import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "MainView",
    component: () => import("@/views/Main.vue"),
  },
  {
    path: "/login",
    name: "LoginView",
    component: () => import("@/views/Login.vue"),
  },
  {
    path: "/logout",
    name: "LogoutView",
    component: () => import("@/views/Logout.vue"),
  },
  {
    path: "/registration",
    name: "RegistrationView",
    component: () => import("@/views/Registration.vue"),
  },
  {
    path: "/cart",
    name: "CardView",
    component: () => import("@/views/Cart.vue"),
  },
  {
    path: "/orders",
    name: "OrdersView",
    component: () => import("@/views/Orders.vue"),
  },
  {
    path: "/products",
    name: "ProductsView",
    component: () => import("@/views/Products.vue"),
  },
  {
    path: "/admin",
    name: "AdminView",
    component: () => import("@/views/admin/Admin.vue"),
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("@/views/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
