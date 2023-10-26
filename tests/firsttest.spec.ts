import { test, expect } from '@playwright/test';

test('basic test', async ({ page }) => {
  await page.goto('/');

  // Page defaults to say "Playwright"
  await expect(page.getByRole('heading', { name: 'Playwright' })).toBeVisible();

  // Fill and submit the form
  await page.getByLabel('Name').fill('Test 1');
  await page.getByLabel('Description').fill('This is a test!');
  await page.getByRole('button').dispatchEvent('click');

  // Check the the page is updated based on the form
  await expect(page.getByRole('heading', { name: 'Test 1' })).toBeVisible();
});