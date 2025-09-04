import { cpSync, existsSync, mkdirSync } from 'node:fs';
import { join } from 'node:path';

const root = process.cwd();
const distDir = join(root, 'dist', 'apps', 'frontend');
const staticDir = join(root, 'apps', 'api', 'static');
const templatesDir = join(root, 'apps', 'api', 'templates');

function ensureDir(p) {
  if (!existsSync(p)) mkdirSync(p, { recursive: true });
}

ensureDir(staticDir);
ensureDir(join(staticDir, 'assets'));
ensureDir(templatesDir);

// Copy assets
cpSync(join(distDir, 'assets'), join(staticDir, 'assets'), { recursive: true });
// Copy favicon
cpSync(join(distDir, 'favicon.ico'), join(staticDir, 'favicon.ico'));
// Copy index.html to templates
cpSync(join(distDir, 'index.html'), join(templatesDir, 'index.html'));

console.log('Copied frontend build into backend:');
console.log(`  assets -> ${join('apps', 'api', 'static', 'assets')}`);
console.log(`  favicon.ico -> ${join('apps', 'api', 'static', 'favicon.ico')}`);
console.log(`  index.html -> ${join('apps', 'api', 'templates', 'index.html')}`);

