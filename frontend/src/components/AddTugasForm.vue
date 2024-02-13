<script setup lang="ts">
import Navbar from "./Navbar.vue";
// import { ref } from "vue";
import Editor from "primevue/editor";
import Calendar from "primevue/calendar";
// const soalText = ref("");
</script>

<template>
  <div class="container">
    <Navbar />
    <div class="button-add">
      <button>Tambah Tugas</button>
    </div>
    <div class="form-container">
      <div class="form-inputs-container">
        <div class="form-inputs">
          <form action="#" method="POST">
            <label for="email">Judul Tugas</label>
            <input type="email" id="email" name="email" />

            <label for="fullName">Dead line</label>
            <!-- <input type="text" id="fullName" name="fullName" /> -->
            <Calendar showIcon :showOnFocus="false" />

            <label for="photo">File</label>
            <input type="file" id="photo" name="photo" />

            <button class="submit-btn" type="submit">Submit</button>
          </form>
        </div>
        <div class="editor">
          <label>Arahan Tugas</label>
          <Editor editorStyle="height: 270px" />
        </div>
      </div>
    </div>
    <!-- <textarea v-model="soalText" name="" id="" cols="30" rows="10" placeholder="Tulis Arahan"></textarea> -->
  </div>
</template>

<style scoped>
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



.editor {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.form-inputs-container {
  margin-top: 100px;
  display: flex;
  justify-content: center;
  /* align-items: center; */
  gap: 17px;
}

label,
button {
  font-size: 16px;
  cursor: pointer;
  margin-top: 10px;
}

input,
textarea {
  border: 3px solid #eee;
  padding: 0.4em 1em;
  border-radius: 6px;
}

form {
  display: grid;
  grid-row-gap: 10px;
  justify-content: center;
}

.submit-btn {
  background: #0000ff;
  padding: 9px 82px;
  font-size: 14px;
  font-weight: 600;
  color: #fff;
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

