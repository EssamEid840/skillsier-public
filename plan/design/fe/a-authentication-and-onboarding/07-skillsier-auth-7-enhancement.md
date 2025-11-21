## AUTH-7 — Password Reset & Account Recovery (Forgot Password, Magic Link / 6‑Digit Code, Forced Reset, Change Password, Bot Protection)

> Enhancement add‑on for the existing AUTH‑7 journey in `skillsier-frontend-journeys-claude-final.md`.
> This deliverable fully completes the **Enhancement Prompt Template** sections: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-structure.md` conventions. Scope covers: *Forgot Password*, *Reset via email link or code*, *Phone SMS fallback*, *Forced password reset interstitials*, *Change Password in Settings*, *bot/abuse protections*, *deep‑link handoffs*, *analytics*, *a11y*, and *security hardening*.

---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(auth)/password/forgot/page.tsx
Purpose: Request password reset via email (default) or choose SMS code if phone on file.
Key Features:
  - Email or username input (non‑enumerating copy)
  - Prefer magic link; switch to 6‑digit code
  - Resend cooldown & captcha after N attempts

Path: apps/web/app/(auth)/password/sent/page.tsx
Purpose: “Check your email/SMS” confirmation with resend countdown and troubleshooting.

Path: apps/web/app/(auth)/password/confirm/page.tsx
Purpose: Magic‑link landing; validate token; forward to Reset form on success; help on error.
Props: token (search param)

Path: apps/web/app/(auth)/password/code/page.tsx
Purpose: Enter 6‑digit code (email or SMS) → proceed to Reset form.

Path: apps/web/app/(auth)/password/reset/page.tsx
Purpose: Set new password with strength meter and policy checklist.
Key Features:
  - New password + confirm
  - Strength meter + hints
  - Show/Hide toggles; paste allowed
  - Password visibility a11y labels

Path: apps/web/app/(settings)/(security)/password/change/page.tsx
Purpose: In‑product change password (requires re‑auth per AUTH‑6).

Path: apps/web/app/(auth)/password/forced-reset/page.tsx
Purpose: Interstitial forcing password reset due to compromise/policy.

Path: apps/web/app/(auth)/layout.tsx
Purpose: Shared auth frame; Toaster; ErrorBoundary + Sentry.

Path: apps/web/app/(auth)/password/*/loading.tsx
Path: apps/web/app/(auth)/password/*/error.tsx
Purpose: Route‑level skeletons and friendly fallbacks.
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(public)/(auth)/forgot-password.tsx
Purpose: Request reset (email default, code alternative).

Path: apps/mobile/app/(public)/(auth)/reset-sent.tsx
Purpose: Confirmation with resend countdown.

Path: apps/mobile/app/(public)/(auth)/reset-confirm.tsx
Purpose: Handle magic‑link deep link and forward to reset form.

Path: apps/mobile/app/(public)/(auth)/reset-code.tsx
Purpose: 6‑digit code entry (email/SMS).

Path: apps/mobile/app/(public)/(auth)/reset.tsx
Purpose: New password form with strength meter & checklist.

Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/change-password.tsx
Purpose: Change password in settings.
```

> Handoff: AUTH‑7 integrates with **AUTH‑6** (re‑auth for change‑password), **AUTH‑4** (if MFA required to confirm reset), and **AUTH‑5** (verified contact preferred).

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: Input, PasswordInput, Button, LinkButton, Checkbox, Tooltip, Alert/Banner, Modal, Spinner
Component: CodeInput (6‑digit)
Path: packages/ui/src/components/forms/CodeInput.tsx

Component: CaptchaWidget
Purpose: Progressive anti‑bot on repeat attempts

Component: Countdown
Purpose: Accessible resend countdown

Component: StrengthMeter
Path: packages/ui/src/components/security/StrengthMeter.tsx
Purpose: Visual strength bar + textual hints

Component: PolicyChecklist
Path: packages/ui/src/components/security/PolicyChecklist.tsx
Purpose: Live checklist (min length, classes, no reuse, no leaks)

Component: MaskedEmail / MaskedPhone
Purpose: Display contact safely (a***@exa…, +1 ***‑***‑1234)
```

**Feature Components** (`packages/lib/src/features/password/components/`):
```
Component: ForgotPasswordForm.tsx
Purpose: Request reset via email (default) or code alternative

