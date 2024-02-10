<script setup lang="ts">
import Swal from "sweetalert2";
import Navbar from "./Navbar.vue";
import { ref, onMounted, computed } from "vue";

interface UserData {
  nis: number;
  name: string;
  passphrase: string;
  email: string;
  gender: string;
  religion: string;
  status: string;
}

const search = ref("");
const data = ref<Array<UserData>>([]);


const fetchData = async () => {
  try {
    const response = await fetch("http://localhost:8080/Auth/");
    const result = await response.json();
    data.value = result;
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};

onMounted(() => {
  fetchData();
});

const tambahData = () => {
  // Logic for adding data
};

const hapusData = async (nis: number) => {
  try {
    const result = await Swal.fire({
      title: "Are you sure?",
      text: "You will not be able to recover this user!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Yes, delete it!",
      cancelButtonText: "Cancel",
    });

    if (result.isConfirmed) {
      const response = await fetch(`http://localhost:8080/Auth/delete/${nis}`, {
        method: "DELETE",
      });

      if (response.ok) {
        // Remove the deleted user from the data array
        data.value = data.value.filter((userData) => userData.nis !== nis);
        Swal.fire("Deleted!", "User has been deleted.", "success");
        console.log(`User with NIS ${nis} deleted successfully.`);
      } else {
        console.error(`Failed to delete user with NIS ${nis}.`);
        Swal.fire("Error!", "Failed to delete user.", "error");
      }
    } else {
      // Handle cancellation
      console.log("Deletion cancelled.");
    }
  } catch (error) {
    console.error("Error deleting user:", error);
    Swal.fire("Error!", "Failed to delete user.", "error");
  }
};

const filteredData = computed(() => {
  return data.value.filter((userData) =>
    userData.name.toLowerCase().includes(search.value.toLowerCase()),
  );
});
</script>

<template>
  <div class="container">
    <Navbar />
    <div class="search">
      <input v-model="search" placeholder="Search..." class="input-search" />
      <button @click="tambahData" class="button-add">Tambah</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>NIS</th>
          <th>Nama</th>
          <th>E-mail</th>
          <th>No-Telp</th>
          <th>Jenis Kelamin</th>
          <th>Agama</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="userData in filteredData" :key="userData.nis">
          <td>{{ userData.nis }}</td>
          <td>{{ userData.name }}</td>
          <td>{{ userData.email }}</td>
          <td>{{ userData.gender }}</td>
          <td>{{ userData.religion }}</td>
          <td>{{ userData.status }}</td>
          <td>
            <button class="button-edit">Edit</button>
            <button class="button-delete" @click="hapusData(userData.nis)">
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
}

td {
  border-color: #fff;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}</style>

