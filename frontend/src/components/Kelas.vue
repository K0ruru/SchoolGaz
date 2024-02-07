<script setup lang="ts">
import Navbar from "./Navbar.vue";
</script>

<!-- <template> -->
<!--   <div class="container"> -->
<!--     <Navbar /> -->
<!--     <div class="kelas-content"></div> -->
<!--   </div> -->
<!-- </template> -->
<!---->
<!-- <style scoped></style> -->
<template>
  <div class="container">
  <Navbar />
    <div>
      <input v-model="search" placeholder="Search...">
      <button @click="tambahData">Tambah</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>NIS</th>
          <th>Nama</th>
          <th>Passphrase</th>
          <th>E-mail</th>
          <th>No-Telp</th>
          <th>Jenis Kelamin</th>
          <th>Agama</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="data in filteredData" :key="data.NIS">
          <td>{{ data.NIS }}</td>
          <td>{{ data.Nama }}</td>
          <td>{{ data.Passphrase }}</td>
          <td>{{ data.ConfirmPassphrase }}</td>
          <td>{{ data.Email }}</td>
          <td>{{ data.NoTelp }}</td>
          <td>{{ data.JenisKelamin }}</td>
          <td>{{ data.Agama }}</td>
          <td>
            <button @click="editData(data.NIS)">Edit</button>
            <button @click="hapusData(data.NIS)">Hapus</button>
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
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}
</style>

