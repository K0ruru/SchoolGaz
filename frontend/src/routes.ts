import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

import Dashboard from "./components/Dashboard.vue";
import Kelas from "./components/Kelas.vue";

const routes: Array<RouteRecordRaw> = [
	{
		path: "/",
		component: Dashboard,
	},
	{
		path: "/kelas",
		component: Kelas,
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;
