<script setup lang="ts">
	import { ref } from "vue";
	import { useRouter } from "vue-router";

	const tabAktif = ref("");
	const router = useRouter();

	const navigasi = (tab: string) => {
		tabAktif.value = tab;
		router.push(`/${tab}`);
	};

	const logout = () => {
		localStorage.removeItem("token");
		router.push("/login");
	};
</script>

<template>
	<nav>
		<div class="container">
			<div class="nav-content">
				<div class="nav-logo">
					<h1>SkoolGaz</h1>
				</div>
				<ul class="nav-links">
					<li
						:class="{ 'nav-link': true, active: tabAktif === 'Dashboard' }"
						@click="navigasi('')"
					>
						Home
					</li>
					<li
						:class="{ 'nav-link': true, active: tabAktif === 'kelas' }"
						@click="navigasi('kelas')"
					>
						Kelas
					</li>
					<li
						:class="{ 'nav-link': true, active: tabAktif === 'tugas' }"
						@click="navigasi('tugas')"
					>
						Tugas
					</li>
				</ul>
				<button class="logout" @click="logout">Log-out</button>
			</div>
		</div>
	</nav>
</template>

<style scoped>
	nav {
		background-color: #333;
		color: white;
	}

	.container {
		max-width: 1200px;
		margin: 0 auto;
	}

	.nav-content {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 10px 0;
	}

	.nav-logo h1 {
		margin: 0;
	}

	.nav-links {
		list-style-type: none;
		display: flex;
	}

	.nav-link {
		cursor: pointer;
		padding: 10px;
	}

	.active {
		font-weight: bold;
		border-bottom: 2px solid white;
	}

	.logout {
		border: none;
		background-color: red;
		color: white;
		padding: 10px 20px;
	}
</style>
