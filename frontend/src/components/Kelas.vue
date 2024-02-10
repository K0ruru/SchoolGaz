<script setup lang="ts">
	import Navbar from "./Navbar.vue";
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
						<tr>
							<td>1</td>
							<td><img src="../assets/Doge hehe.jpg" alt="" /></td>
							<td>Nurfaiz Muhammad Qalbi</td>
							<td>001002003</td>
							<td>nurfaizqalbi@gmail.com</td>
							<td>Pria</td>
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
