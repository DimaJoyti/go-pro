# Learning Dashboard - Code Examples

## 🎨 Key Design Patterns & Code Examples

### 1. Enhanced Stat Card with Gradient & Hover Effects

**Before:**
```tsx
<Card>
  <CardContent className="p-4 text-center">
    <BookOpen className="h-6 w-6 text-blue-500 mx-auto mb-2" />
    <div className="text-2xl font-bold">3</div>
    <div className="text-sm text-muted-foreground">Lessons Completed</div>
  </CardContent>
</Card>
```

**After:**
```tsx
<Card className="group relative overflow-hidden border-blue-200/50 dark:border-blue-800/50 hover:shadow-lg hover:shadow-blue-500/10 transition-all duration-300 hover:-translate-y-1">
  <div className="absolute inset-0 bg-gradient-to-br from-blue-500/5 via-transparent to-transparent" />
  <CardContent className="relative p-6">
    <div className="flex items-center justify-between mb-3">
      <div className="p-3 rounded-xl bg-gradient-to-br from-blue-500/20 to-blue-500/10 group-hover:from-blue-500/30 group-hover:to-blue-500/20 transition-all duration-300">
        <BookOpen className="h-6 w-6 text-blue-600 dark:text-blue-400" />
      </div>
      <Sparkles className="h-4 w-4 text-blue-400 opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
    </div>
    <div className="text-3xl font-bold mb-1 bg-gradient-to-br from-blue-600 to-blue-500 bg-clip-text text-transparent">3</div>
    <div className="text-sm font-medium text-muted-foreground">Lessons Completed</div>
  </CardContent>
</Card>
```

**Key Improvements:**
- ✨ Gradient background overlay
- 🎯 Hover lift effect (`hover:-translate-y-1`)
- 💫 Shadow with color glow (`hover:shadow-blue-500/10`)
- 🎨 Icon in gradient container
- ✨ Sparkle icon appears on hover
- 📊 Gradient text for number
- 🎪 Larger padding and font sizes

---

### 2. Premium Lesson Card with Multiple Layers

**Before:**
```tsx
<div className="group flex items-center justify-between p-4 rounded-xl bg-gradient-to-br from-muted/50 to-muted/30 border border-border/50 hover:border-primary/30 hover:shadow-md transition-all duration-300">
  <div className="flex items-start space-x-4 flex-1">
    <div className="p-3 rounded-xl bg-gradient-to-br from-primary/20 to-primary/10">
      <BookOpen className="h-5 w-5 text-primary" />
    </div>
    <div className="flex-1 space-y-2">
      <h3 className="font-semibold text-base">{lesson.title}</h3>
      <p className="text-sm text-muted-foreground">{lesson.description}</p>
    </div>
  </div>
  <Button size="sm">Continue</Button>
</div>
```

**After:**
```tsx
<div className="group relative overflow-hidden rounded-2xl bg-gradient-to-br from-muted/50 via-muted/30 to-background border border-border/50 hover:border-primary/40 hover:shadow-xl hover:shadow-primary/5 transition-all duration-500 hover:-translate-y-1">
  {/* Gradient overlay that appears on hover */}
  <div className="absolute inset-0 bg-gradient-to-br from-primary/5 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500" />
  
  <div className="relative flex items-center justify-between p-5">
    <div className="flex items-start space-x-4 flex-1">
      <div className="relative">
        <div className="p-4 rounded-2xl bg-gradient-to-br from-primary/20 via-primary/15 to-primary/10 group-hover:from-primary/30 group-hover:via-primary/25 group-hover:to-primary/20 transition-all duration-500 group-hover:scale-110">
          <BookOpen className="h-6 w-6 text-primary" />
        </div>
        {/* Active indicator */}
        <div className="absolute -top-1 -right-1 w-3 h-3 bg-green-500 rounded-full border-2 border-background animate-pulse" />
      </div>
      
      <div className="flex-1 space-y-3">
        <div className="flex items-center gap-2 flex-wrap">
          <span className="text-xs font-bold text-primary/80 tracking-wider uppercase px-2 py-1 rounded-md bg-primary/10">
            Lesson {lesson.id}
          </span>
          <Badge className="text-xs bg-green-500/10 text-green-700 dark:text-green-400 border-green-500/30">
            In Progress
          </Badge>
        </div>
        
        <h3 className="font-bold text-lg leading-tight group-hover:text-primary transition-colors duration-300">
          {lesson.title}
        </h3>
        
        <p className="text-sm text-muted-foreground leading-relaxed line-clamp-2">
          {lesson.description}
        </p>
        
        {/* Progress indicator */}
        <div className="flex items-center gap-2">
          <Progress value={35} className="w-20 h-1.5" />
          <span className="text-xs font-medium text-primary">35%</span>
        </div>
      </div>
    </div>
    
    <Button size="lg" className="group-hover:shadow-lg group-hover:shadow-primary/20 transition-all duration-300 group-hover:scale-105">
      Continue
      <ArrowRight className="ml-2 h-4 w-4 group-hover:translate-x-1 transition-transform" />
    </Button>
  </div>
</div>
```

