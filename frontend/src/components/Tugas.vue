<script setup lang="ts">
	import Navbar from "./Navbar.vue";
</script>

<template>
	<div class="container">
		<Navbar />
		<div class="tugas-content">
			<div class="tugas-container">
				<div class="tugas-heading">
					<div class="tugas-info">
						<h4>Matematika</h4>
						<h1>Fungsi CIN, COS, TAN</h1>
						<p>Deadline 12/02/24</p>
					</div>
					<img src="../assets/Doge hehe.jpg" alt="" class="pp-guru" />
				</div>
				<div class="tugas">
					<h1>Soal</h1>
					<div class="container-content">
						<div class="soal-container">
							<p>
								Kerjakanlah soal di buku matematika kalian hal 201 mengenai CIN,
								COS, TAN buat soalnya dibuku lalu kerjakan dengan membut alur
								cara mengerjakan nya...
							</p>
							<button class="download-attachment">
								<ion-icon name="cloud-download"></ion-icon>
								Arahan tugas.pdf
							</button>
						</div>
						<form class="jawaban-container">
							<div class="jawaban">
								<textarea
									name=""
									id=""
									cols="30"
									rows="10"
									placeholder="Tulis Jawaban..."
								></textarea>
							</div>
							<div class="jawaban-btns">
								<input type="submit" class="submit-btn" />
								<button class="upload-attachment">
									<ion-icon name="cloud-upload"></ion-icon>
								</button>
							</div>
						</form>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
	.tugas-content {
		width: 100vw;
		height: 80vh;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.tugas-container {
		display: flex;
		flex-direction: column;
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.3);
		padding: 52px 50px;
		max-width: 1050px;
		border-radius: 7px;
	}

	.tugas-heading {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 35px;
	}

	.tugas-info h4 {
		font-size: 14px;
		color: grey;
		font-weight: 500;
	}

	.tugas-info h1 {
		font-size: 30px;
		font-weight: 700;
		margin-bottom: 5px;
	}

	.tugas-info p {
		font-size: 12px;
		font-weight: 300;
		color: white;
		background-color: #ff0000;
		max-width: 150px;
		text-align: center;
		padding: 3px 10px;
		border-radius: 25px;
	}

	.pp-guru {
		max-width: 100px;
		border-radius: 50px;
	}

	.tugas {
		background-color: #0000ff;
		display: flex;
		flex-direction: column;
		padding: 34px;
		border-radius: 7px;
	}

	.tugas h1 {
		font-size: 30px;
		font-weight: 700;
		color: #fff;
		margin-bottom: 5px;
	}

	.tugas-content {
		max-width: 100%;
	}

	.container-content {
		display: flex;
		justify-content: space-between;
	}

	.soal-container {
		display: flex;
		flex-direction: column;
		max-width: 40%;
		align-items: center;
		justify-content: center;
	}

	.soal-container p {
		font-size: 14px;
		text-align: justify;
		font-weight: 400;
		margin-bottom: 50px;
		color: white;
	}

	.download-attachment {
		background-color: #00ff75;
		border: none;
		display: flex;
		align-items: center;
		font-size: 10px;
		font-weight: 500;
		gap: 20px;
		padding: 5px 45px;
		border-radius: 20px;
		margin-bottom: 10px;
	}

	.download-attachment ion-icon {
		font-size: 20px;
	}

	.jawaban-container {
		width: 50%;
	}

	.jawaban {
		width: 100%;
		padding: 7px 10px;
		background-color: white;
		margin-bottom: 10px;
	}

	.jawaban textarea {
		width: 100%;
		height: 250px;
		background-color: none;
		border: none;
		resize: none;
		font-size: 14px;
		font-weight: 500;
	}

	.jawaban textarea:focus {
		outline: none;
	}

	.jawaban-btns {
		display: flex;
		justify-content: space-between;
	}

	.submit-btn {
		padding: 9px 82px;
		font-size: 14px;
		font-weight: 600;
		color: #0000ff;
		border-radius: 25px;
		border: none;
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.3);
		cursor: pointer;
		transition: 0.3s all ease-in-out;
	}

	.submit-btn:hover {
		background-color: #d6d6d6;
	}

	.upload-attachment {
		font-size: 20px;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 10px;
		border-radius: 100px;
		border: none;
		transition: 0.3s all ease-in-out;
		color: #0000ff;
	}

	.upload-attachment:hover {
		background-color: #d6d6d6;
	}
</style>
