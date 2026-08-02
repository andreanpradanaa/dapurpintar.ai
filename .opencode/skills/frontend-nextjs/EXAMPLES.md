# Frontend Next.js Examples

## Error mapping to M6 codes

```tsx
const error = await api.getPantryItem(itemId).catch(handleApiError)
if (error.code === "PANTRY_ITEM_NOT_FOUND") {
  // show "This item is no longer in your pantry."
}
```

## Degraded AI surface

```tsx
{recommendation.status === "unavailable" ? (
  <EmptyState message="Recommendation is temporarily unavailable." />
) : (
  <RecommendationCard data={recommendation} />
)}
```

## Typed client usage

```tsx
const { data: items } = await useApi(() => api.listPantryItems({ sort: "expiry", order: "asc" }))
```
