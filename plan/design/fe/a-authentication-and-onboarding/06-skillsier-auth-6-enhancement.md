## AUTH-6 — Step‑Up Authentication, Session Timeout & Account Lock (Re‑Auth Modal, Idle Lock, Suspicious Activity, Keep‑Alive)

> Enhancement add‑on for the existing AUTH‑6 journey in `skillsier-frontend-journeys-claude-final.md`.  
> This deliverable fully completes the three sections per the **Enhancement Prompt Template**: **Frontend Implementation**, **Visual Representations**, and **User Stories** (web + mobile), aligned with `combined-fe-folder-structure.md` conventions. Scope covers: *re‑authentication for sensitive actions*, *idle screen lock*, *session expiration/renewal*, *suspicious‑activity step‑up*, *organization re‑auth policy*, and *global sign‑out*.


---

### Frontend Implementation

#### A) Page Components (Routes)

**Web (Next.js App Router):**
```
Path: apps/web/app/(auth)/reauth/page.tsx
Purpose: Full‑page fallback for re‑authentication (password, passkey, TOTP/SMS, recovery).
Key Features:
  - Starts challenge and renders method tabs
  - Remember‑for‑15‑minutes option
  - Returns to original action on success

Path: apps/web/app/(auth)/lock/page.tsx
Purpose: Idle/session lock screen requiring re‑auth to unlock.
Key Features:
  - Shows masked user info and last active time
  - Supports passkey/password/MFA

Path: apps/web/app/(auth)/session-expired/page.tsx
Purpose: Session expired interstitial; prompts sign‑in or step‑up if refresh still valid.
Key Features:
  - Resume action after login
  - Explains why session expired

Path: apps/web/app/(settings)/(security)/sessions/page.tsx
Purpose: Session management — view and revoke active sessions (device, IP, last seen).
Key Features:
  - Revoke one/all
  - Suspicious flag and recommended action

Path: apps/web/app/(settings)/(security)/policy-reauth/page.tsx
Purpose: Org policy notice (re‑auth cadence); shows remaining time to next re‑auth.

Path: apps/web/app/(auth)/layout.tsx
Purpose: Shared auth frame; Toaster mount; ErrorBoundary + Sentry.

Path: apps/web/app/(auth)/*/loading.tsx
Path: apps/web/app/(auth)/*/error.tsx
Purpose: Route‑level skeletons and error fallbacks.
```

**Mobile (Expo Router):**
```
Path: apps/mobile/app/(tabs)/(authenticated)/reauth.tsx
Purpose: Full‑screen step‑up auth (modal on native).

Path: apps/mobile/app/(tabs)/(authenticated)/lock.tsx
Purpose: Idle lock screen; unlock via passkey/password/MFA.

Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/sessions.tsx
Purpose: Session management; revoke one/all.

Path: apps/mobile/app/(tabs)/(authenticated)/(settings)/security/policy-reauth.tsx
Purpose: Org re‑auth cadence interstitial.
```

> Handoff: AUTH‑6 integrates with **AUTH‑2** (MFA at sign‑in), **AUTH‑4** (enabled factors), and **FIN/CL/ACC** features that require step‑up prior to sensitive operations (payouts, contact change, API keys, delete account, etc.).

---

#### B) Shared Components

**UI Components** (`packages/ui/src/components/`):
```
Component: Modal / Drawer
Purpose: Re‑auth as modal overlay on desktop/mobile

Component: Tabs
Purpose: Switch between Password / Passkey / Code / Recovery

Component: PasswordInput, CodeInput, Button, Checkbox, Spinner, Alert/Banner
Purpose: Inputs and interactions for re‑auth

Component: Badge, Tooltip, Card, Table
Purpose: Sessions list, suspicious labels
```

