exports.apps = [
  {
      name: "auth-svc-http",
      script: "main.go",
      args: "serve-http",
      interpreter: "go",
      interpreter_args: "run",
      watch: "*.go", // auto-restart the application once the source code is changed
      env: {
        APP_ENV: 'staging',
      }
  },
  {
      name: "auth-svc-grpc",
      script: "main.go",
      args: "serve-grpc",
      interpreter: "go",
      interpreter_args: "run",
      watch: "*.go", // auto-restart the application once the source code is changed
      env: {
        APP_ENV: 'staging',
      }
  }
];