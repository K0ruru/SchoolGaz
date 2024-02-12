// main.js or where you configure your Vue app

import { createApp } from "vue";
import App from "./App.vue";
import router from "./routes";
import "./style.css";

import VueSweetalert2 from "vue-sweetalert2";
import "sweetalert2/dist/sweetalert2.min.css";

const app = createApp(App);

// Add a global navigation guard
router.beforeEach((to, _, next) => {
	// Check if the route requires authentication
	if (to.meta.requiresAuth) {
		// Check if the token exists in local storage
		const token = localStorage.getItem("token");

		if (token) {
			// Token exists, proceed to the route
			next();
		} else {
			// Token does not exist, redirect to the login page or another route
			next("/login");
		}
	} else {
		// Route does not require authentication, proceed
		next();
	}
});

app.use(VueSweetalert2);
app.use(router);
app.mount("#app");
