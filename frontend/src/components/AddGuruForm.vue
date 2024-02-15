<script setup lang="ts">
	import { ref, defineProps, defineEmits } from "vue";
	import { useToast } from "primevue/usetoast";
	import axios from "axios";

	const { guruData } = defineProps(["guruData"]);
	const emits = defineEmits(["closeEditForm"]);
	const toast = useToast();

	const nis = ref(guruData.NIS);
	const nama = ref(guruData.NamaGuru);
	const email = ref(guruData.Email);
	const noTelp = ref(guruData.NoTelp);
	const jenkel = ref(guruData.Gender);
	const agama = ref(guruData.Religion);
	const mapel = ref(guruData.Mapel);

	interface Mapel {
		Id_mapel: number;
		Nama_mapel: string;
	}

	const mapelData = ref<Mapel[]>([]);

	axios
		.get("http://localhost:8080/mapel/")
		.then((response) => {
			mapelData.value = response.data;
		})
		.catch((error) => {
			console.error("Error fetching mapel data:", error);
		});

	const saveChanges = async () => {
		try {
			const response = await axios.put(
				`http://localhost:8080/guru/update/${String(nis.value)}`,
				{
					NIS: nis.value,
					NamaGuru: nama.value,
					Email: email.value,
					NoTelp: noTelp.value,
					Gender: jenkel.value,
					Religion: agama.value,
					Mapel: mapel.value,
				}
			);

			if (response.status === 200) {
				emits("closeEditForm"); // Emit the closeEditForm event
				toast.add({
					severity: "info",
					summary: "Confirmed",
					detail: "Data Berhasil Di-edit",
					life: 3000,
				});
			} else {
				// Handle other status codes or errors
				console.error("Failed to update data.");
			}
		} catch (error) {
			console.error("Error updating data:", error);
			toast.add({
				severity: "error",
				summary: "Error",
				detail: "Terjadi sebuah kesalahan :(",
				life: 3000,
			});
		}
	};
</script>

<template>
	<div class="edit-form-overlay">
		<div class="edit-form-container">
			<form @submit.prevent="saveChanges">
				<div class="form-heading">
					<h1>Edit Profile</h1>
					<button class="close-button" @click="emits('closeEditForm')">
						<ion-icon name="close"></ion-icon>
					</button>
				</div>
				<div class="form-inputs">
					<div class="form-input">
						<label for="nis">NIS :</label>
						<input v-model="nis" type="number" id="nis" name="nis" disabled />
					</div>
					<div class="form-input">
						<label for="nama">Nama :</label>
						<input v-model="nama" type="text" id="nama" name="nama" />
					</div>
					<div class="form-input">
						<label for="email">E-mail :</label>
						<input v-model="email" type="email" id="email" name="email" />
					</div>
					<div class="form-input">
						<label for="notelp">No-Telp :</label>
						<input v-model="noTelp" type="number" id="notelp" name="notelp" />
					</div>
					<div class="form-input">
						<label for="jenkel">Jenis Kelamin :</label>
						<input v-model="jenkel" type="text" id="jenkel" name="jenkel" />
					</div>
					<div class="form-input">
						<label for="agama">Agama :</label>
						<select v-model="agama" id="agama" name="agama">
							<option value="Islam">Islam</option>
							<option value="Kristen Protestan">Kristen Protestan</option>
							<option value="Kristen Katolik">Kristen Katolik</option>
							<option value="Hindu">Hindu</option>
							<option value="Buddha">Buddha</option>
							<option value="Khonghucu">Khonghucu</option>
						</select>
					</div>
					<div class="form-input">
						<label for="kelas">Bidang :</label>
						<select v-model="mapel" id="kelas" name="kelas">
							<option
								v-for="mapelItem in mapelData"
								:key="mapelItem.Id_mapel"
								:value="mapelItem.Id_mapel"
							>
								{{ mapelItem.Nama_mapel }}
							</option>
						</select>
					</div>
				</div>
				<button class="button" @click="saveChanges">Save Changes</button>
			</form>
		</div>
	</div>
</template>

<style scoped>
	.edit-form-overlay {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0, 0, 0, 0.5);
		display: flex;
		justify-content: center;
		align-items: center;
		z-index: 100;
	}

	.edit-form-container {
		background: #fff;
		padding: 20px;
		border-radius: 8px;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		width: 40%;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.container-center {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}

	form {
		max-width: 600px;
		width: 100%;
	}

	.form-heading {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.form-inputs {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		gap: 15px;
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
		margin-top: 30px;
		margin-bottom: 10px;
		cursor: pointer;
	}

	.close-button {
		border: none;
		background-color: white;
		font-size: 35px;
		color: gray;
		cursor: pointer;
	}

	label {
		display: block;
		margin-top: 3px;
		margin-bottom: 2px;
		font-size: 12px;
	}

	select {
		width: 100%;
		padding: 8px;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
		font-size: 15px;
		cursor: pointer;
	}
</style>

