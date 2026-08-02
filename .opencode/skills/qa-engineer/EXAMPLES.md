# QA Engineer Examples

## Contract test for an error path

```go
func TestGetPantryItem_ReturnsPantryItemNotFound(t *testing.T) {
    // Arrange: authenticated subject, item id not owned by subject.
    // Act: GET /api/v1/pantry/items/{itemId}
    // Assert: 404, envelope.error.code == "PANTRY_ITEM_NOT_FOUND".
}
```

## Acceptance criterion mapping

```text
Story: As Sarah, I remove an expired ingredient so my pantry is accurate.
  AC-1: Removing a pantry item soft-deletes it.  -> test_pantry_remove_soft_delete
  AC-2: Removed items do not appear in listings. -> test_pantry_list_excludes_removed
  AC-3: Removing someone else's item is rejected.-> test_pantry_remove_forbidden_owner
```

## Deterministic timezone test

```go
// The test fixes the subject timezone to Asia/Jakarta and a known UTC clock,
// then asserts day boundaries are computed in Jakarta local time.
```
