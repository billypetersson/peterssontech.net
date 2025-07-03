# Billy Petersson - Professional Portfolio Website

A modern, responsive portfolio website built with Hugo static site generator, showcasing professional IT expertise, projects, and technical blog content.

## 🌟 Features

### 🏠 Homepage
- Professional hero section with call-to-action buttons
- Skills preview with categorized technologies
- Featured projects showcase
- Latest blog posts preview
- Responsive design optimized for all devices

### 📝 Content Sections
- **About Me**: Comprehensive professional biography and career journey
- **Technical Skills**: Categorized expertise in programming languages, frameworks, and tools
- **Work Experience**: Timeline of professional roles and achievements
- **Projects Portfolio**: Detailed project case studies with code examples
- **Tech Blog**: In-depth technical articles and tutorials
- **Contact**: Professional contact form and social media links

### 🚀 Technical Features
- **SEO Optimized**: Meta tags, structured data, sitemaps, and performance optimization
- **Accessibility**: WCAG 2.1 compliant with proper ARIA labels and keyboard navigation
- **Performance**: Optimized images, minified assets, and fast loading times
- **Responsive Design**: Mobile-first approach with cross-device compatibility
- **Dark Mode Support**: Automatic theme switching based on user preferences
- **Print Styles**: Optimized for professional document printing
- **Search Functionality**: Client-side search with JSON index
- **Progressive Web App**: PWA-ready with web manifest and service worker support
- **Security Headers**: CSP, HSTS, and other security best practices
- **Performance Monitoring**: Core Web Vitals tracking and error handling

## 🛠️ Built With

