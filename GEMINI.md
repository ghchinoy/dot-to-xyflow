# Agent Instructions

This project uses **bd** (beads) for issue tracking. Run `bd onboard` to get started.

## Coding Conventions

- **Language**: Go is the primary language for the translator.
- **Dependency Management**: Always maintain `go.mod` and run `go mod tidy` after adding dependencies.
- **Error Handling**: Use `log.Fatalf` for terminal errors in CLI tools to provide clear feedback.
- **Testing**: Every new feature or mapping logic should have a corresponding test in `*_test.go`.
- **Documentation**: Update `README.md` when changing installation steps or core features.

## Technical Context

- **Coordinate Systems**: 
    - **Graphviz**: Bottom-up (Y increases upwards).
    - **XYFlow**: Top-down (Y increases downwards).
    - **Transformation**: To map coordinates, use the graph's bounding box (`bb`) height: `xyflowY = graphHeight - (gvY + nodeHeight/2)`.
- **Node Positioning**: Graphviz `pos` attributes are center-based. XYFlow positions are top-left based. Adjust accordingly using node `width` and `height` (converted from inches to points: `1in = 72pts`).
- **Node Types**:
    - `input`: Map from shapes like `note`, `invhouse`, or IDs containing "input".
    - `output`: Map from shapes like `component` or `doublecircle`.
    - `default`: Standard fallback.

## Issue Tracking

This project uses **bd (beads)** for issue tracking.
Run `bd prime` for workflow context, or install hooks (`bd hooks install`) for auto-injection.

**Quick reference:**
- `bd ready` - Find unblocked work
- `bd show <id>` - View issue details
- `bd update <id> --status in_progress` - Claim work
- `bd create "Title" --type task --priority 2` - Create issue
- `bd close <id>` - Complete work
- `bd sync` - Sync with git (run at session end)

For full workflow details: `bd prime`

## Landing the Plane (Session Completion)

**When ending a work session**, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **File issues for remaining work** - Create issues for anything that needs follow-up
2. **Run quality gates** (if code changed) - Tests, linters, builds
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   git pull --rebase
   bd sync
   git push
   git status  # MUST show "up to date with origin"
   ```
5. **Clean up** - Clear stashes, prune remote branches
6. **Verify** - All changes committed AND pushed
7. **Hand off** - Provide context for next session

**CRITICAL RULES:**
- Work is NOT complete until `git push` succeeds
- NEVER stop before pushing - that leaves work stranded locally
- NEVER say "ready to push when you are" - YOU must push
- If push fails, resolve and retry until it succeeds