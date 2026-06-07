# peterssontech.net

Personal site for **Billy Petersson** — Azure / Cloud Engineer, Örebro, Sweden.

A hand-built, single-page static site. No build step, no framework — just HTML, CSS and a
little vanilla JavaScript. Fast, accessible, dark/light themed.

## Structure

```
index.html            # the whole page
assets/css/style.css  # design system + layout
assets/js/main.js     # theme toggle, scroll reveals, active nav
CNAME                 # custom domain (peterssontech.net)
robots.txt, sitemap.xml, site.webmanifest
.github/workflows/deploy.yml   # publishes to GitHub Pages
```

## Hosting

Two ways the same files are served:

1. **Self-hosted (primary)** — served by nginx on the home server (ProDesk) and exposed
   through a Cloudflare Tunnel at `peterssontech.net` / `www.peterssontech.net`.
2. **GitHub Pages (mirror)** — `deploy.yml` uploads the repo root as a Pages artifact.

## Local preview

```bash
python3 -m http.server 8080   # then open http://localhost:8080
```

## Editing

It's plain HTML — edit `index.html` directly. Colours, fonts and spacing live as CSS
variables at the top of `assets/css/style.css`.
