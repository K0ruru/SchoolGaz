import { createApp } from "vue";
import App from "./App.vue";
import router from "./routes";
import "./style.css";
import PrimeVue from "primevue/config";
import ToastService from "primevue/toastservice";
import ConfirmService from "primevue/confirmationservice";
import VueSweetalert2 from "vue-sweetalert2";
import "sweetalert2/dist/sweetalert2.min.css";
import "primevue/resources/themes/aura-light-indigo/theme.css";

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

// Use PrimeVue
app.use(PrimeVue, { ripple: true });
app.use(ToastService as any);
app.use(ConfirmService);
app.use(VueSweetalert2);
app.use(router);

app.mount("#app");
