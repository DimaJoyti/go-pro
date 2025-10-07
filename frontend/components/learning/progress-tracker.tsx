"use client";

import { useState, useEffect } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import {
  CheckCircle,
  Clock,
  Trophy,
  Target,
  Star,
  TrendingUp,
  BookOpen,
  Code2,
  Award,
  Zap
} from "lucide-react";

interface ProgressStats {
  totalLessons: number;
  completedLessons: number;
  totalExercises: number;
  completedExercises: number;
  totalProjects: number;
  completedProjects: number;
  totalXP: number;
  currentStreak: number;
  achievements: number;
}

interface ProgressTrackerProps {
  userId?: string;
  className?: string;
}

export function ProgressTracker({ userId = "demo-user", className = "" }: ProgressTrackerProps) {
  const [stats, setStats] = useState<ProgressStats>({
    totalLessons: 15,
    completedLessons: 3,
    totalExercises: 120,
    completedExercises: 18,
    totalProjects: 4,
    completedProjects: 0,
    totalXP: 350,
    currentStreak: 5,
    achievements: 4,
  });

  const [loading, setLoading] = useState(false);

  const progressPercentage = Math.round((stats.completedLessons / stats.totalLessons) * 100);
  const exerciseProgress = Math.round((stats.completedExercises / stats.totalExercises) * 100);

  const progressItems = [
    {
      icon: BookOpen,
      label: "Lessons",
      current: stats.completedLessons,
      total: stats.totalLessons,
      color: "text-blue-500",
      bgColor: "bg-blue-50 dark:bg-blue-950",
      borderColor: "border-blue-200 dark:border-blue-800",
    },
    {
      icon: Code2,
      label: "Exercises",
      current: stats.completedExercises,
      total: stats.totalExercises,
      color: "text-green-500",
      bgColor: "bg-green-50 dark:bg-green-950",
      borderColor: "border-green-200 dark:border-green-800",
    },
    {
      icon: Trophy,
      label: "Projects",
      current: stats.completedProjects,
      total: stats.totalProjects,
      color: "text-yellow-500",
      bgColor: "bg-yellow-50 dark:bg-yellow-950",
      borderColor: "border-yellow-200 dark:border-yellow-800",
    },
    {
      icon: Award,
      label: "Achievements",
      current: stats.achievements,
      total: 20,
      color: "text-purple-500",
      bgColor: "bg-purple-50 dark:bg-purple-950",
      borderColor: "border-purple-200 dark:border-purple-800",
    },
  ];

  return (
    <div className={`space-y-6 ${className}`}>
      {/* Overall Progress */}
      <Card>
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle className="flex items-center">
                <TrendingUp className="mr-2 h-5 w-5 text-primary" />
                Learning Progress
              </CardTitle>
              <CardDescription>Your journey through GO-PRO curriculum</CardDescription>
            </div>
            <div className="text-right">
              <div className="text-2xl font-bold text-primary">{progressPercentage}%</div>
              <div className="text-sm text-muted-foreground">Complete</div>
            </div>
          </div>
        </CardHeader>
        <CardContent>
          <Progress value={progressPercentage} className="mb-4" />
          <div className="flex items-center justify-between text-sm text-muted-foreground">
            <span>{stats.completedLessons} of {stats.totalLessons} lessons completed</span>
            <span>{stats.totalLessons - stats.completedLessons} lessons remaining</span>
          </div>
        </CardContent>
      </Card>

      {/* Progress Grid */}
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
        {progressItems.map((item, index) => (
          <Card key={index} className={`${item.bgColor} ${item.borderColor}`}>
            <CardContent className="p-4 text-center">
              <div className="flex justify-center mb-2">
                <item.icon className={`h-6 w-6 ${item.color}`} />
              </div>
              <div className="text-2xl font-bold">{item.current}</div>
              <div className="text-sm text-muted-foreground">
                of {item.total} {item.label}
              </div>
              <div className="mt-2">
                <Progress 
                  value={(item.current / item.total) * 100} 
                  className="h-1"
                />
              </div>
            </CardContent>
          </Card>
        ))}
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Card className="bg-gradient-to-r from-primary/5 to-primary/10 border-primary/20">
          <CardContent className="p-4 text-center">
            <div className="flex justify-center mb-2">
              <Star className="h-6 w-6 text-yellow-500" />
            </div>
            <div className="text-2xl font-bold">{stats.totalXP}</div>
            <div className="text-sm text-muted-foreground">Total XP</div>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-r from-orange-50 to-orange-100 dark:from-orange-950 dark:to-orange-900 border-orange-200 dark:border-orange-800">
          <CardContent className="p-4 text-center">
            <div className="flex justify-center mb-2">
              <Zap className="h-6 w-6 text-orange-500" />
            </div>
            <div className="text-2xl font-bold">{stats.currentStreak}</div>
            <div className="text-sm text-muted-foreground">Day Streak</div>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-r from-green-50 to-green-100 dark:from-green-950 dark:to-green-900 border-green-200 dark:border-green-800">
          <CardContent className="p-4 text-center">
            <div className="flex justify-center mb-2">
              <Target className="h-6 w-6 text-green-500" />
            </div>
            <div className="text-2xl font-bold">{exerciseProgress}%</div>
            <div className="text-sm text-muted-foreground">Exercise Progress</div>
          </CardContent>
        </Card>
      </div>

      {/* Next Steps */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center">
            <Clock className="mr-2 h-5 w-5 text-primary" />
            Next Steps
          </CardTitle>
          <CardDescription>Continue your learning journey</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            <div className="flex items-center justify-between p-3 rounded-lg bg-muted/50">
              <div className="flex items-center space-x-3">
                <div className="p-2 rounded-lg bg-primary/10">
                  <BookOpen className="h-4 w-4 text-primary" />
                </div>
                <div>
                  <div className="font-medium">Continue Lesson 4</div>
                  <div className="text-sm text-muted-foreground">Arrays, Slices, and Maps</div>
                </div>
              </div>
              <Button size="sm">Continue</Button>
            </div>

            <div className="flex items-center justify-between p-3 rounded-lg bg-muted/50">
              <div className="flex items-center space-x-3">
                <div className="p-2 rounded-lg bg-green-100 dark:bg-green-900">
                  <Code2 className="h-4 w-4 text-green-600 dark:text-green-400" />
                </div>
                <div>
                  <div className="font-medium">Practice Exercise</div>
                  <div className="text-sm text-muted-foreground">Variable Declaration Challenge</div>
                </div>
              </div>
              <Button size="sm" variant="outline">Practice</Button>
            </div>

            <div className="flex items-center justify-between p-3 rounded-lg bg-muted/50 opacity-60">
              <div className="flex items-center space-x-3">
                <div className="p-2 rounded-lg bg-yellow-100 dark:bg-yellow-900">
                  <Trophy className="h-4 w-4 text-yellow-600 dark:text-yellow-400" />
                </div>
                <div>
                  <div className="font-medium">First Project</div>
                  <div className="text-sm text-muted-foreground">CLI Task Manager</div>
                </div>
              </div>
              <Badge variant="outline">Locked</Badge>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

export default ProgressTracker;