**Key Improvements:**
- 🎨 Multi-layer gradient backgrounds
- 🟢 Active indicator with pulse animation
- 📊 Progress bar for in-progress lessons
- 🏷️ Enhanced badges with custom styling
- 🎯 Multiple hover effects (lift, shadow, scale)
- 💫 Smooth transitions (500ms)
- 📱 Better spacing and typography

---

### 3. Activity Timeline with Visual Connection

**Before:**
```tsx
<div className="space-y-4">
  {recentActivities.map((activity) => (
    <div key={activity.id} className="flex items-start space-x-3 p-3 rounded-lg bg-muted/50">
      <div className="p-2 rounded-lg bg-background">
        <activity.icon className={`h-4 w-4 ${activity.color}`} />
      </div>
      <div className="flex-1">
        <div className="font-medium">{activity.title}</div>
        <div className="text-sm text-muted-foreground">{activity.description}</div>
      </div>
    </div>
  ))}
</div>
```

**After:**
```tsx
<div className="relative space-y-4">
  {/* Timeline line */}
  <div className="absolute left-[29px] top-4 bottom-4 w-0.5 bg-gradient-to-b from-primary/50 via-primary/30 to-transparent" />
  
  {recentActivities.map((activity, index) => (
    <div 
      key={activity.id} 
      className="relative flex items-start space-x-4 p-4 rounded-xl bg-gradient-to-br from-muted/50 to-background border border-border/50 hover:border-primary/30 hover:shadow-lg transition-all duration-300 hover:-translate-y-0.5 group"
      style={{ animationDelay: `${index * 100}ms` }}
    >
      <div className="relative z-10">
        <div className={`p-3 rounded-xl bg-gradient-to-br ${
          activity.type === 'lesson' ? 'from-blue-500/20 to-blue-500/10' :
          activity.type === 'exercise' ? 'from-green-500/20 to-green-500/10' :
          'from-yellow-500/20 to-yellow-500/10'
        } group-hover:scale-110 transition-transform duration-300`}>
          <activity.icon className={`h-5 w-5 ${activity.color}`} />
        </div>
      </div>
      
      <div className="flex-1 min-w-0">
        <div className="flex items-start justify-between gap-2 mb-1">
          <div className="font-semibold text-base group-hover:text-primary transition-colors">
            {activity.title}
          </div>
          <Badge variant="outline" className="text-xs shrink-0">
            {activity.type}
          </Badge>
        </div>
        <div className="text-sm text-muted-foreground mb-2 leading-relaxed">
          {activity.description}
        </div>
        <div className="flex items-center gap-2 text-xs text-muted-foreground">
          <Calendar className="h-3.5 w-3.5" />
          <span className="font-medium">{activity.timestamp}</span>
        </div>
      </div>
    </div>
  ))}
</div>
```

**Key Improvements:**
- 📍 Visual timeline line connecting activities
- 🎨 Color-coded gradient backgrounds by type
- 🚀 Staggered entrance animations
- 💫 Icon scale effect on hover
- 🏷️ Type badges for categorization
- 🎯 Better typography and spacing

---

### 4. Enhanced Header with Quick Stats

**Before:**
```tsx
<div className="margin-responsive">
  <h1 className="text-responsive-heading font-bold tracking-tight mb-3 bg-gradient-to-r from-primary to-primary/70 bg-clip-text text-transparent">
    Learning Dashboard
  </h1>
  <p className="text-responsive-body text-muted-foreground max-w-2xl">
    Track your progress and continue your Go programming journey
  </p>
</div>
```

