import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

import Dashboard from "./components/Dashboard.vue";
import KelasList from "./components/KelasList.vue";
import Kelas from "./components/Kelas.vue";
import Siswa from "./components/Siswa.vue";
import Guru from "./components/Guru.vue";
import Signup from "./components/SignupForm.vue";
import Login from "./components/LoginForm.vue";
import Profile from "./components/Profile.vue";
import Tugas from "./components/Tugas.vue";

const routes: Array<RouteRecordRaw> = [
	{
		path: "/",
		component: Dashboard,
		meta: { requiresAuth: true },
	},
	{
		path: "/kelas",
		component: KelasList,
		meta: { requiresAuth: true },
	},
	{
		path: "/siswa",
		component: Siswa,
		meta: { requiresAuth: true },
	},
	{
		path: "/guru",
		component: Guru,
		meta: { requiresAuth: true },
	},
	{
		path: "/signup",
		component: Signup,
	},
	{
		path: "/login",
		component: Login,
	},
	{
		path: "/profile",
		component: Profile,
		meta: { requiresAuth: true },
	},
	{
		path: "/tugas",
		component: Tugas,
		meta: { requiresAuth: true },
	},
	{
		path: "/kelas/:id",
		name: "kelas",
		component: Kelas,
		meta: { requiresAuth: true },
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;
