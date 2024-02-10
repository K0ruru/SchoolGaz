<script setup lang="ts">
	import Navbar from "./Navbar.vue";
	import axios from "axios";
	import { ref, onMounted } from "vue";

	interface Kelas {
		id_kelas: number;
		walas: number;
		nama_kelas: string;
		nama_walas: string;
	}

	const kelasData = ref<Kelas[]>([]);

	onMounted(() => {
		axios
			.get("http://localhost:8080/kelas/show", {
				headers: {
					"Access-Control-Allow-Origin": "http://localhost:7070", // Make sure to replace with your frontend URL
					"Content-Type": "application/json",
				},
			})
			.then((response) => {
				kelasData.value = response.data;
			})
			.catch((error) => {
				console.error("Error fetching kelas data:", error);
			});
	});
</script>

<template>
	<div class="container">
		<Navbar />
		<div class="kelas-content">
			<div v-for="kelas in kelasData" :key="kelas.id_kelas" class="kelas-card">
				<div class="kelas-card-content">
					<img src="../assets/Doge hehe.jpg" alt="" class="pp-walas" />
					<div class="kelas-info">
						<h1>{{ kelas.nama_kelas }}</h1>
						<p>{{ kelas.nama_walas }}</p>
					</div>
				</div>
				<div class="list"></div>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.kelas-content {
		display: flex;
		justify-content: space-evenly;
		max-width: 100%;
		flex-wrap: wrap;
		margin-top: 50px;
	}

	.kelas-card {
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.3);
		display: flex;
		flex-direction: column;
		height: 168px;
		border-radius: 7px;
		margin: 30px 45px;
		transition: 0.3s all ease-in-out;
	}

	.kelas-card-content {
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 17px 33px;
	}

	.pp-walas {
		max-width: 100px;
		border-radius: 50px;
		margin-right: 15px;
	}

	.kelas-card h1 {
		font-size: 25px;
		font-weight: 700;
		transition: 0.3s all ease-in-out;
	}

	.kelas-card p {
		font-size: 10px;
		margin-top: -8px;
		color: grey;
		transition: 0.3s all ease-in-out;
	}

	.list {
		width: 100%;
		height: 100%;
		background-color: #0000ff;
		display: flex;
		justify-self: flex-end;
		border-bottom-left-radius: 7px;
		border-bottom-right-radius: 7px;
		transition: 0.3s all ease-in-out;
	}

	.kelas-card:hover {
		background-color: #0000ff;
	}

	.kelas-card:hover .list {
		background-color: #0000b4;
	}

	.kelas-card:hover h1 {
		color: #fff;
	}

	.kelas-card:hover p {
		color: rgb(231, 231, 231);
	}
</style>
