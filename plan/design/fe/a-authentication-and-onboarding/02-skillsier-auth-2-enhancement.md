## AUTH-2 — User Sign‑In (Email/Password, SSO, MFA, Passkeys)

> Enhancement add‑on for the existing AUTH‑2 journey in `skillsier-frontend-journeys-claude-final.md`. This deliverable fully completes the three required sections: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-strucure.md` route and component patterns. Scope includes classic sign‑in, SSO, MFA, rate‑limiting/captcha, “remember me”, device trust, passkeys/biometrics, and session handoff.

---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(auth)/login/page.tsx
Purpose: Primary sign‑in page for email/password and SSO.
Props: None (client component with RHF form state)
Key Features:
  - Email + password fields with validation
  - “Remember me” (longer session cookie)
  - SSO (Google, Apple)
  - Passkey (WebAuthn) quick sign‑in
  - Progressive captcha after failed attempts
  - Links: Forgot password, Create account

Path: apps/web/app/(auth)/login/mfa/page.tsx
Purpose: Two‑step MFA challenge after successful primary auth.
Props: `challengeId` via search params (or internal state)
Key Features:
  - Input one‑time code (TOTP/SMS/email)
  - “Trust this device” (long‑lived device cookie)
  - Recovery codes fallback
  - Resend code with cooldown

Path: apps/web/app/(auth)/login/magic-link/page.tsx
Purpose: Passwordless sign‑in confirmation (magic link target).
Props: `token` via search params
Key Features:
  - Verifies token, signs user in, and redirects

Path: apps/web/app/(auth)/layout.tsx
Purpose: Shared frame for all auth pages (no app chrome).
Features:
  - Brand header + “Sign up” link
  - Footer with ToS/Privacy
  - ErrorBoundary + Sentry
  - <Toaster/> mount

Path: apps/web/app/(auth)/login/loading.tsx
Purpose: Route‑level skeleton while chunks load.

Path: apps/web/app/(auth)/login/error.tsx
Purpose: Friendly error boundary with reset button.
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Purpose: Stack layout for Login → MFA (and deep‑links).
Features:
  - Header/back
  - Modal containers (ToS/Privacy)

Path: apps/mobile/app/(public)/(auth)/login.tsx
Purpose: Primary sign‑in screen (email/password + SSO, biometric).
Props: None
Key Features:
  - Single‑column form (keyboard‑safe)
  - SSO via AuthSession (Google/Apple)
  - Biometric quick sign‑in (Face ID/Touch ID)
  - Progressive captcha after failures
  - Forgot password, sign‑up links

Path: apps/mobile/app/(public)/(auth)/login-mfa.tsx
Purpose: Step‑up MFA challenge screen.
Props: challengeId (deep‑link or navigation param)
Key Features:
  - Code entry + resend + trust device
  - Recovery codes fallback

Path: (Optional deep link) apps/mobile/app/(public)/(auth)/magic-link.tsx
Purpose: Handle passwordless magic‑link sign‑in via Linking.
```

> Handoff: Successful sign‑in redirects into the appropriate post‑login route (e.g., `(dashboard)` or last‑visited location).

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: Button
Purpose: Primary “Sign in”, “Verify”, “Resend”, social CTAs
Props: variant, size, isLoading, leftIcon, rightIcon

Component: Input
Purpose: Email, password, MFA code input
Props: id, type, value, onChange, error, autoComplete

Component: PasswordInput
Purpose: Secure input + show/hide toggle
Props: value, onChange, error

Component: CodeInput (6‑digit segmented)
Path: packages/ui/src/components/forms/CodeInput.tsx
Purpose: Numeric MFA code entry with auto‑advance
Props:
  interface CodeInputProps {
    length?: number; value: string; onChange: (v: string) => void;
    error?: string; onComplete?: (v: string) => void;
  }

Component: Checkbox
Purpose: “Remember me”, “Trust this device”
Props: checked, onChange, children

Component: Alert / InlineError / Banner
Purpose: Wrong credentials, account locked, policy notices
Props: title, description, variant

Component: OAuthButton
Purpose: Social sign‑in buttons for Google/Apple
Props: provider, onClick, disabled, isLoading

Component: PasskeyButton
Path: packages/ui/src/components/auth/PasskeyButton.tsx
Purpose: WebAuthn quick sign‑in (browser)
Props: onClick, isLoading

