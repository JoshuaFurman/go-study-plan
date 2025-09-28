# Agent Guidelines

## Commands
- **Build**: `npm run build` or `yarn build`
- **Lint**: `npm run lint` or `yarn lint`
- **Test all**: `npm test` or `yarn test`
- **Test single**: `npm test -- --testNamePattern="test name"` or `yarn test --testNamePattern="test name"`
- **Type check**: `npm run typecheck` or `yarn typecheck` (if available)

## Code Style
- **Imports**: Group by external libraries first, then internal modules. Use absolute imports.
- **Formatting**: Use Prettier with default settings. 2-space indentation.
- **Types**: Use TypeScript with strict mode. Prefer interfaces over types for objects.
- **Naming**: camelCase for variables/functions, PascalCase for classes/components, UPPER_SNAKE_CASE for constants.
- **Error handling**: Use try/catch for async operations. Throw descriptive Error objects.
- **Comments**: Add JSDoc for public APIs. Avoid inline comments for obvious code.
- **Testing**: Write unit tests for pure functions. Use descriptive test names with `describe`/`it` blocks.

## Additional Rules
- No Cursor or Copilot rules found in this repository.
- Follow existing patterns in the codebase when making changes.
- Run lint and tests before committing changes.