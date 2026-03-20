import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// Ensure built asset paths are relative so dist/index.html works when opened via file://
export default defineConfig({
  base: './',
  plugins:[react()]
})
