## AUTH-5 — Email/Phone Verification & Account Activation (Links, 6‑Digit Codes, Resend, Bounce Handling, Policy Gate)

> Enhancement add‑on for the existing AUTH‑5 journey in `skillsier-frontend-journeys-claude-final.md`.
> This deliverable fully completes the three sections per the **Enhancement Prompt Template**: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-structure.md` conventions.
> Scope includes: email verification (magic link + code), phone verification (SMS code), resend with cooldown and captcha, change contact before verification, deep‑link handoff, bounce handling, pre‑verification restrictions, analytics, a11y, and security hardening.

---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(auth)/verify/email/request/page.tsx
Purpose: Request/trigger email verification link (or code) after signup/sign‑in.
Key Features:
  - Shows current (masked) email
  - “Send verification email” CTA
  - Option: “Use 6‑digit code instead”
  - Resend cooldown & progressive captcha
  - Change email flow link
  - Links to ONB‑1

Path: apps/web/app/(auth)/verify/email/sent/page.tsx
Purpose: Confirmation screen with resend timer + guidance (check spam, etc.).
Props: masked email via store

Path: apps/web/app/(auth)/verify/email/confirm/page.tsx
Purpose: Magic‑link landing; verifies token then forwards to success or help.
Props: `token` (search param)

Path: apps/web/app/(auth)/verify/email/code/page.tsx
Purpose: Enter 6‑digit code sent via email (alternative to magic link).
Key Features:
  - 6‑cell input with paste support
  - Resend with cooldown

Path: apps/web/app/(auth)/verify/phone/page.tsx
Purpose: Phone verify (SMS). Enter phone → send code → verify.
Key Features:
  - E.164 phone input with country picker
  - Code input + resend
  - Change phone link

Path: apps/web/app/(auth)/verify/success/page.tsx
Purpose: Unified success page; returns users to last destination (ONB‑1/2/3).
Key Features:
  - Toast “Verified”
  - “Continue” CTA

Path: apps/web/app/(settings)/(account)/verification/page.tsx
Purpose: Verification Center inside settings (status for email/phone, actions).
Key Features:
  - “Verified” badges
  - Resend/change contact details
  - Organization policy notices

Path: apps/web/app/(auth)/layout.tsx
Purpose: Shared auth frame (no main app chrome), Toaster, ErrorBoundary.

Path: apps/web/app/(auth)/verify/*/loading.tsx
Path: apps/web/app/(auth)/verify/*/error.tsx
Purpose: Skeletons and route‑level error fallbacks.
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(public)/(auth)/verify-email-request.tsx
Purpose: Request verification email (or choose code option).

Path: apps/mobile/app/(public)/(auth)/verify-email-sent.tsx
Purpose: Confirmation with resend timer.

Path: apps/mobile/app/(public)/(auth)/verify-email-confirm.tsx
Purpose: Handle email magic‑link deep link (Linking).

Path: apps/mobile/app/(public)/(auth)/verify-email-code.tsx
Purpose: Enter 6‑digit code (email).

Path: apps/mobile/app/(public)/(auth)/verify-phone.tsx
Purpose: Phone verify (SMS).

Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/account/verification.tsx
Purpose: Settings Verification Center.
```

> Handoff: Successful verification drops **account_status** from `PENDING_VERIFICATION` → `ACTIVE` and unlocks flows in ONB and the main app.

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: Button / LinkButton
Purpose: Primary/secondary CTAs (Send, Verify, Resend, Change)

Component: Input
Purpose: Email/phone entry; ARIA‑compliant helper/errors

Component: PhoneInput
Path: packages/ui/src/components/forms/PhoneInput.tsx
Purpose: Country picker + E.164 formatting & validation

Component: CodeInput
Path: packages/ui/src/components/forms/CodeInput.tsx
Purpose: 6‑digit input with auto‑advance & paste support

Component: Alert / Banner / InlineError
Purpose: Token invalid, rate‑limit, bounce detected, policy warnings

Component: CaptchaWidget
Purpose: Progressive anti‑bot (Turnstile/reCAPTCHA) after N resends

Component: MaskedEmail / MaskedPhone
Purpose: Safe display of contact (e.g., a***@exa…, +1 ***‑***‑1234)

