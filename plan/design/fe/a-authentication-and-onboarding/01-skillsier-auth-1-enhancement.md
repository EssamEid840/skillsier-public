## AUTH-1 — User Registration

> Enhancement add‑on for the existing AUTH‑1 journey in `skillsier-frontend-journeys-claude-final.md`. This version **fully completes** the three required sections: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-strucure.md` routes and component patterns.

---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(auth)/register/page.tsx
Purpose: Entry page to create a Skillsier account with email + password and/or SSO.
Props: None (client component with internal form state + actions)
Key Features:
  - Responsive 2‑column layout (hero + form)
  - react-hook-form + Zod schema validation
  - Password strength meter + live policy checklist
  - Debounced email duplicate check
  - reCAPTCHA/Turnstile + CSRF token
  - Terms & Privacy consent
  - SSO: Google, Apple

Path: apps/web/app/(auth)/register/verification/page.tsx
Purpose: Email verification callback page (link target from email).
Props: token (via search params)
Key Features:
  - Verifies token with backend
  - Success and error states
  - Auto‑redirect to onboarding (ONB‑1)

Path: apps/web/app/(auth)/layout.tsx
Purpose: Shared frame for all auth pages (no dashboard chrome).
Features:
  - Header (brand + Sign in link)
  - Footer (legal links)
  - ErrorBoundary + Sentry
  - <Toaster/> mount

Path: apps/web/app/(auth)/register/loading.tsx
Purpose: Skeleton UI while route chunk loads.
Features:
  - Form skeleton
  - Disabled SSO button placeholders

Path: apps/web/app/(auth)/register/error.tsx
Purpose: Route‑scoped error display & reset.
Features:
  - Friendly message, “Try again”, Sentry capture
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Purpose: Stack layout for Register → (success inline) and deep link handling.
Features:
  - Header with back navigation
  - Modal for Terms & Privacy

Path: apps/mobile/app/(public)/(auth)/register.tsx
Purpose: Registration screen for anonymous users.
Props: None
Key Features:
  - Single‑column form, secure password input
  - Native autofill hints (email/name)
  - Debounced duplicate check
  - SSO via AuthSession (Google/Apple)
  - Inline success handoff → ONB‑1

(Email verification generally opens via magic link in browser; if app‑linking is configured, it can deep‑link into a mobile verification route.)
```

> Handoff: AUTH‑1 ends after account creation; **ONB‑1** handles verification & role selection. Web has a dedicated `register/verification` callback page; mobile uses deep linking to show verified state or prompts user to open email.

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: Button
Purpose: Primary/secondary/ghost CTAs (Create account, Retry, SSO)
Props: variant, size, isLoading, leftIcon, rightIcon

Component: Input
Purpose: First/Last Name, Email
Props: id, type, value, onChange, error, autoComplete

Component: PasswordInput
Path: packages/ui/src/components/forms/PasswordInput.tsx
Purpose: Secure input with show/hide toggle + meter slot
Props:
  interface PasswordInputProps {
    id: string;
    value: string;
    onChange: (v: string) => void;
    error?: string;
    onToggle?: () => void;
  }

Component: Checkbox
Purpose: Terms consent
Props: checked, onChange, children

Component: Form
Path: packages/ui/src/components/forms/Form.tsx
Purpose: RHF wrapper with Zod resolver & field components
Props: schema, defaultValues, onSubmit

Component: Card, Alert, InlineError, Tooltip, Toaster
Purpose: Layout and feedback patterns used across pages

Component: OAuthButton
Path: packages/ui/src/components/auth/OAuthButton.tsx
Purpose: Provider button with icon & accessibility labels
Props: provider, onClick, disabled, isLoading
```

**Feature Components** (`apps/web/components/auth/` and/or `packages/lib/src/auth/components/`):
```
Component: RegisterForm.tsx
Purpose: Encapsulates full sign‑up form (web wrapper; mobile wrapper reuses logic)
Props: onSuccess?: (payload) => void
State: RHF internal state; delegates side effects to hooks
APIs: useRegisterMutation, usePasswordPolicy, useCheckIdentifier

