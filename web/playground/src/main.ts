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

import init from './golangci-lint-linter.wasm?init';

const app = createApp(App)

registerPlugins(app)

app.mount('#app')

// Load Go WASM
// @ts-expect-error
const go = new Go();
WebAssembly.instantiateStreaming(fetch("src/golangci-lint-linter.wasm"), go.importObject).then(result => {
  go.run(result.instance);
  init(go.importObject);
});
