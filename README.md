# projectMentis

# Commit Guidelines
We follow Conventional Commits to keep our Git history consistent.

# Format
    <type>(<scope>): <message>
    Type: feat, fix, chore, docs, refactor, test
    Scope: area of the code affected (e.g., api, linter, auth)
    Message: short description (≤ 50 chars for the title)
# Examples
Valid:
    feat(api): add login endpoint
    chore(linter): setup pre-commit
    fix(auth): handle nil pointer
Invalid:
    added pre-commit hooks       # ❌ missing type(scope)
    feat: add login endpoint     # ❌ missing scope
    foo(api): something          # ❌ type not allowed
Body (optional): max 72 chars per line.