Component: Countdown
Purpose: Accessible countdown for resend cooldown

Component: EmptyState / Card / Badge / Tooltip / Modal / Spinner
```

**Feature Components** (`packages/lib/src/features/verify/components/`):
```
Component: EmailVerifyRequestForm.tsx
Purpose: Send verification email (and switch to code mode)

Component: EmailCodeForm.tsx
Purpose: Enter & verify 6‑digit email code; resend on cooldown

Component: EmailMagicLinkVerifier.tsx
Purpose: Verify token (magic link) and route to success/help

Component: PhoneVerifyWizard.tsx
Purpose: Phone entry → SMS send → code verify → success

Component: VerificationCenter.tsx (settings)
Purpose: Status overview + actions for email/phone; policy hints
```

**Domain‑Specific Components:**
```
Component: VerifyGateBanner.tsx (web + mobile)
Purpose: In‑app banner when user is not verified; CTA to Verify

Component: OnboardingNudgeCard.tsx
Purpose: Gentle prompt linking ONB‑1 after verification
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/features/verify/hooks/`):
```
Hook: useEmailVerifyStart.ts
POST /v1/auth/verify/email/start
→ { sent: true, cooldownEndsAt: ts }

Hook: useEmailVerifyConfirm.ts
POST /v1/auth/verify/email/confirm { token }
→ { verified: true }

Hook: useEmailCodeStart.ts
POST /v1/auth/verify/email/code/start
→ { challengeId, resendAt }

Hook: useEmailCodeVerify.ts
POST /v1/auth/verify/email/code/verify { challengeId, code }
→ { verified: true }

Hook: usePhoneVerifyStart.ts
POST /v1/auth/verify/phone/start { phone }
→ { challengeId, resendAt }

Hook: usePhoneVerifyConfirm.ts
POST /v1/auth/verify/phone/confirm { challengeId, code }
→ { verified: true, phoneMasked }

Hook: useResendVerification.ts
POST /v1/auth/verify/resend { channel }  // email|sms
→ { sent: true, cooldownEndsAt }

Hook: useBounceStatusQuery.ts
GET /v1/auth/verify/email/bounce-status
→ { bounced: boolean, reason?: string, recommendedAction?: "change_email" }

Hook: useVerificationStatusQuery.ts
GET /v1/auth/verify/status
→ { email: "verified"|"pending", phone: "verified"|"pending" }
```

**State Management (Zustand)** (`packages/lib/src/stores/verification-store.ts`):
```
interface VerificationState {
  pendingEmail?: string;
  pendingPhone?: string;
  challengeId?: string|null;
  cooldownEndsAt?: number|null;
  postVerifyRedirect?: string|null;
  set: (p: Partial<VerificationState>) => void;
  clear: () => void;
}
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/verify/`):
```
emailSchema.ts  → string().email()
phoneSchema.ts  → E.164 validation
codeSchema.ts   → 6 digits numeric
```

**Formatting & Security Utils** (`packages/lib/src/utils/`):
```
format/maskEmail.ts, format/maskPhone.ts
time/cooldown.ts         → server‑synced countdown helpers
security/csrf.ts         → attach CSRF token on web
retry/backoff.ts         → exponential backoff (resend)
link/deeplink.ts         → deep‑link helpers for mobile
```

**Types** (`packages/lib/src/types/verify.ts`):
```
VerificationChannel = "email"|"sms"
VerifyStatus, VerifyStartResponse, VerifyConfirmResponse
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: AuthVerifyStack (Expo Router)
Paths:
  - /(public)/(auth)/verify-email-request
  - /(public)/(auth)/verify-email-sent
  - /(public)/(auth)/verify-email-confirm
  - /(public)/(auth)/verify-email-code
  - /(public)/(auth)/verify-phone
  - /(tabs)/(authenticated)/(settings)/account/verification
```

**Native Features:**
```
- expo-linking to handle magic‑link deep links securely
- haptics on verification success
- secure store for minimal state (no tokens)
```

---

#### F) Layout Components

**Web:**
```
Layout: (auth) Layout
Path: apps/web/app/(auth)/layout.tsx
Features:
  - Brand header/footer
  - ErrorBoundary + Sentry
  - <Toaster/> mount
