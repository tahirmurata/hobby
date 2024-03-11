import { createApp, reactive } from "vue";
import "./style.css";
import { createRouter, createWebHashHistory } from "vue-router";
import Dice from "./components/Dice.vue";
import App from "./App.vue";

const routes = [{ path: "/dice", component: Dice }];

export const state = reactive({ coins: 100 });

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

createApp(App).use(router).mount("#app");
