/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import express from 'express';
import logging from '@fbcnms/logging';

import type {Middleware} from 'express';
const logger = logging.getLogger(module);

type WebpackMiddlewareOptions = {|
  distPath: string,
  devWebpackConfig: Object,
|};

type WebpackSmartMiddlewareOptions = {|
  ...WebpackMiddlewareOptions,
  devMode: boolean,
|};

function webpackDevMiddleware(options: WebpackMiddlewareOptions): Middleware {
  const {devWebpackConfig} = options;
  const webpack = require('webpack');
  const webpackMiddleware = require('webpack-dev-middleware');
  const webpackHotMiddleware = require('webpack-hot-middleware');
  const compiler = webpack(devWebpackConfig);
  const middleware = webpackMiddleware(compiler, {
    publicPath: devWebpackConfig.output.publicPath,
    contentBase: 'src',
    logger,
    stats: {
      colors: true,
      hash: false,
      timings: true,
      chunks: false,
      chunkModules: false,
      modules: false,
    },
  });

  const router = express.Router();
  router.use(middleware);
  webpackHotMiddleware(compiler);
  return router;
}

export default function webpackSmartMiddleware(
  options: WebpackSmartMiddlewareOptions,
): Middleware {
  const {devMode, devWebpackConfig, distPath} = options;

  const router = express.Router();
  if (devMode) {
    router.use(webpackDevMiddleware({devWebpackConfig, distPath}));
  } else {
    // serve built resources from static/dist/ folder
    router.use(devWebpackConfig.output.publicPath, express.static(distPath));
  }
  return router;
}
