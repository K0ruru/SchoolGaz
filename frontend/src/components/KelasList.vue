<script setup lang="ts">
import Navbar from "./Navbar.vue";
import axios from "axios";
import { ref } from "vue";

interface Kelas {
  Id_kelas: number;
  NamaKelas: string;
  Walas: {
    NIS: number;
    NamaGuru: string;
  };
}

const kelasData = ref<Kelas[]>([]);

axios
  .get("http://localhost:8080/kelas/")
  .then((response) => {
    kelasData.value = response.data;
  })
  .catch((error) => {
    console.error("Error fetching kelas data:", error);
  });
</script>

<template>
  <div class="container">
    <Navbar />
    <div class="button-add" @click="showAddKelasForm">
      <button>Tambah Tugas</button>
    </div>

    <div class="kelas-content">
      <router-link v-for="kelas in kelasData" :key="kelas.Id_kelas"
        :to="{ name: 'kelas', params: { id: kelas.Id_kelas } }" class="kelas-card" style="text-decoration: none">
        <div class="kelas-card-content">
          <img src="../assets/Doge hehe.jpg" alt="" class="pp-walas" />
          <div class="kelas-info">
            <h1>{{ kelas.NamaKelas }}</h1>
            <p>{{ kelas.Walas.NamaGuru }}</p>
          </div>
          <div class="setting">
            <ion-icon name="ellipsis-vertical"></ion-icon>
          </div>
        </div>
        <div class="list"></div>
      </router-link>
    </div>
  </div>
</template>

<style scoped>

.setting {
  padding-left: 20px;
  margin-left: 40px;
}
.button-add {
  display: flex;
  width: 100%;
  justify-content: flex-end;
}

.button-add button {
  margin-top: 20px;
  background-color: blue;
  border: none;
  border-radius: 14px;
  color: #fff;
  padding: 10px;
  font-weight: 600;
}

.button-add button:hover {
  margin-top: 20px;
  background-color: darkblue;
  border: none;
  border-radius: 14px;
  color: #fff;
  padding: 10px;
  font-weight: 600;
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
  height: 168px;
  border-radius: 7px;
  margin: 30px 45px;
  transition: 0.3s all ease-in-out;
  width: 310px;
}

.kelas-card-content {
  display: flex;
  /* align-items: center; */
  justify-content: center;
  margin: 17px 0px;
  margin-left: 10px;
}

.pp-walas {
  max-width: 100px;
  border-radius: 50px;
  margin-right: 15px;
}

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
</style>
