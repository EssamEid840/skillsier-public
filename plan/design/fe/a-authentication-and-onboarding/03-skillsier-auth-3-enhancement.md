## AUTH-3 — Password Reset & Account Recovery (Forgot Password, Token/Code Verify, Reset, Sign‑out Others)

> Enhancement add‑on for the existing AUTH‑3 journey in `skillsier-frontend-journeys-claude-final.md`. This deliverable fully completes the three required sections: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-strucure.md` route and component patterns. Scope includes reset request, non‑enumerating responses, captcha & rate limits, token/code verification, password policy enforcement, session/device revocation, and success handoff to sign‑in.

---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(auth)/forgot-password/page.tsx
Purpose: Start password reset by submitting email (non‑enumerating UX).
Props: None
Key Features:
  - Email field + validation
  - Non‑enumeration response (“If an account exists…”)
  - Progressive captcha after failures
  - Links: Sign in, Create account

Path: apps/web/app/(auth)/forgot-password/sent/page.tsx
Purpose: Confirmation screen after request (shows resend/cooldown UI).
Props: email? (from store or masked display)
Key Features:
  - “Check your email” guidance
  - Resend with cooldown + change email link

Path: apps/web/app/(auth)/reset-password/verify/page.tsx
Purpose: Verify token from email (magic link) or accept code input.
Props: token (search param) OR route state for code flow
Key Features:
  - Validates token with backend
  - Error states: invalid/expired/already used
  - CTA to proceed to reset or request new link

Path: apps/web/app/(auth)/reset-password/page.tsx
Purpose: Set a new password (with policy meter and confirm).
Props: token (if required) or challengeId (for code‑verified flow)
Key Features:
  - Password + Confirm with meter & policy checklist
  - Optional: “Sign out of other devices” (recommended)

Path: apps/web/app/(auth)/reset-password/success/page.tsx
Purpose: Success handoff after reset.
Props: None
Key Features:
  - “Password updated” confirmation
  - CTA: Sign in now
  - Security note about device sessions

Path: apps/web/app/(auth)/layout.tsx
Purpose: Shared frame for auth routes.
Features:
  - Brand header/footer
  - ErrorBoundary + Sentry
  - <Toaster/> mount

Path: apps/web/app/(auth)/forgot-password/loading.tsx
Path: apps/web/app/(auth)/forgot-password/error.tsx
Path: apps/web/app/(auth)/reset-password/loading.tsx
Path: apps/web/app/(auth)/reset-password/error.tsx
Purpose: Skeletons and route‑scoped error fallbacks.
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Purpose: Stack layout for Forgot → Sent → (Verify) → Reset → Success.
Features:
  - Header/back
  - Modal containers (ToS/Privacy)

Path: apps/mobile/app/(public)/(auth)/forgot-password.tsx
Purpose: Submit email for reset; non‑enumerating result.

Path: apps/mobile/app/(public)/(auth)/forgot-password-sent.tsx
Purpose: “Check your email” screen with resend/cooldown.

Path: apps/mobile/app/(public)/(auth)/reset-password-verify.tsx (optional)
Purpose: Handle deep‑link token verification or allow manual code entry.

Path: apps/mobile/app/(public)/(auth)/reset-password.tsx
Purpose: New password entry with policy meter; sign‑out devices checkbox.

Path: apps/mobile/app/(public)/(auth)/reset-password-success.tsx
Purpose: Success handoff to sign‑in.
```

> Handoff: After a successful reset, users are redirected to **AUTH‑2 (Sign‑In)**. Auto sign‑in after reset is **not** recommended (security).

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: Button
Purpose: Submit, Resend, Try again
Props: variant, size, isLoading, leftIcon, rightIcon

Component: Input
Purpose: Email field
Props: id, type, value, onChange, error, autoComplete

Component: PasswordInput
Purpose: New password & confirm; show/hide + strength meter slot
Props: value, onChange, error

