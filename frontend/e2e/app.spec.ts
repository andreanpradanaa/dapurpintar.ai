import { test, expect } from "@playwright/test";

const EMAIL = `e2e-${Date.now()}@test.example`;
const PASSWORD = "password123";

test.describe("DapurPintar MVP", () => {
  test("01 - register account", async ({ page }) => {
    await page.goto("/login");
    await expect(page.locator("h1")).toContainText("DapurPintar");

    await page.getByText("Daftar").click();
    await page.getByPlaceholder("Display name").fill("Test User");
    await page.getByPlaceholder("Email").fill(EMAIL);
    await page.getByPlaceholder("Password").fill(PASSWORD);
    await page.getByRole("button", { name: "Create account" }).click();

    await expect(page).toHaveURL("/today");
    await expect(page.locator("h1")).toContainText("Today");
  });

  test("02 - add pantry items", async ({ page }) => {
    // login first
    await page.goto("/login");
    await page.getByPlaceholder("Email").fill(EMAIL);
    await page.getByPlaceholder("Password").fill(PASSWORD);
    await page.getByRole("button", { name: "Log in" }).click();
    await expect(page).toHaveURL("/today");

    // go to pantry
    await page.getByText("Pantry").click();
    await expect(page.locator("h1")).toContainText("Pantry");

    // add item
    await page.getByRole("button", { name: "Add item" }).click();
    await page.getByPlaceholder("Ingredient name").fill("Beras");
    await page.getByPlaceholder("Category").fill("pokok");
    await page.getByPlaceholder("Qty").fill("2");
    await page.getByPlaceholder("Unit").fill("kg");
    await page.getByRole("button", { name: "Save" }).click();

    // verify item appears
    await expect(page.getByText("Beras")).toBeVisible();

    // add another item
    await page.getByRole("button", { name: "Add item" }).click();
    await page.getByPlaceholder("Ingredient name").fill("Telur");
    await page.getByPlaceholder("Category").fill("protein");
    await page.getByText("Save").click();
    await expect(page.getByText("Telur")).toBeVisible();
  });

  test("03 - today dashboard shows stats", async ({ page }) => {
    await page.goto("/login");
    await page.getByPlaceholder("Email").fill(EMAIL);
    await page.getByPlaceholder("Password").fill(PASSWORD);
    await page.getByRole("button", { name: "Log in" }).click();
    await expect(page).toHaveURL("/today");

    // stats should show total items > 0
    await expect(page.getByText("Pantry items")).toBeVisible();
  });

  test("04 - discover recipes", async ({ page }) => {
    await page.goto("/login");
    await page.getByPlaceholder("Email").fill(EMAIL);
    await page.getByPlaceholder("Password").fill(PASSWORD);
    await page.getByRole("button", { name: "Log in" }).click();

    await page.getByText("Discover").click();
    await expect(page.locator("h1")).toContainText("Discover");

    // should show seeded recipes
    await expect(page.getByText("Nasi Goreng")).toBeVisible();
    await expect(page.getByText("Rendang")).toBeVisible();

    // search
    await page.getByPlaceholder("Search recipes...").fill("soto");
    await page.getByRole("button", { name: "Search" }).click();
    await expect(page.getByText("Soto Ayam")).toBeVisible();
  });

  test("05 - meal planner", async ({ page }) => {
    await page.goto("/login");
    await page.getByPlaceholder("Email").fill(EMAIL);
    await page.getByPlaceholder("Password").fill(PASSWORD);
    await page.getByRole("button", { name: "Log in" }).click();

    await page.getByText("Planner").click();
    await expect(page.locator("h1")).toContainText("Planner");

    // create a plan
    await page.getByRole("button", { name: "New plan" }).click();
    const today = new Date().toISOString().slice(0, 10);
    const nextWeek = new Date(Date.now() + 7 * 86400000).toISOString().slice(0, 10);
    await page.locator("#plan-start").fill(today);
    await page.locator("#plan-end").fill(nextWeek);
    await page.getByRole("button", { name: "Create" }).click();

    // should show the weekly grid
    await expect(page.getByText("← Back to plans")).toBeVisible();
  });

  test("06 - shopping list", async ({ page }) => {
    await page.goto("/login");
    await page.getByPlaceholder("Email").fill(EMAIL);
    await page.getByPlaceholder("Password").fill(PASSWORD);
    await page.getByRole("button", { name: "Log in" }).click();

    await page.getByText("Shopping").click();
    await expect(page.locator("h1")).toContainText("Shopping");

    // create list
    await page.getByRole("button", { name: "New list" }).click();
    await page.getByPlaceholder("List title").fill("Belanja Mingguan");
    await page.getByRole("button", { name: "Create" }).click();

    // add item
    await page.getByPlaceholder("Item name").fill("Sabun");
    await page.getByRole("button", { name: "Add" }).click();
    await expect(page.getByText("Sabun")).toBeVisible();
  });

  test("07 - logout", async ({ page }) => {
    await page.goto("/login");
    await page.getByPlaceholder("Email").fill(EMAIL);
    await page.getByPlaceholder("Password").fill(PASSWORD);
    await page.getByRole("button", { name: "Log in" }).click();

    await page.getByText("Keluar").click();
    await expect(page).toHaveURL("/login");
  });
});
