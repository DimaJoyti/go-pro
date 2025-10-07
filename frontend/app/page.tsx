import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  ArrowRight,
  BookOpen,
  Code2,
  Trophy,
  Users,
  Star,
  Play,
  CheckCircle,
  Clock,
  Target,
  Zap,
  Globe,
  Award,
  TrendingUp
} from "lucide-react";
import Link from "next/link";

export default function Home() {
  const stats = [
    { label: "Active Learners", value: "10,000+", icon: Users },
    { label: "Lessons Completed", value: "50,000+", icon: BookOpen },
    { label: "Code Challenges", value: "500+", icon: Code2 },
    { label: "Success Rate", value: "94%", icon: Trophy },
  ];

  const features = [
    {
      icon: BookOpen,
      title: "Interactive Lessons",
      description: "Learn Go through hands-on, interactive lessons that adapt to your pace.",
    },
    {
      icon: Code2,
      title: "Real Code Practice",
      description: "Write actual Go code with instant feedback and automated testing.",
    },
    {
      icon: Trophy,
      title: "Project-Based Learning",
      description: "Build real applications from CLI tools to microservices.",
    },
    {
      icon: Users,
      title: "Community Support",
      description: "Join thousands of Go developers in our supportive community.",
    },
    {
      icon: Target,
      title: "Personalized Path",
      description: "AI-powered learning paths tailored to your goals and experience.",
    },
    {
      icon: Award,
      title: "Industry Recognition",
      description: "Earn certificates recognized by top tech companies.",
    },
  ];

  const testimonials = [
    {
      name: "Sarah Chen",
      role: "Backend Developer at Google",
      avatar: "/avatars/sarah.svg",
      content: "GO-PRO transformed my understanding of Go. The hands-on approach made complex concepts click instantly.",
      rating: 5,
    },
    {
      name: "Marcus Rodriguez",
      role: "Senior Engineer at Uber",
      avatar: "/avatars/marcus.svg",
      content: "Best Go learning platform I've used. The project-based approach prepared me for real-world development.",
      rating: 5,
    },
    {
      name: "Emily Johnson",
      role: "DevOps Engineer at Netflix",
      avatar: "/avatars/emily.svg",
      content: "The microservices module was incredible. I went from beginner to building production systems.",
      rating: 5,
    },
  ];

  return (
    <div className="flex flex-col">
      {/* Hero Section */}
      <section
        className="relative overflow-hidden bg-gradient-to-br from-background via-background to-accent/5 hero-pattern"
        aria-labelledby="hero-title"
      >
        <div className="container-responsive py-16 sm:py-24 lg:py-32 xl:py-40">
          <div className="mx-auto max-w-5xl text-center">
            <Badge variant="secondary" className="mb-6 text-sm">
              🚀 New: Advanced Microservices Course Available
            </Badge>

            <h1
              id="hero-title"
              className="text-3xl font-bold tracking-tight sm:text-5xl lg:text-6xl xl:text-7xl mb-6"
            >
              Master{" "}
              <span className="go-gradient-text">Go Programming</span>
              <br />
              Through Practice
            </h1>

            <p className="text-sm sm:text-base lg:text-lg xl:text-xl text-muted-foreground mb-6 sm:mb-8 max-w-xl sm:max-w-2xl mx-auto leading-relaxed">
              Learn Go from basics to microservices with interactive lessons, real code practice,
              and projects that prepare you for production development.
            </p>

            <div className="flex flex-col sm:flex-row gap-3 sm:gap-4 justify-center mb-8 sm:mb-12">
              <Link href="/learn/lesson-1">
                <Button size="lg" className="go-gradient text-white text-sm sm:text-base lg:text-lg px-6 sm:px-8 py-4 sm:py-6 min-h-[48px] sm:min-h-[56px]">
                  <Play className="mr-2 h-4 w-4 sm:h-5 sm:w-5" />
                  Start Learning Free
                  <ArrowRight className="ml-2 h-5 w-5" />
                </Button>
              </Link>
              <Link href="/curriculum">
                <Button size="lg" variant="outline" className="text-sm sm:text-base lg:text-lg px-6 sm:px-8 py-4 sm:py-6 min-h-[48px] sm:min-h-[56px]">
                  <BookOpen className="mr-2 h-4 w-4 sm:h-5 sm:w-5" />
                  View Curriculum
                </Button>
              </Link>
            </div>

            {/* Stats */}
            <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 sm:gap-6 max-w-xs sm:max-w-2xl lg:max-w-4xl mx-auto">
              {stats.map((stat, index) => (
                <div key={`stat-${stat.label}-${index}`} className="text-center p-2 sm:p-3">
                  <div className="flex justify-center mb-1 sm:mb-2">
                    <stat.icon className="h-5 w-5 sm:h-6 sm:w-6 text-primary" />
                  </div>
                  <div className="text-lg sm:text-xl lg:text-2xl font-bold">{stat.value}</div>
                  <div className="text-xs sm:text-sm text-muted-foreground">{stat.label}</div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 sm:py-24 lg:py-32" aria-labelledby="features-title">
        <div className="container-responsive">
          <div className="mx-auto max-w-3xl text-center margin-responsive">
            <h2
              id="features-title"
              className="text-2xl font-bold tracking-tight sm:text-3xl lg:text-4xl mb-4"
            >
              Why Choose GO-PRO?
            </h2>
            <p className="text-lg text-muted-foreground">
              Our platform combines the best of interactive learning with real-world application
            </p>
          </div>

          <div className="grid-responsive-cards gap-responsive">
            {features.map((feature, index) => (
              <Card key={`feature-${feature.title}-${index}`} className="lesson-card h-full">
                <CardHeader className="pb-3 sm:pb-4">
                  <div className="flex items-center space-x-3">
                    <div className="flex h-8 w-8 sm:h-10 sm:w-10 items-center justify-center rounded-lg bg-primary/10 flex-shrink-0">
                      <feature.icon className="h-4 w-4 sm:h-5 sm:w-5 text-primary" />
                    </div>
                    <CardTitle className="text-base sm:text-lg lg:text-xl leading-tight">{feature.title}</CardTitle>
                  </div>
                </CardHeader>
                <CardContent className="pt-0">
                  <CardDescription className="text-sm sm:text-base leading-relaxed">
                    {feature.description}
                  </CardDescription>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* Learning Path Section */}
      <section className="py-20 sm:py-32 bg-accent/5">
        <div className="container max-w-screen-2xl px-4">
          <div className="mx-auto max-w-2xl text-center mb-16">
            <h2 className="text-3xl font-bold tracking-tight sm:text-4xl mb-4">
              Your Learning Journey
            </h2>
            <p className="text-lg text-muted-foreground">
              Structured curriculum that takes you from beginner to Go expert
            </p>
          </div>

          <Tabs defaultValue="beginner" className="max-w-4xl mx-auto">
            <TabsList className="grid w-full grid-cols-3">
              <TabsTrigger value="beginner">Beginner</TabsTrigger>
              <TabsTrigger value="intermediate">Intermediate</TabsTrigger>
              <TabsTrigger value="advanced">Advanced</TabsTrigger>
            </TabsList>

            <TabsContent value="beginner" className="mt-8">
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <Zap className="mr-2 h-5 w-5 text-primary" />
                    Go Fundamentals
                  </CardTitle>
                  <CardDescription>
                    Master the basics of Go programming language
                  </CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div className="space-y-3">
                      <div className="flex items-center space-x-2">
                        <CheckCircle className="h-4 w-4 text-green-500" />
                        <span className="text-sm">Variables & Data Types</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <CheckCircle className="h-4 w-4 text-green-500" />
                        <span className="text-sm">Control Structures</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <CheckCircle className="h-4 w-4 text-green-500" />
                        <span className="text-sm">Functions & Methods</span>
                      </div>
                    </div>
                    <div className="space-y-3">
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Structs & Interfaces</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Error Handling</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Package Management</span>
                      </div>
                    </div>
                  </div>
                  <Progress value={75} className="mt-4" />
                  <p className="text-sm text-muted-foreground">8 lessons • 2-3 weeks</p>
                </CardContent>
              </Card>
            </TabsContent>

            <TabsContent value="intermediate" className="mt-8">
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <Globe className="mr-2 h-5 w-5 text-primary" />
                    Web Development & APIs
                  </CardTitle>
                  <CardDescription>
                    Build web applications and REST APIs with Go
                  </CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div className="space-y-3">
                      <div className="flex items-center space-x-2">
                        <CheckCircle className="h-4 w-4 text-green-500" />
                        <span className="text-sm">HTTP Server Basics</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">REST API Design</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Database Integration</span>
                      </div>
                    </div>
                    <div className="space-y-3">
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Middleware & Auth</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Testing Strategies</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Deployment</span>
                      </div>
                    </div>
                  </div>
                  <Progress value={25} className="mt-4" />
                  <p className="text-sm text-muted-foreground">12 lessons • 4-5 weeks</p>
                </CardContent>
              </Card>
            </TabsContent>

            <TabsContent value="advanced" className="mt-8">
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <TrendingUp className="mr-2 h-5 w-5 text-primary" />
                    Microservices & Performance
                  </CardTitle>
                  <CardDescription>
                    Build scalable, production-ready Go applications
                  </CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div className="space-y-3">
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Concurrency Patterns</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Microservices Architecture</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">gRPC & Protocol Buffers</span>
                      </div>
                    </div>
                    <div className="space-y-3">
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Performance Optimization</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Monitoring & Observability</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Clock className="h-4 w-4 text-blue-500" />
                        <span className="text-sm">Production Deployment</span>
                      </div>
                    </div>
                  </div>
                  <Progress value={0} className="mt-4" />
                  <p className="text-sm text-muted-foreground">15 lessons • 6-8 weeks</p>
                </CardContent>
              </Card>
            </TabsContent>
          </Tabs>
        </div>
      </section>

      {/* Testimonials Section */}
      <section className="py-20 sm:py-32">
        <div className="container max-w-screen-2xl px-4">
          <div className="mx-auto max-w-2xl text-center mb-16">
            <h2 className="text-3xl font-bold tracking-tight sm:text-4xl mb-4">
              Loved by Developers Worldwide
            </h2>
            <p className="text-lg text-muted-foreground">
              Join thousands of developers who've advanced their careers with GO-PRO
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {testimonials.map((testimonial, index) => (
              <Card key={index} className="lesson-card">
                <CardHeader>
                  <div className="flex items-center space-x-3">
                    <Avatar>
                      <AvatarImage src={testimonial.avatar} alt={testimonial.name} />
                      <AvatarFallback>{testimonial.name.split(' ').map(n => n[0]).join('')}</AvatarFallback>
                    </Avatar>
                    <div>
                      <CardTitle className="text-base">{testimonial.name}</CardTitle>
                      <CardDescription className="text-sm">{testimonial.role}</CardDescription>
                    </div>
                  </div>
                  <div className="flex space-x-1">
                    {[...Array(testimonial.rating)].map((_, i) => (
                      <Star key={i} className="h-4 w-4 fill-yellow-400 text-yellow-400" />
                    ))}
                  </div>
                </CardHeader>
                <CardContent>
                  <p className="text-sm text-muted-foreground">"{testimonial.content}"</p>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 sm:py-32 bg-accent/5">
        <div className="container max-w-screen-2xl px-4">
          <div className="mx-auto max-w-2xl text-center">
            <h2 className="text-3xl font-bold tracking-tight sm:text-4xl mb-4">
              Ready to Master Go?
            </h2>
            <p className="text-lg text-muted-foreground mb-8">
              Join thousands of developers who are already building amazing things with Go.
              Start your journey today - it's completely free!
            </p>

            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Link href="/learn/lesson-1">
                <Button size="lg" className="go-gradient text-white text-lg px-8 py-6">
                  <Play className="mr-2 h-5 w-5" />
                  Start Learning Now
                  <ArrowRight className="ml-2 h-5 w-5" />
                </Button>
              </Link>
              <Link href="/curriculum">
                <Button size="lg" variant="outline" className="text-lg px-8 py-6">
                  <Users className="mr-2 h-5 w-5" />
                  View Curriculum
                </Button>
              </Link>
            </div>

            <p className="text-sm text-muted-foreground mt-6">
              No credit card required • 30-day money-back guarantee
            </p>
          </div>
        </div>
      </section>
    </div>
  );
}
