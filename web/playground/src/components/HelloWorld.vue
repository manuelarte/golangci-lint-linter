<template>
  <v-container class="fill-height" max-width="1400">
    <v-row>
      <v-col cols="6">
        <v-textarea
          v-model="originalContent"
          clearable
          label="Original"
          row-height="30"
          rows="30"
          variant="outlined"
          @input="lint($event.target.value)"
        />
      </v-col>

      <v-col cols="6">
        <v-textarea
          v-model="proposedContent"
          label="Fixed"
          readonly
          row-height="30"
          rows="30"
          variant="outlined"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
  import { ref } from 'vue'

  const originalContent = ref()
  const proposedContent = ref()

  const lint = (str: string) => {
    if (str != null) {
      originalContent.value = str
      proposedContent.value = apply(originalContent.value)
    }
  }

  if (originalContent.value != null) {
    let pc = apply(originalContent)
    console.log(pc)
    proposedContent.value = pc
  }

</script>
