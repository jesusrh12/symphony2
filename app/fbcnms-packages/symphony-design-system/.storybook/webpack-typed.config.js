/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

require('@fbcnms/babel-register');
const createCompiler = require('@storybook/addon-docs/mdx-compiler-plugin');

const autoprefixer = require('autoprefixer');
const paths = require('@fbcnms/webpack-config/paths');
const webpack = require('webpack');

import type {WebpackOptions} from 'webpack';

type BuilderParams = {
  config: WebpackOptions,
};

export default function builder({config}: BuilderParams): WebpackOptions {
  return {
    ...config,
    module: {
      ...config.module,
      rules: [
        {
          test: /\.(js|jsx|mjs)$/,
          include: [paths.appSrc, paths.packagesDir, paths.nodeModulesDir],
          loader: require.resolve('babel-loader'),
          options: {
            configFile: '../../babel.config.js',
            // This is a feature of `babel-loader` for webpack (not Babel
            // itself). It enables caching results in
            // ./node_modules/.cache/babel-loader/ directory for faster
            // rebuilds.
            cacheDirectory: true,
          },
        },
        {
          test: /\.(stories|story)\.mdx$/,
          use: [
            {
              loader: require.resolve('babel-loader'),
              options: {
                configFile: '../../babel.config.js',
                cacheDirectory: true,
              },
            },
            {
              loader: '@mdx-js/loader',
              options: {
                compilers: [createCompiler({})],
              },
            },
          ],
        },
        {
          test: /\.(stories|story)\.[tj]sx?$/,
          use: [
            {
              loader: require.resolve('@storybook/source-loader'),
            },
            {
              loader: require.resolve('babel-loader'),
              options: {
                configFile: '../../babel.config.js',
                cacheDirectory: true,
              },
            },
          ],
          exclude: [/node_modules/],
          enforce: 'pre',
        },
        {
          test: /\.css$/,
          use: [
            {
              loader: 'style-loader',
            },
            {
              loader: 'css-loader',
              options: {
                sourceMap: true,
              },
            },
          ],
        },
      ],
    },
  };
}