Component: SocialButtons.tsx
Purpose: Render Google/Apple buttons; start OAuth flows
Props: onStart(provider), isLoading
APIs: /v1/auth/oauth/start (web), AuthSession.startAsync (mobile)

Component: PasswordPolicyHints.tsx
Purpose: Live checklist (min length, classes, entropy)
Props: policy, value

Component: CaptchaWidget.tsx
Purpose: Render Turnstile/reCAPTCHA and surface token
Props: onToken(token), provider

Component: IdentifierStatus.tsx
Purpose: Inline status for email availability (✓ Available / Already exists)
Props: state ("idle" | "checking" | "ok" | "exists"), message

Component: TermsPrivacyModal.tsx
Purpose: Show ToS/Privacy in modal (web) / sheet (mobile)
```

**Domain‑Specific Components (Mobile/Web):**
```
Component: AuthHero.tsx (web)
Path: apps/web/components/auth/AuthHero.tsx
Purpose: Marketing/benefits in left column (desktop)

Component: AuthKeyboardAvoider.tsx (mobile)
Path: apps/mobile/components/auth/AuthKeyboardAvoider.tsx
Purpose: Avoid keyboard overlap for form
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/auth/hooks/` or similar):
```
Hook: useRegisterMutation.ts
Purpose: Create account with email/password
Endpoint: POST /v1/auth/register  (users-be/auth)
Return: { userId: string; email: string; requiresVerification: boolean }
Usage:
  const { mutate, isPending, error } = useRegisterMutation();
  const onSubmit = (data) => mutate(data, { onSuccess });

Hook: useCheckIdentifier.ts
Purpose: Debounced duplicate check for email/phone
Endpoint: POST /v1/auth/check-identifier
Return: { exists: boolean; canLogin: boolean }

Hook: usePasswordPolicy.ts
Purpose: Fetch password rules to render hints/meter
Endpoint: GET /v1/auth/password-policy
Return: { minLength: number; requireUpper: boolean; requireNumber: boolean; entropyMin: number }

Hook: useCaptchaToken.ts
Purpose: Expose captcha token, auto‑refresh on expiry
Endpoint: client‑side only; token verified server‑side

Hook: useRateLimit.ts
Purpose: Provide visual cooldown if server signals rate‑limit
```

**State Management (Zustand)** (`packages/lib/src/stores/auth-store.ts`):
```
interface AuthState {
  pendingEmail?: string;
  pendingUserId?: string;
  provider?: "google" | "apple" | null;
  setPending: (p: Partial<AuthState>) => void;
  clearPending: () => void;
}
Usage: const { pendingEmail, setPending } = useAuthStore();
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/auth/`):
```
Schema: registerSchema.ts
Validates: firstName, lastName, email, password, confirm, consent
Rules:
  - firstName: string().trim().min(2).max(50)
  - lastName:  string().trim().min(2).max(50)
  - email:     string().email()
  - password:  string().min(minLength) + custom policy checks (upper/number/symbol/entropy)
  - confirm == password
  - consent:   literal(true)

Schema: emailSchema.ts, passwordSchema.ts (reusable primitives)
```

**Formatting Utilities** (`packages/lib/src/utils/`):
```
format/formatName.ts      → Normalize human names
format/normalizeEmail.ts  → Lowercase/trim; provider‑safe normalization
security/csrf.ts          → Attach CSRF to requests (web)
security/xss.ts           → Encode/strip dangerous characters
```

**Type Definitions** (`packages/lib/src/types/auth.ts`):
```
Exports:
  - RegisterRequest, RegisterResponse
  - PasswordPolicy
  - IdentifierCheckRequest/Response
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: AuthStack (Expo Router group)
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Purpose: Stack navigation; modals for ToS/Privacy
Screens:
  - register
  - (deep link to verified state handled via Linking.addEventListener)