```

**Mobile:**
```
Layout: (public)/(auth) Stack Layout
Path: apps/mobile/app/(public)/(auth)/_layout.tsx
Features:
  - Header/back
  - Modals for help/change
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Routes: (auth)/verify/*/error.tsx → friendly fallback + reset
Components: InlineError banners (token invalid, code wrong, rate‑limit)
```

**Loading States:**
```
Routes: (auth)/verify/*/loading.tsx → skeletons
Component: Spinner.tsx → inline busy states
```

---

### Visual Representations

#### Screen 1 — Verify Email (Request)

**Web View**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Verify your email                                                           │
├─────────────────────────────────────────────────────────────────────────────┤
│ We’ll send a link to:  a***@exa…   [ Change email ]                         │
│                                                                             │
│ [ Send verification email ]                                                 │
│                                                                             │
│ Prefer a code?  [ Use 6‑digit code instead ]                                │
│                                                                             │
│ {Captcha may appear after repeated requests}                                │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Mobile View**
```
┌──────────────────────────┐
│ ← Verify email           │
├──────────────────────────┤
│ a***@exa…  [Change]      │
│ ┏━━━━━━━━━━━━━━━━━━━━━┓  │
│ ┃ Send verification   ┃  │
│ ┗━━━━━━━━━━━━━━━━━━━━━┛  │
│ Use a 6‑digit code       │
└──────────────────────────┘
```

---

#### Screen 1a — “Check Your Email” (Sent + Resend)

```
┌──────────────────────────────────────────────┐
│ ✅ If an account exists for a***@exa…,       │
│    we’ve sent a verification link.           │
│ Didn’t get it?  [Resend 54s]  [Change email] │
└──────────────────────────────────────────────┘
```

---

#### Screen 2 — Magic Link Confirmation

```
┌──────────────────────────────────────────────┐
│ Verifying your email…                        │
├──────────────────────────────────────────────┤
│ ✅ Verified!                                  │
│ [ Continue ]                                  │
└──────────────────────────────────────────────┘
```

**Invalid/Expired**
```
┌──────────────────────────────────────────────┐
│ ❌ This link is invalid or expired.          │
│ [ Send a new link ]   [ Use a 6‑digit code ] │
└──────────────────────────────────────────────┘
```

---

#### Screen 3 — 6‑Digit Email Code

```
┌──────────────────────────────────────────────┐
│ Enter the 6‑digit code we sent to a***@exa…  │
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]         │
│ [ Verify ]     Didn’t get it? [Resend 28s]   │
└──────────────────────────────────────────────┘
```

---

#### Screen 4 — Verify Phone (SMS)

```
┌──────────────────────────────────────────────┐
│ Verify your phone                            │
├──────────────────────────────────────────────┤
│ Phone *  [+1] [ (___) ___‑____ ]             │
│ [ Send code ]                                 │
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]         │
│ [ Verify ]     [Resend 25s]                   │
└──────────────────────────────────────────────┘
```

---

#### Screen 5 — Verification Center (Settings)

```
┌──────────────────────────────────────────────┐
│ Verification Center                          │
├──────────────────────────────────────────────┤
│ Email    Verified ✓   a***@exa…   [ Resend ] [ Change ]          │
│ Phone    Pending      +1 ***‑***‑1234   [ Verify ] [ Change ]    │
│ Policy   Required for payouts before Nov 30                        │
└──────────────────────────────────────────────┘
```

---

#### Screen 6 — Verify Gate Banner (In‑App)

```
┌──────────────────────────────────────────────┐
│ 🔒 Verify your email to unlock all features. │
│ [ Verify now ]   [ Later ]                   │
└──────────────────────────────────────────────┘
```

---

## User Stories — Complete

> Abbrev: **W**=Web, **M**=Mobile. Each story is in “As a … I want … so that …” and includes components, routes, APIs, plus **≥5 Happy** & **≥10 Bad** scenarios.

### Story 1 — Verification Gate & Center (Status Overview)

**Story**
```
As a Newly Registered or Returning User
I want to see my verification status and required actions
So that I can complete verification and unlock the platform
```

**Components & Routes**
- **W:** `(settings)/(account)/verification/page.tsx` → `VerificationCenter`
- **M:** `/(tabs)/(authenticated)/(settings)/account/verification.tsx`
- **Gate:** `VerifyGateBanner` appears across restricted routes
- **APIs:** `GET /v1/auth/verify/status`, `GET /v1/auth/verify/email/bounce-status`

**Acceptance Criteria — ✅ Happy**
1. **AC1.1** Status loads < 1.0s (p95) showing email/phone states.
2. **AC1.2** If both verified → no banner; Success badge visible.
3. **AC1.3** If pending → clear CTAs to verify (email, phone).
4. **AC1.4** Gate banner appears only on restricted features.
5. **AC1.5** Policy notice rendered if verification is required before a date.

**Acceptance Criteria — ❌ Bad**
1. **AC1.6** Status fetch 5xx → soft error; retry; banner still shows minimal guidance.
2. **AC1.7** Unauthorized (401) → redirect to sign‑in, then back.
3. **AC1.8** Bounce detected → “Change email” CTA emphasized.
4. **AC1.9** Conflicting states across tabs → refetch on focus to reconcile.
5. **AC1.10** i18n missing keys → English fallback; key logged.
6. **AC1.11** A11y: banner not announced → role=alert + aria‑live.
7. **AC1.12** Policy API fails → hide policy row; log warning.
8. **AC1.13** Stale cache after verify → optimistic UI refresh on success events.
9. **AC1.14** Offline → cached status with offline badge and retry.
10. **AC1.15** Feature access incorrectly blocked after verify → clear gate on event bus.

---

### Story 2 — Request Email Verification (Magic Link Default)

**Story**
```
As a User with an email address
I want to request a verification link
So that I can confirm my email ownership
```

**Components & Routes**
- **W:** `(auth)/verify/email/request/page.tsx` → `EmailVerifyRequestForm`, `CaptchaWidget`, `MaskedEmail`
- **M:** `/(public)/(auth)/verify-email-request.tsx`
- **APIs:** `POST /v1/auth/verify/email/start`

**Acceptance Criteria — ✅ Happy**
1. **AC2.1** Clicking “Send verification email” → server sends link → navigate to “Sent” screen.
2. **AC2.2** Masked email displayed; copy uses `MaskedEmail` util.
3. **AC2.3** Resend disabled until cooldown ends; countdown accurate and accessible.
4. **AC2.4** Analytics `verify.email.start` fires exactly once per request.
5. **AC2.5** Captcha appears only after repeated requests and unblocks when passed.

**Acceptance Criteria — ❌ Bad**
1. **AC2.6** Network/5xx → banner; retry keeps masked email.
2. **AC2.7** Rate‑limit (429) → server‑hinted cooldown displayed.
3. **AC2.8** Captcha provider blocked → alternative challenge path offered.
4. **AC2.9** Unverified and bounced email → suggest “Change email” first.
5. **AC2.10** Double‑click submit → second request ignored (idempotency key).
6. **AC2.11** Low‑JS mode → fallback server render path supported.
7. **AC2.12** Localization mismatch → locale respected in content.
8. **AC2.13** Autofill shows wrong email → warn and offer change.
9. **AC2.14** Session expired → redirect to login, then back.
10. **AC2.15** Tracking blocked → analytics non‑blocking, no UX impact.

---

### Story 3 — Confirm Email via Magic Link

**Story**
```
As a User who received a verification email
I want to click the magic link and be verified
So that I can continue onboarding without friction
```

**Components & Routes**
- **W:** `(auth)/verify/email/confirm/page.tsx` → `EmailMagicLinkVerifier`
- **M:** `/(public)/(auth)/verify-email-confirm.tsx` (deep link)
- **APIs:** `POST /v1/auth/verify/email/confirm`

**Acceptance Criteria — ✅ Happy**
1. **AC3.1** Valid token verifies instantly (<500ms p95) and redirects to Success.
2. **AC3.2** Post‑verify event updates status in Verification Center and removes gate.
3. **AC3.3** “Continue” returns to last attempted route or ONB step.
4. **AC3.4** Analytics `verify.email.confirm.success` single‑fire.
5. **AC3.5** Token not persisted in URL after completion (POST‑redirect pattern).

**Acceptance Criteria — ❌ Bad**
1. **AC3.6** Token invalid/expired/used → friendly error with CTAs (resend / code).
2. **AC3.7** Cross‑origin link → blocked with security explanation.
3. **AC3.8** Multiple clicks → subsequent attempts show “Already verified” info.
4. **AC3.9** Deep link opens without app → offer “Continue in browser”.
5. **AC3.10** Network/5xx → retry; state preserved.
6. **AC3.11** Token leaked in referer → scrubbed via redirect; warning logged.
7. **AC3.12** Clock skew affecting TTL → relative time, server‑based.
8. **AC3.13** Accessibility: status announced via `role=status` live region.
9. **AC3.14** CSRF missing → hard refresh then retry.
10. **AC3.15** User on another device already verified → display success nonetheless.

---

### Story 4 — Verify Email via 6‑Digit Code (Alternative)

**Story**
```
As a User who can’t use the magic link
I want to enter a 6‑digit code from the email
So that I can verify my email anyway
```

**Components & Routes**
- **W:** `(auth)/verify/email/code/page.tsx` → `EmailCodeForm`, `CodeInput`
- **M:** `/(public)/(auth)/verify-email-code.tsx`
- **APIs:** `POST /v1/auth/verify/email/code/start`, `POST /v1/auth/verify/email/code/verify`

**Acceptance Criteria — ✅ Happy**
1. **AC4.1** Requesting a code starts challenge and sends email.
2. **AC4.2** Correct code verifies; user redirected to Success.
3. **AC4.3** Code input supports paste and auto‑advance; accessible.
4. **AC4.4** Resend honors cooldown; countdown localized.
5. **AC4.5** Analytics `verify.email.code.success` single‑fire.

**Acceptance Criteria — ❌ Bad**
1. **AC4.6** Wrong/expired code → inline error; N retries lock challenge.
2. **AC4.7** Spam filter delays → show tips and alternative magic link option.
3. **AC4.8** Pasted non‑digits → sanitized; error explained.
4. **AC4.9** ChallengeId missing → restart flow safely.
5. **AC4.10** Network/5xx → retry w/out losing input.
6. **AC4.11** Screen reader reads cells incorrectly → single‑field mode option.
7. **AC4.12** Tabs verify same challenge → second tab blocked gracefully.
8. **AC4.13** RTL numerals normalize to ASCII.
9. **AC4.14** Rate‑limit on resend → timer & tooltip show next attempt.
10. **AC4.15** Email bounce mid‑flow → prompt to change email.

---

### Story 5 — Verify Phone via SMS

**Story**
```
As a User with a phone number
I want to receive and enter a 6‑digit SMS code
So that I can verify my phone
```

**Components & Routes**
- **W:** `(auth)/verify/phone/page.tsx` → `PhoneVerifyWizard`, `PhoneInput`, `CodeInput`
- **M:** `/(public)/(auth)/verify-phone.tsx`
- **APIs:** `POST /v1/auth/verify/phone/start`, `POST /v1/auth/verify/phone/confirm`, `POST /v1/auth/verify/resend`

**Acceptance Criteria — ✅ Happy**
1. **AC5.1** E.164 phone validation; start returns challengeId.
2. **AC5.2** SMS arrives within SLA; code verifies → phone marked verified.
3. **AC5.3** Masked phone displayed in confirmations.
4. **AC5.4** Resend available after cooldown; accurate countdown.
5. **AC5.5** Analytics `verify.phone.confirm.success` logged (no PII).

**Acceptance Criteria — ❌ Bad**
1. **AC5.6** Invalid phone format → inline error.
2. **AC5.7** Wrong/expired code → error; N tries → restart challenge.
3. **AC5.8** Carrier delay/block → guidance and switch to email option.
4. **AC5.9** Roaming/number recycling → warning & confirmation.
5. **AC5.10** Dual‑SIM routing issues → allow change and resend.
6. **AC5.11** Rate‑limit (429) on start/resend → show cooldown.
7. **AC5.12** Network/5xx → retry with backoff.
8. **AC5.13** Accessibility: resend is keyboard/focus reachable.
9. **AC5.14** Local numerals → normalized.
10. **AC5.15** Session expiry → return to login; state restored post‑login.

---

### Story 6 — Resend with Cooldown & Captcha

**Story**
```
As a User waiting for a code or link
I want to resend after a short cooldown with anti‑abuse protections
So that I can complete verification without friction
```

**Components & Routes**
- **W/M:** “Sent” screens; `Countdown`, `CaptchaWidget`
- **APIs:** `POST /v1/auth/verify/resend` (channel), Anti‑bot verify endpoint

**Acceptance Criteria — ✅ Happy**
1. **AC6.1** Resend disabled until server cooldown elapsed; countdown exact.
2. **AC6.2** After cooldown, resend succeeds and restarts timer.
3. **AC6.3** Captcha only after repeated attempts; passing it unblocks resend.
4. **AC6.4** Analytics `verify.resend` fires once per action.
5. **AC6.5** Cooldown persists across refresh (server timestamp).

**Acceptance Criteria — ❌ Bad**
1. **AC6.6** Captcha provider blocked → alternate path.
2. **AC6.7** Offline → disable resend, show offline banner.
3. **AC6.8** Race conditions (multiple tabs) → server throttles; UI reconciles.
4. **AC6.9** Countdown negative due to skew → clamp to zero.
5. **AC6.10** Abusive IP → extended cooldown with clear message.
6. **AC6.11** Accessibility: countdown announced politely; not spammy.
7. **AC6.12** Localization of time units accurate.
8. **AC6.13** Resend target contact changed → restart to new contact.
9. **AC6.14** CSRF missing → hard refresh and retry.
10. **AC6.15** Analytics failure → non‑blocking; queued retry.

---

### Story 7 — Change Email/Phone Before Verification

**Story**
```
As a User who mistyped my contact
I want to change my email or phone before verifying
So that verification reaches the right destination
```

**Components & Routes**
- **W/M:** Links from request/sent/center to “Change email/phone” (can be handled in Settings profile forms)
- **APIs:** `POST /v1/account/change-email/start`, `POST /v1/account/change-email/confirm` (similar for phone)

**Acceptance Criteria — ✅ Happy**
1. **AC7.1** Start change sends confirm link/code to the **new** contact.
2. **AC7.2** Confirming change updates profile and resets verification status.
3. **AC7.3** Previous pending challenges invalidated.
4. **AC7.4** Masked new contact shown; flows continue seamlessly.
5. **AC7.5** Analytics `verify.contact.change.success` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC7.6** Attempt to change to same value → explain no change needed.
2. **AC7.7** Confirm link expired → CTA to resend.
3. **AC7.8** Concurrent changes → last confirmation wins; prior invalidated.
4. **AC7.9** Organization policy forbids unverified change → show policy reason.
5. **AC7.10** Phone change fails validation → inline error.
6. **AC7.11** Security: re‑auth required (password/MFA) to change contact.
7. **AC7.12** Emails bounce on new contact → suggest another domain.
8. **AC7.13** Rate limits on change start → cooldown messaging.
9. **AC7.14** Network/5xx → retries & safe resume.
10. **AC7.15** A11y: form fields labeled; errors announced.

---

### Story 8 — Pre‑Verification Restrictions

**Story**
```
As a Platform
I want to restrict sensitive actions until verification is complete
So that spam and fraud are reduced
```

**Components & Routes**
- **Gate:** `VerifyGateBanner` on attempts to post jobs, submit proposals, start contracts, withdraw funds, etc.
- **APIs:** N/A (policy enforced server‑side; client reflects via status)

**Acceptance Criteria — ✅ Happy**
1. **AC8.1** Attempting restricted action shows gate with clear messaging.
2. **AC8.2** After successful verify, action resumes without starting over.
3. **AC8.3** Analytics records blocked attempts to monitor friction.
4. **AC8.4** “Verify later” retains context; postVerifyRedirect used.
5. **AC8.5** Deep links to verify flows provided.

**Acceptance Criteria — ❌ Bad**
1. **AC8.6** Gate appears erroneously for verified users → refetch & clear cache.
2. **AC8.7** Server allows action despite pending → client still completes; log warning.
3. **AC8.8** Messaging too vague → copy guidelines enforced (linting).
4. **AC8.9** Excessive gating (innocuous actions) → policy tuned.
5. **AC8.10** A11y: banner focus trap fixed; close accessible.
6. **AC8.11** i18n omissions → fallback; keys logged.
7. **AC8.12** Multiple simultaneous gates → single consolidated banner.
8. **AC8.13** Offline state → defer until back online; show info banner.
9. **AC8.14** Analytics PII leakage → scrubbing guaranteed.
10. **AC8.15** Redirect loops → guard with seen‑token flag.

---

### Story 9 — Bounce Handling & Undeliverable Email

**Story**
```
As a Platform
I want to detect email bounces and guide users to fix their address
So that verification can succeed
```

**Components & Routes**
- **W/M:** Bounced banner on Verification Center and request pages
- **APIs:** `GET /v1/auth/verify/email/bounce-status`

**Acceptance Criteria — ✅ Happy**
1. **AC9.1** Hard bounce → show “Change email” primary CTA.
2. **AC9.2** Soft bounce → allow retry; suggest allowlist/safe sender.
3. **AC9.3** After change, bounce status clears and normal flow resumes.
4. **AC9.4** Analytics records bounce categories (non‑PII).

**Acceptance Criteria — ❌ Bad**
1. **AC9.5** Bounce API down → hide banner; log and allow retry.
2. **AC9.6** False positive → allow manual override with confirm.
3. **AC9.7** Disposable domain policy blocks → show policy reason.
4. **AC9.8** DMARC issue at sender → user gets clear timeline.
5. **AC9.9** Incorrect masking leaks info → fixed by util; test enforced.
6. **AC9.10** A11y: banner announced via live region.
7. **AC9.11** Localization issues → keys reviewed.
8. **AC9.12** Multiple bounce signals conflicting → most severe shown.
9. **AC9.13** Analytics includes domain name → allowed, but no full emails.
10. **AC9.14** Broken help link → inline guidance displayed.
11. **AC9.15** Admin‑forced verify despite bounce → show admin contact path.

---

### Story 10 — Cross‑Device Deep Link & Handoff

**Story**
```
As a Mobile‑first User
I want email links to open in the app when possible
So that verification feels seamless across devices
```

**Components & Routes**
- **M:** Linking handler in `verify-email-confirm.tsx` via expo‑linking
- **APIs:** None beyond the confirm endpoint

**Acceptance Criteria — ✅ Happy**
1. **AC10.1** If app installed, deep link opens app to confirm screen.
2. **AC10.2** If not installed, link falls back to mobile web confirm page.
3. **AC10.3** After verify, user returns to last intent (postVerifyRedirect).
4. **AC10.4** Analytics records app vs web completion.
5. **AC10.5** Security: link domain & scheme validated.

**Acceptance Criteria — ❌ Bad**
1. **AC10.6** Link hijacked by other app → fallback to web with warning.
2. **AC10.7** iOS Universal Links misconfigured → graceful fallback.
3. **AC10.8** Android intent chooser confusion → help text shown.
4. **AC10.9** Invalid path/query → redirect to request with banner.
5. **AC10.10** Token lingering in URL → removed after confirm.
6. **AC10.11** No network → prompt to retry later; keep context.
7. **AC10.12** User denies open‑app prompt → continue in browser.
8. **AC10.13** Multiple clicks on link → no duplicate actions.
9. **AC10.14** Old app version lacks route → open in web.
10. **AC10.15** Analytics blocked → non‑blocking.

---

### Story 11 — Accessibility & Localization

**Story**
```
As an Inclusive Platform
I want all verification flows to be fully accessible and localized
So that all users can complete verification
```

**Acceptance Criteria — ✅ Happy**
1. **AC11.1** All forms have labels, described‑by, and error ARIA.
2. **AC11.2** CodeInput supports screen readers and single‑field mode.
3. **AC11.3** Countdown announced politely without spam.
4. **AC11.4** RTL layouts and localized digits supported.
5. **AC11.5** Contrast and focus states meet WCAG AA.

**Acceptance Criteria — ❌ Bad**
1. **AC11.6** Missing keys → English fallback + log.
2. **AC11.7** Motion respects prefers‑reduced‑motion.
3. **AC11.8** Keyboard traps in modals → fixed and tested.
4. **AC11.9** Time/number formatting not localized → i18n formatters used.
5. **AC11.10** Error messages not announced → live regions added.
6. **AC11.11** Touch targets too small → 44px minimum enforced.
7. **AC11.12** Screen zoom 200% remains usable.
8. **AC11.13** Focus order skips critical actions → corrected.
9. **AC11.14** Color‑only indicators → add text/icons.
10. **AC11.15** A11y CI checks (axe) must pass.

---

### Story 12 — Analytics & Telemetry

**Story**
```
As a Product Team
I want to instrument the verification funnel end‑to‑end
So that we can measure completion and reliability
```

**Events (examples)**
- `verify.email.start` / `.sent` / `.confirm.success` / `.confirm.error`
- `verify.email.code.start` / `.verify.success` / `.verify.error`
- `verify.phone.start` / `.confirm.success` / `.confirm.error`
- `verify.resend` / `verify.cooldown.start` / `.end`
- `verify.bounce.detected`
- `verify.gate.blocked_action`

**Acceptance Criteria — ✅ Happy**
1. **AC12.1** Events debounced/idempotent; no duplicates on rerender.
2. **AC12.2** No PII (emails, full phone) in payloads.
3. **AC12.3** Client queues and retries analytics on failures.
4. **AC12.4** Versioned schemas validated in CI.
5. **AC12.5** Funnels visible per channel (link vs code vs sms).

**Acceptance Criteria — ❌ Bad**
1. **AC12.6** Payload includes tokens/codes → blocked by linter.
2. **AC12.7** Missing success events on deep link → fixed via app‑open hook.
3. **AC12.8** Time skew between client/server → server timestamp included.
4. **AC12.9** Ad‑blockers prevent send → non‑blocking; drop safe.
5. **AC12.10** Event names change without version bump → build fails.

---

### Story 13 — Security Hardening

**Story**
```
As a Security Team
I want robust verification tokens and anti‑abuse controls
So that ownership is proven without exposing users
```

**Acceptance Criteria — ✅ Happy**
1. **AC13.1** Email tokens single‑use, short TTL (≤30m), high entropy.
2. **AC13.2** Code challenges lock after N attempts; cooldown enforced.
3. **AC13.3** CSRF on POST endpoints; Referrer‑Policy/CSP prevent leaks.
4. **AC13.4** Post‑verify rotates session to prevent fixation.
5. **AC13.5** Audit logs redact PII; store masked contact only.

**Acceptance Criteria — ❌ Bad**
1. **AC13.6** Token in URL logged by third‑party → redirect POST pattern.
2. **AC13.7** Brute‑force code attempts unthrottled → tests fail.
3. **AC13.8** Captcha never triggers → threshold misconfig detected.
4. **AC13.9** Bypass via API (no policy) → server middleware enforces.
5. **AC13.10** Contact changed mid‑flow → invalidate prior challenges.
6. **AC13.11** Recovery path leaks account existence → non‑enumerating copy.
7. **AC13.12** Multiple tab races → server idempotency required.
8. **AC13.13** Tokens persisted client‑side → store only ephemeral state.
9. **AC13.14** Logs include raw emails → scrub middleware enforced.
10. **AC13.15** SMS code content contains PII → remove and re‑issue.
 
---

## Non‑Functional Requirements (AUTH‑5 scope)

- **Performance:** Status fetch < 1.0s (p95); confirm verify < 500ms (p95)
- **Security:** Single‑use tokens; CSRF; captcha & rate limits; POST‑redirect; audit logging; Sentry
- **Privacy:** Masked contact in UI/logs; no PII in analytics
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+
- **Testing:** Unit tests (schemas/hooks/utils), e2e (link + code + sms), a11y (axe), anti‑enumeration timing

---

**End of AUTH‑5 Enhancements (Complete).**
