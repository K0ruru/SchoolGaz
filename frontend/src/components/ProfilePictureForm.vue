<template>
  <div class="cropper-overlay">
    <div class="cropper-popup">
      <div class="wrap-button">
      <button class="select-picture">
        <!-- <span>Select Picture</span> -->

        <label for="uploadInput">
          <ion-icon name="cloud-upload"></ion-icon> UPLOAD FILE
        </label>
        <input class="file-upload" id="uploadInput" type="file" accept="image/jpg, image/jpeg, image/png, image/gif"
          @change="selectFile" />
      </button>
        <button class="button-cencel" @click="emits('close-cropper')">cancel</button>
      </div>

      <!-- Crop result preview -->
      <section v-if="result.dataURL">
        <div class="preview">
          <img :src="result.dataURL" alt="Cropped Preview" />
        </div>
      </section>

      <!-- A Popup for cropping -->
      <div v-if="isShowModal" class="modal-wrap">
        <div class="modal-mask"></div>
        <div class="modal-scroll-view">
          <div class="modal">
            <div class="modal-title">
              <span class="title">Crop Image</span>
              <div class="tools">
                <button class="btn" @click="isShowModal = false">Cancel</button>
                <button class="btn" @click="clear">Clear</button>
                <button class="btn" @click="reset">Reset</button>
                <button class="btn primary" @click="getResult">Crop</button>
              </div>
            </div>

            <div class="modal-content">
              <!-- The component imported from `vue-picture-cropper` plugin -->
              <VuePictureCropper :boxStyle="{
                width: '100%',
                height: '500px',
                backgroundColor: '#f8f8f8',
                margin: 'auto',
              }" :img="pic" :options="{
  viewMode: 1,
  dragMode: 'crop',
  aspectRatio: 1 / 1,
}" @ready="ready" />
            </div>
          </div>
        </div>
      </div>

      <!-- Button to save and send cropped image -->
      <button class="btn primary" v-if="result.dataURL" @click="saveAndSendClick">
        Save and Send
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineEmits } from "vue";
import axios from "axios";
import VuePictureCropper, { cropper } from "vue-picture-cropper";
import { useRoute } from "vue-router";

const emits = defineEmits(["close-cropper"]);

const route = useRoute();
const params = route.params;

const isShowModal = ref(false);
const uploadInput = ref<HTMLInputElement | null>(null);
const pic = ref<string>("");
const result = ref({
  dataURL: "",
});

const selectFile = (e: Event) => {
  pic.value = "";
  result.value.dataURL = "";

  const { files } = e.target as HTMLInputElement;
  if (!files || !files.length) return;

  const file = files[0];
  const reader = new FileReader();
  reader.readAsDataURL(file);
  reader.onload = () => {
    pic.value = String(reader.result);
    isShowModal.value = true;

    if (uploadInput.value) {
      uploadInput.value.value = "";
    }
  };
};

const getResult = () => {
  if (!cropper) return;
  const base64 = cropper.getDataURL();
  result.value.dataURL = base64;
  isShowModal.value = false;
};

const clear = () => {
  if (!cropper) return;
  cropper.clear();
};

const reset = () => {
  if (!cropper) return;
  cropper.reset();
};

const ready = () => {
  console.log("Cropper is ready.");
};

const saveAndSendClick = async () => {
  // Extract the user ID from somewhere, e.g., your component's state
  var userId = params.id;

  // Call saveAndSend with the user ID
  if (Array.isArray(userId)) {
    userId = userId[0];
  }

  await saveAndSend(userId);
};

// Your existing saveAndSend function
const saveAndSend = async (userId: string | number) => {
  try {
    if (!result.value.dataURL) {
      console.error("No cropped image to send.");
      return;
    }

    const formData = new FormData();
    const file = dataURLtoFile(result.value.dataURL, "profile.jpg");

    if (!file) {
      console.error("Invalid file format");
      return;
    }

    formData.append("profilePicture", file);

    // Make an Axios request to send the cropped image to the server
    const response = await axios.put(
      `http://localhost:8080/Auth/pp/${userId}`,
      formData,
      {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      }
    );

    console.log("Server response:", response.data);
    window.location.reload();
  } catch (error) {
    console.error("Error sending cropped image:", error);
  }
};

// Helper function to convert Data URL to File
const dataURLtoFile = (dataURL: string, fileName: string): File | null => {
  const arr = dataURL.split(",");
  const mimeMatch = arr[0].match(/:(.*?);/);

  if (!mimeMatch) {
    console.error("Invalid dataURL format");
    return null;
  }

  const mime = mimeMatch[1];
  const bstr = atob(arr[1]);
  let n = bstr.length;
  const u8arr = new Uint8Array(n);

  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }

  return new File([u8arr], fileName, { type: mime });
};
</script>

<style scoped>
input[type="file"] {
  display: none;
}

.wrap-button {
  display: flex;
  align-items: center;
  /* flex-direction: c; */
}
.cropper-overlay {
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

.cropper-popup {
  background: #fff;
  padding: 10px;
  border-radius: 8px;
  text-align: center;
}

.select-picture {
  align-items: center;
  border: none;
  background: #0000ff;
  padding: 5px;
  color: #fff;
  border-radius: 3px;
  display: flex;
  flex-direction: column;
}
.button-cencel {
  padding: 5px;
  margin: 5px;
}

.preview {
  margin-top: 20px;
}

.preview img {
  max-width: 200px;
  border-radius: 50%;
  margin-bottom: 20px;
}

.modal-wrap {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-mask {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
}

.modal-scroll-view {
  overflow-y: auto;
}

.modal {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  margin: 20px;
  position: relative;
}

.modal-title {
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #eee;
}

.modal-content {
  padding: 20px;
}

.btn {
  margin-right: 10px;
  background-color: rgba(0, 0, 0, 0.3);
  /* Blue color */
  color: #fff;
  border: none;
  padding: 10px 15px;
  border-radius: 4px;
  cursor: pointer;
}

.btn.primary {
  background-color: #0000ff;
  /* Green color */
}
</style>

