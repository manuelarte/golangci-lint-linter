/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'

// Styles
import 'unfonts.css'

import wasmUrl from "./golangci-lint-linter.wasm?url";

const app = createApp(App)

registerPlugins(app)

app.mount('#app')

// Load Go WASM
const responsePromise = fetch(wasmUrl);
// @ts-expect-error
const go = new Go();
WebAssembly.instantiateStreaming(responsePromise, go.importObject)
  .then(result => {
    go.run(result.instance);
});
