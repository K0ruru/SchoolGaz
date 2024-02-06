import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

import Dashboard from "./components/Dashboard.vue";
import Kelas from "./components/Kelas.vue";
import Signup from "./components/SignupForm.vue";
import Login from "./components/LoginForm.vue";

const routes: Array<RouteRecordRaw> = [
	{
		path: "/",
		component: Dashboard,
	},
	{
		path: "/kelas",
		component: Kelas,
	},
	{
		path: "/signup",
		component: Signup,
	},
	{
		path: "/login",
		component: Login,
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;