Component: ResetCodeForm.tsx
Purpose: Enter 6‑digit code; resend with cooldown

Component: ResetMagicLinkVerifier.tsx
Purpose: Validate token then route to Reset form

Component: ResetPasswordForm.tsx
Purpose: New password + confirm with strength + checklist

Component: ChangePasswordForm.tsx (settings)
Purpose: Old password + new password (AUTH‑6 re‑auth may short‑circuit old password)
```

**Domain‑Specific Components:**
```
Component: ForcedResetInterstitial.tsx
Purpose: Block access until password is changed
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/features/password/hooks/`):
```
Hook: usePasswordResetStart.ts
POST /v1/auth/password/reset/start { emailOrUsername, channel? }  // channel: email|sms
→ { sent:true, channel, cooldownEndsAt }

Hook: usePasswordResetConfirm.ts
POST /v1/auth/password/reset/confirm { token }
→ { valid:true, resetToken }  // short‑lived internal reset token

Hook: usePasswordCodeStart.ts
POST /v1/auth/password/code/start { emailOrPhone }
→ { challengeId, resendAt }

Hook: usePasswordCodeVerify.ts
POST /v1/auth/password/code/verify { challengeId, code }
→ { valid:true, resetToken }

Hook: usePasswordResetComplete.ts
POST /v1/auth/password/reset/complete { resetToken, newPassword }
→ { success:true }

Hook: useChangePassword.ts
POST /v1/auth/password/change { currentPassword?, newPassword }
→ { success:true }  // currentPassword optional if recent re‑auth

Hook: useForcedResetStatus.ts
GET /v1/auth/password/forced-status
→ { required:boolean, reason?:string }

Hook: useResendReset.ts
POST /v1/auth/password/reset/resend { channel }
→ { sent:true, cooldownEndsAt }
```

**State Management (Zustand)** (`packages/lib/src/stores/password-store.ts`):
```
interface PasswordState {
  emailMasked?: string;
  phoneMasked?: string;
  challengeId?: string|null;
  resetToken?: string|null;
  cooldownEndsAt?: number|null;
  postResetRedirect?: string|null;
  set: (p: Partial<PasswordState>) => void;
  clear: () => void;
}
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/password/`):
```
forgotSchema.ts      → emailOrUsername: string().min(3)
codeSchema.ts        → code: regex(/^\d{6}$/)
newPasswordSchema.ts → min/max length, character classes, deny lists, identical confirm
```

**Security & Formatting Utils** (`packages/lib/src/utils/`):
```
security/passwordStrength.ts  → returns score + hints
security/pwnedCheck.ts        → k‑anon HIBP client (if enabled)
time/cooldown.ts              → server‑synced cooldown helpers
format/maskEmail.ts, maskPhone.ts
link/deeplink.ts              → mobile deep links for magic link
```

**Types** (`packages/lib/src/types/password.ts`):
```
ResetChannel, ResetStartResponse, ResetConfirmResponse, ResetCompleteResponse
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: AuthResetStack (Expo Router)
Paths:
  - /(public)/(auth)/forgot-password
  - /(public)/(auth)/reset-sent
  - /(public)/(auth)/reset-confirm
  - /(public)/(auth)/reset-code
  - /(public)/(auth)/reset
  - /(tabs)/(authenticated)/(settings)/security/change-password
```

**Native Features:**
```
- expo-linking to capture magic links securely
- Haptics on success
- SecureStore for ephemeral reset token (not persisted across restarts)
```

---

#### F) Layout Components

**Web:**
```
Layout: (auth) Layout
Path: apps/web/app/(auth)/layout.tsx
Features:
  - Minimal chrome, brand header/footer
  - ErrorBoundary + Sentry
  - <Toaster/> mount
