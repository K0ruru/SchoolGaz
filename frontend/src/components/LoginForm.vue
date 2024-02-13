<!-- Login.vue -->
<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import Swal from "sweetalert2";
import FloatLabel from 'primevue/floatlabel';

const nis = ref("");
const pass = ref("");

const router = useRouter();

const login = async () => {
  try {
    const response = await axios.post("http://localhost:8080/Auth/login", {
      nis: nis.value,
      passphrase: pass.value,
    });

    const { token, status, nama, role } = response.data;

    if (status != "active") {
      // alert("Akun anda tidak aktif");
      Swal.fire("Akun anda tidak aktif", "", "warning");
      return;
    }

    localStorage.setItem("role", role);
    localStorage.setItem("token", token);
    localStorage.setItem("nama", nama);

    router.push("/");
    // alert("Login successful");
    Swal.fire("Login successful", "", "success");
  } catch (error) {
    if (error.response && error.response.status === 400) {
      // alert("NIS or passphrase is incorrect"); // Display alert for incorrect NIS or passphrase
      Swal.fire("NIS or passphrase is incorrect", "", "error");
    } else {
      console.error("Error logging in:", error);
    }
  }
};
</script>

<template>
  <div class="container-center">
    <form @submit.prevent="login">
      <h1>Login</h1>
      <div class="form-inputs">
        <div class="form-input">
          <FloatLabel>
          <label for="nis">NIS :</label>
          <input v-model="nis" type="number" id="nis" name="nis" />
          </FloatLabel>
        </div>
        <div class="form-input">
          <FloatLabel>
          <label for="pass">Passphrase :</label>
          <input v-model="pass" type="password" id="pass" name="pass" />
          </FloatLabel>
        </div>
      </div>
      <!-- <input type="submit" value="Log-in" /> -->
      <button class="button" type="submit">Log-in</button>
      <p>Belum punya akun? <router-link to="/signup">Sign-up</router-link></p>
    </form>
    <img src="../assets/vector.png" alt="" class="vector" />
  </div>
</template>

<style scoped>
h1 {
  text-align: center;
  margin-bottom: 20px;
}

.form-inputs {
  display: flex;
  flex-direction: column;
  gap: 35px;
  margin-bottom: 20px;
  /* align-items: center; */
}

.form-input {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  /* align-items: center; */
}


.form-input input {
  width: 100%;
  padding: 10px;
  /* margin-bottom: 10px; */
  border: 1px solid #ccc;
  border-radius: 5px;
}

.form-input input:focus::placeholder {
  color: #ccc;
}

.button {
  font-family: Poppins, sans-serif;
  border-radius: 100px;
  box-shadow: 0px 15px 30px 0px rgba(0, 26, 255, 0.2);
  background-color: #001aff;
  color: #fff;
  padding: 12px 17px;
  outline: none;
  border: none;
  width: 100%;
  margin-top: 10px;
  margin-bottom: 10px;
}

.container-center {
  gap: 5rem;
}

form {
  width: 20%;
}

.vector {
  width: 50%;
}

@media (max-width: 991px) {
  .button {
    white-space: initial;
  }
}
</style>
