const { test, expect } = require('@playwright/test');

test.describe('Home page', () => {
  test('loads and displays hero CTA', async ({ page }) => {
    await page.goto('/');
    await expect(page.locator('body')).toBeVisible();

    // The hero section should contain a call-to-action link or button
    const cta = page.locator('text=/scan|get started|try it/i').first();
    await expect(cta).toBeVisible({ timeout: 10000 });
  });

  test('navigation is visible', async ({ page }) => {
    await page.goto('/');
    const nav = page.locator('nav');
    await expect(nav).toBeVisible();
  });
});
