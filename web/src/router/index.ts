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
    path: "/address",
    name: "AddressView",
    component: () => import("@/views/Address.vue"),
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
    children: [
      {
        path: "category",
        name: "CategoryAdmin",
        component: () => import("@/views/admin/CategoryAdmin.vue"),
      },
      {
        path: "product",
        name: "ProductAdmin",
        component: () => import("@/views/admin/ProductAdmin.vue"),
      },
      {
        path: "order",
        name: "OrderAdmin",
        component: () => import("@/views/admin/OrderAdmin.vue"),
      },
    ],
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("@/views/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
