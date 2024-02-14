<script setup lang="ts">
// import { ref } from "vue";
// const soalText = ref("");
	import { defineEmits } from "vue";
	const emits = defineEmits(["closeAddKelasForm"]);
</script>

<template>
  <div class="edit-form-overlay">
    <div class="form-container">
      <div class="form-inputs-container">

        <div class="form-inputs">
        <div class="close-button">
            <button @click="emits('closeAddKelasForm')" class="button-close">
             ✖
          </button>
        </div>
          <form action="#" method="POST">
            <label for="namakelas">Nama Kelas</label>
            <input type="email" id="email" name="email" />

            <label for="namakelas">Wali Kelas</label>
            <select>
              <option>Bu Ratna Sari SUgih</option>
              <option>Asep Lukman AH SKOM</option>
            </select>

            <button class="submit-btn" type="submit">Buat Kelas</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
	.button-close {
		background-color: rgba(1, 1, 1, 0);
		border: none;
		/* padding: 2px 10px; */
		/* border-radius: 100px; */
		font-size: 17px;
    margin: -5px;
	}

	.close-button {
		display: flex;
		justify-content: flex-end;
    padding: 0px;
    margin: 0px;
	}

.edit-form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
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

.form-inputs-container {
  padding: 40px;
  display: flex;
  justify-content: center;
  background-color: #ffffff;
  gap: 17px;
  border-radius: 7px;
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
