<script setup lang="ts">
	import { defineEmits, onMounted } from "vue";
	import axios from "axios";
	import { ref } from "vue";

	const emits = defineEmits(["closeFloatingKelas"]);
	interface Kelas {
		Id_kelas: number;
		WalasNIS: number;
		NamaKelas: string;
		Walas: {
			NamaGuru: string;
		};
	}

	const kelasData = ref<Kelas[]>([]);

	onMounted(async () => {
		try {
			const response = await axios.get("http://localhost:8080/kelas/");
			kelasData.value = response.data;
		} catch (error) {
			console.error("Error fetching kelas data:", error);
		}
	});
</script>
<template>
	<div class="floating-overlay">
		<div class="floating-container fade-in">
			<div class="close-button">
				<button @click="emits('closeFloatingKelas')" class="button-close">
					✖
				</button>
			</div>
			<div class="kelas-content">
				<router-link
					v-for="kelas in kelasData"
					style="text-decoration: none; color: inherit"
					:key="kelas.Id_kelas"
					:to="{ name: 'addtugasform', params: { id: kelas.Id_kelas } }"
					class="kelas-card"
				>
					<div class="kelas-card-content">
						<div class="kelas-info">
							<h1>{{ kelas.NamaKelas }}</h1>
							<p>{{ kelas.Walas.NamaGuru }}</p>
						</div>
					</div>
					<div class="list"></div>
				</router-link>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.fade-in {
		animation: fadeIn 0.5s ease-in-out;
	}

	@keyframes fadeIn {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
	.floating-overlay {
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

	.floating-container {
		width: 80%;
		height: 80%;
		background: #fff;
		padding: 20px;
		border-radius: 8px;
		box-shadow: 0 0 10px rgba(0, 0, 255, 0.1);
	}

	.kelas-content {
		display: flex;
		justify-content: space-evenly;
		max-width: 100%;
		flex-wrap: wrap;
		margin-top: 50px;
	}

	.kelas-card {
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.2);
		display: flex;
		flex-direction: column;
		height: 100px;
		border-radius: 7px;
		margin: 30px 45px;
		transition: 0.3s all ease-in-out;
		width: 310px;
	}

	.kelas-card-content {
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 17px 33px;
	}

	/* .pp-walas { */
	/*   max-width: 100px; */
	/*   border-radius: 50px; */
	/*   margin-right: 15px; */
	/* } */

	.kelas-card a {
		text-decoration: none;
		color: inherit;
	}

	.kelas-card h1 {
		font-size: 25px;
		font-weight: 700;
		transition: 0.3s all ease-in-out;
		color: black;
		text-decoration: none !important;
	}

	.kelas-card p {
		font-size: 10px;
		margin-top: -8px;
		color: grey;
		transition: 0.3s all ease-in-out;
		text-decoration: none !important;
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

	.container-center {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
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
		cursor: pointer;
	}

	.button-close {
		background-color: rgba(1, 1, 1, 0);
		border: none;
		padding: 2px 10px;
		/* border-radius: 100px; */
		font-size: 17px;
	}

	.close-button {
		display: flex;
		justify-content: flex-end;
	}
</style>
