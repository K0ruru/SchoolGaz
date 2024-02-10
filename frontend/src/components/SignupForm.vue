<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";

const nis = ref("");
const nama = ref("");
const passphrase = ref(""); // Change "pass" to "passphrase"
const passConfirm = ref("");
const email = ref("");
const noTelp = ref("");
const jenkel = ref("");
const agama = ref("");

const router = useRouter();

const signUp = async () => {
  if (passphrase.value.length < 8) {
    console.error("Password must be at least 8 characters long");
    return;
  }

  if (passphrase.value !== passConfirm.value) {
    console.error("Passwords do not match");
    return;
  }

  try {
    const response = await axios.post(
      "http://localhost:8080/Auth/add",
      {
        nis: nis.value,
        name: nama.value,
        passphrase: passphrase.value,
        email: email.value,
        no_telp: noTelp.value,
        gender: jenkel.value,
        religion: agama.value,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      },
    );

    console.log(response.data);

    router.push("/login");
  } catch (error) {
    console.error("Error signing up:");
  }
};
</script>

<template>
  <div class="container-center">
    <form @submit.prevent="signUp">
      <h1>Sign-Up</h1>
      <div class="form-inputs">
        <div class="form-input">
          <label for="nis">NIS :</label>
          <input v-model="nis" type="number" id="nis" name="nis" />
        </div>
        <div class="form-input">
          <label for="nama">Nama :</label>
          <input v-model="nama" type="text" id="nama" name="nama" />
        </div>
        <div class="form-input">
          <label for="pass">Passphrase :</label>
          <input v-model="passphrase" type="password" id="pass" name="pass" />
        </div>
        <div class="form-input">
          <label for="pass-confirm">Confirm Passphrase :</label>
          <input v-model="passConfirm" type="password" id="pass-confirm" name="pass-confirm" />
        </div>
        <div class="form-input">
          <label for="email">E-mail :</label>
          <input v-model="email" type="email" id="email" name="email" />
        </div>
        <div class="form-input">
          <label for="no-telp">No-Telp :</label>
          <input v-model="noTelp" type="number" id="no-telp" name="no-telp" />
        </div>
        <div class="form-input">
          <label for="jenkel">Jenis Kelamin :</label>
          <input v-model="jenkel" type="text" id="jenkel" name="jenkel" />
        </div>
        <div class="form-input">
          <label for="agama">Agama :</label>
          <select v-model="agama" id="agama" name="agama">
            <option value="Islam">Islam</option>
            <option value="Kristen Protestan">Kristen Protestan</option>
            <option value="Kristen Katolik">Kristen Katolik</option>
            <option value="Hindu">Hindu</option>
            <option value="Buddha">Buddha</option>
            <option value="Khonghucu">Khonghucu</option>
          </select>
        </div>
      </div>
      <button class="button" type="submit">Sign-Up</button>
      <p>Sudah punya akun? <router-link to="/login">Login</router-link></p>
    </form>
    <img src="../assets/vector.png" alt="" class="vector">
  </div>
</template>

<style scoped>
.container-center {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

form {
  max-width: 400px;
  width: 100%;
}

.form-inputs {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 15px;
}

.form-input {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.form-input label {
  margin-bottom: 10px;
}

.form-input input {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.form-input input:focus {
  outline: none;
}

.form-input input:focus::placeholder {
  color: #ccc;
}

form {
  width: 30%;
}

.container-center {
  gap: 5rem;
}

  .vector {
    width: 50%;
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
  cursor: pointer;
}
</style>
