<script setup lang="ts">
import Navbar from "./Navbar.vue";
import Swal from "sweetalert2";
import EditForm from "./EditForm.vue";
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

const hapusData = async (nis: number) => {
  try {
    const result = await Swal.fire({
      title: "Are you sure?",
      text: "You will not be able to recover this user!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Yes, delete it!",
      cancelButtonText: "Cancel",
      confirmButtonColor: "#ff0000",
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

const filterKey = ref("name");

const filteredData = computed(() => {
  return data.value.filter((userData) => {
    const key = filterKey.value;
    if (
      Object.prototype.hasOwnProperty.call(userData, key) &&
      typeof userData[key as keyof typeof userData] === "string"
    ) {
      return (userData[key as keyof typeof userData] as string)
        .toLowerCase()
        .includes(search.value.toLowerCase());
    }
    return false;
  });
});

const editFormVisible = ref(false);
const selectedUserData = ref<UserData | null>(null);

const showEditForm = (userData: UserData) => {
  // Assigning a new object to selectedUserData to ensure reactivity
  selectedUserData.value = { ...userData };
  editFormVisible.value = true;
};

const closeEditForm = () => {
  editFormVisible.value = false;
};
</script>

<template>
  <EditForm v-if="editFormVisible" :userData="selectedUserData" @closeEditForm="closeEditForm" />
  <div class="container">
    <Navbar />
    <div class="search">
      <input v-model="search" placeholder="Search..." class="input-search" />
      <div class="filter-dropdown">
        <label for="filterKey">Filter by:</label>
        <select v-model="filterKey" id="filterKey">
          <option value="nis">NIS</option>
          <option value="name">Name</option>
          <option value="email">E-mail</option>
          <option value="gender">Jenis Kelamin</option>
          <option value="religion">Agama</option>
          <option value="status">Status</option>
        </select>
      </div>
    </div>
    <table>
      <thead>
        <tr>
          <th>NIS</th>
          <th>Nama</th>
          <th>E-mail</th>
          <th>Jenis Kelamin</th>
          <th>Agama</th>
          <th>Status</th>
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
            <button class="button-edit" @click="showEditForm(userData)">
              Edit
            </button>
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
