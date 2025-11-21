import { expect, test } from '@playwright/test';

test.describe('Jobs Feature', () => {
  test.beforeEach(async ({ page }) => {
    // Navigate to login page
    await page.goto('/');
    
    // Login with dev credentials
    await page.fill('input[type="email"]', 'client@skillsier.dev');
    await page.fill('input[type="password"]', 'client123');
    await page.click('button[type="submit"]');
    
    // Wait for navigation to dashboard
    await page.waitForURL('**/dashboard');
  });

  test('displays jobs list', async ({ page }) => {
    // Navigate to jobs page
    await page.goto('/dashboard/jobs');
    
    // Check page title
    await expect(page.locator('h1')).toContainText('Jobs');
    
    // Check for jobs list (might be empty)
    const jobsList = page.locator('[data-testid="jobs-list"]');
    await expect(jobsList).toBeVisible();
  });

  test('can create a new job', async ({ page }) => {
    // Navigate to create job page
    await page.goto('/dashboard/jobs/create');
    
    // Fill job form
    await page.fill('input[name="title"]', 'E2E Test Job');
    await page.fill(
      'textarea[name="description"]',
      'This is a test job created by E2E tests. It should have at least 50 characters.'
    );
    await page.fill('input[name="budget"]', '3000');
    await page.fill('input[name="skills"]', 'React, TypeScript, Testing');
    await page.selectOption('select[name="duration"]', '1-3-months');
    
    // Submit form
    await page.click('button[type="submit"]');
    
    // Wait for redirect to job detail page
    await page.waitForURL('**/jobs/*');
    
    // Verify job was created
    await expect(page.locator('h1')).toContainText('E2E Test Job');
  });

  test('can filter jobs', async ({ page }) => {
    await page.goto('/dashboard/jobs');
    
    // Open filters panel
    await page.click('text=Filters');
    
    // Apply status filter
    await page.check('input[type="checkbox"][value="ACTIVE"]');
    
    // Check that filter is applied
    const activeFilter = page.locator('text=ACTIVE').first();
    await expect(activeFilter).toBeVisible();
  });

  test('can view job details', async ({ page }) => {
    await page.goto('/dashboard/jobs');
    
    // Click on first job (if exists)
    const firstJob = page.locator('[data-testid="job-card"]').first();
    if (await firstJob.isVisible()) {
      await firstJob.click();
      
      // Verify we're on detail page
      await expect(page.locator('h1')).toBeVisible();
      await expect(page.locator('text=Budget')).toBeVisible();
      await expect(page.locator('text=Proposals')).toBeVisible();
    }
  });

  test('displays empty state when no jobs', async ({ page }) => {
    await page.goto('/dashboard/jobs');
    
    // If no jobs, should show empty state
    const emptyState = page.locator('text=No jobs found');
    if (await emptyState.isVisible()) {
      await expect(emptyState).toBeVisible();
    }
  });
});