import { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  schema: "../graphql/*",
  documents: ["src/**/*.ts", "src/**/*.tsx"],
  ignoreNoDocuments: false,
  generates: {
    "./src/shared/gql/": {
      preset: "client",
    },
  },
};

export default config;
