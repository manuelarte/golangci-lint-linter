<template>
  <v-banner
    class="my-4"
    color="warning"
    icon="$warning"
    lines="two"
    :sticky="true"
  >
    <v-banner-text>
      This linter is experimental and you should check all the fixed proposed.
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
        <p>TODO: List of errors</p>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="40">
        <CodeDiff
          filename=".golangci.yml"
          lang="yaml"
          :new-string="proposedContent"
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

  const originalContent: Ref<string | null> = ref(null)
  const proposedContent: Ref<string | null> = ref(null)

  const lint = (str: string | null) => {
    if (str == null || str == '') {
      proposedContent.value = null
    } else {
      originalContent.value = str
      proposedContent.value = apply(originalContent.value)
    }
  }

</script>
