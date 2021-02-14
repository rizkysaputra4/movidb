<template>
  <div class="q-pa-md">
    <q-card class="my-card">
      <q-card-section class="row">
        <q-avatar size="lg" class="col-shrink q-pr-md q-pt-xs">
          <img src="https://cdn.quasar.dev/img/avatar5.jpg" />
        </q-avatar>

        <div class="col-grow">
          <q-input
            filled
            v-model="text"
            label="Write a comment..."
            autogrow
            type="textarea"
            :dense="dense"
            borderless
          >
          </q-input>

          <div v-if="urlIMG" class="bg-grey-3 q-pa-sm">
            <q-img
              v-if="urlIMG"
              :src="urlIMG"
              style="width: 50px"
              :ratio="1"
              spinner-color="white"
              class="rounded-borders my-img"
            >
              <div class="close-btn absolute-center">
                <q-btn flat dense round icon="close" @click="clearImg" />
              </div>
            </q-img>
          </div>

          <div class="row q-mt-sm">
            <div class="col">
              <input
                ref="file"
                id="fileUpload"
                type="file"
                hidden
                v-on:change="onChangeFile"
              />
              <q-btn
                round
                flat
                dense
                @click="chooseFiles()"
                icon="add_a_photo"
              ></q-btn>
            </div>
            <div class="row justify-end">
              <q-btn
                class="q-mr-sm"
                dense
                flat
                color="white"
                text-color="black"
                label="Cancel"
              />
              <q-btn color="primary" label="post" @click="uploadFile" />
            </div>
          </div>
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      selected_file: "",
      check_if_document_upload: false,
      img: null,
      text: null,
      urlIMG: null,
      files: null,
      model: null,
      dense: false,
    };
  },
  methods: {
    uploadFile() {
      let fd = new FormData();
      console.log(fd);
      fd.append("file", this.img);
      axios
        .post("http://localhost:3001/post/upload", fd, {
          headers: { "Content-Type": "multipart/form-data" },
          withCredentials: true,
        })
        .then(
          function (response) {
            if (response.data.ok) {
              console.log("ok");
            }
          }.bind(this)
        )
        .catch((err) => {
          console.log(err);
        });
    },

    chooseFiles() {
      let fileInputElement = this.$refs.file;
      fileInputElement.click();
    },
    onChangeFile(e) {
      console.log(e.target.files);
      const file = e.target.files[0];
      this.img = file;

      this.urlIMG = URL.createObjectURL(file);
    },
    clearImg() {
      this.urlIMG = null;
      this.img = null;
      this.$refs.file.value = "";
    },
  },
};
</script>

<style>
.my-img .close-btn {
  visibility: hidden;
  opacity: 0;
  transition: 0.3s;
}

.my-img:hover .close-btn {
  visibility: visible;
  opacity: 1;
  transition: 0.3s;
}
</style>