```

**Mobile:**
```
Layout: (public)/(auth)/_layout.tsx
Features:
  - Header/back
  - Modal for troubleshooting/help
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Routes: (auth)/password/*/error.tsx → friendly fallback + reset
Components: Inline banners (invalid token, expired code, rate‑limit)
```

**Loading States:**
```
Routes: (auth)/password/*/loading.tsx → skeletons
Component: Spinner.tsx → inline busy states
```

---

### Visual Representations

#### Screen 1 — Forgot Password (Request)

**Web View**
```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Forgot your password?                                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│ Enter your email or username and we’ll send a reset link.                   │
│                                                                             │
│ Email or username *                                                         │
│ ┌──────────────────────────────────────────────┐                            │
│ │ a***@exa…                                   │                            │
│ └──────────────────────────────────────────────┘                            │
│ [ Send reset link ]                                                           │
│                                                                             │
│ Prefer a code? [ Use 6‑digit code ] • [ Try SMS ]                            │
│ {Captcha may appear after repeated requests}                                 │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Mobile View**
```
┌──────────────────────────┐
│ ← Forgot password        │
├──────────────────────────┤
│ Email or username        │
│ ┌────────────────────┐   │
│ │a***@exa…           │   │
│ └────────────────────┘   │
│ ┏━━━━━━━━━━━━━━━━━━━┓    │
│ ┃ Send reset link   ┃    │
│ ┗━━━━━━━━━━━━━━━━━━━┛    │
│ Use a 6‑digit code       │
└──────────────────────────┘
```

---

#### Screen 1a — “Check Your Email/SMS” (Sent + Resend)

```
┌──────────────────────────────────────────────┐
│ ✅ If an account exists, we’ve sent a link.  │
│ Didn’t get it?  [Resend 54s]  [Troubleshoot] │
└──────────────────────────────────────────────┘
```

---

#### Screen 2 — Magic Link Confirmation

