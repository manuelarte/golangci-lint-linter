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
          @click:clear="lint(null)"
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

  const lint = (str: string | null) => {
    if (str == null || str == '') {
      proposedContent.value = null
    } else {
      originalContent.value = str
      proposedContent.value = apply(originalContent.value)
    }
  }

</script>
