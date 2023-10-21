# Playwright Demo

A small demonstration of [Playwright](https://playwright.dev/) testing for web pages and APIs.

## What's covered by this repository

* Web page tests
	* Simple find one thing on the page
		* Fixtures - https://playwright.dev/docs/test-fixtures
		* Intro assertions - https://playwright.dev/docs/test-assertions
	* Chain of actions: fill input, click button, and check result
		* Intro actions - https://playwright.dev/docs/input
		* Explain auto awaiting - https://playwright.dev/docs/actionability
* [API tests](https://playwright.dev/docs/api-testing)
	* Simple API test
	* Simple test with Bearer authentication - https://playwright.dev/docs/auth
	* Chain of actions: GET, POST based on GET, GET, and check the updated result
* Handy commands/examples
	* VS Code extension
	* Generating code from VS Code
	* Commands for running Playwright code gen and tracing locally without VS Code - `npx playwright test --ui`
	* Commands for running Playwright tests without VS Code - `npx playwright test`
