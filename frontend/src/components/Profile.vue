<script setup lang="ts">
	import axios from "axios";
	import Navbar from "./Navbar.vue";
	import { ref, onMounted } from "vue";
	import { useRouter, useRoute } from "vue-router";
	import Cropper from "./ProfilePictureForm.vue";

	const tabAktif = ref("");
	const router = useRouter();
	const route = useRoute();

	const navigasi = (tab: string) => {
		tabAktif.value = tab;
		router.push(`/${tab}`);
	};

	interface User {
		NIS: number;
		Nama: string;
		Profilepicture: string;
		Email: string;
		No_telp: number;
		Gender: string;
		Religion: string;
		KelasData: {
			NamaKelas: string;
		};
	}

	const userData = ref<User | null>(null);

	onMounted(async () => {
		const NIS = route.params.id;

		try {
			const response = await axios.get(`http://localhost:8080/Auth/get/${NIS}`);

			userData.value = response.data;
		} catch (error) {
			console.log(error);
		}
	});

	const getProfilePictureUrl = (filePath: string) => {
		const baseURL = "http://localhost:8080";
		const normalizedPath = filePath.replace(/\\/g, "/");
		return `${baseURL}/${normalizedPath}`;
	};

	const isCropperVisible = ref(false);

	const openCropper = () => {
		isCropperVisible.value = true;
	};

	const closeCropper = () => {
		isCropperVisible.value = false;
	};

	const showOverlay = ref(false);
	const showOverlayText = ref(false);

	const onMouseEnter = () => {
		showOverlay.value = true;
		showOverlayText.value = true;
	};

	const onMouseLeave = () => {
		showOverlay.value = false;
		showOverlayText.value = false;
	};
</script>