- **[Hugo](https://gohugo.io/)** (v0.147.9) - Static Site Generator
- **HTML5 & CSS3** - Modern web standards
- **JavaScript (ES6+)** - Interactive functionality
- **Font Awesome** - Professional iconography
- **Inter Font** - Clean, modern typography

## 📁 Project Structure

```
peterssontech-portfolio/
├── config.toml              # Hugo configuration
├── content/                 # Content files (Markdown)
│   ├── _index.md           # Homepage content
│   ├── about.md            # About page
│   ├── skills.md           # Technical skills
│   ├── experience.md       # Work experience
│   ├── contact.md          # Contact page
│   ├── projects/           # Project portfolio
│   │   ├── _index.md
│   │   ├── ecommerce-platform.md
│   │   ├── analytics-dashboard.md
│   │   └── task-manager-app.md
│   └── blog/               # Tech blog posts
│       ├── _index.md
│       ├── kubernetes-deployment-strategies.md
│       ├── react-performance-optimization.md
│       └── python-automation-guide.md
├── layouts/                # Hugo templates
│   ├── _default/
│   │   ├── baseof.html     # Base template
│   │   ├── single.html     # Single page template
│   │   └── list.html       # List page template
│   ├── partials/
│   │   ├── header.html     # Site header
│   │   ├── footer.html     # Site footer
│   │   ├── seo.html        # SEO meta tags
│   │   └── breadcrumb.html # Breadcrumb navigation
│   ├── index.html          # Homepage template
│   ├── index.json          # Search index template
│   └── 404.html            # Error page
├── static/                 # Static assets
│   ├── css/
│   │   └── main.css        # Main stylesheet
│   ├── js/
│   │   └── main.js         # JavaScript functionality
│   ├── images/             # Images and media
│   ├── robots.txt          # Search engine directives
│   ├── CNAME               # Custom domain configuration
│   ├── site.webmanifest    # PWA manifest
│   └── .htaccess           # Apache security headers and caching
├── .github/
│   └── workflows/
│       └── hugo.yml        # GitHub Actions deployment
└── README.md               # This file
```

## 🚀 Getting Started

### Prerequisites

- [Hugo](https://gohugo.io/getting-started/installing/) (Extended version recommended)
- [Git](https://git-scm.com/)
- [Node.js](https://nodejs.org/) (optional, for additional tooling)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/billy-petersson/peterssontech-portfolio.git
   cd peterssontech-portfolio
   ```

2. **Start the development server**
   ```bash
   hugo server --buildDrafts --buildFuture
   ```

3. **Open your browser**
   Navigate to `http://localhost:1313` to view the site

### Building for Production

```bash
# Build the static site
hugo --minify

# The generated site will be in the 'public' directory
```

## 🌐 Deployment

### GitHub Pages (Automated)

The site is configured for automatic deployment to GitHub Pages using GitHub Actions:

1. **Push to main branch** - Triggers automatic build and deployment
2. **Custom domain** - Configured for `peterssontech.net`
3. **HTTPS** - Automatically enabled with custom domain

### Manual Deployment

You can deploy the generated `public` folder to any static hosting service:

- **Netlify**: Connect your GitHub repository for automatic deployments
- **Vercel**: Import the project and deploy with zero configuration
- **AWS S3**: Upload the `public` folder to an S3 bucket configured for static hosting
- **Cloudflare Pages**: Connect your repository for automatic builds

## 🎨 Customization

### Content Updates

1. **Personal Information**: Update `hugo.toml` with your details
2. **About Page**: Edit `content/about.md`
3. **Projects**: Add new projects in `content/projects/`
4. **Blog Posts**: Add articles in `content/blog/`
5. **Skills**: Update `content/skills.md` with your expertise

### Design Customization

1. **Colors**: Modify CSS variables in `static/css/main.css`
2. **Typography**: Change font settings in the CSS file
3. **Layout**: Edit templates in the `layouts/` directory
4. **Navigation**: Update menus in `hugo.toml`

### Adding New Features

1. **Contact Form**: Integrate with services like Netlify Forms or Formspree
2. **Analytics**: Add Google Analytics or other tracking services
3. **Comments**: Integrate Disqus or other commenting systems
4. **Search**: ✅ Client-side search already implemented
5. **CMS Integration**: Connect with headless CMS like Contentful or Strapi
6. **Multi-language**: Hugo's i18n features for internationalization

## 📊 Performance

### Core Web Vitals

- **Largest Contentful Paint (LCP)**: < 2.5s
- **First Input Delay (FID)**: < 100ms
- **Cumulative Layout Shift (CLS)**: < 0.1

### Optimization Features

- **Image Optimization**: Responsive images with WebP support
- **CSS/JS Minification**: Automated asset optimization
- **Caching Headers**: Proper browser and CDN caching via .htaccess
- **Code Splitting**: Optimized JavaScript loading
- **Font Loading**: Optimized web font delivery
- **Intersection Observer**: Lazy loading and scroll animations
- **Performance Monitoring**: Real-time Core Web Vitals tracking

## ♿ Accessibility

### WCAG 2.1 Compliance

- **AA Level** color contrast ratios
- **Semantic HTML** structure with proper headings
- **ARIA Labels** for enhanced screen reader support
- **Keyboard Navigation** for all interactive elements
- **Focus Indicators** for improved usability
- **Skip Links** for main content navigation
- **Breadcrumb Navigation** for better page hierarchy understanding
- **High Contrast Support** for users with visual impairments
- **Reduced Motion** support for users with vestibular disorders

### Testing

Regular testing with:
- **Screen readers** (NVDA, JAWS, VoiceOver)
- **Keyboard navigation** testing
- **Color contrast** analysis tools
- **Accessibility audits** with Lighthouse

## 🔧 Development Tools

### Recommended Extensions (VS Code)

- **Hugo Language and Syntax Support**
- **Markdown All in One**
- **Prettier - Code formatter**
- **Live Server** (for local development)

### Testing Tools

- **Hugo's built-in server** with live reload
- **Lighthouse** for performance auditing
- **WAVE** for accessibility testing
- **W3C Markup Validator** for HTML validation

## 📈 SEO Features

### Technical SEO

- **Structured Data**: JSON-LD markup for rich snippets
- **Open Graph**: Social media sharing optimization
- **XML Sitemap**: Automatic generation for search engines
- **Robots.txt**: Search engine crawling directives
- **Canonical URLs**: Duplicate content prevention
- **Search Index**: JSON-based search functionality for improved UX
- **Security Headers**: CSP and other headers for better search rankings

### Content SEO

- **Meta Descriptions**: Unique descriptions for all pages
- **Title Tags**: SEO-optimized page titles
- **Heading Structure**: Proper H1-H6 hierarchy
- **Internal Linking**: Strategic content connections
- **Image Alt Text**: Descriptive alternative text

## 🔒 Security Features

### Security Headers (.htaccess)
- **Content Security Policy (CSP)**: Prevents XSS attacks
- **X-Frame-Options**: Prevents clickjacking
- **X-Content-Type-Options**: Prevents MIME sniffing
- **Referrer Policy**: Controls referrer information
- **HSTS**: HTTP Strict Transport Security (when using HTTPS)

### Performance Security
- **Error Handling**: Global JavaScript error catching
- **Input Validation**: Form validation and sanitization
- **Safe External Links**: `rel="noopener noreferrer"` on external links

## 🎯 Browser Support

### Modern Browsers
- **Chrome**: 88+
- **Firefox**: 84+
- **Safari**: 14+
- **Edge**: 88+

### Progressive Enhancement
- **ES6+ Features**: With graceful degradation
- **CSS Grid/Flexbox**: With fallbacks
- **Intersection Observer**: With polyfill support

## 📊 Analytics & Monitoring

### Built-in Monitoring
- **Core Web Vitals**: LCP, FID, CLS tracking
- **Performance Timing**: Page load metrics
- **Error Tracking**: JavaScript error logging
- **User Interactions**: Navigation and form analytics

### Integration Ready
- **Google Analytics**: Configuration ready
- **Search Console**: Sitemap and structured data
- **PageSpeed Insights**: Optimized for testing

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

While this is a personal portfolio, suggestions and improvements are welcome:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/improvement`)
3. **Commit your changes** (`git commit -am 'Add new feature'`)
4. **Push to the branch** (`git push origin feature/improvement`)
5. **Create a Pull Request**

## 📞 Support

For questions or support:

- **Email**: [billy@peterssontech.net](mailto:billy@peterssontech.net)
- **LinkedIn**: [linkedin.com/in/billy-petersson](https://linkedin.com/in/billy-petersson)
- **GitHub Issues**: [Submit an issue](https://github.com/billy-petersson/peterssontech-portfolio/issues)

## 🏆 Acknowledgments

- **Hugo Community** for the excellent static site generator
- **Font Awesome** for professional icons
- **Inter Font** by Rasmus Andersson for typography
- **GitHub Pages** for reliable hosting

---

## 📈 Project Statistics

- **Total Pages**: 79 generated pages
- **Build Time**: ~99ms (optimized for speed)
- **Static Files**: 6 optimized assets
- **Code Quality**: Zero build warnings
- **Lighthouse Score**: 95+ across all categories
- **Accessibility**: WCAG 2.1 AA compliant

## 🚀 Quick Start Commands

```bash
# Clone and start development
git clone https://github.com/billy-petersson/peterssontech-portfolio.git
cd peterssontech-portfolio
hugo server --buildDrafts

# Build for production
hugo --minify

# Deploy to GitHub Pages (automatic on push to main)
git push origin main
```

## 🎨 Customization Examples

### Changing Colors
```css
/* In static/css/main.css */
:root {
    --primary-color: #your-color;
    --secondary-color: #your-secondary;
}
```

### Adding a New Project
```markdown
<!-- In content/projects/your-project.md -->
---
title: "Your Project Name"
date: 2024-12-25
tags: ["React", "Node.js"]
featured: true
---

Your project description...
```

### Adding Blog Posts
```markdown
<!-- In content/blog/your-post.md -->
---
title: "Your Blog Post"
date: 2024-12-25
readingTime: 5
tags: ["Technology", "Tutorial"]
---

Your blog content...
```

**Built with ❤️ by Billy Petersson** | **Last Updated**: January 2025