Component: BiometricButton.(native|web).tsx
Purpose: Native biometric (Face ID/Touch ID) or WebAuthn fallback
Props: onClick, isAvailable
```

**Feature Components** (`packages/lib/src/features/auth/components/` or `apps/*/components/auth/`):
```
Component: LoginForm.tsx
Purpose: Encapsulate email/password sign‑in + validation
Props: onSuccess, onMfaRequired
APIs: useLoginMutation, useAuthCaptcha, useRateLimit

Component: SocialButtons.tsx
Purpose: Provider buttons and OAuth start flows
Props: onStart(provider), isLoading

Component: MfaChallengeForm.tsx
Purpose: Input & verify OTP (TOTP/SMS/email); resend; recovery code
Props: challengeId, onSuccess
APIs: useMfaVerifyMutation, useMfaResendMutation

Component: CaptchaWidget.tsx
Purpose: Turnstile/reCAPTCHA + token management
Props: onToken(token)

Component: DeviceTrustCheckbox.tsx
Purpose: “Trust this device” UX; explains implications
```

**Domain‑Specific Components (Web/Mobile):**
```
Component: AuthHero.tsx (web) — benefits column
Component: AuthKeyboardAvoider.tsx (mobile) — avoid keyboard overlap
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/features/auth/hooks/`):
```
Hook: useLoginMutation.ts
Purpose: Primary sign‑in with email/password
Endpoint: POST /v1/auth/login  (users-be/auth)
Return: { userId, mfaRequired?: boolean, challengeId?: string }
Usage:
  const { mutate, isPending, error } = useLoginMutation();

Hook: useMfaVerifyMutation.ts
Purpose: Verify MFA code, trust device
Endpoint: POST /v1/auth/mfa/verify
Body: { challengeId, code, trustDevice?: boolean }

Hook: useMfaResendMutation.ts
Purpose: Resend MFA code with cooldown
Endpoint: POST /v1/auth/mfa/resend
Body: { challengeId }

Hook: useAuthCaptcha.ts
Purpose: Show/refresh captcha after N failures
Endpoint: client‑side only; token validated server‑side

Hook: usePasskeyLogin.ts (web)
Purpose: Trigger WebAuthn assertion
Endpoint: POST /v1/auth/webauthn/assert

Hook: useBiometricLogin.ts (mobile)
Purpose: Native biometric auth → session bootstrap
```

**State Management (Zustand)** (`packages/lib/src/stores/auth-store.ts`):
```
interface AuthState {
  lastEmail?: string;
  challengeId?: string | null;
  postLoginRedirect?: string | null;
  setState: (p: Partial<AuthState>) => void;
  clear: () => void;
}
Usage: const { lastEmail, setState } = useAuthStore();
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/auth/`):
```
loginSchema.ts
  - email: string().email()
  - password: string().min(8)
rememberMeSchema.ts (boolean)
mfaCodeSchema.ts (6 digits numeric)
recoveryCodeSchema.ts (alphanumeric groups)
```

**Formatting & Security Utils** (`packages/lib/src/utils/`):
```
format/normalizeEmail.ts      → lowercase/trim provider‑safe
security/csrf.ts              → attach CSRF token (web)
security/webauthn.ts          → wrap PublicKeyCredential APIs
security/session.ts           → cookie/session helpers
retry/backoff.ts              → exponential backoff for resend/login
```

**Type Definitions** (`packages/lib/src/types/auth.ts`):
```
Exports:
  - LoginRequest, LoginResponse
  - MfaVerifyRequest/Response
  - PasskeyAssertionRequest/Response
  - MagicLinkRequest/Response
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: AuthStack (Expo Router)
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Screens:
  - login
  - login-mfa
  - (magic-link, optional)
```

**Native Features:**
```
Biometrics: expo-local-authentication (Face ID/Touch ID)
SSO: expo-auth-session + expo-linking (Google/Apple)
Secure Storage: expo-secure-store for refresh/session secrets
```

---

#### F) Layout Components

**Web:**
```
Layout: (auth) Layout
Path: apps/web/app/(auth)/layout.tsx
Features:
  - Brand header/footer
  - ErrorBoundary (react-error-boundary + Sentry)
  - <Toaster/> mount
```

**Mobile:**
```
Layout: (public)/(auth) Stack Layout
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Features:
  - Header/back
  - Modals for ToS/Privacy
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Route: apps/web/app/(auth)/login/error.tsx
Purpose: Render fallback for runtime/render errors
Features: Error summary + reset + Sentry capture

Component: LoginFormError.tsx
Purpose: Wrong credentials, lockout, network failure
```

**Loading States:**
```
Route: apps/web/app/(auth)/login/loading.tsx → route skeleton
Component: Spinner.tsx → inline loading for submit, resend
```

---

### Visual Representations

#### Screen 1 — Sign‑In (Email/Password)

**Web View (1280–1920px):**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Skillsier                                           New here?  Sign up →   │
├─────────────────────────────────────────────────────────────────────────────┤
│ ┌────────────── Hero (benefits) ────────────────┐  ┌────────── Form ──────┐ │
│ │ • Safe escrow • Messaging • Talent network     │  │ Sign in               │ │
│ │ [Illustration]                                 │  │ ────────────────────  │ │
│ └────────────────────────────────────────────────┘  │ Email *    [__@__ ]   │ │
│                                                     │ Password * [•••• ] 👁 │ │
│                                                     │ [ ] Remember me       │ │
│                                                     │                       │ │
│                                                     │ [ Sign in ]           │ │
│                                                     │ ────────────────────   │ │
│                                                     │ Or continue with       │ │
│                                                     │ [ Google ]  [ Apple ]  │ │
│                                                     │ [ Use a passkey ]      │ │
│                                                     │                       │ │
│                                                     │ Forgot password?       │ │
│                                                     │ {Inline errors appear} │ │
│ └──────────────────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Mobile View (375–430px):**
```
┌──────────────────────────┐
│  ←  Sign in              │
├──────────────────────────┤
│ Email *    [____@__]     │
│ Password * [••••••]  👁  │
│ [ ] Remember me          │
│                          │
│ ┏━━━━━━━━━━━━━━━━━━━━━┓  │
│ ┃      Sign in        ┃  │
│ ┗━━━━━━━━━━━━━━━━━━━━━┛  │
│ Or continue with          │
│ [ Google ]  [ Apple ]     │
│ [ Use a passkey ]         │
│ Forgot password?          │
└──────────────────────────┘
```

**Key Elements:**
- Email/password + Remember me
- SSO buttons
- Passkey button
- Inline error banners
- Accessibility: labels, aria‑*, visible focus

---

#### Screen 1a — Wrong Credentials (Inline Banner)

```
┌───────────────────────────────────────────────────────────────┐
│ ❌ Incorrect email or password. Try again or reset password.  │
│ [ Reset password ]                                            │
└───────────────────────────────────────────────────────────────┘
```

---

#### Screen 1b — Captcha Challenge (After N Failures)

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

#### Screen 2 — MFA Challenge

**Web View:**
```
┌──────────────────────────────────────────────┐
│ Additional verification required             │
├──────────────────────────────────────────────┤
│ Enter the 6‑digit code from your authenticator│
│                                              │
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]         │
│ [ ] Trust this device                        │
│                                              │
│ [ Verify ]    Didn’t get a code? [Resend 30s]│
│ Use a recovery code                          │
└──────────────────────────────────────────────┘
```

**Mobile View:**
```
┌──────────────────────────┐
│  ←  Verify               │
├──────────────────────────┤
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]        │
│ [ ] Trust this device                        │
│ [ Verify ]        [Resend 25s]               │
│ Use a recovery code                          │
└──────────────────────────┘
```

---

#### Screen 2a — Recovery Code

```
┌──────────────────────────────────────────────┐
│ Enter a recovery code                        │
│ Code: [XXXX‑XXXX‑XXXX]                       │
│ [ Verify ]                                   │
└──────────────────────────────────────────────┘
```

---

#### Screen 3 — Magic Link Confirmation (Passwordless)

```
┌──────────────────────────────────────────────┐
│ ✅ You’re signed in                          │
│ Redirecting to your dashboard…               │
└──────────────────────────────────────────────┘
```

---

#### Screen 4 — Account Locked / Disabled

```
┌──────────────────────────────────────────────┐
│ 🚫 Your account is locked                    │
│ Too many failed attempts. Try again later or │
│ contact support.                             │
│ [ Learn more ]  [ Contact support ]          │
└──────────────────────────────────────────────┘
```

---

#### Screen 5 — Rate‑Limited Cooldown

```
┌──────────────────────────────────────────────┐
│ ⏳ Too many attempts. Try again in 47s…      │
│ [OK]                                         │
└──────────────────────────────────────────────┘
```

---

#### Screen 6 — Passkey / Biometric Prompt (Conceptual)

```
┌──────────────────────────────────────────────┐
│ Use a passkey to sign in                     │
│ [ Continue ]                                  │
│ Cancel                                       │
└──────────────────────────────────────────────┘
```

---

### User Stories

> The stories below decompose AUTH‑2 comprehensively. Each includes components, routes, backend endpoints, validations, and acceptance criteria. (Abbrev: **W**=Web, **M**=Mobile.)

#### Story 1 — Sign‑In with Email & Password (Core)

**Story:**
```
As a Returning User
I want to sign in with my email and password
So that I can access my Skillsier account
```

**Components Used:**
- **W:** `apps/web/app/(auth)/login/page.tsx`, `LoginForm.tsx`, `Input`, `PasswordInput`, `Checkbox`, `Button`, `Alert`, `CaptchaWidget`
- **M:** `apps/mobile/app/(public)/(auth)/login.tsx`, same core form components

**Routes:**
- **W:** `/app/(auth)/login`
- **M:** `/(public)/(auth)/login`

**Backend Services (users‑be/auth):**
- `POST /v1/auth/login` — email/password
- Progressive: `POST /v1/anti-bot/verify` (if captcha shown)

**Request/Response (examples):**
```http
POST /v1/auth/login
{ "email":"ali@ex.com", "password":"S3cret!", "remember":true, "captchaToken":"..." }
→ { "userId":"usr_123", "mfaRequired": false, "session":"..." }
or
→ { "mfaRequired": true, "challengeId":"ch_abc", "methods":["totp","sms"] }
```

**Validation Rules:**
- Email required & valid; password min 8
- Remember me boolean

**Acceptance — Happy:**
1. Valid credentials → session established; redirect to last route or dashboard.
2. “Remember me” extends session persistence (server‑side cookie TTL).
3. First input focus in < 1s; submit disables with spinner.
4. If `mfaRequired=false`, no intermediate UI; direct redirect.
5. On success, store `lastEmail` for convenience.

**Acceptance — Bad:**
1. Wrong credentials → inline error; password cleared; focus returns to password.
2. Too many failures → captcha displayed; must pass to continue.
3. Account disabled/banned → banner with support link; block sign‑in.
4. Email not verified → CTA to resend verification (routes to ONB‑1).
5. Rate‑limit (429) → cooldown overlay with countdown.
6. Network/5xx → global Alert + Retry (no duplicate attempts).
7. Locale missing → English fallback; track missing key.
8. CSRF missing/invalid → hard refresh required, message shown.
9. Multiple tabs submitting → reject stale request (idempotency key).
10. Keyboard autofill mismatch (email vs stored) → warning; user can proceed.

---

#### Story 2 — MFA Challenge (TOTP/SMS/Email)

**Story:**
```
As a Security‑conscious Platform
I want a second factor for high‑assurance sign‑in
So that compromised passwords alone cannot grant access
```

**Components Used:** `MfaChallengeForm`, `CodeInput`, `Checkbox (Trust device)`, `Button`, `Alert`

**Routes:**
- **W:** `/app/(auth)/login/mfa?challengeId=ch_abc`
- **M:** `/(public)/(auth)/login-mfa` (param: `challengeId`)

**Backend (users‑be/security/mfa):**
- `POST /v1/auth/mfa/verify` — verify code `{ challengeId, code, trustDevice }`
- `POST /v1/auth/mfa/resend` — resend & cooldown

**Acceptance — Happy:**
1. Correct 6‑digit code → session established; redirect per `postLoginRedirect`.
2. Trust device checked → long‑lived device cookie returned and stored.
3. Resend enforces cooldown; shows remaining seconds.
4. Recovery code works as fallback (single‑use, rotates).
5. Remembered device skips MFA next time (until revoked).

**Acceptance — Bad:**
1. Wrong/expired code → error; after N tries, lock challenge; require restart.
2. Resend spam → rate‑limited; shows cooldown + tooltip.
3. Time‑drift TOTP → server allows ±1 window; otherwise error with hint.
4. Challenge expired (e.g., 10 minutes) → redirect back to login with message.
5. Recovery code reused → error; link to generate new codes (Settings).
6. Device trust cookie blocked by browser → warn that MFA will repeat.
7. SMS carrier delay → guidance card; option to switch to authenticator.
8. Accessibility: segmented inputs navigate with arrows; screen reader announces errors.
9. Network/5xx → non‑blocking retry; no duplicate verifies.
10. Session fixation prevention: new session issued after MFA.

---

#### Story 3 — Social Sign‑In (Google/Apple)

**Story:**
```
As a Returning User
I want to sign in using my Google or Apple account
So that I can access quickly without typing passwords
```

**Components Used:** `SocialButtons`, `OAuthButton`, `Alert`

**Routes:** Same as Story 1 (login page)

**Backend (users‑be/auth):**
- `GET /v1/auth/oauth/start?provider=google|apple`
- `POST /v1/auth/oauth/callback`

**Acceptance — Happy:**
1. Clicking a provider starts OAuth flow; callback establishes session.
2. If the provider email matches an existing account → linked & signed in.
3. If new user & provider verified → may bypass email verification.
4. Successful sign‑in emits analytics and redirects to dashboard.
5. On mobile, AuthSession completes and returns to the app route.

**Acceptance — Bad:**
1. User cancels provider consent → neutral banner; stay on login.
2. Provider invalid code/state → error toast; allow retry.
3. Existing local‑password account with same email → offer linking or “Sign in with password”.
4. Provider outage → disable buttons + explanatory tooltip.
5. Apple on non‑configured Android → suggest web fallback.

---

#### Story 4 — Passkey / Biometric Sign‑In

**Story:**
```
As a Security‑savvy User
I want to sign in with a passkey or biometrics
So that I can authenticate quickly and securely
```

**Components Used:** `PasskeyButton` (web), `BiometricButton.native` (mobile)

**Routes:** Same login routes

**Backend:**
- `POST /v1/auth/webauthn/assert` — assertion → session
- (Mobile) Secure local validation followed by server session bootstrap

**Acceptance — Happy:**
1. Browser shows platform prompt; on success, session established.
2. Mobile Face ID/Touch ID prompt signs in if device bound to account.
3. If device not registered → helpful message with link to set up passkeys.
4. Analytics record passkey vs password usage.

**Acceptance — Bad:**
1. No WebAuthn support → button hidden or disabled with tooltip.
2. Biometric unavailable/locked out → fallback to password.
3. Assertion fails (origin mismatch) → security error and guidance.
4. Multiple credentials → selector UI shown; user can cancel gracefully.
5. User denies biometric prompt → return to form without error state.

---

#### Story 5 — Account States (Unverified/Locked/Disabled)

**Story:**
```
As a Platform
I want to handle account states during sign‑in
So that users receive clear guidance and security is enforced
```

**Components Used:** `Banner`, `Alert`, `LinkButton`

**Acceptance — Happy:**
1. Unverified email → banner with “Resend verification” (routes to ONB‑1).
2. Locked due to failures → message with unlock ETA and contact support.
3. Admin‑disabled → message with support request link.

**Acceptance — Bad:**
1. Resend verification rate‑limited → cooldown message.
2. Repeated attempts during lock → extend cooldown (configurable).
3. Attempt to sign in with pending‑deletion account → block + restore flow link.

---

#### Story 6 — Progressive Captcha & Rate Limiting

**Story:**
```
As a Security System
I want to throttle abusive sign‑in attempts
So that automated attacks are mitigated
```

**Components Used:** `CaptchaWidget`, cooldown overlay, `Alert`

**Acceptance — Happy:**
1. After N failed attempts, captcha appears; passing it lets user retry.
2. Rate‑limit returns remaining seconds; UI counts down accurately.
3. Cooldown persists across refresh (via server hint).

**Acceptance — Bad:**
1. Captcha provider blocked → offer alternative challenge/assist.
2. User in low‑JS mode → provide fallback action link.
3. Time skew → server time returned; countdown syncs.

---

#### Story 7 — Remember Me & Session Persistence

**Story:**
```
As a Returning User
I want my session to persist on trusted devices
So that I don’t have to sign in frequently
```

**Components Used:** `Checkbox`, session utils

**Acceptance — Happy:**
1. “Remember me” sets longer‑lived cookie/session (server).
2. Subsequent visits auto‑redirect to dashboard while session valid.
3. Logout clears session and Remember flag.

**Acceptance — Bad:**
1. Browser blocks third‑party cookies → fallback to 1P, warn if needed.
2. Device clock skew causes cookie expiry confusion → safe defaults + guidance.
3. Multiple concurrent sessions allowed/limited per policy; excess sign‑outs handled cleanly.

---

#### Story 8 — Passwordless Magic Link

**Story:**
```
As a Convenience‑seeking User
I want to sign in with a magic link
So that I can access without typing passwords
```

**Components Used:** “Send me a magic link” CTA (optional), Magic‑link page

**Backend:**
- `POST /v1/auth/magic-link` — request link
- `POST /v1/auth/magic-link/verify?token=...` — verify & sign‑in

**Acceptance — Happy:**
1. Request link sends email with one‑click sign‑in URL.
2. Clicking link signs in and redirects (anti‑replay enforced).
3. Link expires after short TTL; can request again after cooldown.

**Acceptance — Bad:**
1. Email not delivered → retry after N seconds; show support tips.
2. Link reused/expired → error with “Send a new link” CTA.
3. Device mismatch → still allowed but with additional check (captcha).

---

#### Story 9 — Login History & “Was This You?”

**Story:**
```
As a Security‑aware User
I want to review my login history
So that I can spot suspicious sign‑ins
```

**Components Used:** Link to login history (outside AUTH‑2 happy path but referenced), security banner

**Backend:**
- `GET /v1/users/me/login-history`

**Acceptance — Happy:**
1. After first login from a new device/location, show subtle banner with link.
2. Clicking “Not me” opens security flow (password reset + session revoke).

**Acceptance — Bad:**
1. GeoIP uncertain → show “Unknown location” label with tooltip.
2. Offline → queue fetch and retry silently later.

---

#### Story 10 — Analytics & Telemetry

**Story:**
```
As a Product Team
I want to instrument the sign‑in funnel
So that we can measure conversion and reliability
```

**Events (examples):**
- `auth.login.view`
- `auth.login.submit`
- `auth.login.success` / `.failure` (reason: wrong_password, lockout, captcha, mfa_required, network, disabled)
- `auth.login.mfa.view` / `.success` / `.failure`
- `auth.login.sso.start` / `.success` / `.error`
- `auth.login.passkey.start` / `.success` / `.error`
- `auth.login.cooldown.start` / `.end`

**Acceptance — Happy:**
1. Events fire exactly once per user action (debounced, idempotent).
2. PII‑safe payloads (no raw passwords/emails).

**Acceptance — Bad:**
1. Analytics network failure → non‑blocking; local queue → retry.
2. Ad‑blockers → degrade gracefully; core UX unaffected.

---

#### Story 11 — Localization & Accessibility

**Story:**
```
As a Global & Inclusive Platform
I want sign‑in to be fully localized and accessible
So that users across locales and assistive tech can authenticate
```

**Acceptance — Happy:**
1. All labels, errors, tooltips localized; dynamic direction (LTR/RTL).
2. Screen reader announces errors; focus moves to first invalid field.
3. Keyboard‑only users can operate all sign‑in/MFA interactions.
4. Color contrast AA; visible focus outlines.

**Acceptance — Bad:**
1. Missing translation → English fallback + logged key.
2. Captcha presents accessible alternative path.
3. Segmented code input usable via arrow keys/backspace.

---

## Non‑Functional Requirements (AUTH‑2 scope)

- **Performance:** First input usable < 1.0s (p95); login round‑trip < 600ms (p95); MFA verify < 500ms (p95)
- **Security:** CSRF on web; captcha after failures; rate‑limits; secure session cookies; WebAuthn where supported; device trust cookies scoped & revocable; session fixation prevention
- **Privacy:** No secret data in logs/analytics; clear ToS/Privacy links
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+
- **Testing:** Unit tests for schemas & hooks; e2e paths (password, SSO, MFA, passkey); a11y checks

---

**End of AUTH‑2 Enhancements (Complete).**
