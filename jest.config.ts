import type {Config} from 'jest';

const config: Config = {
  clearMocks: true,
  coverageProvider: "v8",
  rootDir: "./",
  roots: [
    "<rootDir>/lang/javascript/test"
  ]
};

export default config;
