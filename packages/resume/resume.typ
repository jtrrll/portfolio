#import "@preview/basic-resume:0.2.9": *

#let name = "Jackson Terrill"
#let location = "Boston, MA"
#let email = "jacksonterrill3@gmail.com"
#let phone = "+1 (314) 651-6907"

#let github = "github.com/jtrrll"
#let linkedin = "linkedin.com/in/jacksonterrill"
#let personal-site = "jtrrll.com"

#show: resume.with(
  author: name,
  location: location,
  email: email,
  github: github,
  linkedin: linkedin,
  phone: phone,
  font: "IBM Plex Serif",
  paper: "us-letter",
  author-position: left,
  personal-info-position: left,
)

== Education

#edu(
  institution: "Northeastern University",
  location: "Boston, MA",
  dates: dates-helper(start-date: "Sep 2021", end-date: "Dec 2024"),
  degree: "Bachelor of Science in Computer Science and Business Administration",
  consistent: true,
)
- Cumulative GPA: 3.95\/4.00 | Dean's List, University Honors Distinction, National Merit Scholarship
- Relevant Coursework: Software Development, Computer Systems, Networks, Algorithms, Object-Oriented Design, Database Design

== Experience

#work(
  title: "Fullstack Software Engineer",
  company: "Proof",
  dates: dates-helper(start-date: "Jan 2024", end-date: "Present"),
  location: "Boston, MA",
)
- Implemented a customer-facing fraud detection suite by integrating with third-party risk analysis APIs to reduce fraudulent activity and bolster transaction trust
- Refactored expensive, polling-based state synchronization systems to use lightweight, event-based partial updates transmitted through WebSockets
- Ported build systems and development environments to Nix to improve build reproducibility and caching, simplify deployment pipelines, and enhance developer experience
- Built and maintained monitoring libraries for two microservices to enable runtime observability in production
- Improved codebase reliability by introducing Sorbet type-checking into Rails and migrating legacy JavaScript to TypeScript, eliminating type errors and preventing regressions while enabling fearless refactoring

#work(
  title: "Chief of Software | Tech Lead",
  company: "Generate@Northeastern",
  dates: dates-helper(start-date: "Sep 2023", end-date: "Dec 2024"),
  location: "Boston, MA",
)
- As a Chief of Software, designed technical assessments, established development best-practices, and hosted knowledge sharing sessions
- As a Tech Lead, led teams to create a 3D printer sharing service and an event discovery app with Go and TypeScript
- Guided project architecture discussions, delegated tasks, and reviewed developer pull requests
- Managed client relations and facilitated timely completion of project milestones

#work(
  title: "Backend Networking Developer (Co-op)",
  company: "VMware",
  dates: dates-helper(start-date: "Jan 2023", end-date: "Aug 2023"),
  location: "Boston, MA",
)
- Refactored REST endpoints in Java to enforce strict type validation, enabling efficient querying and filtering of available network resources in Cloud Director
- Implemented Raw Port-Protocols for Edge Gateways, a new feature allowing anonymous TCP/UDP port traffic configuration on firewall rules
- Developed and documented new firewall synchronization features in response to critical client requests
- Restored backward compatibility for firewall configurations on legacy virtualization platforms (NSX-V)

#work(
  title: "Software Engineer (Intern)",
  company: "Federal Reserve Bank of St. Louis",
  dates: dates-helper(start-date: "Jun 2021", end-date: "Dec 2022"),
  location: "St. Louis, MO",
)
- Created and maintained a certification reporting tool with Python, SQLite, and the SharePoint API
- Built and published an Alexa skill for public educational resources using the ASK SDK for Node.js

== Projects

#project(
  name: "snekcheck",
  dates: dates-helper(start-date: "Oct 2024", end-date: "Present"),
  url: "github.com/jtrrll/snekcheck",
)
- Built an opinionated filename linter that loves snake_case to streamline bulk file operations
- Wrote comprehensive unit tests and end-to-end regression tests
- Released v0.1.0 as a Go application packaged with Nix

#project(
  name: "dotfiles",
  dates: dates-helper(start-date: "Jun 2024", end-date: "Present"),
  url: "github.com/jtrrll/dotfiles",
)
- Used Nix and Home Manager to declaratively configure tools and systems
- Modified frequently to improve development workflow and decrease cognitive load

#project(
  name: "portfolio",
  dates: dates-helper(start-date: "Oct 2025", end-date: "Present"),
  url: "jtrrll.com",
)
- Designed a personal website to showcase software projects, music, and art with Go and Templ
- Rewrote a previous SvelteKit iteration to drastically improve responsiveness and maintainability

#project(
  name: "mona",
  dates: dates-helper(start-date: "Apr 2024", end-date: "Nov 2024"),
)
- Built a terminal image viewer in Go that converts images into low-resolution colorful ASCII art
- Implemented an edge detection kernel for improving legibility and various color palette options
- Supported various color depths, and multiple file formats including PNG, JPEG, and GIF

#project(
  name: "mini.fs",
  dates: dates-helper(start-date: "Dec 2022", end-date: "Dec 2022"),
)
- Built a portable 1MB file system with C and FUSE that can be transferred between systems and mounted at will

#project(
  name: "mini.sh",
  dates: dates-helper(start-date: "Nov 2022", end-date: "Nov 2022"),
)
- Created a basic “POSIX-inspired” shell with redirection, pipes, sequencing, and grouping in C

== Skills
- *Programming Languages*: Go, Nix, Ruby, TypeScript/JavaScript, Bash, Python, Rust, C, Racket, SQL, GraphQL, HCL, KQL, HTML/CSS, Haskell, Gleam
- *Frameworks*: Ruby on Rails, React, Svelte, SvelteKit, Astro, Tailwind CSS, Templ
- *Software*: Git, Linux, Docker, VS Code, Neovim, Sentry, Azure Data Explorer, Postman, Jenkins, GitHub Actions

== Interests
Video Game Design/Development, Type Systems, Workflow Optimizations, Water Polo, Songwriting, Formula One