**Feature Components** (`packages/lib/src/features/session/components/`):
```
Component: ReauthDialog.tsx
Purpose: Start challenge; render method tabs; resolve promise on success

Component: IdleLockScreen.tsx
Purpose: Full‑screen lock; unlock via any available factor

Component: SessionExpiredInterstitial.tsx
Purpose: Explain expiry; route to sign‑in or re‑auth

Component: SessionsTable.tsx
Purpose: List sessions; revoke actions; suspicious indicators

Component: PolicyReauthNotice.tsx
Purpose: Shows required cadence; countdown to next step‑up
```

**Domain‑Specific Components (web/mobile):**
```
Component: SensitiveActionGuard.tsx
Platform: Web
Purpose: Wraps protected UIs; invokes ReauthDialog when needed

Component: SensitiveActionSheet.tsx
Platform: Mobile
Purpose: Native sheet to request step‑up before continuing
```

---

#### C) Hooks & Data Fetching

**TanStack Query / Mutations** (`packages/lib/src/features/session/hooks/`):
```
Hook: useReauthStart.ts
POST /v1/auth/reauth/start { intent, deviceInfo }
→ { challengeId, methods:["password","passkey","totp","sms","recovery"], rememberWindowSec }

Hook: useReauthPassword.ts
POST /v1/auth/reauth/password { challengeId, password }
→ { success:true, rememberUntil?:ts }

Hook: useReauthTotp.ts
POST /v1/auth/reauth/totp { challengeId, code }
→ { success:true, rememberUntil?:ts }

Hook: useReauthSms.ts
POST /v1/auth/reauth/sms { challengeId, code }
→ { success:true, rememberUntil?:ts }

Hook: useWebAuthnAssertStart.ts
POST /v1/auth/webauthn/assert/start { challengeId }
→ { publicKeyCredentialRequestOptions }

Hook: useWebAuthnAssertFinish.ts
POST /v1/auth/webauthn/assert/finish { credential }
→ { success:true, rememberUntil?:ts }

Hook: useRecoveryCode.ts
POST /v1/auth/reauth/recovery { challengeId, code }
→ { success:true }

Hook: useSessionStatus.ts
GET /v1/auth/session/status
→ { user, locked:boolean, expiresAt, idleLockAt?, policyNextReauthAt? }

Hook: useKeepAlive.ts
POST /v1/auth/session/keep-alive
→ { expiresAt }

Hook: useLockSession.ts
POST /v1/auth/session/lock → { locked:true }

Hook: useUnlockSession.ts
POST /v1/auth/session/unlock { factor }
→ { locked:false }

Hook: useSessionsQuery.ts
GET /v1/auth/sessions
→ { sessions:[{ id, ua, ip, lastSeen, createdAt, current:boolean, suspicious?:boolean }] }

Hook: useRevokeSession.ts
DELETE /v1/auth/sessions/{id}

Hook: useRevokeAllSessions.ts
POST /v1/auth/sessions/revoke-all

Hook: useRiskSignals.ts
GET /v1/auth/risk/signals → { ipChange:boolean, newDevice:boolean, impossibleTravel:boolean }
```

**State Management (Zustand)** (`packages/lib/src/stores/session-store.ts`):
```
interface SessionState {
  challengeId?: string|null;
  methods?: string[];
  rememberUntil?: number|null;
  locked?: boolean;
  expiresAt?: number|null;
  gatingIntent?: string|null; // e.g., "payout.withdraw"
  set: (p: Partial<SessionState>) => void;
  clear: () => void;
}
```

---

#### D) Utilities & Helpers

**Validation Schemas (Zod)** (`packages/lib/src/schemas/session/`):
```
reauthPasswordSchema.ts → password: string().min(8)
reauthCodeSchema.ts     → code: string().regex(/^\d{6}$/)
```

**Security & Timing Utils** (`packages/lib/src/utils/`):
```
security/webauthn.ts        → navigator.credentials.get() wrapper
time/idleTimer.ts           → activity listeners + idle thresholds
time/countdown.ts           → server‑synced countdowns
format/maskDevice.ts        → present UA/IP succinctly
risk/policy.ts              → client helpers for showing risk reasons
```

