import { cpSync, existsSync, mkdirSync, rmSync } from 'node:fs';
import { join } from 'node:path';

const root = process.cwd();
const distDir = join(root, 'dist', 'apps', 'frontend');
const staticDir = join(root, 'apps', 'api', 'static');
const assetsDir = join(staticDir, 'assets');
const templatesDir = join(root, 'apps', 'api', 'templates');

function ensureDir(p) {
  if (!existsSync(p)) mkdirSync(p, { recursive: true });
}

rmSync(assetsDir, { recursive: true, force: true });
rmSync(join(staticDir, 'favicon.ico'), { force: true });
rmSync(join(templatesDir, 'index.html'), { force: true });

ensureDir(staticDir);
ensureDir(assetsDir);
ensureDir(templatesDir);

cpSync(join(distDir, 'assets'), assetsDir, { recursive: true });
cpSync(join(distDir, 'favicon.ico'), join(staticDir, 'favicon.ico'));
cpSync(join(distDir, 'index.html'), join(templatesDir, 'index.html'));

console.log('Copied frontend build into backend:');
console.log(`  assets -> ${join('apps', 'api', 'static', 'assets')}`);
console.log(`  favicon.ico -> ${join('apps', 'api', 'static', 'favicon.ico')}`);
console.log(`  index.html -> ${join('apps', 'api', 'templates', 'index.html')}`);

