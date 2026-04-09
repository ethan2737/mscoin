import { readdir, readFile, stat } from 'node:fs/promises'
import path from 'node:path'
import process from 'node:process'

const projectRoot = process.cwd()
const pagesRoot = path.join(projectRoot, 'src', 'pages-vue3')

const assetPattern = /(?:src=|url\()["']?([^"')]+\.(?:png|jpg|jpeg|gif|svg|webp))/g
const importPattern = /(?:import\s+[^;]*?\s+from\s+["']([^"']+)["'])|(?:import\(["']([^"']+)["']\))/g
const assetExtensions = ['', '.js', '.vue', '.mjs', '.json']

async function walk(dir) {
  const entries = await readdir(dir, { withFileTypes: true })
  const results = []

  for (const entry of entries) {
    const fullPath = path.join(dir, entry.name)
    if (entry.isDirectory()) {
      results.push(...await walk(fullPath))
      continue
    }

    if (fullPath.endsWith('.vue') || fullPath.endsWith('.js')) {
      results.push(fullPath)
    }
  }

  return results
}

async function exists(target) {
  try {
    await stat(target)
    return true
  } catch {
    return false
  }
}

async function resolveImport(basePath) {
  const direct = await exists(basePath)
  if (direct) return true

  for (const ext of assetExtensions) {
    if (ext && await exists(basePath + ext)) return true
  }

  if (await exists(path.join(basePath, 'index.js'))) return true
  if (await exists(path.join(basePath, 'index.vue'))) return true

  return false
}

async function main() {
  const files = await walk(pagesRoot)
  const failures = []

  for (const file of files) {
    const content = await readFile(file, 'utf8')

    for (const match of content.matchAll(assetPattern)) {
      const assetRef = match[1]
      if (/^(https?:|data:|\/)/.test(assetRef)) continue

      const resolved = path.resolve(path.dirname(file), assetRef)
      if (!(await exists(resolved))) {
        failures.push(`ASSET  ${path.relative(projectRoot, file)} -> ${assetRef}`)
      }
    }

    for (const match of content.matchAll(importPattern)) {
      const importRef = match[1] ?? match[2]
      if (!importRef || !importRef.startsWith('.')) continue

      const basePath = path.resolve(path.dirname(file), importRef)
      if (!(await resolveImport(basePath))) {
        failures.push(`IMPORT ${path.relative(projectRoot, file)} -> ${importRef}`)
      }
    }
  }

  if (failures.length > 0) {
    console.error('Unresolved asset/import references in src/pages-vue3:')
    for (const failure of failures) {
      console.error(`- ${failure}`)
    }
    process.exit(1)
  }

  console.log('Asset resolution check passed for src/pages-vue3.')
}

await main()