**Types** (`packages/lib/src/types/session.ts`):
```
ReauthMethod, ReauthStartResponse, ReauthVerifyResponse, SessionSummary, RiskSignals
```

---

#### E) Mobile‑Specific Components

**Navigation Components:**
```
Component: StepUpStack (Expo Router)
Path: apps/mobile/app/(tabs)/(authenticated)/_layout.tsx (register sheet/modal routes)
Screens:
  - reauth (sheet/full)
  - lock
  - sessions
  - policy-reauth
```

**Native Features:**
```
- Local biometric prompt (passkey platform authenticator)
- Haptics on unlock
- AppState to trigger idle lock when backgrounded (optional)
```

---

#### F) Layout Components

**Web:**
```
Layout: (auth) Layout
Path: apps/web/app/(auth)/layout.tsx
Features: Minimal chrome, Toaster, Sentry, ErrorBoundary
```

**Mobile:**
```
Layout: (tabs)/(authenticated)/_layout.tsx
Features: Register global modal routes for re‑auth & lock
```

---

#### G) Error Boundaries & Loading States

**Error Handling:**
```
Routes: (auth)/*/error.tsx
Components: Inline banners for wrong code/password, WebAuthn errors
```

**Loading States:**
```
Routes: (auth)/*/loading.tsx
Components: Spinner + dimmed modal overlay during verification
```

---

### Visual Representations

#### Screen 1 — Re‑Auth Modal (Password)

```
┌──────────────────────────────────────────────┐
│ Re‑authenticate to continue                  │
├──────────────────────────────────────────────┤
│ Sensitive action: Withdraw funds             │
│                                              │
│ Password *  [ •••••••• ]   👁                │
│ [ ] Remember this device for 15 minutes      │
│                                              │
│ [ Cancel ]                           [ Verify ] │
│ Methods:  Password | Passkey | Code | Recovery  │
└──────────────────────────────────────────────┘
```

#### Screen 2 — Re‑Auth Modal (Passkey)

```
┌──────────────────────────────────────────────┐
│ Use your passkey                             │
├──────────────────────────────────────────────┤
│ A browser prompt will appear.                │
│ [ Verify with passkey ]                      │
│ Methods:  Password | Passkey | Code | Recovery │
└──────────────────────────────────────────────┘
```

#### Screen 3 — Re‑Auth Modal (Code: TOTP/SMS)

```
┌──────────────────────────────────────────────┐
│ Enter 6‑digit code                           │
├──────────────────────────────────────────────┤
│ Code: [ _ ][ _ ][ _ ][ _ ][ _ ][ _ ]         │
│ Didn’t get it?  [Resend 27s]                 │
│ Methods:  Password | Passkey | Code | Recovery │
└──────────────────────────────────────────────┘
```

#### Screen 4 — Idle Lock Screen

```
┌──────────────────────────────────────────────┐
│ 🔒 Session locked due to inactivity          │
├──────────────────────────────────────────────┤
│ Signed in as:  e***@exa…                     │
│ Last active:  12m ago                        │
│ [ Unlock with passkey ]                      │
│ [ Or use password / code ]                   │
└──────────────────────────────────────────────┘
```

#### Screen 5 — Session Expired Interstitial

```
┌──────────────────────────────────────────────┐
│ Your session expired                         │
├──────────────────────────────────────────────┤
│ For your security, please sign in again.     │
│ [ Sign in ]   [ Use step‑up (if available) ] │
└──────────────────────────────────────────────┘
```

#### Screen 6 — Suspicious Activity Step‑Up

```
┌──────────────────────────────────────────────┐
│ Extra verification required                  │
├──────────────────────────────────────────────┤
│ We noticed a sign‑in from a new location.    │
│ Please verify to continue.                   │
│ [ Verify with passkey ]  [ Use password ]    │
└──────────────────────────────────────────────┘
```

---

## User Stories — Complete

