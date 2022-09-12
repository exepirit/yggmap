import { createServer, request } from 'http';
// eslint-disable-next-line import/no-named-default
import { default as pathLib } from 'path';
import { dirname } from 'path/posix';
import { fileURLToPath } from 'url';
import { createRequire } from 'module';
import { build, serve } from 'esbuild';
import alias from 'esbuild-plugin-alias';

console.clear();

const serverUrl = 'http://localhost:3000';

const currentFile = fileURLToPath(import.meta.url);
const currentDir = dirname(currentFile);
const outDirectory = 'public';

/* alias */
const require = createRequire(import.meta.url);

const aliasesLib = {
  'react-dom': require.resolve('preact/compat'),
  react: require.resolve('preact/compat'),
  hooks: require.resolve('preact/hooks'),
};
/* aliases only on entry level - src */
const paths = ['pages', 'components', 'shared', 'features'].reduce((acc, path) => ({ ...acc, [path]: pathLib.resolve(currentDir, `src/${path}/index.ts`) }), {});

const allAliases = { ...aliasesLib, ...paths };

const successMsg = `\x1b[32m\ Compiled successfully!
You can now view \x1b[36mapp-base\x1b[0m in the browser.
Local:\x1b[37m  \x1b[36m${serverUrl}\x1b[0m
Note that the development build is not optimized.
To create a production build, use yarn s p.`;

const clients = [];

const options = {
  entryPoints: ['src/app.js'],
  outdir: outDirectory,
  loader: {
    '.js': 'js',
    '.png': 'file',
    '.jpg': 'file',
    '.jpeg': 'file',
    '.svg': 'file',
    '.gif': 'file',
  },
  jsxFactory: 'h',
  minify: true,
  sourcemap: true,
  assetNames: 'assets/[name]', // -[hash]
  // chunkNames: 'chunks/[name]-[hash]',
  // entryNames: '[dir]/[name]', // -[hash]
  bundle: true,
  resolveExtensions: ['.tsx', '.ts', '.jsx', '.js', '.css', '.json'],
  watch: {
    onRebuild(error) {
      clients.forEach((res) => res.write('data: update\n\n'));
      clients.length = 0;
      console.clear();
      console.log(error || successMsg);
    },
  },
  banner: { js: ' (() => new EventSource("/esbuild").onmessage = () => location.reload())();' },
  plugins: [alias(allAliases)],
  inject: ['src/react-shim.js'],
};

const mode = process.argv[2];

build(options).then(console.log(successMsg)).catch(({ message }) => console.error(message));

serve({ servedir: outDirectory }, {}).then(() => {
  createServer((req, res) => {
    const {
      url, method, headers,
    } = req;
    if (url === '/esbuild') {
      return clients.push(
        res.writeHead(200, {
          'Content-Type': 'text/event-stream',
          'Cache-Control': 'no-cache',
          Connection: 'keep-alive',
        }),
      );
    }
    // eslint-disable-next-line no-bitwise
    const path = ~url.split('/').pop().indexOf('.') ? url : '/index.html'; // for PWA with router
    req.pipe(
      request({
        hostname: '0.0.0.0', port: 8000, path, method, headers,
      }, (prxRes) => {
        res.writeHead(prxRes.statusCode, prxRes.headers);
        prxRes.pipe(res, { end: true });
      }),
      { end: true },
    );
  }).listen(serverUrl.split(':')[2]);
});
