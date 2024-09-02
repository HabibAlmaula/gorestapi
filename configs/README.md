# `config` Folder: Centralized Configuration Management

In our "Clean Code" architecture boilerplate, the `config` folder serves as the central repository for all configuration settings and parameters needed for your application. It's especially essential during development in a local environment.

## What Belongs in the `config` Folder?

1. **Configuration Files:** The `config` folder houses all configuration files, including settings for different environments (e.g., development, staging, production). These files can include database connection details, API keys, environment-specific settings, and more.

## Why the `config` Folder?

1. **Centralized Management:** Centralizing your configuration settings simplifies maintenance and helps keep all settings organized.

2. **Environment Flexibility:** Having environment-specific configuration files allows you to seamlessly switch between development, testing, and production environments.

3. **Security:** Separating configuration from code reduces the risk of sensitive information (like database credentials) being inadvertently exposed in your codebase.

## How to Use the `config` Folder

1. **Create Configuration Files:** Inside the `config` folder, create different configuration files for each environment (e.g., `development.yaml`, `production.yaml`). These files should contain the settings and parameters specific to each environment.

2. **Load Configuration:** In your application code, use libraries or methods to load configuration settings from the appropriate configuration file based on the current environment.

3. **Environment Variables:** Consider using environment variables to select the appropriate configuration file dynamically. This allows you to switch configurations easily between different environments.

4. **Security Best Practices:** Ensure that sensitive information, such as API keys and database credentials, are stored securely and not exposed in your codebase.

By using the `config` folder, you create a structured approach to managing configuration settings in a way that is secure, adaptable to different environments, and centralized for easy maintenance.