> Abbrev: **W**=Web, **M**=Mobile. Each story follows “As a … I want … so that …” and includes *components, routes, APIs*, with **≥5 Happy** & **≥10 Bad** scenarios.

### Story 1 — Trigger Step‑Up for Sensitive Actions

**Story**
```
As a Platform
I want to require re‑authentication before sensitive actions
So that unauthorized users cannot perform critical operations
```

**Components & Routes**
- **W:** `SensitiveActionGuard` wraps actions; launches `ReauthDialog`
- **M:** `SensitiveActionSheet` in native flows
- **APIs:** `POST /v1/auth/reauth/start`

**Acceptance Criteria — ✅ Happy**
1. **AC1.1** Guard starts a challenge with intent (e.g., `payout.withdraw`).
2. **AC1.2** Available methods reflect user’s enabled factors.
3. **AC1.3** On success, the original action proceeds automatically.
4. **AC1.4** Remember window respected (e.g., 15 min) per device.
5. **AC1.5** Analytics: `reauth.start` and `reauth.success` fire once.

**Acceptance Criteria — ❌ Bad**
1. **AC1.6** Start fails (5xx) → non‑blocking banner; user can retry.
2. **AC1.7** No methods available → show help to enable a factor; action blocked.
3. **AC1.8** Challenge expired while acting → restart step‑up gracefully.
4. **AC1.9** Cross‑tab conflict (two dialogs) → serialize; keep one active.
5. **AC1.10** Idle auto‑lock occurs during step‑up → resume dialog after unlock.
6. **AC1.11** “Remember” disabled by policy → checkbox hidden and enforced.
7. **AC1.12** i18n key missing → English fallback; key logged.
8. **AC1.13** A11y: focus lands on dialog title; escape closes safely.
9. **AC1.14** Back navigation mid‑dialog → action canceled; state cleared.
10. **AC1.15** Analytics fails → non‑blocking; queued retry.

---

### Story 2 — Re‑Auth via Password

**Story**
```
As a User
I want to confirm my password to continue a sensitive action
So that I can proceed securely
```

**Components & Routes**
- **W/M:** `ReauthDialog` → Password tab, `PasswordInput`
- **APIs:** `POST /v1/auth/reauth/password`

**Acceptance Criteria — ✅ Happy**
1. **AC2.1** Valid password verifies; dialog closes; action resumes.
2. **AC2.2** “Remember for 15 minutes” sets `rememberUntil` per server.
3. **AC2.3** Show/Hide toggle accessible and announced.
4. **AC2.4** Keyboard Enter submits; button shows spinner.
5. **AC2.5** Analytics `reauth.password.success` logs once.

**Acceptance Criteria — ❌ Bad**
1. **AC2.6** Wrong password → inline error; rate‑limit after N attempts.
2. **AC2.7** Password manager autofill mismatch → validation explains.
3. **AC2.8** Network/5xx → retry without clearing input.
4. **AC2.9** Session expired between start and verify → redirect to Sign‑in.
5. **AC2.10** CSRF missing → hard refresh and retry.
6. **AC2.11** Caps‑lock warning appears when appropriate.
7. **AC2.12** Field loses focus unexpectedly → keep cursor position.
8. **AC2.13** Clipboard paste disabled by policy → explain if blocked.
9. **AC2.14** A11y: error announced via aria‑live; focus to field on error.
10. **AC2.15** Locale input method issues → normalize and preserve characters.

---

### Story 3 — Re‑Auth via Passkey (WebAuthn)

**Story**
```
As a User with a passkey
I want to verify using my device authenticator
So that I can complete step‑up quickly and securely
```

**Components & Routes**
- **W:** `ReauthDialog` → Passkey tab; `useWebAuthnAssertStart/Finish`
- **M:** Native platform authenticator via passkeys
- **APIs:** `POST /v1/auth/webauthn/assert/start`, `/finish`

