import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    lib: {
      entry: './src/components/web-components/locationWebComponent.tsx',
      name: 'Autofill-Location', 
      fileName: () => 'autofill-location',
      formats: ['es']
    },
    rollupOptions: {
      external: ['react', 'react-dom'],
      output: {
        globals: {
          react: 'React',
          'react-dom': 'ReactDOM',
        },
      },
    }, 
  },
  server: {
    watch: {
      usePolling: true
    },
    port: 5176,
    host: true,
    cors: true
  }
})
