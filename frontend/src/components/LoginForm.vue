<script setup lang="ts">
	import { ref } from "vue";
	import { useRouter } from "vue-router";

	const nis = ref("");
	const pass = ref("");

	const router = useRouter();

	const login = () => {
		router.push("/");
	};
</script>

<template>
	<div class="container-center">
		<form @submit.prevent="login">
			<h1>Login</h1>
			<div class="form-inputs">
				<div class="form-input">
					<label for="nis">NIS :</label>
					<input v-model="nis" type="text" id="nis" name="nis" />
				</div>
				<div class="form-input">
					<label for="pass">Passphrase :</label>
					<input v-model="pass" type="password" id="pass" name="pass" />
				</div>
			</div>
			<input type="submit" value="Log-in" />
			<p>Belum punya akun? <router-link to="/signup">Sign-up</router-link></p>
		</form>
	</div>
</template>

<style scoped>
	.form-inputs {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.form-input {
		display: flex;
		flex-direction: column;
		align-items: center;
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
</style>