**Acceptance Criteria — ✅ Happy**
1. **AC3.1** Browser prompt appears; user verifies; success returns to action.
2. **AC3.2** Unsupported browser disables button with tooltip.
3. **AC3.3** “Remember 15 min” stored when policy allows.
4. **AC3.4** Multiple credentials supported; correct one chosen by authenticator.
5. **AC3.5** Analytics `reauth.passkey.success` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC3.6** User cancels prompt → dialog remains; no error state.
2. **AC3.7** Origin/RP mismatch → hard error; guidance to official domain.
3. **AC3.8** Authenticator locked → prompt to unlock and retry.
4. **AC3.9** Network failure finishing → retry w/ original challenge.
5. **AC3.10** Duplicate assertion → server idempotency returns success.
6. **AC3.11** Attestation policy blocks → explain and allow platform only.
7. **AC3.12** Time skew → re‑start flow seamlessly.
8. **AC3.13** A11y: focus returns to Verify button after cancel.
9. **AC3.14** Mobile lacks support → suggest password/TOTP fallback.
10. **AC3.15** Pop‑up blockers interfere → use WebAuthn without pop‑ups.

---

### Story 4 — Re‑Auth via TOTP/SMS Code

**Story**
```
As a User with codes enabled
I want to enter a 6‑digit code (TOTP or SMS)
So that I can complete step‑up
```

**Components & Routes**
- **W/M:** `ReauthDialog` → Code tab; `CodeInput`
- **APIs:** `POST /v1/auth/reauth/totp`, `POST /v1/auth/reauth/sms`

**Acceptance Criteria — ✅ Happy**
1. **AC4.1** Code input auto‑advances; paste supported; success resumes action.
2. **AC4.2** Resend (SMS) respects cooldown; accessible countdown.
3. **AC4.3** TOTP allows ±1 time window tolerance.
4. **AC4.4** “Remember” option respected if policy allows.
5. **AC4.5** Analytics `reauth.code.success` emitted once.

**Acceptance Criteria — ❌ Bad**
1. **AC4.6** Wrong/expired code → inline error; after N → lockout cooldown.
2. **AC4.7** SMS delivery delay → show tips and retry path.
3. **AC4.8** Pasted non‑digits → sanitized; error if invalid.
4. **AC4.9** Phone number change mid‑flow → restart challenge.
5. **AC4.10** Network/5xx → retry; input preserved.
6. **AC4.11** Screen reader reads cells out of order → single‑field mode option.
7. **AC4.12** Rate‑limit on resend → timer and tooltip show next attempt.
8. **AC4.13** i18n digits normalized to ASCII for verification.
9. **AC4.14** Timezone skew affects TOTP window → rely on server time.
10. **AC4.15** Concurrent verify in two tabs → second tab blocked gracefully.

---

### Story 5 — Recovery Code Fallback

**Story**
```
As a Locked‑out User
I want to use a recovery code as a last resort
So that I can proceed with the sensitive action
```

**Components & Routes**
- **W/M:** `ReauthDialog` → Recovery tab
- **APIs:** `POST /v1/auth/reauth/recovery`

**Acceptance Criteria — ✅ Happy**
1. **AC5.1** Valid recovery code unlocks and consumes the code.
2. **AC5.2** Prompt to regenerate a new set after success.
3. **AC5.3** Analytics logs usage (no code value). 
4. **AC5.4** Action resumes seamlessly; remember window respected if allowed.
5. **AC5.5** Audit entry created with masked metadata.

**Acceptance Criteria — ❌ Bad**
1. **AC5.6** Invalid/used code → inline error; attempts limited.
2. **AC5.7** No codes available → CTA to generate in MFA settings.
3. **AC5.8** Copy/paste includes spaces/dashes → normalized.
4. **AC5.9** Concurrent usage across devices → first wins; others fail.
5. **AC5.10** Network/5xx → retry without losing input.
6. **AC5.11** A11y: code field labeled and announced on error.
7. **AC5.12** Analytics failure → non‑blocking.
8. **AC5.13** Policy forbids recovery codes for step‑up → show reason.
9. **AC5.14** Rate‑limited after many failures → lock and backoff.
10. **AC5.15** Session expires mid‑flow → redirect to Sign‑in and return.

