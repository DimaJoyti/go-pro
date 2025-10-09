package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"ai-agent-platform/internal/agent"
	"ai-agent-platform/internal/languages/common"
	"ai-agent-platform/internal/languages/golang"
	"ai-agent-platform/internal/llm"
	"ai-agent-platform/internal/tools/programming"
	"ai-agent-platform/pkg/types"
)

func main() {
	fmt.Println("🤖 Coding Expert AI Agent - Programming Q&A System")
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println()

	// Get API key from environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ OPENAI_API_KEY environment variable not set")
	}

	// Create LLM provider
	fmt.Println("🔧 Initializing LLM provider...")
	llmProvider, err := llm.NewOpenAIProvider(llm.OpenAIConfig{
		APIKey: apiKey,
		Model:  "gpt-4",
	})
	if err != nil {
		log.Fatalf("❌ Failed to create LLM provider: %v", err)
	}
	fmt.Println("✅ LLM provider initialized")

	// Create language registry
	fmt.Println("🔧 Setting up language support...")
	languageRegistry := common.NewLanguageRegistry()
	
	// Register Go language provider
	goProvider := golang.NewProvider()
	if err := languageRegistry.Register(goProvider); err != nil {
		log.Fatalf("❌ Failed to register Go provider: %v", err)
	}
	fmt.Println("✅ Registered Go language support")

	// Create programming tools
	fmt.Println("🔧 Creating programming tools...")
	tools := []types.Tool{
		programming.NewCodeAnalysisTool(languageRegistry),
		programming.NewCodeExecutionTool(languageRegistry),
		programming.NewDocumentationSearchTool(),
		programming.NewStackOverflowSearchTool(),
		programming.NewGitHubSearchTool(),
	}
	fmt.Println("✅ Created 5 programming tools")

	// Create coding expert agent
	fmt.Println("🔧 Initializing Coding Expert Agent...")
	codingAgent := agent.NewCodingExpertAgent(agent.CodingExpertConfig{
		Name:           "CodingExpert",
		Description:    "Expert programming assistant",
		LLM:            llmProvider,
		Tools:          tools,
		MaxSteps:       5,
		Temperature:    0.7,
		VerboseLogging: true,
		SupportedLangs: []string{"Go", "Python", "JavaScript", "TypeScript", "Rust"},
	})
	fmt.Println("✅ Coding Expert Agent ready")
	fmt.Println()

	// Example questions
	examples := []struct {
		name  string
		query string
	}{
		{
			name:  "Go Goroutines",
			query: "How do I use goroutines in Go to process items concurrently?",
		},
		{
			name:  "Code Analysis",
			query: "Analyze this Go code for issues: package main\n\nfunc main() {\n\tvar x int\n\tif x == 0 {\n\t}\n}",
		},
		{
			name:  "Best Practices",
			query: "What are the best practices for error handling in Go?",
		},
	}

	// Run examples
	ctx := context.Background()

	for i, example := range examples {
		fmt.Printf("📝 Example %d: %s\n", i+1, example.name)
		fmt.Println(string(make([]byte, 60)))
		fmt.Printf("Question: %s\n\n", example.query)

		// Run agent
		input := types.NewAgentInput(example.query)
		result, err := codingAgent.Run(ctx, input)

		if err != nil {
			fmt.Printf("❌ Error: %v\n\n", err)
			continue
		}

		// Display results
		fmt.Println("🎯 Answer:")
		fmt.Println(result.Output)
		fmt.Println()

		// Display metadata
		fmt.Printf("📊 Metadata:\n")
		fmt.Printf("   - Execution Time: %v\n", result.Metadata.Duration)
		fmt.Printf("   - Steps Taken: %d\n", len(result.Steps))
		fmt.Printf("   - Tokens Used: %d (Prompt: %d, Completion: %d)\n",
			result.Metadata.TokensUsed.TotalTokens,
			result.Metadata.TokensUsed.PromptTokens,
			result.Metadata.TokensUsed.CompletionTokens,
		)

		// Display steps if verbose
		if len(result.Steps) > 0 {
			fmt.Println("\n🔍 Reasoning Steps:")
			for _, step := range result.Steps {
				fmt.Printf("   Step %d: %s\n", step.StepNumber, step.Thought)
				if step.Action != "" {
					fmt.Printf("   Action: %s\n", step.Action)
				}
				if step.Observation != "" {
					fmt.Printf("   Observation: %s\n", step.Observation[:min(100, len(step.Observation))])
				}
			}
		}

		fmt.Println()
		fmt.Println(string(make([]byte, 60)))
		fmt.Println()
	}

	// Interactive mode
	fmt.Println("💬 Interactive Mode")
	fmt.Println("Type your programming questions (or 'quit' to exit)")
	fmt.Println(string(make([]byte, 60)))
	fmt.Println()

	// Simple interactive loop
	for {
		fmt.Print("You: ")
		var query string
		fmt.Scanln(&query)

		if query == "quit" || query == "exit" {
			fmt.Println("👋 Goodbye!")
			break
		}

		if query == "" {
			continue
		}

		// Run agent
		input := types.NewAgentInput(query)
		result, err := codingAgent.Run(ctx, input)

		if err != nil {
			fmt.Printf("❌ Error: %v\n\n", err)
			continue
		}

		fmt.Printf("\n🤖 Agent: %s\n\n", result.Output)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