Component: CodeInput
Path: packages/ui/src/components/forms/CodeInput.tsx
Purpose: 6‑digit code entry (if code‑based verify flow used)
Props:
  interface CodeInputProps {
    length?: number; value: string; onChange: (v: string) => void;
    error?: string; onComplete?: (v: string) => void;
  }

Component: Checkbox
Purpose: “Sign out of other devices”
Props: checked, onChange, children

Component: Alert / Banner / InlineError
Purpose: Invalid/expired token, resend confirmation, network errors

Component: CaptchaWidget
Purpose: Turnstile/reCAPTCHA token provider (progressive after N failures)
Props: onToken(token), provider

Component: MaskedEmail
Purpose: Safely display masked email (e.g., a***@example.com)
Props: value (email); handles masking
```

**Feature Components** (`packages/lib/src/features/auth/components/`):
```
Component: ForgotPasswordForm.tsx
Purpose: Encapsulate email submission + captcha + non‑enumeration UX
Props: onSuccess(email)

Component: ResetTokenVerifier.tsx
Purpose: Verify token and display next actions
Props: token, onVerified({ challengeId? })

Component: ResetPasswordForm.tsx
Purpose: New password entry with policy + confirm + sign‑out devices
Props: onSuccess, challengeId?, token?
APIs: usePasswordPolicy, useResetPasswordMutation

Component: ResendResetLink.tsx
Purpose: Resend with server‑driven cooldown & countdown
Props: email

Component: ExpiredTokenHelp.tsx
Purpose: Guidance + CTA to request a new link or switch to code flow
```

**Domain‑Specific Components:**
```
Component: AuthHero.tsx (web)
Path: apps/web/components/auth/AuthHero.tsx
Purpose: Benefits column

Component: AuthKeyboardAvoider.tsx (mobile)
Path: apps/mobile/components/auth/AuthKeyboardAvoider.tsx
Purpose: Prevent keyboard overlap on forms
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/features/auth/hooks/`):
```
Hook: useForgotPasswordMutation.ts
Purpose: Request a reset link (non‑enumerating)
Endpoint: POST /v1/auth/forgot-password  (users-be/auth)
Return: { accepted: true } always (non‑enumerating)

Hook: useVerifyResetTokenQuery.ts
Purpose: Validate reset token from email
Endpoint: POST /v1/auth/reset-password/verify
Body: { token }
Return: { valid: boolean, challengeId?: string, reason?: "expired"|"used"|"invalid" }

Hook: useResetPasswordMutation.ts
Purpose: Submit new password (+ signOutOthers flag)
Endpoint: POST /v1/auth/reset-password
Body: { token? , challengeId?, newPassword, signOutOthers?: boolean }

Hook: useResendResetMutation.ts
Purpose: Resend link with cooldown
Endpoint: POST /v1/auth/forgot-password/resend
Body: { email }

Hook: usePasswordPolicy.ts
Purpose: Fetch policy for meter/checklist
Endpoint: GET /v1/auth/password-policy
```

**State Management (Zustand)** (`packages/lib/src/stores/recovery-store.ts`):
```
interface RecoveryState {
  pendingEmail?: string;
  challengeId?: string | null;
  cooldownEndsAt?: number | null;
  setState: (p: Partial<RecoveryState>) => void;
  clear: () => void;
}
Usage: const { pendingEmail, setState } = useRecoveryStore();
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/auth/`):
```
Schema: forgotPasswordSchema.ts
  - email: string().email()

Schema: resetPasswordSchema.ts
  - newPassword: password policy rules
  - confirm: equals(newPassword)
  - signOutOthers: boolean().optional()

Schema: codeSchema.ts (if code flow used)
  - code: string().regex(/^\d{6}$/)
