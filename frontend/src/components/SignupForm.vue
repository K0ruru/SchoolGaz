<script setup lang="ts">
	import { ref } from "vue";
	import { useRouter } from "vue-router";

	const nis = ref("");
	const nama = ref("");
	const pass = ref("");
	const passConfirm = ref("");
	const email = ref("");
	const noTelp = ref("");
	const jenkel = ref("");
	const agama = ref("");

	const router = useRouter();

	const signUp = () => {
		router.push("/");
	};
</script>

<template>
	<div class="container-center">
		<form @submit.prevent="signUp">
			<h1>Sign-Up</h1>
			<div class="form-inputs">
				<div class="form-input">
					<label for="nis">NIS :</label>
					<input v-model="nis" type="text" id="nis" name="nis" />
				</div>
				<div class="form-input">
					<label for="nama">Nama :</label>
					<input v-model="nama" type="text" id="nama" name="nama" />
				</div>
				<div class="form-input">
					<label for="pass">Passphrase :</label>
					<input v-model="pass" type="password" id="pass" name="pass" />
				</div>
				<div class="form-input">
					<label for="pass-confirm">Confirm Passphrase :</label>
					<input
						v-model="passConfirm"
						type="password"
						id="pass-confirm"
						name="pass-confirm"
					/>
				</div>
				<div class="form-input">
					<label for="email">E-mail :</label>
					<input v-model="email" type="text" id="email" name="email" />
				</div>
				<div class="form-input">
					<label for="no-telp">No-Telp :</label>
					<input v-model="noTelp" type="text" id="no-telp" name="no-telp" />
				</div>
				<div class="form-input">
					<label for="jenkel">Jenis Kelamin :</label>
					<input v-model="jenkel" type="text" id="jenkel" name="jenkel" />
				</div>
				<div class="form-input">
					<label for="agama">Agama :</label>
					<input v-model="agama" type="text" id="agama" name="agama" />
				</div>
			</div>
			<input type="submit" value="Sign-Up" />
			<p>Sudah punya akun? <router-link to="/login">Login</router-link></p>
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
