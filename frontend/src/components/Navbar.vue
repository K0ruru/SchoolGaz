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
		localStorage.removeItem("nama");
		router.push("/login");
	};
</script>

<template>
	<div class="header-container">
		<div class="brand-logo" @click="navigasi('')">
			<img src="../assets/SG_.png" alt="SG" />
		</div>
		<div class="menu-container">
			<div class="menu-items">
				<a href="#" class="menu-item"
					><ion-icon name="ellipse"></ion-icon> Online</a
				>
				<!-- <a href="#" class="menu-item" :class="{ active: tabAktif === 'Dashboard' }" @click="navigasi('')">Home</a> -->
				<!-- <a href="#" class="menu-item" :class="{ active: tabAktif === 'siswa' }" @click="navigasi('siswa')">Students</a> -->
				<!-- <a href="#" class="menu-item" :class="{ active: tabAktif === 'guru' }" @click="navigasi('guru')">Teachers</a> -->
				<!-- <a href="#" class="menu-item" :class="{ active: tabAktif === 'tugas' }" @click="navigasi('tugas')">Assignments</a> -->
			</div>
		</div>
		<!-- <div class="profile-section">PROFILE</div> -->
		<div class="profile-dropdown">
			<div class="profile-section">PROFILE</div>
			<div class="dropdown-content">
				<a href="#">Profile</a>
				<a href="#">Settings</a>
				<a href="#" @click="logout">Logout</a>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.profile-dropdown {
		position: relative;
		display: inline-block;
	}

	.profile-section {
		background-color: #0000ff;
		color: #fff;
		padding: 10px;
		cursor: pointer;
	}

	.dropdown-content {
		display: none;
		position: absolute;
		background-color: #f9f9f9;
		min-width: 113px;
		/* box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2); */
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
		z-index: 1;
	}

	.dropdown-content a {
		color: #333;
		padding: 12px 16px;
		text-decoration: none;
		display: block;
		text-align: center;
		transition: 0.2s;
	}

	.dropdown-content a:hover {
		background-color: #0000ff;
		color: #fff;
	}

	.profile-dropdown:hover .dropdown-content {
		display: block;
	}

	.header-container {
		align-self: stretch;
		display: flex;
		justify-content: space-between;
		gap: 20px;
		font-size: 16px;
		color: #000;
		font-weight: 600;
		white-space: nowrap;
		padding: 0 1px;
	}

	@media (max-width: 991px) {
		.header-container {
			max-width: 100%;
			flex-wrap: wrap;
			white-space: initial;
		}
	}

	.brand-logo {
		align-self: start;
		margin-top: 8px;
		flex-grow: 1;
		flex-basis: auto;
	}

	.brand-logo img {
		width: 150px;
		margin-bottom: 7px;
	}
	.brand-logo:hover img {
		width: 150px;
		margin-bottom: 7px;
	}

	.menu-container {
		justify-content: center;
		align-items: center;
		display: flex;
		flex-grow: 1;
		flex-basis: 0%;
		color: rgba(0, 0, 0, 0.8);
		font-weight: 400;
		margin: auto 0;
		padding: 0 60px;
	}

	@media (max-width: 991px) {
		.menu-container {
			max-width: 100%;
			white-space: initial;
			padding: 0 20px;
		}
	}

	.menu-items {
		display: flex;
		width: 333px;
		max-width: 100vw;
		justify-content: center;
		gap: 20px;
		/* margin-right: 290px; */
		position: absolute;
		/* top: 50%; */
		left: 50%;
		transform: translateX(-50%);
	}

	@media (max-width: 991px) {
		.menu-items {
			white-space: initial;
			/* margin-right: 290px; */
			position: absolute;
			/* top: 50%; */
			left: 50%;
			transform: translateX(-50%) translateY(-190%);
		}
	}

	.menu-item {
		color: #000;
		font-family: Poppins, sans-serif;
		font-weight: 500;
		text-decoration: none;
		cursor: pointer;
	}

	.profile-section {
		font-family: Poppins, sans-serif;
		justify-content: center;
		border-radius: 3px;
		border: 1px solid rgba(0, 0, 0, 0.2);
		padding: 10px 25px;
		margin-top: 10px;
	}

	@media (max-width: 991px) {
		.profile-section {
			white-space: initial;
			/* padding: 0 20px; */
		}
	}
</style>