```

**Native Features:**
```
Feature: SSO via Expo AuthSession
Implementation: WebAuth + PKCE; deep link back to app
Libraries: expo-auth-session, expo-linking, expo-secure-store

Feature: Secure password entry + clipboard guard
Implementation: native secureTextEntry + platform hints
```

---

#### F) Layout Components

**Web Layouts:**
```
Layout: (auth) Layout
Path: apps/web/app/(auth)/layout.tsx
Purpose: Wraps all auth routes
Features:
  - Header/brand
  - Footer/legal
  - ErrorBoundary (react-error-boundary + Sentry)
  - <Toaster/> mount
```

**Mobile Layouts:**
```
Layout: (public)/(auth) Stack Layout
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Purpose: Header/back; gestures; modal containers
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Route: apps/web/app/(auth)/register/error.tsx
Purpose: Render fallback for runtime/render errors
Features:
  - Error summary + reset button
  - Sentry capture

Component: RegisterFormError.tsx
Purpose: API‑level errors (duplicate email, banned domain, captcha fail)
```

**Loading States:**
```
Route: apps/web/app/(auth)/register/loading.tsx
Purpose: Route skeleton
Features:
  - Inputs/labels skeleton
  - Disabled OAuth buttons

Component: Spinner.tsx
Purpose: Inline spinners for submit + SSO
```

---

### Visual Representations

#### Screen 1 — Create Account (Registration Form)

**Web View (1280–1920px):**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Skillsier                                                 Already have an   │
│                                                         account?  Sign in → │
├─────────────────────────────────────────────────────────────────────────────┤
│ ┌────────────── Hero (benefits) ────────────────┐  ┌────────── Form ──────┐ │
│ │ • Find great work • Hire experts • Safe escrow │  │ Create your account  │ │
│ │ • Secure messaging • Dispute protection        │  │ ──────────────────── │ │
│ │ [Illustration]                                 │  │ First name * [____]  │ │
│ └────────────────────────────────────────────────┘  │ Last name  * [____]  │ │
│                                                     │ Email *     [__@__]  │ │
│                                                     │          (✓ Available)│ │
│                                                     │ Password *  [••••] 👁 │ │
│                                                     │ Strength: ███○○ Weak  │ │
│                                                     │ Confirm *   [••••]    │ │
│                                                     │ [ ] I agree to Terms  │ │
│                                                     │    and Privacy        │ │
│                                                     │                       │ │
│                                                     │ [ Create account ]    │ │
│                                                     │ ────────────────────  │ │
│                                                     │ Or sign up with       │ │
│                                                     │ [ Google ]  [ Apple ] │ │
│                                                     │ reCAPTCHA • Protected │ │
│                                                     │ {Inline field errors} │ │
│ └──────────────────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Mobile View (375–430px):**
```
┌──────────────────────────┐
│  ←  Create account       │
├──────────────────────────┤
│ First name * [_______]   │
│ Last name  * [_______]   │
│ Email *     [____@__] ✓  │
│ Password *  [••••••] 👁  │
│ Strength: ███○○          │
│ Confirm  *  [••••••]     │
│ [ ] I agree to Terms     │
│                          │
│ ┏━━━━━━━━━━━━━━━━━━━━━┓  │
│ ┃   Create account    ┃  │
│ ┗━━━━━━━━━━━━━━━━━━━━━┛  │
│ Or sign up with          │
│ [ Google ]   [ Apple ]   │
│ reCAPTCHA • Protected    │
└──────────────────────────┘
```

**Key UI Elements:**
- Inputs + inline validation; email availability status (✓/✗)
- Password meter & live policy checklist (popover)
- Consent checkbox (ToS/Privacy links open modal/sheet)
- SSO buttons; disabled when provider unavailable
- Submit button with loading state
- Captcha widget (token persisted for submit)

---

#### Screen 1a — Terms & Privacy (Modal / Sheet)

**Web (Modal):**
```
┌──────────────────────────────────────────┐
│ Terms of Service                        ✕│
├──────────────────────────────────────────┤
│ [Scrollable legal content…]              │
│                                          │
│ [ Cancel ]                    [ Agree ]  │
└──────────────────────────────────────────┘
```

**Mobile (Bottom Sheet):**
```
┌──────────────────────────┐
│ ───── Terms of Service ─ │
│ [Scrollable legal…]      │
│ [ Cancel ]  [ Agree ]    │
└──────────────────────────┘
```

---

#### Screen 1b — Email Already Exists (Inline Banner)

**Web/Mobile:**
```
┌───────────────────────────────────────────────────────────────┐
│ ⚠ An account with this email exists.                          │
│   Continue to Sign in or reset your password.  [Sign in] [→]  │
└───────────────────────────────────────────────────────────────┘
```

---

#### Screen 1c — Captcha Challenge (Web Widget / Mobile Interop)

**Web (Widget area inside form):**
```
┌────────────── CAPTCHA ──────────────┐
│ [ Turnstile / reCAPTCHA widget ]    │
│ Token: ******** (auto‑refresh)      │
└─────────────────────────────────────┘
```

**Mobile (Challenge handoff):**
```
┌──────────────────────────────────────┐
│ Complete verification to continue    │
│ [ Open challenge in browser ]        │
└──────────────────────────────────────┘
```

---

#### Screen 1d — Rate Limited (Cooldown Overlay)

```
┌──────────────────────────────────────┐
│ ⏳ Too many attempts.                 │
│ Try again in 27s…                    │
│ [OK]                                 │
└──────────────────────────────────────┘
```

---

#### Screen 2 — Verification Callback (Web)

**Success:**
```
┌──────────────────────────────────────────────┐
│ ✅ Email verified                             │
├──────────────────────────────────────────────┤
│ Your account is ready.                       │
│                                              │
│ [ Continue to Onboarding ]                   │
└──────────────────────────────────────────────┘
```

**Error (invalid/expired token):**
```
┌──────────────────────────────────────────────┐
│ ❌ Verification failed                        │
├──────────────────────────────────────────────┤
│ Link invalid or expired.                      │
│ [ Resend verification email ]  [ Support ]    │
└──────────────────────────────────────────────┘
```

---

### User Stories

> The following stories decompose AUTH‑1 comprehensively. Each story references specific components, routes, and backend endpoints. (Abbrev: **W**=Web, **M**=Mobile.)

#### Story 1 — Register with Email & Password (Core)

**Story:**
```
As an Anonymous User
I want to create an account with my name, email, and a secure password
So that I can use Skillsier after verifying my email
```

**Components Used:**
- **W:** `apps/web/app/(auth)/register/page.tsx`, `RegisterForm.tsx`, `PasswordInput`, `Form`, `Input`, `Checkbox`, `Button`, `Alert`, `CaptchaWidget`, `PasswordPolicyHints`, `IdentifierStatus`
- **M:** `apps/mobile/app/(public)/(auth)/register.tsx` (+ same logical form components)

**Routes:**
- **W:** `/app/(auth)/register`
- **M:** `/(public)/(auth)/register`

**Backend:**
- `POST /v1/auth/register`
- `GET /v1/auth/password-policy`
- `POST /v1/auth/check-identifier`
- (Security) CSRF, captcha verification server‑side

**API Examples:**
```http
GET /v1/auth/password-policy

