
<script setup lang="ts">
import Navbar from "./Navbar.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
import FloatingKelas from "./FloatingKelas.vue";

const tabAktif = ref("");
const router = useRouter();

const navigasi = (tab: string) => {
  tabAktif.value = tab;
  router.push(`/${tab}`);
};

const tugasForm =ref(false);

const showFloatingKelas =()=>{
  tugasForm.value = true;
}
const closeFloatingKelas =()=>{
  tugasForm.value = false;
}


</script>

<template>
  <FloatingKelas v-if="tugasForm" @closeFloatingKelas='closeFloatingKelas'/>
  <div class="container">
    <Navbar />
      <div class="button-add" @click="showFloatingKelas">
        <button>Tambah Tugas</button>
      </div>
    <div class="cont">
      <!-- box untuk siswa -->
      <div class="box" @click="navigasi('tugas')">
        <div class="content">
          <!-- <div class="icon-1"><ion-icon name="people-outline"></ion-icon></div> -->
          <div class="text">
            <div class="title">PBO Inheritence</div>
            <div class="mapel">Pemograman Berorientasi Objek</div>
            <div class="deadline"><p>Dea Line 23/02/24</p></div>
          </div>
          <div class="info">
            <div class="kelas">TKI 1</div>
          </div>
        </div>
      </div>


      <div class="box" @click="navigasi('')">
        <div class="content">
          <!-- <div class="icon-1"><ion-icon name="people-outline"></ion-icon></div> -->
          <div class="text">
            <div class="title">PBO Inheritence</div>
            <div class="mapel">Pemongraman Berorientasi Objek</div>
            <div class="deadline"><p>Dea Line 23/02/24</p></div>
          </div>
          <div class="info">
            <div class="kelas">TKI 1</div>
          </div>
        </div>
      </div>

      <div class="box" @click="navigasi('')">
        <div class="content">
          <!-- <div class="icon-1"><ion-icon name="people-outline"></ion-icon></div> -->
          <div class="text">
            <div class="title">PBO Inheritence</div>
            <div class="mapel">Pemongraman Berorientasi Objek</div>
            <div class="deadline"><p>Dea Line 23/02/24</p></div>
          </div>
          <div class="info">
            <div class="kelas">TKI 1</div>
          </div>
        </div>
      </div>

      <div class="box" @click="navigasi('')">
        <div class="content">
          <!-- <div class="icon-1"><ion-icon name="people-outline"></ion-icon></div> -->
          <div class="text">
            <div class="title">PBO Inheritence</div>
            <div class="mapel">Pemongraman Berorientasi Objek</div>
            <div class="deadline"><p>Dea Line 23/02/24</p></div>
          </div>
          <div class="info">
            <div class="kelas">TKI 1</div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
.cont {
  display: flex;
  gap: 20px;
  margin-top: 40px;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
}

@media (max-width: 991px) {
  .cont {
    flex-direction: column;
  }
}

.box {
  border-radius: 8px;
  box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
  /* background-color: #fff; */
  display: flex;
  flex-direction: column;
  width: 45%;
  padding: 25px;
  height: 150px;
  transition: 0.3s all ease-in-out;
  cursor: pointer;
  justify-content: center;
}

.box:hover {
  border-radius: 8px;
  box-shadow: 0 30px 60px 0 rgba(0, 26, 255, 0.1);
  background-color: #001aff;
  display: flex;
  flex-direction: column;
  min-width: 30%;
  padding: 25px;
  height: 170px;
  z-index: 10;
}

.box:hover .title,
.box:hover .deadline,
.box:hover .kelas,
.box:hover .mapel {
  color: #fff;
}


.content {
  display: flex;
  justify-content: space-between;
  gap: 13px;
}

.text {
  /* display: flex; */
  flex-direction: column;
  /* min-width: 70%; */
  display: flex;
  margin-right: 200px;
  transition: 0.3s all ease-in-out;
}

.title {
  color: #222;
  font-weight: 600;
  transition: 0.3s all ease-in-out;
}

.mapel {
  color: #222;
  /* color: rgba(255, 255, 255, 0.8); */
  transition: 0.3s all ease-in-out;
}

.deadline p{
  background-color: rgba(255, 0, 0, 1);
  color: #fff;
  border-radius: 100px;
  padding: 5px 5px;
  font-size: 10px;
  width: 120px;
  text-align: center;
  margin-top: 10px;
}


.button-add {
  display: flex;
  width: 100%;
  justify-content: flex-end;
}

.button-add  button{
  margin-top: 20px;
  background-color: blue;
  border: none;
  border-radius: 14px;
  color: #fff;
  padding: 10px;
  font-weight: 600;
}
.button-add  button:hover{
  margin-top: 20px;
  background-color: darkblue;
  border: none;
  border-radius: 14px;
  color: #fff;
  padding: 10px;
  font-weight: 600;
}

.kelas {
  font-size: 10px;
}

</style>