```
┌──────────────────────────────────────────────┐
│ Verifying your reset link…                   │
├──────────────────────────────────────────────┤
│ ✅ Link valid                                 │
│ [ Continue to reset ]                         │
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

#### Screen 3 — 6‑Digit Code (Email/SMS)

```
┌──────────────────────────────────────────────┐
│ Enter the 6‑digit code                       │
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]         │
│ [ Verify ]     Didn’t get it? [Resend 28s]   │
└──────────────────────────────────────────────┘
```

---

#### Screen 4 — Reset Password Form

```
┌──────────────────────────────────────────────┐
│ Create a new password                        │
├──────────────────────────────────────────────┤
│ New password *      [ •••••••• ]  👁         │
│ Strength:  ████░  Good                      │
│ Checklist: ✓ 8+ chars  ✓ number  ✗ symbol    │
│ Confirm password *  [ •••••••• ]  👁         │
│ [ Save new password ]                        │
└──────────────────────────────────────────────┘
```

---

#### Screen 5 — Success

```
┌──────────────────────────────────────────────┐
│ ✅ Password updated                           │
│ You can now sign in.                          │
│ [ Continue ]                                  │
└──────────────────────────────────────────────┘
```

---

#### Screen 6 — Change Password (Settings)

```
┌──────────────────────────────────────────────┐
│ Change password                              │
├──────────────────────────────────────────────┤
│ Current password [ •••••••• ]  👁             │
│ New password     [ •••••••• ]  👁   Strength  │
│ Confirm new      [ •••••••• ]  👁             │
│ [ Save ]                                       │
└──────────────────────────────────────────────┘
```

---

#### Screen 7 — Forced Reset Interstitial

```
┌──────────────────────────────────────────────┐
│ For your security, please reset your password│
├──────────────────────────────────────────────┤
│ Reason: unusual activity detected             │
│ [ Continue to reset ]                         │
└──────────────────────────────────────────────┘
```

---

## User Stories — Complete

> Abbrev: **W**=Web, **M**=Mobile. Each story follows “As a … I want … so that …” and includes components, routes, APIs, with **≥5 Happy** & **≥10 Bad** scenarios.

### Story 1 — Request Password Reset (Email Default)

**Story**
```
As a User who forgot my password
I want to request a reset via my email
So that I can securely regain access
```

**Components & Routes**
- **W:** `(auth)/password/forgot/page.tsx` → `ForgotPasswordForm`, `CaptchaWidget`, `MaskedEmail`
- **M:** `/(public)/(auth)/forgot-password.tsx`
- **APIs:** `POST /v1/auth/password/reset/start`

**Acceptance Criteria — ✅ Happy**
1. **AC1.1** Submitting a valid email starts reset and routes to “Sent” screen.
2. **AC1.2** Copy is non‑enumerating (“If an account exists…”).
3. **AC1.3** Resend disabled until cooldown; countdown is accurate and accessible.
4. **AC1.4** Analytics `password.reset.start` fires once per request.
5. **AC1.5** Troubleshooting opens modal with actionable tips.

**Acceptance Criteria — ❌ Bad**
1. **AC1.6** Network/5xx → banner; retry retains email value.
2. **AC1.7** Rate‑limit (429) → show cooldown (server‑hinted).
3. **AC1.8** Captcha required after N attempts → must complete to continue.
4. **AC1.9** Unverified email → show resend verify path (AUTH‑5).
5. **AC1.10** Typo address → allow change before re‑request.
6. **AC1.11** Session present → sign out not required; continue flow cleanly.
7. **AC1.12** Localization keys missing → English fallback + log.
8. **AC1.13** Autofill with another account → clear instructions.
9. **AC1.14** Offline → disable submit; explain and allow retry.
10. **AC1.15** Analytics blocked → non‑blocking; queued retry.

---

### Story 2 — Confirm via Magic Link (Email)

**Story**
```
As a User who received a reset email
I want to click the magic link and proceed
So that I can set a new password
```

**Components & Routes**
- **W:** `(auth)/password/confirm/page.tsx` → `ResetMagicLinkVerifier`
- **M:** `/(public)/(auth)/reset-confirm.tsx` (deep link)
- **APIs:** `POST /v1/auth/password/reset/confirm`

**Acceptance Criteria — ✅ Happy**
1. **AC2.1** Valid token returns a short‑lived `resetToken`; route to Reset form.
2. **AC2.2** Post‑redirect pattern removes token from URL.
3. **AC2.3** Analytics `password.reset.confirm.success` logged once.
4. **AC2.4** ResetToken stored only in memory (not localStorage).
5. **AC2.5** If already used, show success guidance instead of error.

**Acceptance Criteria — ❌ Bad**
1. **AC2.6** Invalid/expired token → friendly error with CTAs to resend or use code.
2. **AC2.7** Cross‑origin link → blocked; guidance to official domain.
3. **AC2.8** Multiple clicks → subsequent opens show “Already handled” info.
4. **AC2.9** Network/5xx → retry; maintain safe state.
5. **AC2.10** Token leaked in referer → scrub via redirect; log warning.
6. **AC2.11** Clock skew → rely on server validation; copy explains.
7. **AC2.12** CSRF missing → full refresh then retry.
8. **AC2.13** Deep link opens without app → “Continue in browser”.
9. **AC2.14** i18n issues → fallback strings.
10. **AC2.15** Disabled JS → server fallback page explains next steps.

---

### Story 3 — Confirm via 6‑Digit Code (Email/SMS)

**Story**
```
As a User who can’t use the link
I want to enter a 6‑digit code
So that I can proceed to reset my password
```

**Components & Routes**
- **W:** `(auth)/password/code/page.tsx` → `ResetCodeForm`, `CodeInput`
- **M:** `/(public)/(auth)/reset-code.tsx`
- **APIs:** `POST /v1/auth/password/code/start`, `POST /v1/auth/password/code/verify`

**Acceptance Criteria — ✅ Happy**
1. **AC3.1** Starting a code challenge sends the code and returns `challengeId`.
2. **AC3.2** Correct code yields `resetToken`; routes to Reset form.
3. **AC3.3** Resend respects cooldown; countdown accessible.
4. **AC3.4** Paste + auto‑advance supported; RTL digits normalized.
5. **AC3.5** Analytics `password.code.verify.success` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC3.6** Wrong/expired code → inline error; after N → lockout cooldown.
2. **AC3.7** ChallengeId missing → restart flow safely.
3. **AC3.8** Delivery delay → tips + alternative magic link option.
4. **AC3.9** Rate‑limit on resend → timer shows next attempt.
5. **AC3.10** Network/5xx → retry without clearing input.
6. **AC3.11** Screen reader issues → offer single‑field mode.
7. **AC3.12** Phone number recycled → warn and suggest verify phone first.
8. **AC3.13** Pasted non‑digits → sanitize; show helper text.
9. **AC3.14** Concurrent verify in tabs → second tab blocked gracefully.
10. **AC3.15** i18n keys missing → fallback strings.

---

### Story 4 — Reset Password (New + Confirm with Strength & Policy)

**Story**
```
As a User with a valid reset token
I want to set a strong new password that meets policy
So that my account stays secure
```

**Components & Routes**
- **W:** `(auth)/password/reset/page.tsx` → `ResetPasswordForm`, `StrengthMeter`, `PolicyChecklist`
- **M:** `/(public)/(auth)/reset.tsx`
- **APIs:** `POST /v1/auth/password/reset/complete`

**Acceptance Criteria — ✅ Happy**
1. **AC4.1** Strength meter updates live; hints rendered.
2. **AC4.2** Checklist shows real‑time pass/fail for policy requirements.
3. **AC4.3** Confirm match required; show/hide toggles accessible.
4. **AC4.4** On success, sessions are rotated; user brought to Success.
5. **AC4.5** Analytics `password.reset.complete.success` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC4.6** Weak password blocked; hints guide to pass policy.
2. **AC4.7** Confirm mismatch → clear message and field focus.
3. **AC4.8** Known leaked password (pwned) → disallowed with link to guidance.
4. **AC4.9** Reset token expired/used → redirect to start with message.
5. **AC4.10** Network/5xx → retry; input preserved.
6. **AC4.11** A11y: labels, aria‑describedby, error announcements.
7. **AC4.12** i18n keys missing → fallback strings.
8. **AC4.13** Paste blocked by browser policy → allow or explain.
9. **AC4.14** Mobile keyboard overlap → auto‑scroll into view.
10. **AC4.15** Password managers conflict → ensure correct autocomplete attrs.

---

### Story 5 — Change Password (Settings, Post‑Login)

**Story**
```
As a Signed‑in User
I want to change my password from settings
So that I can improve my account security
```

**Components & Routes**
- **W:** `(settings)/(security)/password/change/page.tsx` → `ChangePasswordForm`
- **M:** `…/settings/security/change-password.tsx`
- **APIs:** `POST /v1/auth/password/change` (requires AUTH‑6 re‑auth or old password)

**Acceptance Criteria — ✅ Happy**
1. **AC5.1** If recent re‑auth exists, old password field can be omitted.
2. **AC5.2** Strength meter & checklist apply; success rotates sessions (optional: keep current).
3. **AC5.3** Success toast and email notification sent.
4. **AC5.4** Analytics `password.change.success` logged.
5. **AC5.5** A11y and i18n conforming.

**Acceptance Criteria — ❌ Bad**
1. **AC5.6** Wrong current password → inline error; rate‑limit attempts.
2. **AC5.7** Weak password → blocked with guidance.
3. **AC5.8** Network/5xx → retry without clearing inputs.
4. **AC5.9** Session expired → redirect to re‑auth then resume.
5. **AC5.10** CSRF missing → hard refresh and retry.
6. **AC5.11** Password reuse against recent history → blocked.
7. **AC5.12** A11y issues → errors announced; focus handled.
8. **AC5.13** i18n fallbacks present and logged.
9. **AC5.14** Manager autofill mismatch → helper text clarifies.
10. **AC5.15** Concurrent change on another device → last write wins; user notified.

---

### Story 6 — Forced Password Reset Interstitial

**Story**
```
As a Platform
I want to force a password reset after suspected compromise
So that users are protected
```

**Components & Routes**
- **W/M:** `(auth)/password/forced-reset/page.tsx` → `ForcedResetInterstitial`
- **APIs:** `GET /v1/auth/password/forced-status`

**Acceptance Criteria — ✅ Happy**
1. **AC6.1** Interstitial blocks access until reset complete.
2. **AC6.2** Clear reason and guidance; CTA to Reset form.
3. **AC6.3** After reset, routes back to intended destination.
4. **AC6.4** Email security alert sent.
5. **AC6.5** Analytics `password.forced_reset.comply` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC6.6** Status endpoint down → cached state used; soft warning.
2. **AC6.7** User tries to bypass via URL → redirected back.
3. **AC6.8** Multiple tabs inconsistent → refetch on focus.
4. **AC6.9** A11y: interstitial focus management correct.
5. **AC6.10** i18n missing → fallback.
6. **AC6.11** Offline → blocked with clear copy; retry on reconnect.
7. **AC6.12** Analytics blocked → non‑blocking.
8. **AC6.13** Account locked too → show support path.
9. **AC6.14** Too frequent forced resets → policy cool‑down respected.
10. **AC6.15** Link loops after completion → mark as completed flag.

---

### Story 7 — SMS Fallback (If Phone on File)

**Story**
```
As a User without email access
I want to receive a 6‑digit SMS code
So that I can reset my password
```

**Components & Routes**
- **W/M:** Paths from Story 3; phone option visible only if phone exists & verified
- **APIs:** `POST /v1/auth/password/code/start` with channel=sms

**Acceptance Criteria — ✅ Happy**
1. **AC7.1** Phone is masked; option shown when verified phone exists.
2. **AC7.2** Code verifies and yields resetToken.
3. **AC7.3** Resend respects cooldown; accessible countdown.
4. **AC7.4** Analytics `password.code.sms.success` logged.
5. **AC7.5** A11y and i18n compliant.

**Acceptance Criteria — ❌ Bad**
1. **AC7.6** No phone on file → option hidden.
2. **AC7.7** Phone unverified → prompt to verify (AUTH‑5).
3. **AC7.8** Delivery issues → guidance and alternative email link.
4. **AC7.9** Dual‑SIM confusion → allow change before resend.
5. **AC7.10** Rate‑limit on resend → feedback and timer.
6. **AC7.11** Network/5xx → retry w/out losing input.
7. **AC7.12** Screen reader issues → single‑field mode available.
8. **AC7.13** Local numerals normalization.
9. **AC7.14** Carrier short code blocked → use long code fallback.
10. **AC7.15** Security: enumerate phone existence → copy avoids disclosure.

---

### Story 8 — Bot & Abuse Protections

**Story**
```
As a Security Team
I want captchas, rate limits, and non‑enumerating flows
So that attackers can’t exploit password reset
```

**Acceptance Criteria — ✅ Happy**
1. **AC8.1** Captcha appears progressively after repeated attempts.
2. **AC8.2** Response copy never confirms account existence.
3. **AC8.3** IP/device rate limits enforced; cooldown UI accurate.
4. **AC8.4** Audit log records attempts (no PII in payloads).
5. **AC8.5** Analytics trends monitored without sensitive data.

**Acceptance Criteria — ❌ Bad**
1. **AC8.6** Captcha provider blocked → alternative challenge flow.
2. **AC8.7** Rate‑limit messages leak policy → generic copy used.
3. **AC8.8** Enumeration via timing → constant‑time responses used.
4. **AC8.9** Over‑aggressive limits block legit users → appeal path.
5. **AC8.10** i18n missing → fallback strings.
6. **AC8.11** Accessibility of captcha ensured; audio alternative.
7. **AC8.12** Analytics includes emails → blocked by linter.
8. **AC8.13** CDN cache leaks responses → no‑store headers enforced.
9. **AC8.14** IP reputation false positives → allowlist for support.
10. **AC8.15** Session fixation post‑reset → session rotated on complete.

---

### Story 9 — Cross‑Device Deep Link & Handoff

**Story**
```
As a Mobile‑first User
I want reset links to open in the app when possible
So that I can complete reset seamlessly
```

**Components & Routes**
- **M:** `reset-confirm.tsx` handles deep links
- **APIs:** none beyond confirm endpoint

**Acceptance Criteria — ✅ Happy**
1. **AC9.1** Installed app opens to confirm; otherwise mobile web fallback.
2. **AC9.2** After confirm, route directly to Reset form in app.
3. **AC9.3** Analytics logs app vs web completion.
4. **AC9.4** Domain/scheme validated; no open‑redirects.
5. **AC9.5** Token removed from URL after processing.

**Acceptance Criteria — ❌ Bad**
1. **AC9.6** Universal Links misconfigured → graceful fallback.
2. **AC9.7** Android intent chooser confusion → help text.
3. **AC9.8** Invalid path/query → redirect to request with banner.
4. **AC9.9** No network → retry later; context preserved.
5. **AC9.10** User declines open‑app → continue in browser.
6. **AC9.11** Multiple clicks → no duplicate actions.
7. **AC9.12** Old app version lacks route → open in web.
8. **AC9.13** Analytics blocked → non‑blocking.
9. **AC9.14** Token leaks in logs → scrubbed via redirect.
10. **AC9.15** Deep‑link hijack by other app → security warning & web fallback.

---

### Story 10 — Analytics & Telemetry for Reset Funnel

**Story**
```
As a Product Team
I want to instrument the reset funnel end‑to‑end
So that we can measure drop‑offs and reliability
```

**Events (examples)**
- `password.reset.start` / `.sent` / `.confirm.success` / `.confirm.error`
- `password.code.start` / `.verify.success` / `.verify.error`
- `password.reset.complete.success` / `.error`
- `password.change.success`
- `password.forced_reset.comply`

**Acceptance Criteria — ✅ Happy**
1. **AC10.1** Events debounced/idempotent; no duplicates on rerender.
2. **AC10.2** No PII (emails/phones/codes) in payloads.
3. **AC10.3** Client queues and retries analytics on failures.
4. **AC10.4** Versioned schemas validated in CI.
5. **AC10.5** Dashboard shows link vs code vs SMS conversion.

**Acceptance Criteria — ❌ Bad**
1. **AC10.6** Payload includes tokens/codes → build fails; linter catches.
2. **AC10.7** Missing success on deep link → hook ensures send on open.
3. **AC10.8** Time skew → server timestamp included.
4. **AC10.9** Ad‑blockers drop beacons → non‑blocking; drop safe.
5. **AC10.10** Event names change without version → build fails.
6. **AC10.11** Over‑instrumentation → perf budget enforced.
7. **AC10.12** Analytics library crashes → sandbox & try/catch.
8. **AC10.13** Duplicates from multiple tabs → storage‑based dedupe.
9. **AC10.14** PII logged via console → disabled in prod.
10. **AC10.15** GDPR: unconsented tracking → respect consent gates.
 
---

### Story 11 — Accessibility & Localization

**Story**
```
As an Inclusive Platform
I want the reset flows to be fully accessible and localized
So that all users can recover their accounts
```

**Acceptance Criteria — ✅ Happy**
1. **AC11.1** Forms have labels, described‑by, and error ARIA.
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

### Story 12 — Security Hardening for Reset

**Story**
```
As a Security Team
I want robust tokens, session rotation, and enumeration‑safe UX
So that account recovery is safe
```

**Acceptance Criteria — ✅ Happy**
1. **AC12.1** Tokens are single‑use, short TTL (≤30m), high entropy.
2. **AC12.2** Session rotated on completion; all other sessions revoked (configurable).
3. **AC12.3** Non‑enumerating copy and constant‑time responses.
4. **AC12.4** Audit logs redact PII; only masked contact stored.
5. **AC12.5** Optional HIBP check blocks known leaked passwords.

**Acceptance Criteria — ❌ Bad**
1. **AC12.6** Tokens persisted client‑side → forbidden (memory only).
2. **AC12.7** Reset endpoint allows token reuse → tests fail.
3. **AC12.8** Logs include full emails/phones → scrub middleware enforced.
4. **AC12.9** Reset allows weak passwords → policy enforced in API and UI.
5. **AC12.10** Redirect to arbitrary URL → allowlist original intents only.
6. **AC12.11** Email links reveal account in subject → neutral templates.
7. **AC12.12** SMS code contains PII → removed.
8. **AC12.13** Multiple tab races → server idempotency required.
9. **AC12.14** Mixed content (http) → CSP/HTTPS enforced.
10. **AC12.15** Token in referer logs → POST‑redirect pattern required.

---

## Non‑Functional Requirements (AUTH‑7 scope)

- **Performance:** Start/confirm < 600ms p95; reset complete < 700ms p95
- **Security:** Single‑use tokens; CSRF; captcha & rate limits; POST‑redirect; session rotation; audit logging; Sentry
- **Privacy:** Masked contact; no PII in analytics; neutral copy
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+
- **Testing:** Unit tests (schemas/hooks/utils), e2e (link + code + sms), a11y (axe), anti‑enumeration timing

---

**End of AUTH‑7 Enhancements (Complete).**
