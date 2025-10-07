"use client";

import { useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import ProgressTracker from "./progress-tracker";
import {
  BookOpen,
  Code2,
  Trophy,
  Play,
  ArrowRight,
  Clock,
  CheckCircle,
  Star,
  Calendar,
  TrendingUp,
  Target,
  Zap
} from "lucide-react";
import Link from "next/link";

interface RecentActivity {
  id: string;
  type: "lesson" | "exercise" | "achievement";
  title: string;
  description: string;
  timestamp: string;
  icon: any;
  color: string;
}

interface LearningDashboardProps {
  userId?: string;
}

export function LearningDashboard({ userId = "demo-user" }: LearningDashboardProps) {
  const [activeTab, setActiveTab] = useState("overview");

  const recentActivities: RecentActivity[] = [
    {
      id: "1",
      type: "lesson",
      title: "Completed Lesson 2",
      description: "Variables, Constants, and Functions",
      timestamp: "2 hours ago",
      icon: CheckCircle,
      color: "text-green-500",
    },
    {
      id: "2",
      type: "exercise",
      title: "Solved Exercise",
      description: "Function Implementation Challenge",
      timestamp: "3 hours ago",
      icon: Code2,
      color: "text-blue-500",
    },
    {
      id: "3",
      type: "achievement",
      title: "Earned Achievement",
      description: "Code Warrior - Write 100 lines of Go code",
      timestamp: "1 day ago",
      icon: Trophy,
      color: "text-yellow-500",
    },
    {
      id: "4",
      type: "lesson",
      title: "Started Lesson 3",
      description: "Control Structures and Loops",
      timestamp: "2 days ago",
      icon: Play,
      color: "text-primary",
    },
  ];

  const upcomingLessons = [
    {
      id: 4,
      title: "Arrays, Slices, and Maps",
      description: "Data structures, manipulation, memory considerations",
      duration: "5-6 hours",
      difficulty: "Beginner",
      locked: false,
    },
    {
      id: 5,
      title: "Pointers and Memory Management",
      description: "Pointer basics, memory allocation, garbage collection",
      duration: "4-5 hours",
      difficulty: "Beginner",
      locked: true,
    },
    {
      id: 6,
      title: "Structs and Methods",
      description: "Struct definition, methods, receivers, method sets",
      duration: "5-6 hours",
      difficulty: "Intermediate",
      locked: true,
    },
  ];

  const weeklyGoal = {
    target: 5,
    completed: 3,
    description: "Complete 5 lessons this week",
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-background via-background to-accent/5">
      <div className="container-responsive padding-responsive-y">
        <div className="margin-responsive">
          <h1 className="text-responsive-heading font-bold tracking-tight mb-3 bg-gradient-to-r from-primary to-primary/70 bg-clip-text text-transparent">
            Learning Dashboard
          </h1>
          <p className="text-responsive-body text-muted-foreground max-w-2xl">
            Track your progress and continue your Go programming journey
          </p>
        </div>

      <Tabs value={activeTab} onValueChange={setActiveTab} className="space-y-6">
        <TabsList className="grid w-full grid-cols-3 lg:w-[400px]">
          <TabsTrigger value="overview">Overview</TabsTrigger>
          <TabsTrigger value="progress">Progress</TabsTrigger>
          <TabsTrigger value="activity">Activity</TabsTrigger>
        </TabsList>

        {/* Overview Tab */}
        <TabsContent value="overview" className="space-y-6">
          {/* Quick Stats */}
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
            <Card>
              <CardContent className="p-4 text-center">
                <BookOpen className="h-6 w-6 text-blue-500 mx-auto mb-2" />
                <div className="text-2xl font-bold">3</div>
                <div className="text-sm text-muted-foreground">Lessons Completed</div>
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-4 text-center">
                <Code2 className="h-6 w-6 text-green-500 mx-auto mb-2" />
                <div className="text-2xl font-bold">18</div>
                <div className="text-sm text-muted-foreground">Exercises Solved</div>
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-4 text-center">
                <Star className="h-6 w-6 text-yellow-500 mx-auto mb-2" />
                <div className="text-2xl font-bold">350</div>
                <div className="text-sm text-muted-foreground">XP Earned</div>
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-4 text-center">
                <Zap className="h-6 w-6 text-orange-500 mx-auto mb-2" />
                <div className="text-2xl font-bold">5</div>
                <div className="text-sm text-muted-foreground">Day Streak</div>
              </CardContent>
            </Card>
          </div>

          {/* Weekly Goal */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center">
                <Target className="mr-2 h-5 w-5 text-primary" />
                Weekly Goal
              </CardTitle>
              <CardDescription>{weeklyGoal.description}</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="flex items-center justify-between mb-2">
                <span className="text-sm font-medium">
                  {weeklyGoal.completed} of {weeklyGoal.target} lessons
                </span>
                <span className="text-sm text-muted-foreground">
                  {Math.round((weeklyGoal.completed / weeklyGoal.target) * 100)}%
                </span>
              </div>
              <Progress value={(weeklyGoal.completed / weeklyGoal.target) * 100} />
            </CardContent>
          </Card>

          {/* Continue Learning */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center">
                <Play className="mr-2 h-5 w-5 text-primary" />
                Continue Learning
              </CardTitle>
              <CardDescription>Pick up where you left off</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-3">
                {upcomingLessons.slice(0, 2).map((lesson) => (
                  <div key={lesson.id} className="flex items-center justify-between p-3 rounded-lg bg-muted/50">
                    <div className="flex items-center space-x-3">
                      <div className="p-2 rounded-lg bg-primary/10">
                        <BookOpen className="h-4 w-4 text-primary" />
                      </div>
                      <div>
                        <div className="font-medium">Lesson {lesson.id}: {lesson.title}</div>
                        <div className="text-sm text-muted-foreground">{lesson.description}</div>
                        <div className="flex items-center space-x-2 mt-1">
                          <Badge variant="outline" className="text-xs">
                            {lesson.difficulty}
                          </Badge>
                          <div className="flex items-center space-x-1 text-xs text-muted-foreground">
                            <Clock className="h-3 w-3" />
                            <span>{lesson.duration}</span>
                          </div>
                        </div>
                      </div>
                    </div>
                    <Link href={lesson.locked ? "#" : `/learn/lesson-${lesson.id}`}>
                      <Button size="sm" disabled={lesson.locked}>
                        {lesson.locked ? "Locked" : "Continue"}
                        {!lesson.locked && <ArrowRight className="ml-2 h-4 w-4" />}
                      </Button>
                    </Link>
                  </div>
                ))}
              </div>
              <div className="mt-4 text-center">
                <Link href="/curriculum">
                  <Button variant="outline">
                    View Full Curriculum
                    <ArrowRight className="ml-2 h-4 w-4" />
                  </Button>
                </Link>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        {/* Progress Tab */}
        <TabsContent value="progress">
          <ProgressTracker userId={userId} />
        </TabsContent>

        {/* Activity Tab */}
        <TabsContent value="activity" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center">
                <TrendingUp className="mr-2 h-5 w-5 text-primary" />
                Recent Activity
              </CardTitle>
              <CardDescription>Your learning activity over the past week</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {recentActivities.map((activity) => (
                  <div key={activity.id} className="flex items-start space-x-3 p-3 rounded-lg bg-muted/50">
                    <div className="p-2 rounded-lg bg-background">
                      <activity.icon className={`h-4 w-4 ${activity.color}`} />
                    </div>
                    <div className="flex-1">
                      <div className="font-medium">{activity.title}</div>
                      <div className="text-sm text-muted-foreground">{activity.description}</div>
                      <div className="flex items-center space-x-1 text-xs text-muted-foreground mt-1">
                        <Calendar className="h-3 w-3" />
                        <span>{activity.timestamp}</span>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
      </div>
    </div>
  );
}

export default LearningDashboard;