POST /v1/auth/check-identifier
{ "email": "user@example.com" } → { "exists": false, "canLogin": false }

POST /v1/auth/register
{
  "firstName": "Ava",
  "lastName": "Lee",
  "email": "ava@example.com",
  "password": "S3cur3P@ss!",
  "consent": true,
  "captchaToken": "cf-turnstile-token"
}
→ { "userId": "usr_123", "email": "ava@example.com", "requiresVerification": true }
```

**State (Zustand):**
```ts
interface AuthState {
  pendingEmail?: string;
  pendingUserId?: string;
  setPending: (p: Partial<AuthState>) => void;
  clearPending: () => void;
}
```

**Validation Rules:**
- First/Last name: 2–50 chars
- Email: RFC, normalized
- Password: per policy (length/classes/entropy)
- Confirm matches
- Consent checked

**Acceptance Criteria — Happy:**
1. Valid form → submit → account created; show success inline and store `pendingEmail`.
2. Password meter & checklist update live with keystrokes.
3. Debounced (300–500ms) email check shows “Available” when free.
4. Submit button disables during request; spinner visible.
5. After success, user sees “Check your email” prompt and CTA to ONB‑1.

**Acceptance Criteria — Bad:**
1. Duplicate email → inline banner w/ “Sign in” link; submit blocked.
2. Weak password → checklist shows unmet rules; submit blocked.
3. Consent unchecked → error “Accept Terms & Privacy to continue.”
4. Captcha missing/invalid → error toast; widget resets.
5. Network/5xx → global Alert + Retry; no duplicate account created.
6. Rate‑limit (429) → cooldown overlay with countdown.
7. Banned/disposable domain → error “Email domain not permitted.”
8. CSRF invalid → hard refresh & retry required.
9. Form navigation away with dirty state → confirm modal to prevent loss.
10. Localization missing for current locale → default to English fallback.

---

#### Story 2 — Sign Up with Google/Apple (SSO)

**Story:**
```
As an Anonymous User
I want to create an account using my Google or Apple identity
So that I can register quickly without creating a password
```

**Components:** `SocialButtons`, `OAuthButton`, `Alert`
- **W:** OAuth redirect flow; callback handled server‑side
- **M:** `expo-auth-session` + deep linking

**Routes:** same as Story 1 register routes

**Backend:**
- `GET /v1/auth/oauth/start?provider=google|apple`
- `POST /v1/auth/oauth/callback`

**Acceptance — Happy:**
1. Click “Continue with Google/Apple” starts provider flow.
2. Callback success → account created/linked; continue to ONB‑1.
3. Verified provider email → local email verification marked satisfied.
4. Provider avatar/email shown in confirmation.

**Acceptance — Bad:**
1. User cancels provider consent → neutral message; remain on form.
2. Callback error (invalid code) → error toast; allow retry.
3. Existing local account with same email → offer linking or Sign in.
4. Provider outage → disable buttons + tooltip; fallback to email.
5. Apple on Android misconfigured → guide to web fallback.

**NFR:** PKCE, state/nonce verification, secure cookie/session bootstrap.

---

#### Story 3 — Email Duplicate Detection & Guidance

**Story:**
```
As a Prospective User
I want to know if my email already has an account
So that I can sign in instead of re‑registering
```

**Components:** `IdentifierStatus`, `Alert`, `InlineHelp`

**Acceptance — Happy:**
1. On email blur, check runs; shows ✓ “Available” if free.
2. If exists & canLogin → shows “Sign in instead” link (prefills email).

**Acceptance — Bad:**
1. Endpoint fails → hide availability badge; allow submit if form valid.
2. Debounce/cancel in‑flight requests on quick edits.
3. Privacy: avoid detailed enumeration (generic copy only).

---

#### Story 4 — Password Policy Compliance & Meter

**Story:**
```
As a Security‑conscious Platform
I want users to meet password strength requirements
So that accounts are protected against common attacks
```

**Components:** `PasswordInput`, `PasswordPolicyHints`, `Tooltip`

**Acceptance — Happy:**
1. Checklist reflects policy rules (min length/classes/entropy).
2. Meter updates instantly with each keystroke.
3. Tooltip explains why a rule is unmet.

**Acceptance — Bad:**
1. Copy/paste weak password → blocked with clear reasons.
2. Attempt to submit while unmet → prevents submit + focuses password.
3. Browser password manager inserts weak value → still validated.

---

#### Story 5 — Terms & Privacy Consent

**Story:**
```
As a Business
I need explicit ToS/Privacy consent
So that we comply with legal requirements
```

**Components:** `Checkbox`, `TermsPrivacyModal`

**Acceptance — Happy:**
1. Clicking ToS/Privacy opens modal/sheet; closing returns focus to checkbox.
2. Consent stored as boolean + timestamp in register payload.

**Acceptance — Bad:**
1. Consent not checked → submit blocked with inline error.
2. Modal focus trap broken → cannot escape (accessibility test).

---

#### Story 6 — Anti‑Bot (Captcha) Verification

**Story:**
```
As a Platform
I want to prevent automated sign‑ups
So that abuse and spam are reduced
```

**Components:** `CaptchaWidget`

**Acceptance — Happy:**
1. Captcha loads and returns token; submit passes when token present.
2. Token auto‑refreshes on expiry without losing form state.

**Acceptance — Bad:**
1. Captcha provider blocked → show fallback guidance + “Open in new window”.
2. Token invalid at server → reset widget and surface error toast.
3. Accessibility mode → provide alternative challenge link.

---

#### Story 7 — Rate Limiting & Cooldown UX

**Story:**
```
As a Security System
I want to limit repeated sign‑up attempts
So that automated abuse is throttled
```

**Components:** `Alert`, cooldown overlay

**Acceptance — Happy:**
1. On 429 response, overlay shows countdown; submit disabled until zero.
2. After cooldown, interactions restored automatically.

**Acceptance — Bad:**
1. User refreshes during cooldown → remaining time restored from server hint.
2. Multiple tabs → only one tab allowed to retry; others show notice.

---

#### Story 8 — Resumable Registration State

**Story:**
```
As a User
I want the app to remember my pending email after submit
So that the verification step can prefill my address
```

**Components:** Zustand `auth-store`

**Acceptance — Happy:**
1. After successful submit, `pendingEmail` saved; ONB‑1 can read it.
2. Clearing pending info on complete onboarding.

**Acceptance — Bad:**
1. Clearing browser storage doesn’t break flow (fallback to manual entry).
2. Private window mode still works without persistent storage.

---

#### Story 9 — Localization & Accessibility

**Story:**
```
As a Global & Inclusive Platform
I want copy and ARIA attributes to be localized and accessible
So that users across locales and assistive tech can register
```

**Components:** i18n provider, `SkipLink`, labels/aria‑*

**Acceptance — Happy:**
1. Form labels, errors, and help texts localized.
2. Screen reader announces errors on submit; focus moves to first error.
3. Keyboard navigation: tab order logical; visible focus rings.

**Acceptance — Bad:**
1. Missing translation key → English fallback with tracking.
2. Low‑contrast theme fails AA → automated check fails build (lint/test).

---

#### Story 10 — Analytics & Telemetry

**Story:**
```
As a Product Team
I want to track registration funnel events
So that we can improve conversion and detect issues
```

**Events (examples):**
- `auth.register.view`
- `auth.register.email_checked` (exists/available)
- `auth.register.submit`
- `auth.register.success`
- `auth.register.failure` (reason)
- `auth.register.sso.start` / `.success` / `.error`
- `auth.register.cooldown.start` / `.end`

**Acceptance — Happy:**
1. Events fire exactly once per user action (debounced).
2. PII‑safe payloads; no raw passwords/emails in logs.

**Acceptance — Bad:**
1. Network failure on analytics → non‑blocking; queued retry.
2. Ad‑blockers → degrade gracefully; core UX unaffected.

---

## Non‑Functional Requirements (AUTH‑1 scope)

- **Performance:** First input usable < 1.0s (p95); submit ↔ response < 700ms (p95)
- **Security:** CSRF on web; captcha; server‑side normalization; password hashing (server); strict rate limits
- **Privacy:** Explicit consent; clear ToS/Privacy links; telemetry opt‑in where required
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+
- **Testing:** Unit tests for schema & hooks; e2e (happy + error); a11y checks (axe)

---

**End of AUTH‑1 Enhancements (Complete).**
