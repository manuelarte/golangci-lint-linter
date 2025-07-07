<template>
  <v-container class="fill-height" max-width="1400">
    <v-row>
      <v-col cols="6">
        <v-textarea
          clearable
          label="Original"
          row-height="30"
          rows="30"
          variant="outlined"
        />
      </v-col>

      <v-col cols="6">
        <v-textarea label="Fixed" row-height="30" rows="30" variant="outlined" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
</script>
