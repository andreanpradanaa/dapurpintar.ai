import { test, expect } from "@playwright/test";

const EMAIL = `e2e-${Date.now()}@test.example`;
const PASSWORD = "password123";

test("DapurPintar MVP full flow", async ({ page }) => {
  // 1. Register
  await page.goto("/login", { waitUntil: "networkidle" });
  await page.getByText("Daftar").click();
  await page.getByPlaceholder("Display name").fill("E2E User");
  await page.getByPlaceholder("Email").fill(EMAIL);
  await page.getByPlaceholder("Password").fill(PASSWORD);
  await page.getByRole("button", { name: "Create account" }).click();
  await page.waitForURL("/today", { timeout: 15000 });
  await expect(page.locator("h1")).toContainText("Today");

  // 2. Add pantry items
  await page.getByRole("link", { name: "Pantry" }).click();
  await page.waitForTimeout(2000);
  await expect(page.locator("h1")).toContainText("Pantry");
  await page.getByRole("button", { name: "Add item" }).click();
  await page.getByPlaceholder("Ingredient name").fill("Beras");
  await page.getByPlaceholder("Category").fill("pokok");
  await page.getByPlaceholder("Qty").fill("2");
  await page.getByPlaceholder("Unit").fill("kg");
  await page.getByRole("button", { name: "Save" }).click();
  await page.waitForTimeout(3000);
  await expect(page.getByText("Beras")).toBeVisible({ timeout: 10000 });

  // 3. Discover recipes
  await page.getByRole("link", { name: "Discover" }).click();
  await page.waitForTimeout(2000);
  await expect(page.locator("h1")).toContainText("Discover");
  await expect(page.getByText("Nasi Goreng")).toBeVisible({ timeout: 10000 });

  // 4. Meal planner
  await page.getByRole("link", { name: "Planner" }).click();
  await page.waitForTimeout(2000);
  await expect(page.locator("h1")).toContainText("Planner");
  await page.getByRole("button", { name: "New plan" }).click();
  const today = new Date().toISOString().slice(0, 10);
  const nextWeek = new Date(Date.now() + 7 * 86400000).toISOString().slice(0, 10);
  await page.locator("#plan-start").fill(today);
  await page.locator("#plan-end").fill(nextWeek);
  await page.getByRole("button", { name: "Create" }).click();
  await expect(page.getByText("← Back to plans")).toBeVisible({ timeout: 10000 });

  // 5. Shopping list
  await page.getByRole("link", { name: "Shopping" }).click();
  await page.waitForTimeout(2000);
  await expect(page.locator("h1")).toContainText("Shopping");
  await page.getByRole("button", { name: "New list" }).click();
  await page.getByPlaceholder("List title").fill("Belanja");
  await page.getByRole("button", { name: "Create" }).click();
  await page.getByPlaceholder("Item name").fill("Sabun");
  await page.getByRole("button", { name: "Add" }).click();
  await expect(page.getByText("Sabun")).toBeVisible({ timeout: 5000 });

  // 6. Logout
  await page.getByText("Keluar").click();
  await page.waitForURL("/login", { timeout: 10000 });
});