```

**Formatting & Security Utils** (`packages/lib/src/utils/`):
```
format/normalizeEmail.ts   → lowercase/trim
format/maskEmail.ts        → a***@example.com
security/csrf.ts           → attach CSRF token (web)
retry/backoff.ts           → resend cool‑downs
time/cooldown.ts           → server time sync + countdown helpers
```

**Type Definitions** (`packages/lib/src/types/auth.ts`):
```
Exports:
  - ForgotPasswordRequest/Response
  - ResetPasswordVerifyRequest/Response
  - ResetPasswordRequest/Response
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: AuthStack (Expo Router)
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Screens:
  - forgot-password
  - forgot-password-sent
  - reset-password-verify (optional)
  - reset-password
  - reset-password-success
```

**Native Features:**
```
Keyboard avoidance & scrolling
Clipboard detection may be disabled on password fields for safety
Deep linking for token verification via expo-linking
```

---

#### F) Layout Components

**Web:**
```
Layout: (auth) Layout
Path: apps/web/app/(auth)/layout.tsx
Features:
  - Header/footer
  - ErrorBoundary + Sentry
  - <Toaster/> mount
```

**Mobile:**
```
Layout: (public)/(auth) Stack Layout
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Features:
  - Header/back
  - Modal containers
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Routes: (auth)/forgot-password/error.tsx, (auth)/reset-password/error.tsx
Purpose: Render route fallback; reset action; Sentry capture

