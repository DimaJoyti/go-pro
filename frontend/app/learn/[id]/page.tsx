"use client";

import { useState, useEffect } from "react";
import { useParams, useRouter } from "next/navigation";
import { api, type LessonDetail, APIError } from "@/lib/api";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import CodeEditor from "@/components/learning/code-editor";
import ExerciseSubmission from "@/components/learning/exercise-submission";
import {
  BookOpen,
  Code2,
  Trophy,
  Play,
  ArrowRight,
  ArrowLeft,
  Clock,
  Target,
  CheckCircle,
  Home,
  List,
  ChevronRight,
  Lightbulb,
  FileText,
  Zap
} from "lucide-react";
import Link from "next/link";

interface LessonData {
  id: number;
  title: string;
  description: string;
  duration: string;
  difficulty: "Beginner" | "Intermediate" | "Advanced";
  phase: string;
  objectives: string[];
  theory: string;
  codeExample: string;
  solution: string;
  exercises: Exercise[];
  nextLessonId?: number;
  prevLessonId?: number;
}

interface Exercise {
  id: string;
  title: string;
  description: string;
  requirements: string[];
  initialCode: string;
  solution: string;
}

export default function LessonPage() {
  const params = useParams();
  const router = useRouter();
  const [activeTab, setActiveTab] = useState("lesson");
  const [lessonData, setLessonData] = useState<LessonData | null>(null);
  const [loading, setLoading] = useState(true);

  const lessonId = parseInt(params.id as string);

  useEffect(() => {
    const loadLessonData = async () => {
      setLoading(true);

      try {
        const lesson = await api.getLessonDetail(lessonId);

        // Convert API response to component format
        const lessonData: LessonData = {
          id: lesson.id,
          title: lesson.title,
          description: lesson.description,
          duration: lesson.duration,
          difficulty: lesson.difficulty.charAt(0).toUpperCase() + lesson.difficulty.slice(1) as "Beginner" | "Intermediate" | "Advanced",
          phase: lesson.phase,
          objectives: lesson.objectives,
          theory: lesson.theory,
          codeExample: lesson.code_example,
          solution: lesson.solution,
          exercises: lesson.exercises.map(ex => ({
            id: ex.id,
            title: ex.title,
            description: ex.description,
            requirements: ex.requirements,
            initialCode: ex.initial_code,
            solution: ex.solution,
          })),
          nextLessonId: lesson.next_lesson_id,
          prevLessonId: lesson.prev_lesson_id,
        };

        setLessonData(lessonData);
      } catch (error) {
        console.error('Failed to load lesson:', error);
        setLessonData(null);
      } finally {
        setLoading(false);
      }
    };

    loadLessonData();
  }, [lessonId]);

  // Fallback mock data for development
  const getMockLessonData = (): LessonData => {
    return {
      id: 1,
      title: "Go Syntax and Basic Types",
      description: "Learn the fundamental syntax of Go and work with basic data types including integers, floats, strings, and booleans.",
      duration: "3-4 hours",
      difficulty: "Beginner",
      phase: "Foundations",
      objectives: [
        "Set up a Go development environment",
        "Understand Go's basic syntax and program structure",
        "Work with primitive data types (int, float, string, bool)",
        "Declare and use constants",
        "Perform type conversions",
        "Use the iota identifier for enumerated constants"
      ],
          theory: `
# Go Program Structure

Every Go program starts with a package declaration, followed by imports, and then the program code:

\`\`\`go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
\`\`\`

## Basic Types

Go has several built-in basic types:

### Numeric Types
- **Integers**: int, int8, int16, int32, int64
- **Unsigned integers**: uint, uint8, uint16, uint32, uint64
- **Floating point**: float32, float64
- **Complex numbers**: complex64, complex128

### Other Types
- **Boolean**: bool (true or false)
- **String**: string (UTF-8 encoded)
- **Byte**: byte (alias for uint8)
- **Rune**: rune (alias for int32, represents Unicode code points)

## Variable Declarations

\`\`\`go
// Explicit type declaration
var name string = "Go"
var age int = 10

// Type inference
var language = "Go"
var version = 1.21

// Short variable declaration (inside functions only)
message := "Hello, World!"
count := 42
\`\`\`

## Constants

\`\`\`go
const Pi = 3.14159
const Language = "Go"

// Enumerated constants with iota
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
\`\`\`
          `,
          codeExample: `package main

import "fmt"

func main() {
    // Basic variable declarations
    var name string = "Go Programming"
    var version float64 = 1.21
    var isAwesome bool = true
    
    // Short variable declaration
    year := 2024
    
    // Constants
    const MaxUsers = 1000
    
    // Type conversion
    var x int = 42
    var y float64 = float64(x)
    
    // Print values
    fmt.Printf("Language: %s\\n", name)
    fmt.Printf("Version: %.2f\\n", version)
    fmt.Printf("Year: %d\\n", year)
    fmt.Printf("Is Awesome: %t\\n", isAwesome)
    fmt.Printf("Max Users: %d\\n", MaxUsers)
    fmt.Printf("Converted: %.1f\\n", y)
}`,
          solution: `package main

import "fmt"

func main() {
    // Basic variable declarations
    var name string = "Go Programming"
    var version float64 = 1.21
    var isAwesome bool = true
    
    // Short variable declaration
    year := 2024
    
    // Constants
    const MaxUsers = 1000
    
    // Type conversion
    var x int = 42
    var y float64 = float64(x)
    
    // Print values
    fmt.Printf("Language: %s\\n", name)
    fmt.Printf("Version: %.2f\\n", version)
    fmt.Printf("Year: %d\\n", year)
    fmt.Printf("Is Awesome: %t\\n", isAwesome)
    fmt.Printf("Max Users: %d\\n", MaxUsers)
    fmt.Printf("Converted: %.1f\\n", y)
    
    // Additional examples
    
    // Multiple variable declaration
    var (
        firstName = "John"
        lastName  = "Doe"
        age       = 30
    )
    
    fmt.Printf("Full Name: %s %s, Age: %d\\n", firstName, lastName, age)
    
    // Enumerated constants
    const (
        Red = iota
        Green
        Blue
    )
    
    fmt.Printf("Colors: Red=%d, Green=%d, Blue=%d\\n", Red, Green, Blue)
}`,
          exercises: [
            {
              id: "basic-variables",
              title: "Variable Declaration Practice",
              description: "Practice declaring variables of different types and using type conversions.",
              requirements: [
                "Declare a string variable for your name",
                "Declare an integer variable for your age",
                "Declare a boolean variable for whether you like programming",
                "Use short variable declaration for the current year",
                "Convert an integer to float64 and print both values"
              ],
              initialCode: `package main

import "fmt"

func main() {
    // TODO: Declare your variables here
    
    // TODO: Print the values
    
}`,
              solution: `package main

import "fmt"

func main() {
    // Variable declarations
    var name string = "Alice"
    var age int = 25
    var likesProgramming bool = true
    currentYear := 2024
    
    // Type conversion
    var score int = 95
    var percentage float64 = float64(score)
    
    // Print values
    fmt.Printf("Name: %s\\n", name)
    fmt.Printf("Age: %d\\n", age)
    fmt.Printf("Likes Programming: %t\\n", likesProgramming)
    fmt.Printf("Current Year: %d\\n", currentYear)
    fmt.Printf("Score: %d, Percentage: %.1f%%\\n", score, percentage)
}`
            }
          ],
      nextLessonId: 2
    };
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-background via-background to-accent/5">
        <div className="container-responsive padding-responsive-y">
          <div className="flex items-center justify-center min-h-[60vh]">
            <div className="text-center">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
              <p className="text-responsive text-muted-foreground">Loading lesson...</p>
            </div>
          </div>
        </div>
      </div>
    );
  }

  if (!lessonData) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-background via-background to-accent/5">
        <div className="container-responsive padding-responsive-y">
          <div className="text-center py-16">
            <h1 className="text-responsive-heading font-bold mb-4">Lesson Not Found</h1>
            <p className="text-responsive text-muted-foreground mb-6">The lesson you're looking for doesn't exist.</p>
            <Link href="/curriculum">
              <Button size="lg">
                <ArrowLeft className="mr-2 h-4 w-4" />
                Back to Curriculum
              </Button>
            </Link>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen animated-gradient">
      <div className="container-responsive padding-responsive-y">
        {/* Breadcrumb Navigation */}
        <div className="flex items-center space-x-2 text-sm text-muted-foreground mb-6 lg:mb-8 animate-in fade-in slide-in-right duration-500">
        <Link href="/" className="hover:text-primary transition-colors">
          <Home className="h-4 w-4" />
        </Link>
        <ChevronRight className="h-4 w-4" />
        <Link href="/curriculum" className="hover:text-primary transition-colors">
          Curriculum
        </Link>
        <ChevronRight className="h-4 w-4" />
        <span className="text-foreground font-medium">Lesson {lessonData.id}</span>
      </div>

      {/* Lesson Header */}
      <div className="mb-8 glass-card p-6 lg:p-8 rounded-2xl animate-in fade-in slide-in-bottom duration-700">
        <div className="flex items-start justify-between mb-4">
          <div className="flex-1">
            <div className="flex items-center flex-wrap gap-2 mb-3">
              <Badge variant="outline" className="shadow-sm">
                <BookOpen className="mr-1 h-3 w-3" />
                Lesson {lessonData.id}
              </Badge>
              <Badge variant="secondary" className="shadow-sm">{lessonData.phase}</Badge>
              <Badge
                variant={lessonData.difficulty === 'Beginner' ? 'success' :
                        lessonData.difficulty === 'Intermediate' ? 'warning' : 'info'}
                className="shadow-sm"
              >
                {lessonData.difficulty}
              </Badge>
            </div>
            <h1 className="text-3xl lg:text-4xl font-bold tracking-tight mb-3 bg-gradient-to-r from-foreground to-foreground/70 bg-clip-text text-transparent">
              {lessonData.title}
            </h1>
            <p className="text-muted-foreground text-base lg:text-lg mb-4 leading-relaxed">{lessonData.description}</p>
            <div className="flex items-center flex-wrap gap-4 text-sm text-muted-foreground">
              <div className="flex items-center space-x-2 bg-background/50 px-3 py-1.5 rounded-lg">
                <Clock className="h-4 w-4 text-primary" />
                <span>{lessonData.duration}</span>
              </div>
              <div className="flex items-center space-x-2 bg-background/50 px-3 py-1.5 rounded-lg">
                <Target className="h-4 w-4 text-primary" />
                <span>{lessonData.objectives.length} objectives</span>
              </div>
            </div>
          </div>
          <div className="ml-6 hidden lg:block">
            <Link href="/curriculum">
              <Button variant="outline" size="sm" className="shadow-sm">
                <List className="mr-2 h-4 w-4" />
                All Lessons
              </Button>
            </Link>
          </div>
        </div>
      </div>

      {/* Lesson Content */}
      <Tabs value={activeTab} onValueChange={setActiveTab} className="space-y-6 animate-in fade-in duration-1000">
        <TabsList className="grid w-full grid-cols-3 lg:w-[450px] bg-card/50 backdrop-blur-sm border border-border/50 shadow-lg">
          <TabsTrigger value="lesson" className="data-[state=active]:bg-primary data-[state=active]:text-primary-foreground">
            <Lightbulb className="mr-2 h-4 w-4" />
            Theory
          </TabsTrigger>
          <TabsTrigger value="practice" className="data-[state=active]:bg-primary data-[state=active]:text-primary-foreground">
            <Code2 className="mr-2 h-4 w-4" />
            Practice
          </TabsTrigger>
          <TabsTrigger value="exercise" className="data-[state=active]:bg-primary data-[state=active]:text-primary-foreground">
            <Trophy className="mr-2 h-4 w-4" />
            Exercise
          </TabsTrigger>
        </TabsList>

        {/* Theory Tab */}
        <TabsContent value="lesson" className="space-y-6">
          <Card className="glass-card border-2">
            <CardHeader className="bg-gradient-to-r from-primary/10 to-primary/5 border-b">
              <CardTitle className="flex items-center text-xl">
                <div className="p-2 rounded-lg bg-primary/10 mr-3">
                  <Lightbulb className="h-5 w-5 text-primary" />
                </div>
                Learning Objectives
              </CardTitle>
            </CardHeader>
            <CardContent className="pt-6">
              <ul className="space-y-3">
                {lessonData.objectives.map((objective) => (
                  <li key={objective} className="flex items-start space-x-3 group">
                    <div className="p-1 rounded-full bg-green-500/10 group-hover:bg-green-500/20 transition-colors">
                      <CheckCircle className="h-4 w-4 text-green-500 flex-shrink-0" />
                    </div>
                    <span className="text-sm leading-relaxed">{objective}</span>
                  </li>
                ))}
              </ul>
            </CardContent>
          </Card>

          <Card className="glass-card border-2">
            <CardHeader className="bg-gradient-to-r from-primary/10 to-primary/5 border-b">
              <CardTitle className="flex items-center text-xl">
                <div className="p-2 rounded-lg bg-primary/10 mr-3">
                  <FileText className="h-5 w-5 text-primary" />
                </div>
                Theory & Concepts
              </CardTitle>
            </CardHeader>
            <CardContent className="pt-6">
              <div className="prose dark:prose-invert max-w-none prose-headings:text-foreground prose-p:text-muted-foreground prose-strong:text-foreground prose-code:text-primary">
                <div dangerouslySetInnerHTML={{ __html: lessonData.theory.replace(/\n/g, '<br>') }} />
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        {/* Practice Tab */}
        <TabsContent value="practice" className="space-y-6">
          <div className="glass-card p-6 rounded-2xl border-2">
            <CodeEditor
              title="Interactive Code Example"
              description="Try modifying this Go code to see how the concepts work. Experiment with different values and see the results!"
              initialCode={lessonData.codeExample}
              solution={lessonData.solution}
              language="go"
              onCodeChange={(code) => console.log("Code changed:", code)}
            />
          </div>
        </TabsContent>

        {/* Exercise Tab */}
        <TabsContent value="exercise" className="space-y-6">
          {lessonData.exercises.map((exercise) => (
            <div key={exercise.id} className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <CodeEditor
                title={exercise.title}
                description={exercise.description}
                initialCode={exercise.initialCode}
                solution={exercise.solution}
                language="go"
                onCodeChange={(code) => console.log("Exercise code:", code)}
              />
              
              <ExerciseSubmission
                exerciseId={exercise.id}
                title={exercise.title}
                description={exercise.description}
                requirements={exercise.requirements}
                code={exercise.initialCode}
                previousSubmissions={[]}
              />
            </div>
          ))}
        </TabsContent>
      </Tabs>

      {/* Navigation */}
      <div className="glass-card p-6 rounded-2xl mt-12 border-2 animate-in fade-in duration-1000">
        <div className="flex flex-col sm:flex-row items-center justify-between gap-4">
          <div className="w-full sm:w-auto">
            {lessonData.prevLessonId ? (
              <Link href={`/learn/${lessonData.prevLessonId}`} className="block w-full sm:w-auto">
                <Button variant="outline" size="lg" className="w-full sm:w-auto shadow-sm hover:shadow-md">
                  <ArrowLeft className="mr-2 h-4 w-4" />
                  Previous Lesson
                </Button>
              </Link>
            ) : (
              <div className="w-full sm:w-auto" />
            )}
          </div>

          <div className="flex flex-col sm:flex-row items-center gap-3 w-full sm:w-auto">
            <Link href="/curriculum" className="w-full sm:w-auto">
              <Button variant="outline" size="lg" className="w-full sm:w-auto shadow-sm hover:shadow-md">
                <List className="mr-2 h-4 w-4" />
                View All Lessons
              </Button>
            </Link>

            {lessonData.nextLessonId && (
              <Link href={`/learn/${lessonData.nextLessonId}`} className="w-full sm:w-auto">
                <Button size="lg" className="go-gradient text-white w-full sm:w-auto shadow-lg hover:shadow-2xl">
                  Next Lesson
                  <ArrowRight className="ml-2 h-4 w-4" />
                </Button>
              </Link>
            )}
          </div>
        </div>
      </div>
      </div>
    </div>
  );
}