**After:**
```tsx
<div className="margin-responsive mb-8">
  <div className="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-6">
    <div className="space-y-3">
      <h1 className="text-responsive-heading font-bold tracking-tight bg-gradient-to-r from-primary via-primary/80 to-primary/60 bg-clip-text text-transparent animate-in fade-in slide-in-from-bottom-4 duration-700">
        Learning Dashboard
      </h1>
      <p className="text-responsive-body text-muted-foreground max-w-2xl animate-in fade-in slide-in-from-bottom-5 duration-700 delay-100">
        Track your progress and continue your Go programming journey
      </p>
    </div>
    
    {/* Quick Stats Preview */}
    <div className="flex items-center gap-4 animate-in fade-in slide-in-from-right-4 duration-700 delay-200">
      <div className="flex items-center gap-2 px-4 py-2 rounded-full bg-gradient-to-r from-orange-500/10 to-orange-500/5 border border-orange-500/20">
        <Flame className="h-4 w-4 text-orange-500" />
        <span className="text-sm font-bold text-orange-600 dark:text-orange-400">5 Day Streak</span>
      </div>
      <div className="flex items-center gap-2 px-4 py-2 rounded-full bg-gradient-to-r from-yellow-500/10 to-yellow-500/5 border border-yellow-500/20">
        <Star className="h-4 w-4 text-yellow-500" />
        <span className="text-sm font-bold text-yellow-600 dark:text-yellow-400">350 XP</span>
      </div>
    </div>
  </div>
</div>
```

**Key Improvements:**
- 🎭 Entrance animations with staggered delays
- 🔥 Quick stats badges in header
- 📱 Responsive flex layout
- 🎨 Gradient backgrounds for badges
- ✨ Smooth fade-in effects

---

## 🎯 Reusable Design Patterns

### Pattern 1: Gradient Icon Container
```tsx
<div className="p-3 rounded-xl bg-gradient-to-br from-primary/20 to-primary/10 group-hover:from-primary/30 group-hover:to-primary/20 transition-all duration-300 group-hover:scale-110">
  <Icon className="h-5 w-5 text-primary" />
</div>
```

### Pattern 2: Card with Gradient Overlay
```tsx
<Card className="group relative overflow-hidden">
  <div className="absolute inset-0 bg-gradient-to-br from-primary/5 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
  <CardContent className="relative">
    {/* Content */}
  </CardContent>
</Card>
```

### Pattern 3: Gradient Text
```tsx
<div className="text-3xl font-bold bg-gradient-to-br from-primary to-primary/70 bg-clip-text text-transparent">
  {value}
</div>
```

### Pattern 4: Hover Lift with Shadow
```tsx
<div className="hover:-translate-y-1 hover:shadow-lg hover:shadow-primary/10 transition-all duration-300">
  {/* Content */}
</div>
```

### Pattern 5: Pulse Indicator
```tsx
<div className="absolute -top-1 -right-1 w-3 h-3 bg-green-500 rounded-full border-2 border-background animate-pulse" />
```

---

## 🚀 Performance Tips

1. **Use GPU-accelerated properties:**
   - `transform` instead of `top/left`
   - `opacity` for fades
   - `scale` for size changes

2. **Optimize animations:**
   - Use `transition-all` sparingly
   - Specify exact properties when possible
   - Keep durations reasonable (300-500ms)

3. **Reduce motion support:**
   - Tailwind automatically handles `prefers-reduced-motion`
   - Animations are disabled for users who prefer reduced motion

---

## 📚 Resources

- **Tailwind CSS**: https://tailwindcss.com
- **Lucide Icons**: https://lucide.dev
- **shadcn/ui**: https://ui.shadcn.com
- **Next.js**: https://nextjs.org

---

## ✨ Summary

These code examples demonstrate how to create modern, beautiful UI components with:
- 🎨 Rich gradient effects
- 💫 Smooth animations
- 🎯 Clear visual hierarchy
- 📱 Responsive design
- ♿ Accessibility support
- 🚀 Performance optimization

All improvements maintain clean, readable code while delivering a premium user experience!

