# URL Shortener Web

## Tech Features

- Vue 3
- **Fully typed** [Vuex 4](https://next.vuex.vuejs.org/) store
- Routing using [vue-router 4](https://next.router.vuejs.org/)
- TypeScript 4.4
- Tailwind CSS 2.2 w/ JIT compiler + following plugins preinstalled
  - `@tailwindcss/aspect-ratio`
  - `@tailwindcss/line-clamp`
  - `@tailwindcss/typography`
  - `@tailwindcss/forms`
  - `firefox`-variant
- PostCSS 8 w/ `postcss-nesting` plugin
- Eslint
- Prettier
- Alias `@` to `<project_root>/src`
- Manually configured global components in `main.ts`
- Predefined and fully typed global variables:
  - `VITE_APP_VERSION` is read from `package.json` version at build time
  - `VITE_APP_BUILD_EPOCH` is populated as `new Date().getTime()` at build time
- Using newest `script setup` syntax w/ Ref sugar (see the official [Script Setup documentation](https://v3.vuejs.org/api/sfc-script-setup.html) and [Ref Sugar RFC](https://github.com/vuejs/rfcs/discussions/369) discussion)
- Cypress.io e2e tests (configured similarly to `vue-cli`)
- Cypress.io component tests
- GitHub workflows
  - Dependabot
  - Automated e2e tests
  - Automated component tests
- GitLab CI
  - Automated e2e tests
  - Automated component tests 
