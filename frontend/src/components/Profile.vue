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
</script>

<template>
	<div class="container">
		<Navbar />
		<div class="profile-content">
			<div class="profile-information">
				<img src="../assets/Doge hehe.jpg" alt="" class="profile-picture" />
				<p class="nis">001002003</p>
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
					<p>Role :</p>
					<p>Siswa</p>
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
		background-color: rgb(238, 238, 238);
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
		background-color: rgb(238, 238, 238);
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
