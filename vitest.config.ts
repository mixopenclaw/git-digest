import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    include: [
      'backend/**/*.test.{ts,js}',
      'scanner/**/*.test.{ts,js}',
    ],
  },
})
