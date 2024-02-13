<script setup lang="ts">
	import Navbar from "./Navbar.vue";
	import GuruEditForm from "./EditFormGuru.vue";
	import Toast from "primevue/toast";
	import ConfirmPopup from "primevue/confirmpopup";
	import axios from "axios";
	import { ref, onMounted, computed } from "vue";
	import { useConfirm } from "primevue/useconfirm";
	import { useToast } from "primevue/usetoast";

	interface GuruData {
		NIS: number;
		NamaGuru: string;
		Email: string;
		Gender: string;
		Religion: string;
		NoTelp: number;
		Status: string;
		// Mapel: string;
		[key: string]: number | string;
	}

	const search = ref("");
	const data = ref<Array<GuruData>>([]);
	const sorting = ref<{ column: string; direction: string | null }>({
		column: "NIS",
		direction: null,
	});

	const fetchData = async () => {
		try {
			const response = await fetch("http://localhost:8080/guru/");
			const result = await response.json();
			data.value = result;
		} catch (error) {
			console.error("Error fetching data:", error);
		}
	};

	onMounted(() => {
		fetchData();
	});

	const confirm = useConfirm();
	const toast = useToast();

	const hapusData = (event: MouseEvent, nis: number) => {
		confirm.require({
			target: event.currentTarget as HTMLElement,
			message: "Apa anda yaking ingin menghapus ini?",
			icon: "pi pi-info-circle",
			rejectClass: "p-button-secondary p-button-outlined p-button-sm",
			acceptClass: "p-button-danger p-button-sm",
			rejectLabel: "Cancel",
			acceptLabel: "Delete",
			accept: async () => {
				try {
					const response = await axios.delete(
						`http://localhost:8080/guru/del/${nis}`
					);

					if (response.status === 200) {
						data.value = data.value.filter((userData) => userData.NIS !== nis);
						toast.add({
							severity: "info",
							summary: "Confirmed",
							detail: "Data dihapus",
							life: 3000,
						});
						console.log(`User with NIS ${nis} deleted successfully.`);
					} else {
						console.error(`Failed to delete user with NIS ${nis}.`);
						toast.add({
							severity: "error",
							summary: "Error",
							detail: "Failed to delete user.",
							life: 3000,
						});
					}
				} catch (error) {
					console.error("Error deleting user:", error);
					toast.add({
						severity: "error",
						summary: "Error",
						detail: "Failed to delete user.",
						life: 3000,
					});
				}
			},
			reject: () => {},
		});
	};

	const editFormVisible = ref(false);
	const selectedGuruData = ref<GuruData | null>(null);

	const showEditForm = (guruData: GuruData) => {
		selectedGuruData.value = { ...guruData };
		editFormVisible.value = true;
	};

	const closeEditForm = () => {
		editFormVisible.value = false;
	};

	const filteredData = computed(() => {
		const searchTerm = search.value.toLowerCase();
		return data.value.filter(
			(guruData) =>
				guruData.NIS.toString().includes(searchTerm) ||
				guruData.NamaGuru.toLowerCase().includes(searchTerm) ||
				guruData.Email.toLowerCase().includes(searchTerm) ||
				guruData.Gender.toLowerCase().includes(searchTerm) ||
				guruData.Religion.toLowerCase().includes(searchTerm) ||
				guruData.NoTelp.toString().includes(searchTerm) // Corrected property name
			// guruData.Mapel.toLowerCase().includes(searchTerm)
		);
	});

	const sortTable = (column: string) => {
		const isAsc =
			sorting.value.column === column && sorting.value.direction === "asc";
		sorting.value.column = column;
		sorting.value.direction = isAsc ? "desc" : "asc";

		data.value.sort((a, b) => {
			const modifier = isAsc ? 1 : -1;
			if (a[column] < b[column]) return -modifier;
			if (a[column] > b[column]) return modifier;
			return 0;
		});
	};
</script>

<template>
	<GuruEditForm
		v-if="editFormVisible"
		:guruData="selectedGuruData"
		@closeEditForm="closeEditForm"
	/>
	<Toast />
	<ConfirmPopup></ConfirmPopup>
	<div class="container">
		<Navbar />
		<div class="search">
			<input v-model="search" placeholder="Search..." class="input-search" />
		</div>
		<table>
			<thead>
				<tr>
					<th @click="sortTable('NIS')">NIS</th>
					<th @click="sortTable('NamaGuru')">Nama</th>
					<th @click="sortTable('Email')">E-mail</th>
					<th @click="sortTable('NoTelp')">No-Telpon</th>
					<th @click="sortTable('Gender')">Jenis Kelamin</th>
					<th @click="sortTable('Religion')">Agama</th>
					<!-- <th @click="sortTable('Status')">Status</th> -->
					<th>Action</th>
				</tr>
			</thead>
			<tbody>
				<tr v-for="guruData in filteredData" :key="guruData.NIS">
					<td>{{ guruData.NIS }}</td>
					<td>{{ guruData.NamaGuru }}</td>
					<td>{{ guruData.Email }}</td>
					<td>{{ guruData.NoTelp }}</td>
					<td>{{ guruData.Gender }}</td>
					<td>{{ guruData.Religion }}</td>
					<!-- <td>{{ userData.Status }}</td> -->
					<td>
						<button class="button-edit" @click="showEditForm(guruData)">
							Edit
						</button>
						<button
							class="button-delete"
							@click="hapusData($event, guruData.NIS)"
						>
							Hapus
						</button>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
</template>

<style scoped>
	table {
		width: 100%;
		border-collapse: collapse;
		margin-top: 50px;
		box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.2);
	}

	th,
	td {
		border: 1px solid #ddd;
		padding: 30px 15px;
		text-align: center;
		border: none;
	}

	th:first-child {
		width: 10%; /* NIS */
	}

	th:nth-child(2) {
		width: 15%; /* Nama */
	}

	th:nth-child(3),
	td:nth-child(3) {
		width: 15%; /* Email */
	}

	th:nth-child(4),
	td:nth-child(4) {
		width: 10%; /* No-Telpon */
	}

	th:nth-child(5) {
		width: 10%; /* Kelas */
	}

	th:nth-child(6),
	td:nth-child(6) {
		width: 10%; /* Jenis Kelamin */
	}

	th:nth-child(7),
	td:nth-child(7) {
		width: 10%; /* Agama */
	}

	th:nth-child(8),
	td:nth-child(8) {
		width: 10%; /* Status */
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

	.search {
		margin-top: 50px;
	}

	.input-search {
		padding: 10px;
		margin-right: 25px;
	}

	.button-add {
		padding: 9px;
		background-color: #0000ff;
		color: #fff;
		border: none;
		border-radius: 7px;
		font-weight: 600;
	}

	.button-delete {
		border: none;
		outline: none;
		padding: 7px;
		background-color: #ff0000;
		color: #fff;
		margin-left: 10px;
		border-radius: 4px;
	}

	.button-edit {
		border: none;
		outline: none;
		padding: 7px;
		background-color: #222;
		color: #fff;
		border-radius: 4px;
	}

	th {
		background-color: #0000ff;
		color: #fff;
		cursor: pointer;
		transition: 0.3s all ease-in-out;
	}

	th:hover {
		background-color: rgb(0, 0, 211);
	}

	td {
		border-color: #fff;
	}

	tr:nth-child(even) {
		background-color: #f2f2f2;
	}

	label {
		display: block;
		margin-top: 3px;
		margin-bottom: 2px;
		font-size: 12px;
	}

	select {
		width: 100xp;
		padding: 8px;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
		font-size: 15px;
		cursor: pointer;
	}
</style>