<template>
	<div class="container">
		<Navbar />
		<div class="profile-content">
			<div class="profile-information">
				<Cropper v-if="isCropperVisible" @close-cropper="closeCropper" />
				<div
					class="profile-picture-container"
					@mouseenter="onMouseEnter"
					@mouseleave="onMouseLeave"
				>
					<img
						v-if="userData?.Profilepicture && !isCropperVisible"
						:src="getProfilePictureUrl(userData?.Profilepicture)"
						alt=""
						class="profile-picture"
					/>
					<img
						v-else
						src="../assets/Doge hehe.jpg"
						alt=""
						class="profile-picture"
					/>
					<div v-if="showOverlay" class="overlay" @click="openCropper"></div>
					<div v-if="showOverlayText" class="overlay-text" @click="openCropper">
						<ion-icon name="create"></ion-icon>
					</div>
				</div>
				<p class="nis">001002003</p>
				<div class="profile-info">
					<p>Nama :</p>
					<p>{{ userData?.Nama }}</p>
				</div>
				<div class="profile-info">
					<p>Kelas :</p>
					<p>{{ userData?.KelasData.NamaKelas }}</p>
				</div>
				<div class="profile-info">
					<p>Email :</p>
					<p>{{ userData?.Email }}</p>
				</div>
				<div class="profile-info">
					<p>No-Telpon :</p>
					<p>{{ userData?.No_telp }}</p>
				</div>
				<div class="profile-info">
					<p>Gender :</p>
					<p>{{ userData?.Gender }}</p>
				</div>
				<div class="profile-info">
					<p>Agama :</p>
					<p>{{ userData?.Religion }}</p>
				</div>
			</div>
			<div class="tugas-container">
				<div class="belum-selesai">
					<h2>Segera Kerjakan!</h2>
					<div class="tugas-card" @click="navigasi('tugas')">
						<div class="tugas-info">
							<h4>Matematika</h4>
							<h3>Fungsi CIN, COS, TAN</h3>
							<p>Dead Line 23/02/24</p>
						</div>
						<img src="../assets/Doge hehe.jpg" alt="" class="pp-guru" />
					</div>
					<div class="tugas-card telat">
						<div class="tugas-info">
							<h4>Pemodelan</h4>
							<h3>Use Case</h3>
							<p>Dead Line 12/02/24</p>
						</div>
						<img src="../assets/Doge hehe.jpg" alt="" class="pp-guru" />
					</div>
					<div class="tugas-card telat">
						<div class="tugas-info">
							<h4>Bahasa Indonesia</h4>
							<h3>Membuat cerpen</h3>
							<p>Dead Line 09/02/24</p>
						</div>
						<img src="../assets/Doge hehe.jpg" alt="" class="pp-guru" />
					</div>
				</div>
				<div class="selesai"></div>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.profile-content {
		margin-top: 70px;
		display: flex;
		justify-content: space-around;
	}
	.profile-information {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		width: 30%;
		/* background-color: rgb(238, 238, 238); */
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
		border-radius: 7px;
		padding: 20px;
	}

	.profile-picture {
		max-width: 200px;
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
		width: 80%;
		margin: 15px 0px;
	}
	.tugas-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 60%;
	}

	.belum-selesai {
		display: flex;
		flex-direction: column;
		width: 90%;
	}

	.tugas-card {
		display: flex;
		margin: 8px 0px;
		justify-content: space-between;
		align-items: center;
		padding: 15px;
		border-radius: 7px;
		max-width: 100%;
		/* background-color: rgb(238, 238, 238); */
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
		transition: 0.3s all ease-in-out;
	}

	.tugas-card:hover {
		background-color: #0000ff;
	}

	.tugas-card h4 {
		color: grey;
		font-weight: 500;
		font-size: 0.8rem;
		margin-bottom: 6px;
		transition: 0.3s all ease-in-out;
	}

	.tugas-card:hover h4 {
		color: rgb(230, 230, 230);
	}

	.tugas-card h3 {
		font-size: 1.5rem;
		transition: 0.3s all ease-in-out;
		width: 300px;
	}

	.tugas-card:hover h3 {
		color: #fff;
	}

	.tugas-card p {
		margin-top: 10px;
		font-size: 0.8rem;
		background-color: #ff0000;
		max-width: 55%;
		padding: 6px;
		border-radius: 20px;
		text-align: center;
		color: white;
		font-weight: 600;
		transition: 0.3s all ease-in-out;
	}

	.tugas-card:hover p {
		background-color: #ff0000;
		color: #fff;
	}

	.pp-guru {
		max-width: 100px;
		border-radius: 50%;
	}

	.telat {
		background-color: #ff0000;
	}

	.telat h4 {
		color: rgb(230, 230, 230);
	}

	.telat h3 {
		color: #fff;
	}

	.telat p {
		background-color: #fff;
		color: #ff0000;
	}

	.telat:hover {
		background-color: rgb(212, 0, 0);
	}

	.profile-picture-container {
		position: relative;
		cursor: pointer;
	}

	.overlay {
		position: absolute;
		top: 0;
		left: 0;
		width: 200px;
		height: 195px;
		background-color: rgba(128, 128, 128, 0.5);
		border-radius: 50%;
		opacity: 0;
		transition: opacity 0.3s ease-in-out;
	}

	.overlay-text {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		color: #fff;
		font-size: 3rem;
		font-weight: bold;
		text-align: center;
		opacity: 0;
		transition: opacity 0.3s ease-in-out;
	}

	.profile-picture-container:hover .overlay,
	.profile-picture-container:hover .overlay-text {
		opacity: 1;
	}

	@media (max-width: 1300px) {
		.profile-content {
			flex-direction: column;
			align-items: center;
		}

		.profile-information {
			width: 500px;
			margin-bottom: 100px;
		}

		.tugas-conatainer {
			width: 500px;
		}

		.belum-selesai {
			width: 500px;
		}

		.tugas-card {
			width: 500px;
		}

		.pp-guru {
			max-width: 15%;
		}

		.tugas-card h3 {
			width: 280px;
		}

		.tugas-card p {
			width: 280px;
		}
	}
</style>
