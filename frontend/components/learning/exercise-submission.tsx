"use client";

import { useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import { Textarea } from "@/components/ui/textarea";
import { 
  Send, 
  CheckCircle, 
  AlertCircle, 
  Clock,
  Trophy,
  Star,
  MessageSquare,
  ThumbsUp,
  ThumbsDown
} from "lucide-react";

interface Feedback {
  score: number;
  maxScore: number;
  passed: boolean;
  comments: string[];
  suggestions: string[];
  strengths: string[];
}

interface ExerciseSubmissionProps {
  exerciseId: string;
  title: string;
  description: string;
  requirements: string[];
  code: string;
  onSubmit?: (code: string, notes?: string) => Promise<Feedback>;
  previousSubmissions?: Array<{
    id: string;
    timestamp: string;
    score: number;
    feedback: string;
  }>;
}

const ExerciseSubmission = ({
  exerciseId,
  title,
  description,
  requirements,
  code,
  onSubmit,
  previousSubmissions = []
}: ExerciseSubmissionProps) => {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [feedback, setFeedback] = useState<Feedback | null>(null);
  const [notes, setNotes] = useState("");
  const [showPreviousSubmissions, setShowPreviousSubmissions] = useState(false);

  const handleSubmit = async () => {
    if (!onSubmit) {
      // Demo feedback
      setIsSubmitting(true);
      setTimeout(() => {
        setFeedback({
          score: 85,
          maxScore: 100,
          passed: true,
          comments: [
            "Great job implementing the core functionality!",
            "Your code is well-structured and readable.",
            "Consider adding more error handling for edge cases."
          ],
          suggestions: [
            "Add input validation for empty strings",
            "Consider using more descriptive variable names",
            "Add unit tests to verify your implementation"
          ],
          strengths: [
            "Clean and readable code structure",
            "Proper use of Go idioms",
            "Efficient algorithm implementation"
          ]
        });
        setIsSubmitting(false);
      }, 3000);
      return;
    }

    setIsSubmitting(true);
    try {
      const result = await onSubmit(code, notes);
      setFeedback(result);
    } catch (error) {
      console.error("Submission failed:", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  const getScoreColor = (score: number, maxScore: number) => {
    const percentage = (score / maxScore) * 100;
    if (percentage >= 90) return "text-green-500";
    if (percentage >= 70) return "text-blue-500";
    if (percentage >= 50) return "text-yellow-500";
    return "text-red-500";
  };

  const getScoreBadge = (score: number, maxScore: number) => {
    const percentage = (score / maxScore) * 100;
    if (percentage >= 90) return { text: "Excellent", color: "bg-green-500 text-white" };
    if (percentage >= 70) return { text: "Good", color: "bg-blue-500 text-white" };
    if (percentage >= 50) return { text: "Needs Work", color: "bg-yellow-500 text-white" };
    return { text: "Poor", color: "bg-red-500 text-white" };
  };

  return (
    <div className="space-y-6">
      {/* Exercise Info */}
      <Card>
        <CardHeader>
          <CardTitle className="text-xl">{title}</CardTitle>
          <CardDescription className="text-base">
            {description}
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div>
              <h4 className="font-medium mb-2">Requirements:</h4>
              <ul className="space-y-1">
                {requirements.map((req, index) => (
                  <li key={index} className="flex items-start space-x-2 text-sm">
                    <div className="w-1.5 h-1.5 rounded-full bg-primary mt-2 flex-shrink-0" />
                    <span>{req}</span>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Submission Form */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center">
            <Send className="mr-2 h-5 w-5" />
            Submit Your Solution
          </CardTitle>
          <CardDescription>
            Submit your code for automated review and feedback
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div>
            <label className="text-sm font-medium mb-2 block">
              Additional Notes (Optional)
            </label>
            <Textarea
              placeholder="Add any notes about your implementation, challenges faced, or questions..."
              value={notes}
              onChange={(e) => setNotes(e.target.value)}
              className="min-h-[100px]"
            />
          </div>

          <div className="flex items-center justify-between">
            <div className="text-sm text-muted-foreground">
              Code will be automatically evaluated against test cases
            </div>
            <Button
              onClick={handleSubmit}
              disabled={isSubmitting || !code.trim()}
              className="go-gradient text-white"
            >
              {isSubmitting ? (
                <>
                  <Clock className="mr-2 h-4 w-4 animate-spin" />
                  Submitting...
                </>
              ) : (
                <>
                  <Send className="mr-2 h-4 w-4" />
                  Submit Solution
                </>
              )}
            </Button>
          </div>
        </CardContent>
      </Card>

      {/* Feedback */}
      {feedback && (
        <Card className="border-primary/20 bg-primary/5">
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle className="flex items-center">
                {feedback.passed ? (
                  <CheckCircle className="mr-2 h-5 w-5 text-green-500" />
                ) : (
                  <AlertCircle className="mr-2 h-5 w-5 text-red-500" />
                )}
                Submission Feedback
              </CardTitle>
              <div className="flex items-center space-x-2">
                <Badge className={getScoreBadge(feedback.score, feedback.maxScore).color}>
                  {getScoreBadge(feedback.score, feedback.maxScore).text}
                </Badge>
                <div className={`text-2xl font-bold ${getScoreColor(feedback.score, feedback.maxScore)}`}>
                  {feedback.score}/{feedback.maxScore}
                </div>
              </div>
            </div>
          </CardHeader>
          <CardContent className="space-y-6">
            {/* Score Progress */}
            <div className="space-y-2">
              <div className="flex justify-between text-sm">
                <span>Score</span>
                <span>{Math.round((feedback.score / feedback.maxScore) * 100)}%</span>
              </div>
              <Progress value={(feedback.score / feedback.maxScore) * 100} className="h-2" />
            </div>

            {/* Feedback Sections */}
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              {/* Comments */}
              <div className="space-y-3">
                <h4 className="font-medium flex items-center">
                  <MessageSquare className="mr-2 h-4 w-4" />
                  Comments
                </h4>
                <div className="space-y-2">
                  {feedback.comments.map((comment, index) => (
                    <div key={index} className="text-sm p-3 bg-background rounded-lg border">
                      {comment}
                    </div>
                  ))}
                </div>
              </div>

              {/* Suggestions */}
              <div className="space-y-3">
                <h4 className="font-medium flex items-center">
                  <ThumbsUp className="mr-2 h-4 w-4" />
                  Suggestions
                </h4>
                <div className="space-y-2">
                  {feedback.suggestions.map((suggestion, index) => (
                    <div key={index} className="text-sm p-3 bg-yellow-50 dark:bg-yellow-950 rounded-lg border border-yellow-200 dark:border-yellow-800">
                      {suggestion}
                    </div>
                  ))}
                </div>
              </div>

              {/* Strengths */}
              <div className="space-y-3">
                <h4 className="font-medium flex items-center">
                  <Star className="mr-2 h-4 w-4" />
                  Strengths
                </h4>
                <div className="space-y-2">
                  {feedback.strengths.map((strength, index) => (
                    <div key={index} className="text-sm p-3 bg-green-50 dark:bg-green-950 rounded-lg border border-green-200 dark:border-green-800">
                      {strength}
                    </div>
                  ))}
                </div>
              </div>
            </div>

            {/* Action Buttons */}
            <div className="flex items-center justify-between pt-4 border-t">
              <div className="flex items-center space-x-2">
                {feedback.passed && (
                  <Badge className="bg-green-500 text-white">
                    <Trophy className="mr-1 h-3 w-3" />
                    Exercise Completed!
                  </Badge>
                )}
              </div>
              <div className="flex items-center space-x-2">
                <Button variant="outline" size="sm">
                  <ThumbsUp className="mr-2 h-3 w-3" />
                  Helpful
                </Button>
                <Button variant="outline" size="sm">
                  <ThumbsDown className="mr-2 h-3 w-3" />
                  Not Helpful
                </Button>
              </div>
            </div>
          </CardContent>
        </Card>
      )}

      {/* Previous Submissions */}
      {previousSubmissions.length > 0 && (
        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle className="text-lg">Previous Submissions</CardTitle>
              <Button
                variant="outline"
                size="sm"
                onClick={() => setShowPreviousSubmissions(!showPreviousSubmissions)}
              >
                {showPreviousSubmissions ? "Hide" : "Show"} History
              </Button>
            </div>
          </CardHeader>
          {showPreviousSubmissions && (
            <CardContent>
              <div className="space-y-3">
                {previousSubmissions.map((submission) => (
                  <div
                    key={submission.id}
                    className="flex items-center justify-between p-3 bg-muted/50 rounded-lg"
                  >
                    <div>
                      <div className="font-medium text-sm">
                        Submission {submission.id}
                      </div>
                      <div className="text-xs text-muted-foreground">
                        {submission.timestamp}
                      </div>
                    </div>
                    <div className="flex items-center space-x-2">
                      <Badge variant="outline">
                        Score: {submission.score}
                      </Badge>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          )}
        </Card>
      )}
    </div>
  );
};

export default ExerciseSubmission;
