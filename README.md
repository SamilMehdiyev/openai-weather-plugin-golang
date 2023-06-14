# OpenAI Plugin Development with Golang: Weather Plugin

This README file provides a step-by-step guide on how to develop your own OpenAI plugin using Golang.

## Overview

OpenAI plugins are a way to extend the capabilities of ChatGPT by bridging the gap between ChatGPT and third-party applications. By developing a plugin, you can enable ChatGPT to perform a wide variety of actions by interacting with APIs.

o develop the plugin, we will follow these steps:

1. Create a plugin manifest file
2. Build an OpenAPI specification file
3. Develop the APIs
4. Register the plugin
5. Activate the plugin

## Step 1: Create a Plugin Manifest File

Every plugin requires an `ai-plugin.json` file hosted on the domain of the API. This file serves as the manifest for the plugin. When installing the plugin via the ChatGPT UI, ChatGPT checks for the manifest file at `/.well-known/ai-plugin.json` on your domain.

The manifest file should include the following key parameters:

- `schema_version`: The version of the manifest schema.
- `name_for_model`: The name of the model that will target the plugin.
- `name_for_human`: A human-readable name for the plugin.
- `description_for_model`: A detailed description of the model.
- `description_for_human`: A brief description of the plugin.
- `auth`: Authentication schema for interacting with the plugin.
- `api`: The URL path where the OpenAPI specification file is located.
- `logo_url`: The URL to fetch the plugin's logo.

## Step 2: Build an OpenAPI Specification File

The next step is to construct the OpenAPI specification file, which documents the API for the plugin. The OpenAPI Specification (OAS) provides a standard, language-agnostic interface to HTTP APIs.

In the specification file, you need to define at least one of the following: `paths`, `components`, or `webhooks`. These define the endpoints, reusable components, or webhook-based interactions of your plugin's API.

The specification file should include the necessary paths, request/response schemas, and other relevant details about the API endpoints.

## Step 3: Develop APIs

With the manifest and OpenAPI specification files in place, you can proceed to develop the APIs. In our Weather Plugin example, we will implement the following APIs:

- `/openapi.yaml`: A GET endpoint to fetch the OpenAPI specification file.
- `/.well-known/ai-plugin.json`: A GET endpoint to retrieve the manifest file.
- `/logo.png`: A GET endpoint to retrieve the plugin's logo image.
- `/wttr`: A POST endpoint that integrates with ChatGPT to provide weather information.

## Step 4: Plugin Registration

Before using the plugin in the ChatGPT UI, you need to register it. To register a plugin, you must have a paid subscription (Chat GPT Plus) and an account with access to register a plugin.

To register the plugin, follow these steps:

1. Execute the command `go run main.go` in the terminal to start the HTTP server.
2. Confirm that the server is running by visiting `http://localhost:5004` or using telnet.
3. In the ChatGPT UI, click on "Develop your own plugin" and add your domain (e.g., `localhost:5004`).
4. Click the "Find manifest file" button to fetch the manifest file.

Once registered, you can use the plugin within the ChatGPT UI.

## Conclusion

Feel free to explore, leave comments, and provide feedback. Stay safe and happy coding!
