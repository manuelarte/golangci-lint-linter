<template>
  <v-toolbar border>
    <v-toolbar-title text="Golangci-Lint Linter Playground" />

    <template #append>
      <v-btn icon="mdi-github" @click="redirect()" />
    </template>
  </v-toolbar>
  <v-main>
    <router-view />
  </v-main>

  <AppFooter />
</template>

<script lang="ts" setup>
  const redirect = () => {
    window.location.href = 'https://github.com/manuelarte/golangci-lint-linter'
  }
</script>