---

### Story 6 — Remember Device/Window (Grace Period)

**Story**
```
As a Frequent User
I want to remember step‑up for a short window
So that I don’t have to re‑auth for every action
```

**Components & Routes**
- **W/M:** `ReauthDialog` with “Remember 15 min” checkbox
- **APIs:** Included in verify responses (`rememberUntil`)

**Acceptance Criteria — ✅ Happy**
1. **AC6.1** When checked and policy allows, server returns `rememberUntil`.
2. **AC6.2** Subsequent guarded actions within window skip step‑up.
3. **AC6.3** Different intents reuse the same window (configurable).
4. **AC6.4** Clearing cookies resets window; UI reflects change.
5. **AC6.5** Analytics `reauth.remember.enabled` recorded.

**Acceptance Criteria — ❌ Bad**
1. **AC6.6** Policy forbids remember → checkbox hidden; no client storage.
2. **AC6.7** Clock skew shows negative time → clamp to zero; refetch.
3. **AC6.8** Private mode blocks storage → remember silently disabled.
4. **AC6.9** Multiple tabs → shared storage ensures consistent skip.
5. **AC6.10** Cross‑device attempt → remember not transferable.
6. **AC6.11** Security downgrade (new IP/device) → force step‑up despite window.
7. **AC6.12** User manually locks session → window cleared.
8. **AC6.13** Analytics PII leak (IP/UA) → scrubbed summary only.
9. **AC6.14** “Remember” persists after logout → must be cleared.
10. **AC6.15** UI indicates remembered when not → reconcile on focus.

---

### Story 7 — Idle Screen Lock

**Story**
```
As a Security‑conscious User
I want my session to lock after inactivity
So that others can’t use my account if I step away
```

**Components & Routes**
- **W:** `IdleLockScreen`, route `(auth)/lock/page.tsx`
- **M:** `lock.tsx`
- **APIs:** `POST /v1/auth/session/lock`, `POST /v1/auth/session/unlock`

**Acceptance Criteria — ✅ Happy**
1. **AC7.1** Idle timer triggers lock after configured threshold.
2. **AC7.2** Unlock accepts passkey/password/MFA.
3. **AC7.3** Backgrounding the app can trigger lock (configurable).
4. **AC7.4** Successful unlock restores app state.
5. **AC7.5** Analytics: `session.lock` and `.unlock` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC7.6** False lock during media playback → user activity keeps alive.
2. **AC7.7** Network offline → local lock still protects; unlock with local UI then server confirm.
3. **AC7.8** A11y: focus stuck on hidden elements → fixed by trap.
4. **AC7.9** Aggressive timer locks while typing → keystrokes counted as activity.
5. **AC7.10** Unlock attempts rate‑limited after failures.
6. **AC7.11** Time drift → use monotonic timers; sync with server on unlock.
7. **AC7.12** Mobile keyboard overlap → scroll into view.
8. **AC7.13** Passkey prompt canceled → remain locked; can switch method.
9. **AC7.14** Dark mode contrast insufficient → AA compliant.
10. **AC7.15** Auto‑lock disabled by policy → setting hidden; always lock.

---

### Story 8 — Session Expiration & Renewal

**Story**
```
As a Platform
I want sessions to expire and renew safely
So that accounts remain secure without surprising users
```

**Components & Routes**
- **W/M:** `SessionExpiredInterstitial`, `KeepAlive` hook
- **APIs:** `GET /v1/auth/session/status`, `POST /v1/auth/session/keep-alive`

**Acceptance Criteria — ✅ Happy**
1. **AC8.1** Sliding expiration extends with activity via keep‑alive.
2. **AC8.2** On expiry, interstitial explains and routes to Sign‑in.
3. **AC8.3** After re‑login, user returns to last state/intended action.
4. **AC8.4** Refresh token rotation done server‑side; no double logins.
5. **AC8.5** Analytics: `session.expire` and `session.renew` captured.

