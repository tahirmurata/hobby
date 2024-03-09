import { createApp } from "vue";
import "./style.css";
import { createRouter, createWebHashHistory } from "vue-router";
import HelloWorld from "./components/HelloWorld.vue";
import Dice from "./components/Dice.vue";
import App from "./App.vue";

const routes = [
  { path: "/hello", component: HelloWorld },
  { path: "/dice", component: Dice },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

createApp(App).use(router).mount("#app");
