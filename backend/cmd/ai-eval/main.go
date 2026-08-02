// Command ai-eval runs the AI evaluation harness against the configured
// provider (M8-003, M4-DEC-012). It evaluates the promoted policy revision (or
// a pinned revision with -schema-rev) over the representative scenarios and
// exits non-zero when the gate does not pass.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/ai"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/ai/openai"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/config"
	"github.com/andreanpradanaa/dapurpintar.ai/backend/internal/platform/logger"
)

func main() {
	schemaRev := flag.String("schema-rev", "", "schema revision to evaluate (empty = promoted)")
	flag.Parse()

	cfg := config.Load()
	log := logger.New(cfg.AppEnv)

	if cfg.AIProvider == "" || cfg.AIAPIKey == "" {
		log.Error("AI provider is not configured (AI_PROVIDER, AI_API_KEY)")
		os.Exit(2)
	}

	adapter, err := openai.New(openai.Config{
		APIKey:     cfg.AIAPIKey,
		Timeout:    cfg.AITimeout,
		MaxRetries: cfg.AIMaxRetries,
		Log:        log,
	})
	if err != nil {
		log.Error("ai gateway setup failed", "error", err)
		os.Exit(2)
	}

	profile := ai.DefaultProfile()
	profile.Name = cfg.AIModel

	eval := ai.NewEvaluator(adapter, ai.DefaultRubric())
	registry := ai.SeedRegistry()

	report, err := eval.Evaluate(context.Background(), registry, ai.PurposeKitchenRecommendation, *schemaRev, ai.SeedScenarios())
	if err != nil {
		log.Error("evaluation failed", "error", err)
		os.Exit(2)
	}

	printReport(report)
	if !report.Pass {
		fmt.Fprintf(os.Stderr, "EVALUATION GATE FAILED: %d/%d scenarios passed\n", report.Passed, report.Total)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "EVALUATION GATE PASSED: %d/%d scenarios passed\n", report.Passed, report.Total)
}

func printReport(r ai.Report) {
	for _, res := range r.Results {
		fmt.Printf("scenario %s: pass=%v overall=%.2f\n", res.ScenarioID, res.Pass, res.Overall)
		for _, s := range res.Scores {
			fmt.Printf("  %s: value=%.2f pass=%v", s.Dimension, s.Value, s.Pass)
			for _, n := range s.Notes {
				fmt.Printf(" [%s]", n)
			}
			fmt.Println()
		}
	}
	fmt.Printf("gate: %d/%d scenarios passed\n", r.Passed, r.Total)
}