**Acceptance Criteria — ❌ Bad**
1. **AC8.6** Keep‑alive fails silently → warn and reduce polling.
2. **AC8.7** Clock skew causes premature expiry → rely on server timestamps.
3. **AC8.8** Multiple tabs race keep‑alive → leader election pattern.
4. **AC8.9** Token refresh error → sign‑in required; unsaved work preserved.
5. **AC8.10** Interstitial loops on redirect → break with one‑time token.
6. **AC8.11** A11y: interstitial readable by screen readers.
7. **AC8.12** Offline mode → pause keep‑alive; resume on reconnect.
8. **AC8.13** Analytics blocked → non‑blocking.
9. **AC8.14** Stale cached status → refetch on visibility change.
10. **AC8.15** Excessive polling drains battery → adaptive intervals.

---

### Story 9 — Suspicious Activity Step‑Up

**Story**
```
As a Platform
I want to require step‑up when risk signals are detected
So that risky sessions are re‑verified
```

**Components & Routes**
- **W/M:** `SensitiveActionGuard` listens to `useRiskSignals`
- **APIs:** `GET /v1/auth/risk/signals`

**Acceptance Criteria — ✅ Happy**
1. **AC9.1** New device/IP/impossible‑travel triggers a re‑auth prompt.
2. **AC9.2** After success, signals reset; action continues.
3. **AC9.3** Email/security alert sent to user.
4. **AC9.4** Remember window ignored when risk is high.
5. **AC9.5** Analytics records signal types (no PII).

**Acceptance Criteria — ❌ Bad**
1. **AC9.6** False positives → user can mark device trusted; reduces prompts.
2. **AC9.7** Signals API unavailable → default to cautious (prompt once).
3. **AC9.8** Repeated prompts loop → cooldown between prompts.
4. **AC9.9** VPN/CGNAT misclassifies → explain and allow trust device.
5. **AC9.10** Notification email bounces → still allow step‑up.
6. **AC9.11** Excess prompts on mobile roaming → heuristics toned down.
7. **AC9.12** A11y: prompt announced; focus managed.
8. **AC9.13** Multiple tabs prompt at once → dedupe to one.
9. **AC9.14** Policy forces prompt but no methods available → block and explain.
10. **AC9.15** Analytics leaks IP/geo → only categorical flags captured.

---

### Story 10 — Session Management (View & Revoke)

**Story**
```
As a User
I want to view my active sessions and revoke access
So that I can control which devices stay logged in
```

**Components & Routes**
- **W:** `(settings)/(security)/sessions/page.tsx` → `SessionsTable`
- **M:** `…/settings/security/sessions.tsx`
- **APIs:** `GET /v1/auth/sessions`, `DELETE /v1/auth/sessions/{id}`, `POST /v1/auth/sessions/revoke-all`

**Acceptance Criteria — ✅ Happy**
1. **AC10.1** Table shows device, location (approx), IP, last seen, current flag.
2. **AC10.2** Revoke one prompts confirm and logs out that session promptly.
3. **AC10.3** Revoke all logs out everywhere except current (configurable).
4. **AC10.4** Email notification summarizing revocations sent.
5. **AC10.5** Analytics: `sessions.revoke_one` / `.revoke_all` logged.

**Acceptance Criteria — ❌ Bad**
1. **AC10.6** Revoke fails → error; state unchanged.
2. **AC10.7** Already revoked → idempotent success.
3. **AC10.8** GeoIP unavailable → label “Unknown location”.
4. **AC10.9** Pagination and search accessible and performant.
5. **AC10.10** Offline → action queued and retried.
6. **AC10.11** Session expired mid‑action → re‑auth then retry.
7. **AC10.12** Rate‑limit on revoke‑all → throttle with feedback.
8. **AC10.13** A11y: table headers labeled; buttons have aria‑labels.
9. **AC10.14** i18n missing → English fallback.
10. **AC10.15** Concurrent revoke and logout race → serialize on server.

