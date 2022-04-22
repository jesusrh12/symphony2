/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @format
 */

const fs = require('fs');
const path = require('path');
const {buildSchema, printSchema} = require('graphql');

// enforces copyright header to be present in every file
// eslint-disable-next-line max-len
const openSourcePattern = /\*\n \* Copyright 20\d{2}-present Facebook\. All Rights Reserved\.\n \*\n \* This source code is licensed under the BSD-style license found in the\n \* LICENSE file in the root directory of this source tree\.\n \*\n/;
// eslint-disable-next-line max-len
const newOpenSourcePattern = /Copyright \(c\) Facebook, Inc\. and its affiliates\./;
const combinedOpenSourcePattern = new RegExp(
  '(' + newOpenSourcePattern.source + ')|(' + openSourcePattern.source + ')',
);

module.exports.extends = ['eslint-config-fbcnms', 'plugin:relay/recommended'];
module.exports.plugins = ['relay', 'graphql'];
module.exports.rules = {
  // Relay Plugin
  'relay/unused-fields': 'off',
};

function graphqlSchemaObject() {
  const schemaPath = path.resolve(
    __dirname,
    '../graph/graphql/schema/symphony.graphql',
  );
  return buildSchema(fs.readFileSync(schemaPath, {encoding: 'utf8'}));
}

const schemaObject = graphqlSchemaObject();
if (Object.keys(schemaObject).length > 0) {
  module.exports.rules['graphql/no-deprecated-fields'] = [
    'error',
    {
      env: 'relay',
      schemaString: printSchema(schemaObject),
      tagName: 'graphql',
    },
  ];
}

require('eslint-plugin-lint').load(path.join(__dirname, 'scripts/lint/rules'));

module.exports.overrides = [
  {
    files: ['*'],
    rules: {
      'prettier/prettier': [
        2,
        {
          singleQuote: true,
          trailingComma: 'all',
          bracketSpacing: false,
          jsxBracketSameLine: true,
          parser: 'flow',
        },
      ],
      'flowtype/require-valid-file-annotation': 'off',
    },
  },
  {
    files: ['*.js'],
    excludedFiles: [
      '**/__generated__/*',
      '**/flow-typed/*',
      '**/magmalte/**/*.js',
    ],
    rules: {
      'lint/type-imports-block': 'error',
      'lint/flow-exact-props': 'error',
      'lint/flow-readonly-props': 'error',
    },
  },
  {
    files: ['*.mdx'],
    extends: ['plugin:mdx/overrides'],
    rules: {
      'flowtype/require-valid-file-annotation': 'off',
      'prettier/prettier': [
        2,
        {
          parser: 'mdx',
        },
      ],
    },
  },
  {
    files: ['.eslintrc.js'],
    rules: {
      quotes: ['warn', 'single'],
    },
  },
  {
    env: {
      jest: true,
      node: true,
    },
    files: [
      '**/__mocks__/**/*.js',
      '**/__tests__/**/*.js',
      '**/tests/*.js',
      'testHelpers.js',
      'testData.js',
    ],
  },
  {
    files: [
      'fbcnms-packages/**/*.js',
      'fbcnms-projects/inventory/**/*.js',
      'fbcnms-projects/platform-server/**/*.js',
    ],
    rules: {
      'header/header': [2, 'block', {pattern: combinedOpenSourcePattern}],
    },
  },
  {
    env: {
      node: true,
    },
    files: [
      '.eslintrc.js',
      'babel.config.js',
      'jest.config.js',
      'jest.*.config.js',
      'fbcnms-packages/eslint-config-fbcnms/**/*.js',
      'fbcnms-packages/fbcnms-auth/**/*.js',
      'fbcnms-packages/fbcnms-babel-register/**/*.js',
      'fbcnms-packages/fbcnms-express-middleware/**/*.js',
      'fbcnms-packages/fbcnms-logging/**/*.js',
      'fbcnms-packages/fbcnms-relay/**/*.js',
      'fbcnms-packages/fbcnms-sequelize-models/**/*.js',
      'fbcnms-packages/fbcnms-ui/stories/**/*.js',
      'fbcnms-packages/fbcnms-util/**/*.js',
      'fbcnms-packages/fbcnms-webpack-config/**/*.js',
      'fbcnms-projects/*/config/webpack.*.js',
      'fbcnms-projects/*/scripts/**/*.js',
      'fbcnms-projects/*/server/**/*.js',
      'fbcnms-projects/platform-server/**/*.js',
      'scripts/fb/fbt/*.js',
    ],
    rules: {
      'no-console': 'off',
    },
  },
];
module.exports.settings = {
  react: {
    version: 'detect',
  },
};
