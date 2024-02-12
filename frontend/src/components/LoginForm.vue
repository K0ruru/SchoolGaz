<script setup lang="ts">
	import { ref } from "vue";
	import { useRouter } from "vue-router";
	import axios, { AxiosError } from "axios";
	import { useToast } from "primevue/usetoast";
	import Toast from "primevue/toast";

	const toast = useToast();

	const nis = ref("");
	const pass = ref("");

	const router = useRouter();

	const login = async () => {
		try {
			const response = await axios.post("http://localhost:8080/Auth/login", {
				nis: nis.value,
				passphrase: pass.value,
			});

			const { token, status, nama, NIS } = response.data;

			if (status !== "active") {
				toast.add({
					severity: "error",
					summary: "Error",
					detail: "Maaf akun anda belum aktif.",
					life: 3000,
				});
				return;
			}

			localStorage.setItem("NIS", NIS);
			localStorage.setItem("token", token);
			localStorage.setItem("nama", nama);

			router.push("/");
		} catch (error: any) {
			const axiosError = error as AxiosError<{ error: string }>;

			if (axiosError.response && axiosError.response.status === 400) {
				const errorData = axiosError.response.data;

				if (errorData.error === "User not found") {
					toast.add({
						severity: "error",
						summary: "Error",
						detail: "Tidak ada user dengan nis itu :(",
						life: 3000,
					});
				} else if (errorData.error === "Invalid password") {
					toast.add({
						severity: "error",
						summary: "Error",
						detail: "Password yang anda masukkan salah.",
						life: 3000,
					});
				} else {
					console.error("Error logging in:", error);
				}
			} else {
				console.error("Error logging in:", error);
			}
		}
	};
</script>

<template>
	<div class="container-center">
		<Toast position="bottom-left" />
		<form @submit.prevent="login">
			<h1>Login</h1>
			<div class="form-inputs">
				<div class="form-input">
					<label for="nis">NIS :</label>
					<input v-model="nis" type="number" id="nis" name="nis" />
				</div>
				<div class="form-input">
					<label for="pass">Passphrase :</label>
					<input v-model="pass" type="password" id="pass" name="pass" />
				</div>
			</div>
			<!-- <input type="submit" value="Log-in" /> -->
			<button class="button" type="submit">Log-in</button>
			<p>Belum punya akun? <router-link to="/signup">Sign-up</router-link></p>
		</form>
		<img src="../assets/vector.png" alt="" class="vector" />
	</div>
</template>

<style scoped>
	h1 {
		text-align: center;
		margin-bottom: 17px;
	}

	.form-inputs {
		display: flex;
		flex-direction: column;
		/* align-items: center; */
	}

	.form-input {
		display: flex;
		flex-direction: column;
		/* align-items: center; */
	}

	.form-input label {
		margin-bottom: 10px;
	}

	.form-input input {
		width: 100%;
		padding: 10px;
		margin-bottom: 10px;
		border: 1px solid #ccc;
		border-radius: 5px;
	}

	.form-input input:focus {
		outline: none;
	}

	.form-input input:focus::placeholder {
		color: #ccc;
	}

	.button {
		font-family: Poppins, sans-serif;
		border-radius: 100px;
		box-shadow: 0px 15px 30px 0px rgba(0, 26, 255, 0.2);
		background-color: #001aff;
		color: #fff;
		padding: 12px 17px;
		outline: none;
		border: none;
		width: 100%;
		margin-top: 10px;
		margin-bottom: 10px;
	}

	.container-center {
		gap: 5rem;
	}

	form {
		width: 20%;
	}

	.vector {
		width: 50%;
	}

	@media (max-width: 991px) {
		.button {
			white-space: initial;
		}
	}
</style>
