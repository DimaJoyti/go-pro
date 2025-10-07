"use client";

import { useState, useRef } from "react";
import Editor from "@monaco-editor/react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { 
  Play, 
  RotateCcw, 
  Copy, 
  Check, 
  AlertCircle,
  CheckCircle,
  Clock,
  Terminal
} from "lucide-react";

interface TestResult {
  name: string;
  passed: boolean;
  message?: string;
}

interface CodeEditorProps {
  title: string;
  description: string;
  initialCode: string;
  solution?: string;
  tests?: TestResult[];
  language?: string;
  theme?: string;
  readOnly?: boolean;
  onCodeChange?: (code: string) => void;
  onRun?: (code: string) => Promise<{ output: string; error?: string; tests?: TestResult[] }>;
}

const CodeEditor = ({
  title,
  description,
  initialCode,
  solution,
  tests = [],
  language = "go",
  theme = "vs-dark",
  readOnly = false,
  onCodeChange,
  onRun
}: CodeEditorProps) => {
  const [code, setCode] = useState(initialCode);
  const [output, setOutput] = useState("");
  const [error, setError] = useState("");
  const [testResults, setTestResults] = useState<TestResult[]>(tests);
  const [isRunning, setIsRunning] = useState(false);
  const [copied, setCopied] = useState(false);
  const [showSolution, setShowSolution] = useState(false);
  const editorRef = useRef<any>(null);

  const handleEditorDidMount = (editor: any) => {
    editorRef.current = editor;
  };

  const handleCodeChange = (value: string | undefined) => {
    const newCode = value || "";
    setCode(newCode);
    onCodeChange?.(newCode);
  };

  const handleRun = async () => {
    if (!onRun) {
      // Simulate running code for demo
      setIsRunning(true);
      setOutput("Running your Go code...\n");
      
      setTimeout(() => {
        setOutput("Hello, World!\nProgram executed successfully!");
        setTestResults([
          { name: "Test 1: Basic functionality", passed: true },
          { name: "Test 2: Edge cases", passed: true },
          { name: "Test 3: Performance", passed: false, message: "Timeout exceeded" }
        ]);
        setIsRunning(false);
      }, 2000);
      return;
    }

    setIsRunning(true);
    setError("");
    setOutput("");
    
    try {
      const result = await onRun(code);
      setOutput(result.output);
      if (result.error) {
        setError(result.error);
      }
      if (result.tests) {
        setTestResults(result.tests);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setIsRunning(false);
    }
  };

  const handleReset = () => {
    setCode(initialCode);
    setOutput("");
    setError("");
    setTestResults(tests);
    setShowSolution(false);
  };

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(code);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch (err) {
      console.error("Failed to copy code:", err);
    }
  };

  const toggleSolution = () => {
    if (solution) {
      if (showSolution) {
        setCode(initialCode);
      } else {
        setCode(solution);
      }
      setShowSolution(!showSolution);
    }
  };

  const passedTests = testResults.filter(test => test.passed).length;
  const totalTests = testResults.length;

  return (
    <div className="space-y-6">
      {/* Header */}
      <Card>
        <CardHeader>
          <div className="flex items-start justify-between">
            <div>
              <CardTitle className="text-xl">{title}</CardTitle>
              <CardDescription className="mt-2 text-base">
                {description}
              </CardDescription>
            </div>
            <div className="flex items-center space-x-2">
              <Badge variant="outline" className="capitalize">
                {language}
              </Badge>
              {totalTests > 0 && (
                <Badge 
                  variant={passedTests === totalTests ? "default" : "secondary"}
                  className={passedTests === totalTests ? "bg-green-500 text-white" : ""}
                >
                  {passedTests}/{totalTests} Tests
                </Badge>
              )}
            </div>
          </div>
        </CardHeader>
      </Card>

      {/* Editor and Output */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Code Editor */}
        <Card>
          <CardHeader className="pb-3">
            <div className="flex items-center justify-between">
              <CardTitle className="text-lg">Code Editor</CardTitle>
              <div className="flex items-center space-x-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={handleCopy}
                  className="h-8"
                >
                  {copied ? (
                    <Check className="h-3 w-3" />
                  ) : (
                    <Copy className="h-3 w-3" />
                  )}
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={handleReset}
                  className="h-8"
                >
                  <RotateCcw className="h-3 w-3" />
                </Button>
                {solution && (
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={toggleSolution}
                    className="h-8"
                  >
                    {showSolution ? "Hide" : "Show"} Solution
                  </Button>
                )}
              </div>
            </div>
          </CardHeader>
          <CardContent className="p-0">
            <div className="border rounded-lg overflow-hidden">
              <Editor
                height="400px"
                language={language}
                theme={theme}
                value={code}
                onChange={handleCodeChange}
                onMount={handleEditorDidMount}
                options={{
                  readOnly,
                  minimap: { enabled: false },
                  fontSize: 14,
                  lineNumbers: "on",
                  roundedSelection: false,
                  scrollBeyondLastLine: false,
                  automaticLayout: true,
                  tabSize: 2,
                  insertSpaces: false,
                }}
              />
            </div>
            <div className="p-4 border-t">
              <Button
                onClick={handleRun}
                disabled={isRunning || readOnly}
                className="go-gradient text-white"
              >
                {isRunning ? (
                  <>
                    <Clock className="mr-2 h-4 w-4 animate-spin" />
                    Running...
                  </>
                ) : (
                  <>
                    <Play className="mr-2 h-4 w-4" />
                    Run Code
                  </>
                )}
              </Button>
            </div>
          </CardContent>
        </Card>

        {/* Output and Tests */}
        <Card>
          <CardHeader className="pb-3">
            <CardTitle className="text-lg flex items-center">
              <Terminal className="mr-2 h-5 w-5" />
              Output & Tests
            </CardTitle>
          </CardHeader>
          <CardContent>
            <Tabs defaultValue="output" className="w-full">
              <TabsList className="grid w-full grid-cols-2">
                <TabsTrigger value="output">Output</TabsTrigger>
                <TabsTrigger value="tests">
                  Tests {totalTests > 0 && `(${passedTests}/${totalTests})`}
                </TabsTrigger>
              </TabsList>
              
              <TabsContent value="output" className="mt-4">
                <div className="code-block min-h-[300px] p-4 font-mono text-sm">
                  {error ? (
                    <div className="text-red-500">
                      <div className="flex items-center mb-2">
                        <AlertCircle className="mr-2 h-4 w-4" />
                        Error
                      </div>
                      <pre className="whitespace-pre-wrap">{error}</pre>
                    </div>
                  ) : output ? (
                    <pre className="whitespace-pre-wrap">{output}</pre>
                  ) : (
                    <div className="text-muted-foreground italic">
                      Click "Run Code" to see the output here...
                    </div>
                  )}
                </div>
              </TabsContent>
              
              <TabsContent value="tests" className="mt-4">
                <div className="space-y-3">
                  {testResults.length > 0 ? (
                    testResults.map((test, index) => (
                      <div
                        key={index}
                        className={`flex items-start space-x-3 p-3 rounded-lg border ${
                          test.passed
                            ? 'bg-green-50 border-green-200 dark:bg-green-950 dark:border-green-800'
                            : 'bg-red-50 border-red-200 dark:bg-red-950 dark:border-red-800'
                        }`}
                      >
                        {test.passed ? (
                          <CheckCircle className="h-4 w-4 text-green-500 mt-0.5" />
                        ) : (
                          <AlertCircle className="h-4 w-4 text-red-500 mt-0.5" />
                        )}
                        <div className="flex-1">
                          <div className="font-medium text-sm">{test.name}</div>
                          {test.message && (
                            <div className="text-sm text-muted-foreground mt-1">
                              {test.message}
                            </div>
                          )}
                        </div>
                      </div>
                    ))
                  ) : (
                    <div className="text-muted-foreground italic text-center py-8">
                      No tests available for this exercise
                    </div>
                  )}
                </div>
              </TabsContent>
            </Tabs>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default CodeEditor;
