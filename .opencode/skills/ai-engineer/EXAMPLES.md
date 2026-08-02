# AI Engineer Examples

## Gateway port (domain-facing)

```go
type Gateway interface {
    // Recommend returns a structured, validated proposal. It never
    // mutates business state.
    Recommend(ctx context.Context, in RecommendInput) (*Recommendation, error)
}
```

## Provider adapter with bounded retry

```go
func (a *OpenAIAdapter) Recommend(ctx context.Context, in RecommendInput) (*Recommendation, error) {
    var resp Recommendation
    err := a.retryer.Do(ctx, func() error {
        raw, err := a.client.CreateChatCompletion(ctx, a.request(in))
        if err != nil {
            return err
        }
        return json.Unmarshal(raw.Choices[0].Message.Content, &resp)
    })
    if err != nil {
        return nil, apperr.Wrap(apperr.CodeAIUnavailable, "Recommendation is temporarily unavailable.", err)
    }
    if err := resp.Validate(); err != nil {
        return nil, apperr.New(apperr.CodeAIOutputInvalid, "The model produced an invalid proposal.")
    }
    return &resp, nil
}
```

## Fail-closed on timeout

```go
ctx, cancel := context.WithTimeout(parent, 20*time.Second)
defer cancel()
// A missing context never becomes an invented answer.
```

## Evaluation scenario

```go
func TestRecommendation_RefusesUnsupportedCategory(t *testing.T) {
    // Given a user without sufficient pantry confidence, the harness
    // asserts the output declares low certainty and declines a concrete plan.
}
```
