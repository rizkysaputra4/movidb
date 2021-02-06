<template>
  <div class="window-width row justify-center items-center q-pt-md">
    <q-input
      style="width: 300px"
      color="black"
      filled
      outlined
      bottom-slots
      v-model="userName"
      label="Target username"
    >
      <template v-if="userName" v-slot:append>
        <q-icon
          name="cancel"
          @click.stop="userName = null"
          class="cursor-pointer"
        />
      </template>
      <template v-slot:after>
        <q-btn round dense flat icon="search" />
      </template>
    </q-input>
  </div>
</template>
<script>
export default {
  data() {
    return {
      userName: "",
    };
  },
};
</script>
