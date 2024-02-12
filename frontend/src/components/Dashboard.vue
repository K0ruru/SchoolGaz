<script setup lang="ts">
	import Navbar from "./Navbar.vue";
	import { ref } from "vue";
	import { useRouter } from "vue-router";

	const tabAktif = ref("");
	const router = useRouter();

	const navigasi = (tab: string) => {
		tabAktif.value = tab;
		router.push(`/${tab}`);
	};

	const username = localStorage.getItem("nama");
</script>

<template>
	<div class="container">
		<Navbar />
		<div class="content-container">
			<div class="welcome-message">Welcome to SchoolGaz, {{ username }}</div>
			<div class="tutorial-message">
				&quot;Unsure how to navigate this website? Visit our tutorial webpage
				for guidance.&quot;
			</div>
			<div class="feature-section">
				<div class="feature-button">explore features</div>
				<div class="how-it-works">How it works?</div>
			</div>
		</div>

		<div class="cont">
			<!-- box untuk siswa -->
			<div class="box" @click="navigasi('siswa')">
				<div class="content">
					<div class="icon-1"><ion-icon name="people"></ion-icon></div>
					<div class="text">
						<div class="title">Siswa</div>
						<div class="subtitle">List Siswa</div>
					</div>
				</div>
				<div class="line"></div>
				<div class="info">
					<div class="value">170</div>
					<div class="red">
						<ion-icon name="radio-button-on-sharp"></ion-icon>
					</div>
				</div>
			</div>
			<!-- box untuk guru -->
			<div class="box" @click="navigasi('guru')">
				<div class="content">
					<div class="icon-2">
						<ion-icon name="accessibility"></ion-icon>
					</div>
					<div class="text">
						<div class="title">Guru</div>
						<div class="subtitle">List Guru</div>
					</div>
				</div>
				<div class="line"></div>
				<div class="info">
					<div class="value">170</div>
					<div class="green">
						<ion-icon name="radio-button-on-sharp"></ion-icon>
					</div>
				</div>
			</div>
			<!-- box untuk kelas -->
			<div class="box" @click="navigasi('kelas')">
				<div class="content">
					<div class="icon-3"><ion-icon name="school"></ion-icon></div>
					<div class="text">
						<div class="title">Kelas</div>
						<div class="subtitle">Daftar Kelas</div>
					</div>
				</div>
				<div class="line"></div>
				<div class="info">
					<div class="value">170</div>
					<div class="blue">
						<ion-icon name="radio-button-on-sharp"></ion-icon>
					</div>
				</div>
			</div>
			<!-- box 4 -->
			<div class="box" @click="navigasi('addtugas')">
				<div class="content">
					<div class="icon-3"><ion-icon name="book"></ion-icon></div>
					<div class="text">
						<div class="title">Tugas</div>
						<div class="subtitle">Tugas</div>
					</div>
				</div>
				<div class="line"></div>
				<div class="info">
					<div class="value">170</div>
					<div class="blue">
						<ion-icon name="radio-button-on-sharp"></ion-icon>
					</div>
				</div>
			</div>
			<!-- box-5 -->
			<div class="box" @click="navigasi('kelas')">
				<div class="content">
					<div class="icon-3"><ion-icon name="people-outline"></ion-icon></div>
					<div class="text">
						<div class="title">Kelas</div>
						<div class="subtitle">Daftar Kelas</div>
					</div>
				</div>
				<div class="line"></div>
				<div class="info">
					<div class="value">170</div>
					<div class="blue">
						<ion-icon name="radio-button-on-sharp"></ion-icon>
					</div>
				</div>
			</div>
			<!-- box-6 -->
			<div class="box" @click="navigasi('kelas')">
				<div class="content">
					<div class="icon-3"><ion-icon name="people-outline"></ion-icon></div>
					<div class="text">
						<div class="title">Kelas</div>
						<div class="subtitle">Daftar Kelas</div>
					</div>
				</div>
				<div class="line"></div>
				<div class="info">
					<div class="value">170</div>
					<div class="blue">
						<ion-icon name="radio-button-on-sharp"></ion-icon>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.cont {
		display: flex;
		gap: 20px;
		margin-top: 40px;
		align-items: center;
		justify-content: center;
		flex-wrap: wrap;
	}

	@media (max-width: 991px) {
		.cont {
			flex-direction: column;
		}
	}

	.box {
		border-radius: 8px;
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
		/* background-color: #fff; */
		display: flex;
		flex-direction: column;
		min-width: 30%;
		padding: 25px;
		height: 170px;
		transition: 0.3s all ease-in-out;
		cursor: pointer;
	}

	.box:hover {
		border-radius: 8px;
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
		background-color: #001aff;
		display: flex;
		flex-direction: column;
		min-width: 30%;
		padding: 25px;
		height: 170px;
	}

	.box:hover .title {
		color: #fff;
	}

	.box:hover .subtitle {
		color: #fff;
	}

	.box:hover .value {
		background-color: rgba(255, 255, 255, 0.15);
		color: #fff;
	}

	.box:hover .red {
		color: #fff;
	}

	.box:hover .green {
		color: #fff;
	}

	.box:hover .blue {
		color: #fff;
	}

	.box:hover .icon-1 {
		background-color: #fff;
		color: #000fff;
	}

	.box:hover .icon-2 {
		background-color: #fff;
		color: #000fff;
	}

	.box:hover .icon-3 {
		background-color: #fff;
		color: #000fff;
	}

	.content {
		display: flex;
		justify-content: space-between;
		gap: 13px;
	}

	.icon-1 {
		border-radius: 8px;
		background-color: #0000ff;
		color: white;
		width: 50px;
		height: 50px;
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 10px;
		font-size: 29px;
	}

	.icon-2 {
		border-radius: 8px;
		background-color: #0000ff;
		color: white;
		width: 50px;
		height: 50px;
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 10px;
		font-size: 29px;
	}

	.icon-3 {
		border-radius: 8px;
		background-color: #0000ff;
		color: white;
		width: 50px;
		height: 50px;
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 10px;
		font-size: 29px;
	}

	.text {
		/* display: flex; */
		flex-direction: column;
		min-width: 70%;
		display: flex;
		margin-right: 200px;
		transition: 0.3s all ease-in-out;
	}

	.title {
		color: #222;
		font-weight: 600;
		transition: 0.3s all ease-in-out;
	}

	.subtitle {
		color: #222;
		/* color: rgba(255, 255, 255, 0.8); */
		transition: 0.3s all ease-in-out;
	}

	.line {
		background-color: #e9e9e9;
		/* background-color: rgba(255, 255, 255, 0.15); */
		margin-top: 20px;
		margin-bottom: 20px;
		height: 1px;
		transition: 0.3s all ease-in-out;
	}

	.info {
		color: #222;
		display: flex;
		justify-content: space-between;
		gap: 20px;
		color: #fff;
		font-weight: 600;
		white-space: nowrap;
		transition: 0.3s all ease-in-out;
	}

	.value {
		color: #222;
		border-radius: 100px;
		/* background-color: rgba(255, 255, 255, 0.15); */
		background-color: rgba(0, 0, 0, 0.07);
		/* background-color: #e4e4e4; */
		aspect-ratio: 1.75;
		justify-content: center;
		padding: 5px 14px;
		font: 12px Poppins, sans-serif;
		transition: 0.3s all ease-in-out;
	}

	.red {
		color: #ff0000;
	}

	.green {
		color: #00ff00;
	}

	.blue {
		color: #0000ff;
	}

	.content-container {
		display: flex;
		width: 100%;
		/* max-width: 1141px; */
		max-width: 100vw;
		flex-direction: column;
		align-items: center;
		/* margin-left: 47px; */
		/* margin: 10px 0 146px; */
	}

	@media (max-width: 991px) {
		.content-container {
			max-width: 100%;
			margin-bottom: 40px;
			margin-left: 7px;
		}
	}

	.welcome-message {
		color: #000;
		text-align: center;
		margin-top: 135px;
		width: 653px;
		font: 600 36px Poppins, sans-serif;
	}

	@media (max-width: 991px) {
		.welcome-message {
			max-width: 100%;
			margin-top: 40px;
		}
	}

	.tutorial-message {
		color: rgba(0, 0, 0, 0.6);
		text-align: center;
		margin-top: 15px;
		width: 553px;
		font: 400 16px/32px Poppins, sans-serif;
	}

	@media (max-width: 991px) {
		.tutorial-message {
			max-width: 100%;
		}
	}

	.feature-section {
		display: flex;
		margin-top: 40px;
		justify-content: space-between;
		gap: 20px;
		font-size: 16px;
		font-weight: 600;
		max-width: 100vw;
	}

	.feature-button {
		font-family: Poppins, sans-serif;
		border-radius: 100px;
		box-shadow: 0px 15px 30px 0px rgba(0, 26, 255, 0.2);
		background-color: #001aff;
		flex-grow: 1;
		justify-content: center;
		color: #fff;
		white-space: nowrap;
		padding: 12px 17px;
		transition: 0.9s;
	}

	.feature-button:hover {
		box-shadow: 0px 15px 30px 0px rgba(0, 26, 255, 0.2);
		background-color: #fff;
		flex-grow: 1;
		justify-content: center;
		color: #000fff;
		white-space: nowrap;
		padding: 12px 17px;
	}

	@media (max-width: 991px) {
		.feature-button {
			white-space: initial;
		}
	}

	.how-it-works {
		color: #001aff;
		text-decoration-line: underline;
		flex-grow: 1;
		flex-basis: auto;
		margin: auto 0;
		transition: 0.9s;
	}

	.how-it-works:hover {
		color: #222;
		text-decoration-line: underline;
		flex-grow: 1;
		flex-basis: auto;
		margin: auto 0;
		cursor: pointer;
	}
</style>
