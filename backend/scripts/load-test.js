// k6 Load Testing Script for GO-PRO Backend API
// Run with: k6 run --vus 10 --duration 30s load-test.js

import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate, Trend, Counter } from 'k6/metrics';

// Custom metrics
export const errorRate = new Rate('errors');
export const responseTime = new Trend('response_time');
export const requests = new Counter('requests');

// Test configuration
export const options = {
  stages: [
    { duration: '1m', target: 5 },   // Ramp-up
    { duration: '3m', target: 10 },  // Stay at 10 users
    { duration: '1m', target: 20 },  // Ramp-up to 20 users
    { duration: '2m', target: 20 },  // Stay at 20 users
    { duration: '1m', target: 0 },   // Ramp-down
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% of requests must complete below 500ms
    http_req_failed: ['rate<0.05'],   // Error rate must be below 5%
    errors: ['rate<0.05'],            // Custom error rate must be below 5%
  },
};

const BASE_URL = 'http://localhost:8080';

// Test data
const testUsers = ['user1', 'user2', 'user3', 'user4', 'user5'];
const testCourses = ['go-pro'];
const testLessons = ['lesson-01', 'lesson-02', 'lesson-03'];
const testExercises = ['exercise-01-01', 'exercise-01-02'];

export default function () {
  const userId = testUsers[Math.floor(Math.random() * testUsers.length)];
  const courseId = testCourses[Math.floor(Math.random() * testCourses.length)];
  const lessonId = testLessons[Math.floor(Math.random() * testLessons.length)];
  const exerciseId = testExercises[Math.floor(Math.random() * testExercises.length)];

  // Test health endpoint
  testHealthCheck();

  // Test course endpoints
  testGetCourses();
  testGetCourse(courseId);
  testGetCourseLessons(courseId);

  // Test lesson endpoints
  testGetLesson(lessonId);

  // Test exercise endpoints
  testGetExercise(exerciseId);
  testSubmitExercise(exerciseId);

  // Test progress endpoints
  testGetProgress(userId);
  testUpdateProgress(userId, lessonId);

  sleep(1); // Wait 1 second between iterations
}

function testHealthCheck() {
  const response = http.get(`${BASE_URL}/api/v1/health`);

  const success = check(response, {
    'health check status is 200': (r) => r.status === 200,
    'health check has correct content-type': (r) =>
      r.headers['Content-Type'].includes('application/json'),
    'health check response time < 100ms': (r) => r.timings.duration < 100,
    'health check has status field': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.status === 'healthy';
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testGetCourses() {
  const response = http.get(`${BASE_URL}/api/v1/courses`);

  const success = check(response, {
    'get courses status is 200': (r) => r.status === 200,
    'get courses has correct content-type': (r) =>
      r.headers['Content-Type'].includes('application/json'),
    'get courses response time < 200ms': (r) => r.timings.duration < 200,
    'get courses has data array': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && Array.isArray(body.data);
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testGetCourse(courseId) {
  const response = http.get(`${BASE_URL}/api/v1/courses/${courseId}`);

  const success = check(response, {
    'get course status is 200': (r) => r.status === 200,
    'get course response time < 200ms': (r) => r.timings.duration < 200,
    'get course has valid data': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && body.data && body.data.id === courseId;
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testGetCourseLessons(courseId) {
  const response = http.get(`${BASE_URL}/api/v1/courses/${courseId}/lessons`);

  const success = check(response, {
    'get course lessons status is 200': (r) => r.status === 200,
    'get course lessons response time < 200ms': (r) => r.timings.duration < 200,
    'get course lessons has data array': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && Array.isArray(body.data);
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testGetLesson(lessonId) {
  const response = http.get(`${BASE_URL}/api/v1/lessons/${lessonId}`);

  const success = check(response, {
    'get lesson status is 200': (r) => r.status === 200,
    'get lesson response time < 200ms': (r) => r.timings.duration < 200,
    'get lesson has valid data': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && body.data && body.data.id === lessonId;
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testGetExercise(exerciseId) {
  const response = http.get(`${BASE_URL}/api/v1/exercises/${exerciseId}`);

  const success = check(response, {
    'get exercise status is 200': (r) => r.status === 200,
    'get exercise response time < 200ms': (r) => r.timings.duration < 200,
    'get exercise has valid data': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && body.data && body.data.id === exerciseId;
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testSubmitExercise(exerciseId) {
  const payload = {
    code: 'package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("Hello, World!")\n}',
    language: 'go'
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const response = http.post(
    `${BASE_URL}/api/v1/exercises/${exerciseId}/submit`,
    JSON.stringify(payload),
    params
  );

  const success = check(response, {
    'submit exercise status is 200': (r) => r.status === 200,
    'submit exercise response time < 500ms': (r) => r.timings.duration < 500,
    'submit exercise has valid result': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && body.data && typeof body.data.score === 'number';
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testGetProgress(userId) {
  const response = http.get(`${BASE_URL}/api/v1/progress/${userId}`);

  const success = check(response, {
    'get progress status is 200': (r) => r.status === 200,
    'get progress response time < 200ms': (r) => r.timings.duration < 200,
    'get progress has data array': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && Array.isArray(body.data);
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function testUpdateProgress(userId, lessonId) {
  const payload = {
    completed: true,
    score: Math.floor(Math.random() * 100) + 1
  };

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const response = http.post(
    `${BASE_URL}/api/v1/progress/${userId}/lesson/${lessonId}`,
    JSON.stringify(payload),
    params
  );

  const success = check(response, {
    'update progress status is 200': (r) => r.status === 200,
    'update progress response time < 300ms': (r) => r.timings.duration < 300,
    'update progress has valid data': (r) => {
      try {
        const body = JSON.parse(r.body);
        return body.success === true && body.data && body.data.completed === payload.completed;
      } catch (e) {
        return false;
      }
    },
  });

  recordMetrics(response, success);
}

function recordMetrics(response, success) {
  requests.add(1);
  responseTime.add(response.timings.duration);
  errorRate.add(!success);
}

// Setup function to run once before the test
export function setup() {
  console.log('Starting load test for GO-PRO Backend API');
  console.log(`Target URL: ${BASE_URL}`);

  // Check if API is available
  const response = http.get(`${BASE_URL}/api/v1/health`);
  if (response.status !== 200) {
    throw new Error(`API is not available. Status: ${response.status}`);
  }

  console.log('API health check passed, starting load test...');
}

// Teardown function to run once after the test
export function teardown() {
  console.log('Load test completed');
}