Component: ResetTokenErrorCard.tsx
Purpose: Show invalid/expired/used token guidance
```

**Loading States:**
```
Routes: (auth)/*/loading.tsx → skeletons
Component: Spinner.tsx → inline loading for submit/resend/verify
```

---

### Visual Representations

#### Screen 1 — Forgot Password (Email Submit)

**Web View (1280–1920px):**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Skillsier                                          Remembered it? Sign in → │
├─────────────────────────────────────────────────────────────────────────────┤
│ ┌────────────── Hero (tips/security) ────────────┐  ┌────────── Form ─────┐ │
│ │ • We’ll email a link to reset your password     │  │ Forgot your password │ │
│ │ • For security, we won’t say if an email exists │  │ ───────────────────  │ │
│ │ [Illustration]                                  │  │ Email *  [__@__ ]    │ │
│ └─────────────────────────────────────────────────┘  │ [ Send reset link ]  │ │
│                                                       │ ───────────────────  │ │
│                                                       │ {Captcha shown after │ │
│                                                       │  too many attempts}  │ │
│                                                       │ {Inline errors here} │ │
│ └───────────────────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Mobile View (375–430px):**
```
┌──────────────────────────┐
│  ←  Forgot password      │
├──────────────────────────┤
│ Email *   [____@__]      │
│                          │
│ ┏━━━━━━━━━━━━━━━━━━━━━┓  │
│ ┃  Send reset link    ┃  │
│ ┗━━━━━━━━━━━━━━━━━━━━━┛  │
│                          │
│ {Captcha appears after N failures} │
└──────────────────────────┘
```

**Key Elements & Interactions:**
- Non‑enumerating confirmation, captcha after failures, links to Sign in/Sign up.

---

#### Screen 1a — Forgot Password (Confirmation / “Check your email”)

**Web/Mobile:**
```
┌──────────────────────────────────────────────┐
│ ✅ If an account exists for a***@exa…        │
│    we’ve sent a reset link.                  │
│                                              │
│ Didn’t get it? [Resend 58s]  [Change email]  │
└──────────────────────────────────────────────┘
```

---

#### Screen 1b — Captcha Challenge (Progressive)

**Web Widget:**
```
┌────────────── CAPTCHA ──────────────┐
│ [ Turnstile / reCAPTCHA widget ]    │
│ Token: ********                     │
└─────────────────────────────────────┘
```

**Mobile Interop:**
```
┌──────────────────────────────────────┐
│ Complete verification to continue    │
│ [ Open challenge in browser ]        │
└──────────────────────────────────────┘
```

---

#### Screen 2 — Verify Reset (Token/Code)

**Token Success (Web):**
```
┌──────────────────────────────────────────────┐
│ 🔐 Link verified                             │
├──────────────────────────────────────────────┤
│ You can now set a new password.              │
│ [ Continue ]                                 │
└──────────────────────────────────────────────┘
```

**Token Error (Invalid/Expired):**
```
┌──────────────────────────────────────────────┐
│ ❌ This link is invalid or has expired.      │
├──────────────────────────────────────────────┤
│ [ Send a new link ]   [ Use a 6‑digit code ] │
└──────────────────────────────────────────────┘
```

**Code Entry (Alternative):**
```
┌──────────────────────────────────────────────┐
│ Enter the 6‑digit code                       │
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]         │
│ [ Verify ]       Didn’t get it? [Resend 35s] │
└──────────────────────────────────────────────┘
```

---

#### Screen 3 — Reset Password (Policy + Confirm + Sign‑out Others)

**Web View:**
```
┌──────────────────────────────────────────────┐
│ Set a new password                           │
├──────────────────────────────────────────────┤
│ New password *  [•••••••• ]  👁               │
│ Strength: ████○  Good                         │
│ ✓ 8+ chars  ✓ Upper  ✓ Number  ✗ Symbol       │
│ Confirm *      [•••••••• ]                    │
│ [ ] Sign out of other devices                 │
│                                              │
│ [ Update password ]                           │
└──────────────────────────────────────────────┘
```

**Mobile View:**
```
┌──────────────────────────┐
│  ←  Set new password     │
├──────────────────────────┤
│ New password * [•••• ]👁  │
│ Strength: ███○○           │
│ Confirm *     [•••• ]     │
│ [ ] Sign out of others    │
│ ┏━━━━━━━━━━━━━━━━━━━━━┓   │
│ ┃  Update password    ┃   │
│ ┗━━━━━━━━━━━━━━━━━━━━━┛   │
└──────────────────────────┘
```

---

#### Screen 4 — Success Handoff

```
┌──────────────────────────────────────────────┐
│ ✅ Password updated                          │
│ You can now sign in with your new password.  │
│ [ Go to Sign in ]                            │
└──────────────────────────────────────────────┘
```

---

#### Screen 5 — Rate‑Limited Cooldown / Resend Limit

```
┌──────────────────────────────────────────────┐
│ ⏳ Too many requests. Try again in 47s…      │
│ [OK]                                         │
└──────────────────────────────────────────────┘
```

---

### User Stories

> Abbrev: **W**=Web, **M**=Mobile.  
> Each story includes components, routes, APIs, and **at least 5 Happy** and **10 Bad** acceptance criteria.

##### Story 1 — Request Password Reset (Non‑Enumerating)

**Story**
```
As a User who forgot my password
I want to request a reset link without revealing whether my email exists
So that my privacy is protected while I regain account access
```

**Components & Routes**
- **W:** `(auth)/forgot-password/page.tsx`, `ForgotPasswordForm`, `Input`, `Button`, `CaptchaWidget`, `Alert`
- **M:** `/(public)/(auth)/forgot-password.tsx` with same logical form
- **APIs:** `POST /v1/auth/forgot-password` (always `{accepted:true}`), progressive captcha

**Acceptance Criteria — ✅ Happy**
1. **AC1.1** Valid email format → submission succeeds → navigate to “sent” screen.
2. **AC1.2** Response content/time indistinguishable regardless of existence.
3. **AC1.3** `pendingEmail` masked (e.g., `a***@exa…`) is shown on the “sent” screen.
4. **AC1.4** First input focused on load; Enter key submits; button shows spinner.
5. **AC1.5** Analytics event `auth.forgot.submit` fires exactly once; PII‑safe.

**Acceptance Criteria — ❌ Bad**
1. **AC1.6** Invalid email format → inline error + ARIA announcement.
2. **AC1.7** Network/5xx → non‑blocking banner with “Try again” keeps form state.
3. **AC1.8** Rate‑limit (429) → cooldown overlay shows accurate countdown.
4. **AC1.9** Captcha required (after N attempts) → must pass before submit allowed.
5. **AC1.10** Captcha provider blocked → offer “Open challenge in new window” flow.
6. **AC1.11** Low‑JS mode → fallback link to server‑rendered request page.
7. **AC1.12** Rapid double‑click submit → second request ignored (idempotency key).
8. **AC1.13** RTL locale → field alignment, punctuation, and masking respect RTL.
9. **AC1.14** Form resize/mobile keyboard → no layout shift causing tap‑target errors.
10. **AC1.15** Autofill inserts invalid text → validation still catches and explains.

---

##### Story 2 — “Check Your Email” + Resend with Cooldown

**Story**
```
As a Waiting User
I want to see clear confirmation and a timed resend option
So that I know what to do if the email doesn’t arrive
```

**Components & Routes**
- **W:** `(auth)/forgot-password/sent/page.tsx`, `ResendResetLink`, `MaskedEmail`
- **M:** `/(public)/(auth)/forgot-password-sent.tsx`
- **APIs:** `POST /v1/auth/forgot-password/resend` (server provides cooldown/ETA)

**Acceptance Criteria — ✅ Happy**
1. **AC2.1** Confirmation copy never states whether the email exists.
2. **AC2.2** Resend button disabled until cooldown ends; countdown ticks per second.
3. **AC2.3** After cooldown, clicking Resend shows toast “Sent!” and restarts timer.
4. **AC2.4** “Change email” returns to request form with previous value prefilled.
5. **AC2.5** Analytics: `auth.forgot.resend.submit` emits once per click.

**Acceptance Criteria — ❌ Bad**
1. **AC2.6** Excess resends → extended cooldown with explanatory tooltip.
2. **AC2.7** Server clock skew → UI syncs to server‑hinted expiry timestamp.
3. **AC2.8** Email provider delay → guidance card (check spam, filters, allowlist).
4. **AC2.9** Offline state → resend disabled; banner “You’re offline” appears.
5. **AC2.10** Resend API 5xx → graceful error; button re‑enabled after brief backoff.
6. **AC2.11** Masking mis‑renders in RTL → formatting corrected and tested.
7. **AC2.12** Screen reader not announcing countdown → adds polite live region.
8. **AC2.13** Keyboard focus lost after resend → focus returns to status message.
9. **AC2.14** Multiple tabs → only one tab allowed to resend; others show notice.
10. **AC2.15** Browser blocks notifications link (if offered) → provide alternative tips.

---

##### Story 3 — Verify Reset via Magic Link (Token)

**Story**
```
As a Recovering User
I want my emailed reset link to verify quickly
So that I can proceed to set a new password
```

**Components & Routes**
- **W:** `(auth)/reset-password/verify/page.tsx`, `ResetTokenVerifier`, `ExpiredTokenHelp`
- **M:** `/(public)/(auth)/reset-password-verify.tsx` (deep‑link optional)
- **APIs:** `POST /v1/auth/reset-password/verify`

**Acceptance Criteria — ✅ Happy**
1. **AC3.1** Valid token → verified → auto‑advance to reset form.
2. **AC3.2** Single‑use token is invalidated after first successful verify.
3. **AC3.3** Verification completes < 500ms (p95); skeleton shown during load.
4. **AC3.4** Analytics: `auth.reset.verify.success` fired once.
5. **AC3.5** Secure handoff data (challengeId/token) stored in memory only (no URL).

**Acceptance Criteria — ❌ Bad**
1. **AC3.6** Invalid/expired/used → error card; CTA: “Send new link” and “Use code”.
2. **AC3.7** Cross‑origin/mixed content link → blocked; clear security message.
3. **AC3.8** Token leaked in referer → immediately redacted via POST‑redirect pattern.
4. **AC3.9** Multiple clicks on same link → “Already used” explanation.
5. **AC3.10** Deep‑link opens on mobile without app → offer “Continue in browser”.
6. **AC3.11** Network/5xx → retry button (idempotent) retains context.
7. **AC3.12** Browser back incorrect cache → force fresh verify call.
8. **AC3.13** Clock skew affects TTL display → show relative, server‑time‑based text.
9. **AC3.14** Accessibility: status is announced to screen readers.
10. **AC3.15** Token param missing → redirect to Forgot with info banner.

---

##### Story 4 — Verify with 6‑Digit Code (Alternative)

**Story**
```
As a Recovering User without link access
I want to verify using a 6‑digit code
So that I can continue even if I can’t click the email link
```

**Components & Routes**
- **W:** `(auth)/reset-password/verify/page.tsx` (code mode), `CodeInput`, `Button`
- **M:** `/(public)/(auth)/reset-password-verify.tsx`
- **APIs:** `POST /v1/auth/reset-password/verify-code`, `POST /v1/auth/reset-password/resend-code`

**Acceptance Criteria — ✅ Happy**
1. **AC4.1** Entering correct 6‑digit code verifies challenge and advances.
2. **AC4.2** Code input auto‑advances between cells; supports paste.
3. **AC4.3** Resend code enforces cooldown with accessible countdown.
4. **AC4.4** Analytics: `auth.reset.code.verify.success` single‑fire.
5. **AC4.5** Time‑window tolerance ±1 step for TOTP‑style codes (if used).

**Acceptance Criteria — ❌ Bad**
1. **AC4.6** Wrong code → inline error; after N tries → lock challenge (message).
2. **AC4.7** Resend spam → rate‑limit; show next available time.
3. **AC4.8** Delivery failures (spam filters) → guidance + switch‑to‑link option.
4. **AC4.9** Code expired → error and CTA to resend.
5. **AC4.10** Pasted non‑digits → sanitized & error explained.
6. **AC4.11** Screen readers read cells in order; if not, provide “single field mode”.
7. **AC4.12** Keyboard nav/backspace across cells works consistently.
8. **AC4.13** Network/5xx during verify → retry without losing entered code.
9. **AC4.14** Multiple tabs verifying same challenge → second tab blocked gracefully.
10. **AC4.15** Locale digits (e.g., Arabic‑Indic) → normalized correctly.

---

##### Story 5 — Set New Password (Policy Enforcement)

**Story**
```
As a Security‑aware Platform
I want new passwords to meet strength and reuse policies
So that accounts remain protected
```

**Components & Routes**
- **W:** `(auth)/reset-password/page.tsx`, `ResetPasswordForm`, `PasswordInput`, `PasswordPolicyHints`, `Checkbox`
- **M:** `/(public)/(auth)/reset-password.tsx`
- **APIs:** `GET /v1/auth/password-policy`, `POST /v1/auth/reset-password`

**Acceptance Criteria — ✅ Happy**
1. **AC5.1** Meter & checklist update with each keystroke.
2. **AC5.2** Submit enabled only when policy satisfied and confirm matches.
3. **AC5.3** Option “Sign out of other devices” default ON (configurable).
4. **AC5.4** Successful reset → success page with “Go to Sign in” CTA.
5. **AC5.5** Analytics: `auth.reset.submit.success` emitted once.

**Acceptance Criteria — ❌ Bad**
1. **AC5.6** Weak password → unmet rules highlighted; submit blocked.
2. **AC5.7** Reuse of last N passwords → rejected with clear guidance.
3. **AC5.8** Token/challenge expired mid‑flow → redirect to verify with banner.
4. **AC5.9** Clipboard paste blocked on confirm (optional) → message explains reason.
5. **AC5.10** Network/5xx → form retains input; retry safe; no duplicate resets.
6. **AC5.11** Show/Hide eye not accessible → keyboard + aria fixes required.
7. **AC5.12** Password managers auto‑fill mismatch → detect & require confirm.
8. **AC5.13** Latency > 2s (p95) → loading indicator + non‑blocking tips.
9. **AC5.14** Form lost focus on submit → prevent navigation; confirm dialog.
10. **AC5.15** Non‑Latin characters mishandled → UTF‑8 preserved end‑to‑end.

---

##### Story 6 — Sign Out of Other Devices

**Story**
```
As a Security‑minded User
I want to sign out other active sessions when I reset my password
So that any compromised devices lose access
```

**Components & Routes**
- **W/M:** Reset form checkbox
- **APIs:** `POST /v1/auth/reset-password` with `signOutOthers: true`

**Acceptance Criteria — ✅ Happy**
1. **AC6.1** Server revokes other device sessions/refresh tokens.
2. **AC6.2** Email notification sent summarizing affected devices.
3. **AC6.3** Next requests from those devices require login.

**Acceptance Criteria — ❌ Bad**
1. **AC6.4** Partial revocation failure → user warned; server retries later.
2. **AC6.5** No device list → proceed best‑effort; note in audit log.
3. **AC6.6** Revocation takes time → banner clarifies may take a few minutes.
4. **AC6.7** User unchecks box → keep current sessions; banner clarifies risk.
5. **AC6.8** Audit logging unavailable → local queue until service back.
6. **AC6.9** Notification email bounce → still consider reset successful.
7. **AC6.10** Session fixation → server issues fresh session after reset.

---

##### Story 7 — Non‑Enumeration & Privacy Guarantees

**Story**
```
As a Privacy‑first Platform
I want non‑enumerating responses and copy
So that attackers cannot probe for valid accounts
```

**Components & Routes**
- **W/M:** Forgot + Sent screens copy, timers, masking

**Acceptance Criteria — ✅ Happy**
1. **AC7.1** Same response time/content regardless of account existence.
2. **AC7.2** UI copy avoids existence hints in all locales.
3. **AC7.3** Masked email never reveals more than allowed pattern.

**Acceptance Criteria — ❌ Bad**
1. **AC7.4** A/B copy reintroduces enumeration → test fails the build.
2. **AC7.5** Timing side‑channels detected → add artificial jitter window.
3. **AC7.6** Error messages differ by existence → consolidated generic copy.
4. **AC7.7** Analytics labels existence → remove/rename event fields.
5. **AC7.8** Logs include raw emails or tokens → scrub before write.
6. **AC7.9** Support link auto‑fills email → ensure masking or opt‑in.
7. **AC7.10** Local caches store PII → encrypt or avoid persisting.

---

##### Story 8 — Captcha & Rate Limiting

**Story**
```
As a Security System
I want to throttle abusive reset requests
So that automated attacks are mitigated without hurting UX
```

**Components & Routes**
- **W/M:** `CaptchaWidget`, cooldown overlay, banners

**Acceptance Criteria — ✅ Happy**
1. **AC8.1** Captcha appears only after N attempts; passing it unblocks request.
2. **AC8.2** Cooldown persists across refresh based on server hint.
3. **AC8.3** Countdown is accessible and localized.

**Acceptance Criteria — ❌ Bad**
1. **AC8.4** Captcha blocked by network → alternate link provided.
2. **AC8.5** Clock skew causes negative countdown → clamp to zero.
3. **AC8.6** VPN/CGNAT produce false positives → show appeal/help text.
4. **AC8.7** Retry storm due to flaky network → exponential backoff.
5. **AC8.8** Captcha token expired silently → auto‑refresh without data loss.
6. **AC8.9** Low‑JS devices → fallback verification flow.
7. **AC8.10** Screen reader can’t reach widget → keyboard path documented.
8. **AC8.11** Mobile browser keyboard overlaps → scroll into view automatically.
9. **AC8.12** Analytics leaks captcha score → never logged.

---

##### Story 9 — Accessibility & Localization

**Story**
```
As an Inclusive Platform
I want the recovery flow to be fully accessible and localized
So that all users can complete it successfully
```

**Acceptance Criteria — ✅ Happy**
1. **AC9.1** All labels/inputs/alerts have ARIA semantics and are announced.
2. **AC9.2** Tab order is logical; visible focus rings present.
3. **AC9.3** Copy localized; right‑to‑left layouts supported.
4. **AC9.4** Color contrast meets WCAG AA; automated lint passes.
5. **AC9.5** CodeInput supports screen readers via single‑field mode toggle.

**Acceptance Criteria — ❌ Bad**
1. **AC9.6** Missing translation key → English fallback + logged key.
2. **AC9.7** Motion/animation reduces on prefers‑reduced‑motion.
3. **AC9.8** Error text color fails contrast → build fails with a11y test.
4. **AC9.9** Focus trapped incorrectly in modal → escape/close works via keyboard.
5. **AC9.10** Live regions spam updates → throttled to polite frequency.
6. **AC9.11** Date/time in countdown not localized → i18n formatters used.
7. **AC9.12** Screen zoom 200% → layout remains usable (no overlap).

---

##### Story 10 — Analytics & Telemetry

**Story**
```
As a Product Team
I want to measure the recovery funnel
So that we can improve deliverability and reduce drop‑offs
```

**Acceptance Criteria — ✅ Happy**
1. **AC10.1** Events: `auth.forgot.*`, `auth.reset.verify.*`, `auth.reset.submit.*` fire once per action.
2. **AC10.2** No raw PII (emails, tokens, codes) in analytics.
3. **AC10.3** Client retries analytics on failure; non‑blocking.

**Acceptance Criteria — ❌ Bad**
1. **AC10.4** Duplicate events from double‑clicks → idempotency guard.
2. **AC10.5** Ad‑blockers prevent send → silently degrade; core UX unaffected.
3. **AC10.6** Sampling misconfigured → product dashboard warns on gaps.
4. **AC10.7** Event names change without version bump → validator fails build.
5. **AC10.8** Timezones skew funnel timing → all events UTC with client offset.
6. **AC10.9** Personally identifying metadata sneaks in via `context` → scrubbed.

---

##### Story 11 — Security Hardening (Tokens, Sessions, Logs)

**Story**
```
As a Security Team
I want strong token policies and safe defaults
So that the reset flow cannot be abused
```

**Acceptance Criteria — ✅ Happy**
1. **AC11.1** Tokens are single‑use, short TTL (≤30m), high entropy.
2. **AC11.2** Verify via POST; token never logged or stored in URL after verify.
3. **AC11.3** New session issued after reset; session fixation prevented.
4. **AC11.4** Audit logs record reset with scrubbed identifiers.
5. **AC11.5** CSP/Referrer‑Policy prevent token exfiltration.

**Acceptance Criteria — ❌ Bad**
1. **AC11.6** Token visible in client logs → remove logging, ship hotfix.
2. **AC11.7** Mixed content/link from HTTP → blocked and reported.
3. **AC11.8** Reuse token succeeds → test fails; rotate keys and invalidate.
4. **AC11.9** Reset endpoint lacks CSRF → build fails security checks.
5. **AC11.10** Session not rotated after reset → security test fails.

---

##### Story 12 — Escalation Without Email Access

**Story**
```
As a Locked‑out User
I want a secure escalation path when I can’t access my email
So that I can regain access after identity verification
```

**Acceptance Criteria — ✅ Happy**
1. **AC12.1** Clear guidance to contact support; identity checks explained.
2. **AC12.2** No self‑service flow exposes more data or weakens security.
3. **AC12.3** Ticket created with minimal metadata; no PII beyond user‑supplied.

**Acceptance Criteria — ❌ Bad**
1. **AC12.4** Automated “backdoor” path exists → explicitly disabled.
2. **AC12.5** Support form auto‑fills private info → removed; user‑entered only.
3. **AC12.6** Triage emails leak tokens/links → scrubbed before send.

---

## Non‑Functional Requirements (AUTH‑3 scope)

- **Performance:** First input usable < 1.0s (p95); verify < 500ms (p95); reset submit < 700ms (p95)  
- **Security:** Non‑enumeration; captcha; rate limits; CSRF; single‑use tokens; short TTL; password policy; session revocation; Sentry  
- **Privacy:** No tokens/PII in logs/analytics; masked email display  
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+  
- **Testing:** Unit tests (schemas/hooks/utils), e2e (happy + error), a11y checks (axe), non‑enumeration timing tests

---

**End of AUTH‑3 — User Stories (Rewritten & Complete).**

