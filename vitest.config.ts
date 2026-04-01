import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    include: [
      'src/**/*.test.{ts,js,tsx,jsx}',
      'cli/**/*.test.{ts,js,tsx,jsx}',
      'test/**/*.test.{ts,js,tsx,jsx}',
    ],
    exclude: ['frontend/**'],
  },
})
