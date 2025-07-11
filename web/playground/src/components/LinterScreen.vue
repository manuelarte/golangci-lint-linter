<template>
  <v-banner
    class="my-4"
    color="warning"
    icon="$warning"
    lines="two"
    :sticky="true"
  >
    <v-banner-text>
      This is an experimental linter, and you should review all suggested fixes.
    </v-banner-text>
  </v-banner>
  <v-container class="fill-height" max-width="1400">
    <v-row>
      <v-col cols="6">
        <v-textarea
          v-model="originalContent"
          clearable
          label="Paste here your .golangci.yml"
          row-height="5"
          rows="5"
          variant="outlined"
          @click:clear="lint(null)"
          @input="lint($event.target.value)"
        />
      </v-col>

      <v-col cols="6">
        <v-banner
          v-if="isError()"
          class="my-4"
          color="error"
          icon="error"
          lines="two"
          :sticky="true"
        >
          <v-banner-text>
            {{ applyResponse?.error }}
          </v-banner-text>
        </v-banner>

        <v-banner
          v-if="isOriginalContent() && applyResponse?.output == originalContent"
          class="my-4"
          color="blue"
          icon="mdi-check-circle"
          lines="one"
        >
          <v-banner-text>
            Looks good!
          </v-banner-text>
        </v-banner>

        <p v-if="isReports()">TODO: List of errors</p>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="40">
        <CodeDiff
          context="100"
          filename=".golangci.yml"
          lang="yaml"
          :new-string="applyResponse?.output || ''"
          :old-string="originalContent"
          output-format="side-by-side"
          theme="dark"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
  import { CodeDiff } from 'v-code-diff'
  import { ref } from 'vue'

  interface GoResponse {
    output: string
    reports: string
    error: string
  }

  const originalContent: Ref<string | null> = ref(null)
  const applyResponse: Ref<GoResponse | null> = ref(null)

  const lint = (str: string | null) => {
    console.log(str)
    if (str == null || str == '') {
      applyResponse.value = null
    } else {
      originalContent.value = str
      console.log(originalContent.value == str)
      // @ts-ignore // func defined in wasm
      applyResponse.value = apply(originalContent.value)
    }
  }
  const isError = (): boolean => {
    if (applyResponse.value == null) {
      return false
    }
    if (applyResponse.value?.error == null) {
      return false
    }
    return applyResponse.value?.error.length > 0
  }

  const isReports = (): boolean => {
    if (applyResponse.value == null) {
      return false
    }
    if (applyResponse.value?.reports == null) {
      return false
    }
    return applyResponse.value?.reports.length > 0
  }

  const isOriginalContent = (): boolean => {
    return originalContent != null && originalContent.value != null && originalContent.value.length > 0
  }

</script>
