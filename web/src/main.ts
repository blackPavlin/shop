import { createApp } from "vue";
import ElementPlus from "element-plus";
import ruRU from "element-plus/es/locale/lang/ru";
import "element-plus/dist/index.css";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";

createApp(App)
  .use(ElementPlus, { locale: ruRU })
  .use(createPinia())
  .use(router)
  .mount("#app");