---

### Story 11 — Organization Re‑Auth Cadence

**Story**
```
As an Organization Admin
I want members to re‑authenticate every N hours
So that long‑lived sessions are periodically verified
```

**Components & Routes**
- **W/M:** `PolicyReauthNotice`
- **APIs:** `GET /v1/auth/session/status` (policy fields), server enforcement middleware

**Acceptance Criteria — ✅ Happy**
1. **AC11.1** Countdown to required re‑auth displays in settings and banner.
2. **AC11.2** On deadline, guarded routes require step‑up before access.
3. **AC11.3** After success, countdown resets and access restored.
4. **AC11.4** Admin policy updates reflected within 60s.
5. **AC11.5** Analytics: cadence completions recorded.

**Acceptance Criteria — ❌ Bad**
1. **AC11.6** Policy fetch fails → default to last known; log warning.
2. **AC11.7** Deadline passed while offline → prompt immediately on return.
3. **AC11.8** Multiple banners overlap → consolidate to one.
4. **AC11.9** Users without factors → require password; if disabled, block with guidance.
5. **AC11.10** Time skew confusion → show absolute UTC + local.
6. **AC11.11** Accessibility of countdown and banner ensured.
7. **AC11.12** i18n keys missing → fallback.
8. **AC11.13** Inconsistent state across tabs → refetch on focus.
9. **AC11.14** API throttling → backoff; don’t spam server.
10. **AC11.15** Analytics payloads leak policy values → aggregate only.

---

### Story 12 — Global Sign‑Out (Security Incident)

**Story**
```
As a Security Team
I want to remotely sign out users or revoke all sessions
So that risk is contained quickly
```

**Components & Routes**
- **W/M:** Admin‑triggered server action; user UI receives event to sign out
- **APIs:** `POST /v1/auth/sessions/revoke-all` (server/admin), client listens to SSE/websocket or polling

**Acceptance Criteria — ✅ Happy**
1. **AC12.1** Client receives global‑logout signal → redirects to Sign‑in.
2. **AC12.2** Local data cleared; protected routes blocked.
3. **AC12.3** Banner explains reason (generic, no sensitive details).
4. **AC12.4** Current work preserved where possible (local draft).
5. **AC12.5** Analytics `sessions.global_logout` recorded (no PII).

**Acceptance Criteria — ❌ Bad**
1. **AC12.6** Signal missed offline → caught at next API call.
2. **AC12.7** Local storage not cleared → fix and add regression test.
3. **AC12.8** Re‑login loop due to cache → force full reload.
4. **AC12.9** SSE/WebSocket unavailable → poll as fallback.
5. **AC12.10** A11y: banner focus management on redirect.
6. **AC12.11** i18n fallback on incident copy.
7. **AC12.12** Analytics blocked → non‑blocking.
8. **AC12.13** Partial revoke (server) → client still logs out; server retries.
9. **AC12.14** Race with revoke‑all initiated by user → idempotent.
10. **AC12.15** Phishing risk (fake banner) → only server‑signed signals honored.

---

## Non‑Functional Requirements (AUTH‑6 scope)

- **Performance:** Start/verify roundtrips < 600ms p95; lock/unlock < 400ms p95
- **Security:** CSRF, rate limits, WebAuthn verification, single‑use challenges, risk‑based prompts, refresh rotation, audit logging; Sentry
- **Privacy:** No secrets in logs/analytics; masked device/IP labels; POST‑redirect for tokens
- **Compatibility:** Web (Chrome/Firefox/Safari/Edge), iOS 14+, Android 10+; passkeys where supported
- **Testing:** Unit tests (schemas/hooks/utils), e2e (password/passkey/TOTP/SMS/recovery), a11y (axe), risk‑prompt and cadence policy tests

---

**End of AUTH‑6 Enhancements (Complete).**
