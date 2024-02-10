<script setup lang="ts">
	import Navbar from "./Navbar.vue";
	import { ref, onMounted } from "vue";
	import axios from "axios";
	import { useRoute } from "vue-router";

	interface Siswa {
		nis: number;
		name: string;
		profilepicture: string;
		email: string;
		gender: string;
	}

	const siswaData = ref<Siswa[]>([]);
	const route = useRoute();

	onMounted(async () => {
		// Get kelas ID from the route parameter
		const kelasId = route.params.id;

		try {
			// Fetch data from the API using Axios
			const response = await axios.get(
				`http://localhost:8080/Auth/siswa/${kelasId}`
			);
			// Update the siswaData with the received data
			siswaData.value = response.data;
		} catch (error) {
			console.error("Error fetching data:", error);
		}
	});
</script>

<template>
	<div class="container">
		<Navbar />
		<div class="kelas-content">
			<div class="walas-container">
				<div class="profile-information">
					<img src="../assets/Doge hehe.jpg" alt="" class="profile-picture" />
					<p class="nis">001002003</p>
					<div class="profile-info-container">
						<div class="profile-info">
							<p>Nama :</p>
							<p>Doge</p>
						</div>
						<div class="profile-info">
							<p>Kelas :</p>
							<p>XI - RPL</p>
						</div>
						<div class="profile-info">
							<p>Email :</p>
							<p>Dogeisthebest@gmail.com</p>
						</div>
						<div class="profile-info">
							<p>Gender :</p>
							<p>Doge</p>
						</div>
						<div class="profile-info">
							<p>Agama :</p>
							<p>The Religion Of Doge</p>
						</div>
						<div class="profile-info">
							<p>Bidang :</p>
							<p>Matematika</p>
						</div>
					</div>
				</div>
			</div>
			<div class="siswa-container">
				<table>
					<thead>
						<tr>
							<th>NO</th>
							<th>Foto</th>
							<th>Nama</th>
							<th>Nis</th>
							<th>Email</th>
							<th>Gender</th>
						</tr>
					</thead>
					<tbody>
						<tr v-for="(siswa, index) in siswaData" :key="index">
							<td>{{ index + 1 }}</td>
							<!-- Use a conditional check for profile_picture -->
							<td>
								<img
									:src="
										siswa.profilepicture !== ''
											? siswa.profilepicture
											: require('../assets/Doge hehe.jpg')
									"
									alt=""
								/>
							</td>
							<td>{{ siswa.name }}</td>
							<td>{{ siswa.nis }}</td>
							<td>{{ siswa.email }}</td>
							<td>{{ siswa.gender }}</td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.kelas-content {
		margin-top: 60px;
		display: flex;
		justify-content: space-between;
		align-items: flex-start;
	}

	.walas-container {
		width: 30%;
	}

	.siswa-container {
		width: 60%;
		display: flex;
		align-items: flex-start;
	}

	.profile-information {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		/* background-color: rgb(238, 238, 238); */
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
		border-radius: 7px;
		padding: 25px;
	}

	.profile-picture {
		max-width: 150px;
		border-radius: 50%;
	}

	.nis {
		margin-bottom: 20px;
		margin-top: 10px;
		color: grey;
	}

	.profile-info {
		display: flex;
		justify-content: space-between;
		margin: 40px 0px;
		color: white;
	}

	.profile-info-container {
		width: 80%;
		background-color: #0000ff;
		padding: 20px;
		border-radius: 7px;
	}

	table {
		width: 100%;
		border-collapse: collapse;
		margin-top: 50px;
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.2);
	}

	th,
	td {
		border: 1px solid #ddd;
		padding: 20px 8px;
		text-align: center;
		border: none;
	}

	th {
		background-color: #f2f2f2;
	}

	th:first-child {
		border-top-left-radius: 10px;
	}

	th:last-child {
		border-top-right-radius: 10px;
	}

	th {
		background-color: #0000ff;
		color: #fff;
	}

	td {
		border-color: #fff;
	}

	td img {
		max-width: 80px;
		border-radius: 50px;
	}

	tr:nth-child(even) {
		background-color: #f2f2f2;
	}
</style>
